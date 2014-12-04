package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func get(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Authorization", user+":"+apiToken)

	client := &http.Client{Transport: transportConfig}
	resp, errs := client.Do(req)

	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}
	body, errs := ioutil.ReadAll(resp.Body)
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}
	return body
}

func post(url string, json []byte) []byte {

	contentReader := bytes.NewReader(json)

	req, _ := http.NewRequest("POST", url, contentReader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Authorization", user+":"+apiToken)

	client := &http.Client{Transport: transportConfig}

	resp, errs := client.Do(req)

	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}

	body, errs := ioutil.ReadAll(resp.Body)
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}
	return body
}

func delete(url string) []byte {

	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Authorization", user+":"+apiToken)

	fmt.Println(req)

	client := &http.Client{Transport: transportConfig}
	resp, errs := client.Do(req)

	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}
	body, errs := ioutil.ReadAll(resp.Body)
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}
	return body
}
