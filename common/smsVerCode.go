package common

import (
	"time"
)

var SMSVerCodes = map[string]SMSVerCode{}
var IPSendSMS = map[string]string{}

type SMSVerCode struct {
	Phone         string
	Code          string
	EffectiveTime time.Time
	FailureTime   time.Time
}
