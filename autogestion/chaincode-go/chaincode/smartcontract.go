package chaincode

import (
	"encoding/json"
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
	ID       int         `json:"ID"`
	Name     string      `json:"Name"`
	LastName string      `json:"LastName"`
	Year     YearOfStudy `json:"Year"`
}

type GradeInstance int64

const (
	Exam GradeInstance = iota
	Lab
	Presentation
)

type Grade struct {
	ID           string        `json:"ID"`
	Value        float64       `json:"Value"`
	Timestamp    time.Time     `json:"Timestamp"`
	Student      *Student      `json:"Student"`
	Instance     GradeInstance `json:"GradeInstance"`
	Observations string        `json:"Observations"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	student := Student{79581, "Luis", "Navarro", Fifth}

	grades := []Grade{
		{ID: "1", Value: 9, Timestamp: time.Date(2022, time.Month(time.April), 14, 12, 30, 0, 0, time.UTC), Student: &student, Instance: Exam, Observations: "great!"},
		{ID: "2", Value: 5, Timestamp: time.Date(2022, time.Month(time.April), 15, 12, 30, 0, 0, time.UTC), Student: &student, Instance: Presentation, Observations: "not great!"},
	}

	for _, grade := range grades {
		assetJSON, err := json.Marshal(grade)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(grade.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Grade, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
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
