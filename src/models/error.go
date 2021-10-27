package models

type ErrorMessage struct {
	Err string
}

func SendErr(str string) ErrorMessage {
	return ErrorMessage{Err: str}
}