package main

//http 超时设置问题
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var tr *http.Transport

func init() {
	tr = &http.Transport{
		MaxIdleConns: 100,
		//设置的是连接超时时间
		//Dial: func(netw, addr string) (net.Conn, error) {
		//	conn, err := net.DialTimeout(netw, addr, time.Second*2) //建立连接
		//	if err != nil {
		//		return nil, err
		//	}
		//	err = conn.SetDeadline(time.Now().Add(time.Second * 3)) //超时
		//	if err != nil {
		//		return nil, err
		//	}
		//	return conn, err
		//},
	}
}

func Get(url string) ([]byte, error) {
	m := make(map[string]interface{})
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(data)

	req, _ := http.NewRequest("Get", url, body)
	req.Header.Add("content-type", "application/json")
	client := &http.Client{
		Transport: tr,
		//设置的是请求的超时时间
		Timeout:   time.Second * 2,
	}
	res, err := client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}

func main() {
	for {
		_, err := Get("https://www.baidu.com")
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
