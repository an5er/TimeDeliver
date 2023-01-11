package request

type Todoing struct {
	Is_Done  int    `json:"done"`
	Deadline string `json:"deadline"`
	Thing    string `json:"thing"`
	Remark   string `json:"remark"`
}

type TodoingMonth struct {
	DeadlineMonth string `json:"month"`
	DeadlineYear  string `json:"year"`
}
