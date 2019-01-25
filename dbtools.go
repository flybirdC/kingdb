package kingdb

import (
	"fmt"
	"github.com/satori/go.uuid"
)

//create random uuid as data(k,v) id,_id = id
func IDRandomCreate() string  {

	kingUUID,err := uuid.NewV4()
	if err != nil {
		fmt.Printf("create uuid failed=%s\n",err)
	}

	return  kingUUID.String()
}


