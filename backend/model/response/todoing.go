package response

import "time"

type Todoing struct {
	Is_Done  int       `grom:"column:is_done"`
	Deadline time.Time `grom:"type:date;column:deadline"`
	Thing    string    `grom:"column:thing"`
	Remark   string    `grom:"column:remark"`
}
