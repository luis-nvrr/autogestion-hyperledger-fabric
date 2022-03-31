package chaincode

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type YearOfStudy int64

const (
	First YearOfStudy = iota
	Second
	Third
	Fourth
	Fifth
)

type Student struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	LastName string      `json:"lastName"`
	Year     YearOfStudy `json:"year"`
}
type Grade struct {
	Id           string    `json:"id"`
	Grade        float64   `json:"grade"`
	Date         time.Time `json:"date"`
	Student      *Student  `json:"student"`
	Instance     string    `json:"instance"`
	Observations string    `json:"observations"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) (*[]Grade, error) {
	mspid, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, err
	}

	if mspid == "Org3MSP" {
		return nil, errors.New("invalid user")
	}

	student := Student{79581, "Luis", "Navarro", Fifth}

	grades := []Grade{
		{"grade1", 9, time.Date(2022, time.Month(time.April), 14, 12, 30, 0, 0, time.UTC), &student, "exam", "great!"},
		{"grade2", 5, time.Date(2022, time.Month(time.April), 14, 12, 30, 0, 0, time.UTC), &student, "lab", "not so great!"},
	}

	for _, grade := range grades {
		assetJSON, err := json.Marshal(grade)
		if err != nil {
			return nil, err
		}

		err = ctx.GetStub().PutState(grade.Id, assetJSON)
		if err != nil {
			return nil, fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return &grades, nil
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
	gradeValue float64,
	date time.Time,
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

	gradeId := fmt.Sprintf("%d-%d-%s-%.2f", studentId, studentYear, instance, gradeValue)
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
		&Student{studentId, studentName, studentLastName, YearOfStudy(studentYear)},
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
