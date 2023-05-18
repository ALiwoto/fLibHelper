package fLibHelper

import (
	"crypto/tls"
	"net/http"
	"strings"
)

func NewFLibClient(pass string, config *FLibOptions) *FLibClient {
	if config == nil {
		config = DefaultOptions
	}

	client := &FLibClient{
		PassString:  pass,
		HostString:  config.HostString,
		UserAgent:   config.UserAgent,
		BaseUrl:     strings.Trim(config.BaseUrl, "/"),
		ReqDataName: config.ReqDataName,
		HttpClient:  config.CustomHttpClient,
	}

	if client.HttpClient == nil {
		client.HttpClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	}

	return client
}
