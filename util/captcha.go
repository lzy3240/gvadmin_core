package util

import (
	"github.com/mojocn/base64Captcha"
	"gvadmin_v3/core/config"
	"image/color"
	"time"
)

// 设置自带的store
var store = base64Captcha.NewMemoryStore(20240, 3*time.Minute)

// CaptMake 生成验证码
func CaptMake() (string, string, error) {
	var driver base64Captcha.Driver
	switch config.Instance().App.CaptchaMode {
	case "arithmetic":
		driver = mathCaptcha()
	case "letter":
		driver = stringCaptcha()
	default:
		driver = stringCaptcha()
	}

	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := captcha.Generate()
	return id, b64s, err
}

// 数字运算验证码
func mathCaptcha() base64Captcha.Driver {
	return base64Captcha.NewDriverMath(
		50,
		100,
		0,
		0,
		&color.RGBA{0, 0, 0, 0},
		nil,
		[]string{"RitaSmith.ttf"},
	)
}

// 字符验证码
func stringCaptcha() base64Captcha.Driver {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	captchaConfig := base64Captcha.DriverString{
		Height:          50,
		Width:           100,
		NoiseCount:      0,     // 干扰字母
		ShowLineOptions: 1 | 3, // 干扰线
		Length:          4,
		Source:          "qwertyuioplkjhgfdsazxcvbnm23456789",
	}
	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	return driver
}

// CaptVerify 验证captcha是否正确
func CaptVerify(id string, capt string) bool {
	return store.Verify(id, capt, true)
}
