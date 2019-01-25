package kingdb

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func (dbroot *DBroot) AdminCABusiness(stub shim.ChaincodeStubInterface,args []string) peer.Response  {

	return shim.Success(nil)

}

