package utils

import (
	"../log"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func DoGet(url string, data map[string]string, headers map[string]string) []byte {
	return doRequest("GET", url, data, headers)
}

func DoPost(url string, data map[string]string, headers map[string]string) []byte {
	return doRequest("POST", url, data, headers)
}

func doRequest(
	method string,
	url string,
	data map[string]string,
	headers map[string]string) []byte {
	client := http.Client{}
	bytesData, _ := json.Marshal(data)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bytesData))
	if err != nil {
		log.Log("http error:"+err.Error(), log.ERROR)
		return nil
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Log("read error:"+err.Error(), log.ERROR)
			return nil
		}
		return body
	}
	return nil
}
