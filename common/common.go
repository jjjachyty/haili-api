package common

type Data struct {
	Status bool
	Mesg   string
	// ErrorMesg string
	Data interface{}
}

func ReturnMesg(status bool, mesg string, data interface{}) Data {

	return Data{status, mesg, data}
}
