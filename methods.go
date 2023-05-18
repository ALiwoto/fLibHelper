package fLibHelper

import (
	"io"
	"net/http"
	"strings"

	"github.com/AnimeKaizoku/ssg/ssg"
)

func (c *FLibClient) CollectGold(opt *ReqOptions) (string, error) {
	var data *strings.Reader
	wholeData := ""
	if opt == nil {
		wholeData = c.ReqDataName + "=" + _defaultCollectGData
	} else {
		if opt.ReqData != "" {
			wholeData = c.ReqDataName + "=" + opt.ReqData
		} else if opt.ReqTotalData != "" {
			wholeData = opt.ReqTotalData
		}
	}

	if wholeData != "" {
		data = strings.NewReader(wholeData)
	}
	req, err := http.NewRequest(http.MethodPost, c.BaseUrl+"/cards/collectgold", data)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Content-Length", ssg.ToBase10(len(wholeData)))
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Host", c.HostString)
	req.Header.Set("Connection", "close")
	// req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Cookie", _defaultPassDataName+"="+c.PassString)
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyText), nil
}
