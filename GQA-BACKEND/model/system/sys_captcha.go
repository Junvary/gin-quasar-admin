package system

type ResponseCaptcha struct {
	CaptchaId string `json:"captchaId"`
	PicPath   string `json:"picPath"`
}
