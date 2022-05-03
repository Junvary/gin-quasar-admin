package model

type ResponseCaptcha struct {
	CaptchaImage string `json:"captcha_image"`
	CaptchaId    string `json:"captcha_id"`
}
