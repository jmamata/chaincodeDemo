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

type Prescription  struct {
	Disease            	string 		`json:"disease"`
	Medication          string 		`json:"medication"`
	Duration      		string 		`json:"duration"`
}

type Patient struct {
	Name            	string 			`json:"name"`
	dob           		string 			`json:"dob"`
	CurrentProblem      string 			`json:"currentproblem"`
	allergies      		string 			`json:"currentproblem"`
	Prescriptions 		[]Prescription 	`json:"prescriptions"`
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
	}else if function == "enter_patient_prescription" {
		return t.enter_patient_prescription(stub, args)
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

	if len(args) != 4 { 
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}


	patient := Patient{}
	patient.Name = args[0]
	patient.dob = args[1]
	patient.CurrentProblem = args[2]
	patient.allergies =  args[3]
	bytes, err := json.Marshal(&patient)

	if err != nil { 
		return nil, errors.New("Error converting Patient record")
	}

	err = stub.PutState(PREFIX_PATIENT + args[0], bytes)

	if err != nil { 
		return nil, errors.New("Error storing Patient record") 
	}

	return nil, nil
}


func (t *SimpleChaincode) get_patient_details(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 { 
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	patient, err := stub.GetState(PREFIX_PATIENT + args[0])
	if err != nil {
		return nil, errors.New("Failed to get state for " + args[0])
	}

	return patient, nil
}

func (t *SimpleChaincode) enter_patient_prescription(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	//		0		 1		   2		  3
	//	   Name   Disease  Medication  Duration

	if len(args) != 4 { 
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	bytes, err := stub.GetState(PREFIX_PATIENT + args[0])
	if err != nil { 
		return nil, errors.New("No Patient with name " + args[0])
	}


	prescription := Prescription{}
	prescription.Disease = args[1]
	prescription.Medication = args[2]
	prescription.Duration = args[3]

	var patient Patient
 	err = json.Unmarshal(bytes,&patient)
 	patient.Prescriptions = append(patient.Prescriptions, prescription)

	bytes, err = json.Marshal(&patient)
	if err != nil { 
		return nil, errors.New("Error converting Patient record") 
	}

	err = stub.PutState(PREFIX_PATIENT + args[0], bytes)
	if err != nil { 
		return nil, errors.New("Error storing Patient record") 
	}

	return nil, nil
}
