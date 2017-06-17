package main

import(
	"fmt"
	"errors"
	"encoding/json"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const   PREFIX_PATIENT =  "patient"

type SimpleChaincode struct {
}

type Patient struct {
	Name            	string 		`json:"name"`
	Model           	string 		`json:",string"`
	CurrentProblem      	string 		`json:"currentproblem"`
	allergies      		string 		`json:"currentproblem"`
}



func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    return nil, nil
}
 
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "get_patient_details" {
		return t.get_patient_details(stub, args)
	}

	return nil, errors.New("Received unknown function invocation " + function)

}
 
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "enter_patient_details" {
		return t.enter_patient_details(stub, args)
	}

	return nil, errors.New("Received unknown function invocation " + function)
   
}
 
func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Println("Could not start SampleChaincode")
    } else {
        fmt.Println("SampleChaincode successfully started")
    }
 
}




func (t *SimpleChaincode) enter_patient_details(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {



	return nil, nil
}


func (t *SimpleChaincode) get_patient_details(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 { 
		return false, errors.New("Incorrect number of arguments. Expecting 1")
	}

	patient, err := stub.GetState(PREFIX_PATIENT + args[0])
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + args[0] + "\"}"
		return nil, errors.New(jsonResp)
	}

	return patient, nil
}
