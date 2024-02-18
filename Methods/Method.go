package Methods

import (
	"CityUntils/Untils"
	url2 "net/url"
)

// GetCaptcha 根据规则生成验证码
func GetCaptcha(Net Untils.Config) {
	url := "/user/captcha"
	value := url2.Values{}
	value.Add("uuid", Untils.MD5IsOK())
	Net.GetNetRequest(url, false, value)
}
