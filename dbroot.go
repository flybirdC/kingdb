package kingdb

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type DBroot struct {

}


func (dbroot *DBroot)Init(stub shim.ChaincodeStubInterface) peer.Response  {

	args := stub.GetStringArgs()
	if len(args) != 0 {
		return shim.Error("init parameter error!")
	}


	return shim.Success(nil)
}

func (dbroot *DBroot)Invoke(stub shim.ChaincodeStubInterface) peer.Response  {

	functiondb,args := stub.GetFunctionAndParameters()

	switch functiondb {
	case "example" :
		return shim.Success([]byte(args[0]))

	case "createntity":
		return dbroot.CreateEntity(stub,args)
	case "createDB":
		return dbroot.CreateDB(stub,args)


	default:

	}



	return shim.Success(nil)
}

