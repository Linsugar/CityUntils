package main

import (
	"CityUntils/Data"
	"CityUntils/Methods"
	"CityUntils/Untils"
)

var Net = Untils.Config{}

func main() {
	login := Data.Login{}
	login.AccountNumber = Net.User
	login.Password = Net.Pwd
	Net.PostNetRequest("/user/login", login, false)
	Methods.GetCaptcha(Net)
}

func init() {
	Net.Ip = "10.10.8.222"
	Net.Port = "80"
	Net.User = "root"
	Net.Pwd = "ee5db33d265e8cecc0546c69c3bc92f9"
	Net.Pix = "/optimus/api/v1"
}
