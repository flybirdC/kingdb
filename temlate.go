package kingdb

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type MyEntity struct {

	UserA string `json:"user_a"`
	UserB string `json:"user_b"`
	FileHash string `json:"file_hash"`
	FileHttp string `json:"file_http"`

}
//users db :username,userpass,dbname
//	response := stub.InvokeChaincode("contractdbcc",[][]byte{[]byte(args[0]),[]byte(args[1])},"mychannel")
func GetDB(stub shim.ChaincodeStubInterface,args []string) peer.Response {

	response := stub.InvokeChaincode("dbcc",[][]byte{[]byte("imagefile"),[]byte("wangchong"),[]byte("19861120wang")},"mychannel")

	jsonbyte := response.GetPayload()

	return shim.Success(jsonbyte)
}

//user entity


