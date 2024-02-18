package Untils

import (
	Data2 "CityUntils/Data"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
	"strings"
)

var Client http.Client

type NetWorkTips interface {
	GetNetRequest(url string, isToken bool, Data url2.Values) (data any)
	PostNetRequest(url string, Data any, isToken bool) (data any)
}

type Config struct {
	Ip    string
	Port  string
	Pix   string
	Token string
	User  string
	Pwd   string
}

func (c *Config) GetNetRequest(url string, isToken bool, Data url2.Values) (Res any) {
	urlIp := fmt.Sprintf("http://%s:%s%s%s", c.Ip, c.Port, c.Pix, url)
	Info.Println(urlIp)
	uri, err := url2.ParseRequestURI(urlIp)
	if err != nil {
		return nil
	}
	uri.RawQuery = Data.Encode()
	urlPath := uri.String()
	request, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		Error.Print(err.Error())
		return nil
	}
	request.Header.Add("Content-Type", "application/json")
	if isToken {
		request.Header.Add("Token", c.Token)
	}
	do, err := Client.Do(request)
	if err != nil {
		Error.Print(err.Error())
		return nil
	}
	all, err := io.ReadAll(do.Body)
	if err != nil {
		Error.Print(err.Error())
		return nil
	}
	Info.Println(string(all))
	return Data2.Response{
		Data:  string(all),
		State: 0,
	}
}

func (c *Config) PostNetRequest(url string, Data any, isToken bool) (Res any) {
	marshal, err := json.Marshal(Data)
	var NetResData Data2.NetResponse
	if err != nil {
		return
	}
	url = fmt.Sprintf("http://%s:%s%s%s", c.Ip, c.Port, c.Pix, url)
	Info.Println(url)
	request, err := http.NewRequest("POST", url, strings.NewReader(string(marshal)))
	if err != nil {
		Error.Print(err.Error())
		return nil
	}
	request.Header.Add("Content-Type", "application/json")
	if isToken {
		request.Header.Add("Token", c.Token)
	}
	do, err := Client.Do(request)
	if err != nil {
		Error.Print(err.Error())
		return nil
	}
	all, err := io.ReadAll(do.Body)
	if err != nil {
		Error.Print(err.Error())
		return nil
	}
	err2 := json.Unmarshal(all, &NetResData)
	if err2 != nil {
		Error.Println("解析失败：", err2.Error())
		return nil
	}
	if NetResData.Status == 200 && NetResData.Success {
		Info.Println("开始正常返回数据")
		return Data2.Response{
			Data:  NetResData,
			State: 0,
		}
	} else {
		Error.Println(NetResData)
		return nil
	}

}
