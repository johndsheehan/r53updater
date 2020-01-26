package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// IPFetch interface
type IPFetch interface {
	Fetch() (string, error)
}

// IPIfy implementation of IPFetch for ipify
type IPIfy struct {
	url string
}

// NewIPIfy return new instance
func NewIPIfy() *IPIfy {
	return &IPIfy{
		url: "https://api.ipify.org?format=text",
	}
}

// Fetch return ip or error
func (i *IPIfy) Fetch() (string, error) {
	rsp, err := http.Get(i.url)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer rsp.Body.Close()

	ip, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(ip), nil
}
