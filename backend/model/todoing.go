package model

import "time"

type Todoing struct {
	Tid      int       `grom:"column:tid;primaryKey;autoIncrement"`
	Uid      int       `grom:"column:uid"`
	Is_Done  int       `grom:"column:is_done"`
	Deadline time.Time `grom:"type:date;column:deadline"`
	Thing    string    `grom:"column:thing"`
	Remark   string    `grom:"column:remark"`
}
