package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
