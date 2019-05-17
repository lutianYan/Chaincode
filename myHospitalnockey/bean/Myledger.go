package bean

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"strconv"
	//"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type MyLedgerno struct {

	IdNumber string `json:"idNumber"`

	Department string `json:"department"`

	FinalInfo string `json:"finalInfo"`
}

//插入账本数据
func InsertData(stub shim.ChaincodeStubInterface, args []string) error {
	var my1 MyLedgerno
	my1.IdNumber = args[0]
	my1.Department = args[1]
	my1.FinalInfo = args[2]

	my1JsonBytes, err := json.Marshal(&my1) // Json序列化
	if err != nil {
		return fmt.Errorf("Json serialize Compact fail while Loan")
	}

	key := args[0]


	err = stub.PutState(key, my1JsonBytes)
	if err != nil {
		return fmt.Errorf("Failed to PutState while InsertData")
	}
	message := "Event send data is here!"
	err = stub.SetEvent("evtsender", []byte(message))
	if err != nil {
		return fmt.Errorf("Event send data is not here!")
	}
	return nil
}

//查询一个数据记录
func ReadRecord(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error
	fmt.Print(len(args))
	if len(args) != 2 {
		return nil, fmt.Errorf("11111111111111111111111111111111111")
	}


	key = args[0]

	//获取key的value值
	valAsbytes, err := stub.GetState(key) //get the record from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, fmt.Errorf(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Record does not exist: " + key + "\"}"
		return nil, fmt.Errorf(jsonResp)
	}

	return valAsbytes, nil
}

func GetHistoryRecords(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var my1 MyLedgerno
	if len(args) < 1 {
		return nil, fmt.Errorf("Incorrect number of arguments. Expecting 2")
	}

	key := args[0]

	fmt.Printf("- start GetHistoryRecords: %s\n", key)

	writeKey,_:=stub.GetKeyForKey(key)
	fmt.Println("chachachachulailema")

	resultsIterator, err := stub.GetHistoryForKey(string(writeKey))
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	//buffer.WriteString("[")

	//bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		err = json.Unmarshal(response.Value, &my1) // Json序列化
		if err != nil {
			return buffer.Bytes(), fmt.Errorf("fanxuliehuawenti")
		}

		buffer.Write(response.Value)

	}

	return buffer.Bytes(), nil
}
