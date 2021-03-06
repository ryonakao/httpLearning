package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func proxy() {
	proxyURL, err := url.Parse("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	resp, err := client.Get("http://github.com")
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
