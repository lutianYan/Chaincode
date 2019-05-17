package utils

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//获取当前操作智能合约成员的具体名称
func GetCreatorName(stub shim.ChaincodeStubInterface) (string, error) {
	name, err := GetCreator(stub)
	if err != nil {
		return "", err
	}
	//格式化当前操作的智能合约操作成员名称
	memberName := name[(strings.Index(name, "@") + 1):strings.LastIndex(name, ".example.com")]
	return memberName, nil
}

//获取操作成员
func GetCreator(stub shim.ChaincodeStubInterface) (string, error) {
	creatorByte, _ := stub.GetCreator()
	certStart := bytes.IndexAny(creatorByte, "------BEGIN")
	if certStart == -1 {
		fmt.Errorf("NO certificate found")
	}
	certText := creatorByte[certStart:]
	b1, _ := pem.Decode(certText)
	if b1 == nil {
		fmt.Errorf("Could not decode the PEM structure")
	}
	cert, err := x509.ParseCertificate(b1.Bytes)
	if err != nil {
		fmt.Errorf("ParseCertficate failed")
	}
	uname := cert.Subject.CommonName
	return uname, nil

}
