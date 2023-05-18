package fLibHelper

import "net/http"

type FLibClient struct {
	PassString  string
	HostString  string
	UserAgent   string
	BaseUrl     string
	ReqDataName string
	HttpClient  *http.Client
}

type FLibOptions struct {
	HostString       string
	UserAgent        string
	BaseUrl          string
	ReqDataName      string
	CustomHttpClient *http.Client
}

type ReqOptions struct {
	ReqData      string
	ReqTotalData string
}
