package models

import (
	"context"
	pbCaptcha "ginshop/proto/captcha"

	"go-micro.dev/v4/util/log"
)

//调用验证码微服务
func MakeCaptcha(height int, width int, length int) (string, string, error) {

	// Create client
	captchaClient := pbCaptcha.NewCaptchaService("captcha", CaptchaClient)
	// Call service
	res, err := captchaClient.MakeCaptcha(context.Background(), &pbCaptcha.MakeCaptchaRequest{
		Height: int32(height),
		Width:  int32(width),
		Length: int32(length),
	})
	if err != nil {
		log.Fatal(err)
	}

	return res.Id, res.B64S, err
}

//验证验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	// Create client
	captchaClient := pbCaptcha.NewCaptchaService("captcha", CaptchaClient)
	// Call service
	res, err := captchaClient.VerifyCaptcha(context.Background(), &pbCaptcha.VerifyCaptchaRequest{
		Id:          id,
		VerifyValue: VerifyValue,
	})
	if err != nil {
		log.Fatal(err)
	}
	return res.VerifyResult
}
