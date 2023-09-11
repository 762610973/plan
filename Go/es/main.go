package main

import (
	"crypto/tls"
	"log/slog"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

const (
	_account  = "elastic"
	_password = "123456"
	_url      = "http://192.168.31.22:9200"
)

func main() {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{_url},
		Username:  _account,
		Password:  _password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	})
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info(elasticsearch.Version)
	info, err := es.Info()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info(info.String())
}
