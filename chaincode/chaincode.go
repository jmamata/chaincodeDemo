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
	CurrentProblem      string 		`json:"currentproblem"`
	allergies      		string 		`json:"currentproblem"`
}



func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    return nil, nil
}
 
func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "get_patient_details" {
		return t.get_patient_details(stub, args)
	}

	return nil, errors.New("Received unknown function invocation " + function)

}
 
func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
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




func (t *SimpleChaincode) enter_patient_details(stub shim.ChaincodeStubInterface, args []string) (bool, error) {



	return true, nil
}


func (t *SimpleChaincode) get_patient_details(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	return nil, nil
}
