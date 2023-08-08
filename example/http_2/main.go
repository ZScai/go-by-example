/*
http请求添加header
*/
package main

import (
        "fmt"
        "io"
        "net/http"
        "net/url"
)

func main() {
        params := url.Values{}
        params.Set("ip", "121.35.46.110") // 添加查询参数
        Url, _ := url.ParseRequestURI("http://cip.cc/")
        Url.RawQuery = params.Encode()
        resp, err := http.NewRequest("GET", Url.String(), nil)
        if err != nil {
                fmt.Println("newRequest err: ", err)
                return
        }
        resp.Header.Add("User-Agent", "curl/7.58.0") // 添加请求头
        res, err := http.DefaultClient.Do(resp)
        if err != nil {
                fmt.Println("request err: ", err)
                return
        }
        defer res.Body.Close()
        body, _ := io.ReadAll(res.Body)
        fmt.Println(string(body))
}
