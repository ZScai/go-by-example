
/*解析http返回的json字符串*/
package main

import (
        "encoding/json"
        "fmt"
        "io"
        "net/http"
        "net/url"
)

type result struct {
        Args    map[string]string `json:"args"`
        Headers map[string]string `json:"headers"`
        Origin  string            `json:"origin"`
        Url     string            `json:"url"`
}

func res_param() {
        params := url.Values{}
        Url, _ := url.Parse("http://httpbin.org/get")
        params.Set("name", "小明")
        params.Set("age", "18")
        Url.RawQuery = params.Encode()
        resp, err := http.Get(Url.String())
        if err != nil {
                fmt.Println(err)
                return
        }
        defer resp.Body.Close()
        body, _ := io.ReadAll(resp.Body)
        var res result
        if err := json.Unmarshal(body, &res); err != nil {
                fmt.Println("json analysis error: ", err)
        }
        fmt.Printf("%v\n", res)

}

func main() {
        res_param()
}
