package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"kingdb"
)

type Test struct {

	Username string `json:"username"`
	Userpass string `json:"userpass"`
}
	func main() {
		err := shim.Start(new(kingdb.DBroot))
		if err != nil {
			fmt.Errorf("contract chaincode start error = %s",err)
		}

		////test
		//var t Test = Test{
		//		"wangchong",
		//		"123456",
		//}
		//
		//jsonbytes, _ := json.Marshal(&t)
		//fmt.Println(string(jsonbytes))

	}
