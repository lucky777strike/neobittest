package shodan

type ReqError struct {
	Status int
	Msg    string
}

func (e *ReqError) Error() string {
	return e.Msg
}
