package model

import "time"

type Done struct {
	Did      int       `grom:"column:did"`
	Uid      int       `grom:"column:uid"`
	Thing    string    `grom:"column:thing"`
	Remark   string    `grom:"column:remark"`
	DoneTime time.Time `grom:"type:date;column:done_time"`
}
