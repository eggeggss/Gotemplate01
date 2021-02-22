package models

type UserTable []struct {
	IDUserTable int    `json:"id_user_table"`
	EmpEname    string `json:"emp_ename"`
	Company     int    `json:"company"`
	CreateBy    string `json:"create_by"`
	DtCreate    string `json:"dt_create"`
	StatVoid    int    `json:"stat_void"`
}
