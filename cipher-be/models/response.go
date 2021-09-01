package models

type BaseResponse struct {
	Content string `json:"content"`
}

type AffineResponse struct {
	BaseResponse
	Key    int `json:"key"`
	Offset int `json:"offset"`
}

type PlayfairResponse struct {
	BaseResponse
	Key string `json:"key"`
}

type VigenereResponse struct {
	BaseResponse
	Key string `json:"key"`
}

type VigenereExtendedResponse struct {
	Content []byte `json:"content"`
	Key     string `json:"key"`
}
