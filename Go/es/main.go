package main

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net/http"
)

const (
	_account  = "elastic"
	_password = "*LcQx0*_AB3UrVGlJcl5"
	_url      = "https://192.168.31.22:9200"
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
		log.Fatalln(err)
	}
	log.Println(elasticsearch.Version)
	log.Println(es.Info())
}
