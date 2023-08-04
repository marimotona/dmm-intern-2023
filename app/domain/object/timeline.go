package object

type TimeLine struct {
	Elements []TimeLineElement `json:"elements"`
}

type TimeLineElement struct {
	Account *Account `json:"account"`
	Status  *Status  `json:"status"`
}
