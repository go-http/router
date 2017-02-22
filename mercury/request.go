package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

//获取流量统计数据(系统工具-流量统计)
func (cli *Client) request(path string, query url.Values) ([]byte, error) {
	uri := cli.baseUri + path
	if query != nil {
		uri += "?" + query.Encode()
	}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, fmt.Errorf("请求创建失败: %s", err)
	}

	req.Header.Add("Referer", uri)
	req.SetBasicAuth(cli.Username, cli.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求执行失败: %s", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("响应读取失败: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("错误响应[%d]: %s", resp.StatusCode, data)
	}

	return data, nil
}

//从返回的HTML文档中检查是否包含错误码
func getError(data []byte) error {
	r := regexp.MustCompile(`var errCode = "([0-9]+)"`)
	slice := r.FindSubmatch(data)

	if len(slice) != 2 {
		return nil
	}

	errCode, _ := strconv.Atoi(string(slice[1]))
	if errCode == 0 {
		return nil
	}

	msg, found := errorStringMap[errCode]
	if found {
		return fmt.Errorf("[%d]%s", errCode, msg)
	}

	return fmt.Errorf("未知错误码%d", errCode)
}

//从返回的HTML文档中获取指定名称的Javascript数组
func getJsArray(data []byte, varName string) []string {
	r := regexp.MustCompile(`(?s)var ` + varName + `=new Array\(([^;]*) \);`)
	match := r.FindSubmatch(data)
	if len(match) < 2 {
		return nil
	}

	str := string(match[1])
	str = strings.Replace(str, ",", "", -1)
	str = strings.Replace(str, `"`, "", -1)

	return strings.Split(str[1:], "\n")
}
