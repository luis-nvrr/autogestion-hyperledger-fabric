package chaincode

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}
type Student struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Year     int    `json:"year"`
}
type Grade struct {
	Id           string   `json:"id"`
	Grade        int      `json:"grade"`
	Date         string   `json:"date"`
	Student      *Student `json:"student"`
	Instance     string   `json:"instance"`
	Observations string   `json:"observations"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	mspid, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return err
	}

	if mspid == "Org3MSP" {
		return errors.New("invalid user")
	}

	student := Student{79581, "Luis", "Navarro", 5}

	grades := []Grade{
		{"grade1", 9, "2020-05-12", &student, "exam", "great!"},
		{"grade2", 5, "2020-05-12", &student, "lab", "not so great!"},
	}

	for _, grade := range grades {
		assetJSON, err := json.Marshal(grade)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(grade.Id, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

func (s *SmartContract) GetAllGrades(ctx contractapi.TransactionContextInterface) ([]*Grade, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var grades []*Grade
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var grade Grade
		err = json.Unmarshal(queryResponse.Value, &grade)
		if err != nil {
			return nil, err
		}
		grades = append(grades, &grade)
	}

	return grades, nil
}

func (s *SmartContract) CreateGrade(ctx contractapi.TransactionContextInterface,
	gradeValue int,
	date string,
	studentId int,
	studentName string,
	studentLastName string,
	studentYear int,
	instance string,
	observations string) (*Grade, error) {

	mspid, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, err
	}

	if mspid == "Org3MSP" {
		return nil, errors.New("invalid user")
	}

	gradeId := fmt.Sprintf("%d-%d-%s-%d", studentId, studentYear, instance, gradeValue)
	exists, err := s.GradeExists(ctx, gradeId)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("the grade %s already exists", gradeId)
	}

	grade := Grade{
		gradeId,
		gradeValue,
		date,
		&Student{studentId, studentName, studentLastName, studentYear},
		instance,
		observations,
	}

	gradeJSON, err := json.Marshal(grade)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(gradeId, gradeJSON)
	if err != nil {
		return nil, err
	}

	return &grade, nil
}

func (s *SmartContract) GradeExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}
