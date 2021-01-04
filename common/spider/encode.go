package spider

import (
	"compress/gzip"
	"gopkg.in/resty.v1"
	"net/http"
)

func UnZipRetryResp(response *resty.Response, size int) ([]byte, error) {
	rawBody := response.RawBody()
	defer rawBody.Close()
	switch response.Header().Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(rawBody)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		buff, err := ReadAll(reader, size)
		if err != nil {
			return nil, err
		}
		return buff, nil
	default:
		buff, err := ReadAll(rawBody, size)
		if err != nil {
			return nil, err
		}
		return buff, nil
	}

}

func UnZipHttpResp(response *http.Response, size int) ([]byte, error) {
	rawBody := response.Body
	defer rawBody.Close()
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(rawBody)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		buff, err := ReadAll(reader, size)
		if err != nil {
			return nil, err
		}
		return buff, nil
	default:
		buff, err := ReadAll(rawBody, size)
		if err != nil {
			return nil, err
		}
		return buff, nil
	}

}
