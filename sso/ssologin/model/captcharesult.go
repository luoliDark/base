package ssomodel

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"`

	Image string `json:"image"`
}
