package kingdb

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strings"
	"time"
)

//system entity:user db list
//create db:args[0]=dbname,args[1]=username,args[2]=userpass
//return instance:{"username":"wangchong","userpass":"123456"} json string
/*
args:
	0:username
	1:userpass
	2:DBname
	3:DBpass
 */
func (dbroot *DBroot) CreateDB(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var dbbuild DBBuild

	dbbuild.UserName = args[0]
	dbbuild.UserPass = args[1]
	dbbuild.DBName = args[2]
	dbbuild.DBPass = args[3]

	dbbuild.DBId = IDRandomCreate()

	if strings.ContainsAny(dbbuild.DBName,"-") ||strings.ContainsAny(dbbuild.DBPass,"-") || strings.ContainsAny(dbbuild.UserName,"-") || strings.ContainsAny(dbbuild.UserPass,"-"){

		return shim.Error("can't contain - charactor!")
	}

	//check user and db
	querybytes := fmt.Sprintf("{\"selector\":{\"king_user_name\":\"%s\",\"king_user_pass\":\"%s\",\"king_db_name\":\"%s\",\"king_db_pass\":\"%s\"}}",dbbuild.UserName,dbbuild.UserPass,dbbuild.DBName,dbbuild.DBPass)
	resultInterator, err := stub.GetQueryResult(querybytes)
	if err != nil {
		return shim.Error("db check error!")
	}

	defer resultInterator.Close()
	if resultInterator.HasNext() {
		fmt.Println("user and db have create,open!")
		queryResult, err := resultInterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		

		return shim.Success(queryResult.Value)
	}

	//put user and db
	jsonbytes, err := json.Marshal(&dbbuild)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(dbbuild.DBId,jsonbytes)
	if err != nil {
		fmt.Errorf("put state error!")
	}

	return shim.Success(jsonbytes)

}

//login db
/*
args:
	0:username
	1:userpass
	2:DBname
	3:DBpass
 */
func (dbroot *DBroot)loginDB(stub shim.ChaincodeStubInterface,args []string) peer.Response  {

	var dbbuild DBBuild

	dbbuild.UserName = args[0]
	dbbuild.UserPass = args[1]
	dbbuild.DBName = args[2]
	dbbuild.DBPass = args[3]

	//check user and db
	querybytes := fmt.Sprintf("{\"selector\":{\"king_user_name\":\"%s\",\"king_user_pass\":\"%s\",\"king_db_name\":\"%s\",\"king_db_pass\":\"%s\"}}",dbbuild.UserName,dbbuild.UserPass,dbbuild.DBName,dbbuild.DBPass)
	resultInterator, err := stub.GetQueryResult(querybytes)
	if err != nil {
		return shim.Error("db check error!")
	}

	defer resultInterator.Close()
	if resultInterator.HasNext() {
		fmt.Println("user and db have create,open!")
		queryResult, err := resultInterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}


		return shim.Success(queryResult.Value)
	}

	return shim.Error("db not found!")

}

//param:dbID
func (dbroot *DBroot)DeleteDB(stub shim.ChaincodeStubInterface, args []string) peer.Response  {

	dbID := args[0]

	err := stub.DelState(dbID)
	if err != nil {
		return shim.Error(err.Error())
	}


	return shim.Success([]byte("Delete my db success!"))

}
//parameter:dbid   tablename  tablepass(option) return tableID
func (dbroot *DBroot) CreateTable(stub shim.ChaincodeStubInterface,args []string) peer.Response {

	dbID := args[0]
	tablename := args[1]
	tablepass := args[2]
	checkID := fmt.Sprintf(`{"selector:{"king_db_id":"%s","king_table_name":"%s"}"}`,dbID,tablename)
	result, _:= stub.GetQueryResult(checkID)
	if result.HasNext() {
		shim.Error("table has create")
	}

	var entity DBEntity
	entity.TableName = tablename
	entity.DBid = dbID
	entity.TableID = IDRandomCreate()
	entity.TimeEntity = time.Now().UnixNano()
	entity.TablePass = tablepass

	entitybytes, _ :=json.Marshal(entity)
	stub.PutState(dbID,entitybytes)

	return  shim.Success([]byte(entity.TableID))

}
//return table id
func (dbroot *DBroot)OpenTable(stub shim.ChaincodeStubInterface,args []string) peer.Response {

	//dbid,tablename,
	dbID :=args[0]
	tablename := args[1]
	tablepass := args[2]

	var tableID string

	querybytes := fmt.Sprintf(`{"selector":{"king_db_id":"%s","king_table_name":"%s"}}`,dbID,tablename)
	result, _:= stub.GetQueryResult(querybytes)
	if !result.HasNext() {
		shim.Error("table hasn't create")
	} else {
		queryresult, _ := result.Next()

		var entity DBEntity
		err := json.Unmarshal(queryresult.Value,&entity)
		if err != nil {
			shim.Error(err.Error())
		}
		if entity.TablePass == tablepass{

			tableID = entity.TableID

		} else {

			return shim.Error("tablepass error!")
		}

	}

	return shim.Success([]byte(tableID))
}



//paramters:
/*
	arg[0] = tableid
	entityparams:(key-value):  key(typedefine)  value(paramname)
 */

func (dbroot *DBroot) UpdateEntity(stub shim.ChaincodeStubInterface, args []string) peer.Response  {



	return shim.Success(nil)
}


func (dbroot *DBroot) DeleteEntity(stub shim.ChaincodeStubInterface,args []string) peer.Response  {

	return shim.Success(nil)
}



func (dbroot *DBroot)ModifyUserPass(stub shim.ChaincodeStubInterface,args []string) peer.Response  {

	return shim.Success(nil)
}

func (dbroot *DBroot)ModifyDBPsss(stub shim.ChaincodeStubInterface,args []string) peer.Response  {

	return shim.Success(nil)
}
//query start
//------------------------------------------------------------------------------------------------------------


//------------------------------------------------------------------------------------------------------------
//query end


