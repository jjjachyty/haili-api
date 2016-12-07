package services

import (
	"errors"
	"fmt"
	"haili/common"
	"regexp"
	"time"
)

type SmsVerCodeService struct{}

func (SmsVerCodeService) Get(phone string) (string, error) {
	return "HAILI", nil
}

func (svs SmsVerCodeService) SendSMSVerCode(phone string) error {
	var smscode common.SMSVerCode
	e := svs.beforSend(phone)
	if nil == e {
		code, err := svs.Get(phone)
		if nil != err {
			return err
		}
		smscode.Phone = phone
		smscode.Code = code
		smscode.EffectiveTime = time.Now()
		smscode.FailureTime = smscode.EffectiveTime.Add(1 * time.Minute)
		common.SMSVerCodes[phone] = smscode
		return nil
	}
	return e
}

func (svs SmsVerCodeService) CheckSMSVerCode(phone string, code string) bool {
	codeStr := common.SMSVerCodes[phone]
	fmt.Println(codeStr.Code == code)
	fmt.Println(time.Now().Before(codeStr.FailureTime))
	if codeStr.Code == code && time.Now().Before(codeStr.FailureTime) {
		return true
	}
	return false
}

func (svs SmsVerCodeService) beforSend(phone string) error {

	reg := regexp.MustCompile(`^1(3|4|5|7|8)\d{9}$`)
	isPhone := reg.MatchString(phone)
	fmt.Println("phone", phone, "isPhone", isPhone)
	if !isPhone {
		return errors.New("请输入正确的手机号码")
	}
	codeStruc := common.SMSVerCodes[phone]
	if "" != codeStruc.Code {
		if time.Now().Before(codeStruc.FailureTime) {
			return errors.New("1分钟之内只能发送一次短信")
		}
	}
	return nil
}
