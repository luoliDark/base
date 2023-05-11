package captcha

// 生产验证码类
import (
	"image/color"

	ssomodel "github.com/luoliDark/base/sso/ssologin/model"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

// 生成验证码，验证码只能校验一次，不管成功还是失败
// 验证码还需要考虑分布式问题，在A服务器生成的验证码可能在B服务器不存在，这样会造成困难，需要把验证码和captchaid 对应写入到redis
func CreateCaptchaCode(typeStr string) (*ssomodel.CaptchaResponse, error) {
	var driver base64Captcha.Driver
	//create base64 encoding captcha
	switch typeStr {
	case "audio":
		driverAutdio := new(base64Captcha.DriverAudio)
		driverAutdio.Length = 6
		driverAutdio.Language = "zh"
		driver = driverAutdio
	case "string":
		driverString := new(base64Captcha.DriverString)
		driverString.Height = 60
		driverString.Width = 240
		driverString.ShowLineOptions = 0
		driverString.NoiseCount = 0
		driverString.Source = "1234567890qwertyuioplkjhgfdsazxcvbnm"
		driverString.Length = 6
		driverString.Fonts = []string{"wqy-microhei.ttc"}
		rgba := new(color.RGBA)
		rgba.R = 255
		rgba.G = 255
		rgba.B = 255
		rgba.A = 255
		driverString.BgColor = rgba
		driverString.ConvertFonts()
		driver = driverString
	case "math":
		driverMath := new(base64Captcha.DriverMath)
		driverMath.Height = 60
		driverMath.Width = 240
		driverMath.ShowLineOptions = 0
		driverMath.NoiseCount = 0
		//driverMath.Length= 6
		driverMath.Fonts = []string{"wqy-microhei.ttc"}
		rgba := new(color.RGBA)
		rgba.R = 255
		rgba.G = 255
		rgba.B = 255
		rgba.A = 255
		driverMath.BgColor = rgba
		driverMath.ConvertFonts()
		driver = driverMath
	case "chinese":
		driverChinese := new(base64Captcha.DriverChinese)
		driverChinese.Height = 60
		driverChinese.Width = 320
		driverChinese.ShowLineOptions = 0
		driverChinese.NoiseCount = 0
		driverChinese.Source = "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,,不想要,的值"
		driverChinese.Length = 2
		driverChinese.Fonts = []string{"wqy-microhei.ttc"}
		rgba := new(color.RGBA)
		rgba.R = 255
		rgba.G = 255
		rgba.B = 255
		rgba.A = 255
		driverChinese.BgColor = rgba
		driverChinese.ConvertFonts()
		driver = driverChinese
	default:
		driverDigit := new(base64Captcha.DriverDigit)
		driverDigit.Height = 80
		driverDigit.Width = 240
		driverDigit.Length = 4
		driverDigit.MaxSkew = 0.7
		driverDigit.DotCount = 80
		driver = driverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	if err != nil {
		return nil, err
	}
	captchaResponse := new(ssomodel.CaptchaResponse)
	captchaResponse.CaptchaId = id
	captchaResponse.Image = b64s
	return captchaResponse, nil
}

func Verify(captchaId string, code string) bool {
	return store.Verify(captchaId, code, true)
}
