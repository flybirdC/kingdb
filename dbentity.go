package kingdb


type DBAdmin struct {

	UserId string `json:"king_admin_id"`
	UserName string `json:"king_admin_name"`
	UserPass string `json:"king_admin_pass"`
	UserTel string `json:"king_admin_tel,omitempty"`
}


type DBUser struct {

	UserId string `json:"king_user_id"`
	UserName string `json:"king_user_name"`
	UserPass string `json:"king_user_pass"`
	UserTel string `json:"king_user_tel,omitempty"`

}

type DBBuild struct {

	DBId string `json:"king_db_id"`
	DBName string `json:"king_db_name"`
	DBPass string `json:"king_db_pass"`
	UserName string `json:"king_user_name"`
	UserPass string `json:"king_user_pass"`

}


type DBEntity struct {


	TableID string `json:"king_table_id"`
	DBid string `json:"king_db_id"`
	TableName string `json:"king_table_name"`
	TimeEntity int64 `json:"king_time_entity"`
	TablePass string `json:"king_table_pass,omitempty"`

}
 

type DBEntityData struct {

	DataID string `json:"king_data_id"`
	DBentity DBEntity `json:"king_d_bentity"`
	DataJson string `json:"king_data_json"`
}


type CABusiness struct {

	AdminAll bool `json:"king_admin_all,omitempty"`

}






