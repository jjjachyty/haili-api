package actions

import (
	"haili/services"

	"haili/common"

	"github.com/lunny/tango"
)

type SMSVerCodeAction struct {
	tango.Ctx
	tango.Json
}

var smssvs services.SmsVerCodeService

func (act SMSVerCodeAction) Post() interface{} {

	phone := act.Ctx.Form("phone")

	if "" != phone {
		e := smssvs.SendSMSVerCode(phone)
		if nil != e {
			return common.ReturnMesg(false, error.Error(), e)
		}
	}
	return common.ReturnMesg(true, "短信发送成功", nil)
}
func (act SMSVerCodeAction) Get() interface{} {
	smscode := act.Param(":smscode")
	phone := act.Param(":phone")
	flag := smssvs.CheckSMSVerCode(phone, smscode)
	if flag {
		return common.ReturnMesg(true, "短信验证成功", flag)
	}
	return common.ReturnMesg(false, "短信验证失败", flag)
}
