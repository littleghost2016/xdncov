package main

import "encoding/json"

// HTTPResponse 解析返回的json格式的消息
type HTTPResponse struct {
	E int         `json:"e"`
	M string      `json:"m"`
	D EmptyStruct `json:"d"`
}

// EmptyStruct HTTPResponse.D
type EmptyStruct struct {
}

// unmarshalHTTPResponse 解析json消息
func UnmarshalHTTPResponse(response []byte) (newHTTPResponse HTTPResponse) {
	newHTTPResponse = HTTPResponse{}
	json.Unmarshal(response, &newHTTPResponse)

	return
}
