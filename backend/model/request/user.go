package request

type User struct {
	Uname  string `json:"uname"`
	Passwd string `json:"passwd"`
}

type NewUser struct {
	User      User   `json:"user"`
	NewPasswd string `json:"new_passwd"`
}
