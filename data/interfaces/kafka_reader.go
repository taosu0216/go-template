package interfaces

import (
	"encoding/json"
	"go-template/data/cons"
)

func (dr *DataRepo) SendEmail(value []byte) {
	var message cons.SendEmailModel
	err := json.Unmarshal(value, &message)
	if err != nil {
		// 处理错误
		dr.ErrorfInDataKafka("SendEmail [Unmarshal] err is: %v", err)
		return
	}
	err = dr.SendVerifyCodeToEmailByKafka(message.VerifyCode, message.Email, message.Subject)
	if err != nil {
		dr.ErrorfInDataKafka("SendEmail [SendVerifyCodeToEmailByKafka] err is: %v", err)
	}
	return
}
