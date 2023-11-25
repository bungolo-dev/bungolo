package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	QUERY_INFO string = "query/device-info"
	PRESS_HOME string = "keypress/home"
)

var client *http.Client = createClient()

func createClient() *http.Client {
	tr := &http.Transport{}

	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second}

	return client
}

func makeUrl(ip string, port int) string {
	return fmt.Sprintf("http://%s:%d/", ip, port)
}

func queryInfo(ip string) string {
	resp, err := client.Get(makeUrl(ip, 8060) + QUERY_INFO)

	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	b, readErr := io.ReadAll(resp.Body)

	if readErr != nil {
		return readErr.Error()
	}
	return string(b)
}

func pressHome(ip string) {
	_, err := client.Get(makeUrl(ip, 8060) + PRESS_HOME)

	if err != nil {

	}
}

func KillClient() {
	client.CloseIdleConnections()
	client = nil
}
