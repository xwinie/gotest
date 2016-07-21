package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/base64"
)

//----------------------------------
// 移动联通基站调用示例代码 － 聚合数据
// 在线接口文档：http://www.juhe.cn/docs/8
//----------------------------------

const APPKEY = "*******************" //您申请的APPKEY

func main() {

	//1.基站定位
	Request1()

}

//1.基站定位
func Request1() {
	//请求地址
	juheURL := "http://127.0.0.1:8080/v1/user222/user_11111"

	//初始化参数
	param := url.Values{}

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("appid", "app1")
	//param.Set("uid", "user_11111")
	//param.Set("timestamp", "2016-07-06 11:04:05")
	//param.Set("signature", "hhaUY3y3YAy8KjW/A3PaFRLAAozehDdK6BNfG70eVn8=")

	param.Set("data", base64.StdEncoding.EncodeToString([]byte( `{"Name":"junbin", "Age":21, "Gender":true}`)))
	//发送请求
	data, err := Get(juheURL, param)
	if err != nil {
		fmt.Errorf("请求失败,错误信息:\r\n%v", err)
	} else {
		//var netReturn map[string]interface{}
		//json.Unmarshal(data,&netReturn)
		//fmt.Printf("接口返回result字段是:\r\n%v",netReturn)

		fmt.Printf("33:"+ string(data))

	}
}

// get 网络请求
func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
