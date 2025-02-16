package cons

func init() {
	KafkaTopics = append(KafkaTopics, EmailRegisterNotificationTopic, EmailPasswordResetNotificationTopic)
}

const (
	EmailRegisterNotificationTopic      = "email-service.register.notification"
	EmailPasswordResetNotificationTopic = "email-service.password-reset.notification"
	KafkaTopicTest                      = "test"
)

var (
	KafkaTopics []string
)

type SendEmailModel struct {
	Email      string `json:"email"`
	VerifyCode string `json:"verify_code"`
	Subject    string `json:"subject"`
}
