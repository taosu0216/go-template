package types

type SendRegisterEmailReq struct {
	Email string `json:"email"`
}

type SendRegisterEmailResp struct {
	BasicResp BasicResp `json:"basic_resp"`
}

type SendChangePasswdEmailReq struct {
	Email string `json:"email"`
}

type SendChangePasswdEmailResp struct {
	BasicResp BasicResp `json:"basic_resp"`
}

type UserLoginReq struct {
	UserName string `json:"username"`
	PassWd   string `json:"passwd"`
}

type UserLoginResp struct {
	BasicResp BasicResp `json:"basic_resp"`
	Token     string    `json:"token"`
}

type RegisterUserReq struct {
	UserName   string `json:"username"`
	PassWd     string `json:"passwd"`
	Email      string `json:"email"`
	VerifyCode string `json:"verifyCode"`
}

type RegisterUserResp struct {
	BasicResp BasicResp `json:"basic_resp"`
}

type ChangePasswordReq struct {
	PassWd     string `json:"passwd"`
	Email      string `json:"email"`
	VerifyCode string `json:"verifyCode"`
}

type ChangePasswordResp struct {
	BasicResp BasicResp `json:"basic_resp"`
}
