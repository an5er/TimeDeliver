package model

type Users struct {
	Uid    int    `grom:"column:uid"`
	Uname  string `grom:"column:uname"`
	Passwd string `grom:"column:passwd"`
}
