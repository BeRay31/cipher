package models

import (
	"encoding/base64"
)

type InputRequest struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type BaseRequest struct {
	Input InputRequest `json:"input"`
}

type AffineRequest struct {
	BaseRequest
	Key    int `json:"key"`
	Offset int `json:"offset"`
}

type PlayfairRequest struct {
	BaseRequest
	Key string `json:"key"`
}

type VigenereRequest struct {
	BaseRequest
	Key string `json:"key"`
}

type VigenereExtendedRequest struct {
	Content string `json:"content"`
	Key     string `json:"key"`
}

func (req *InputRequest) ParseContent() (string, error) {
	// Parse file
	if req.Type == "FILE" {
		decoded, err := base64.StdEncoding.DecodeString(req.Content)
		if err != nil {
			return "", err
		}
		return string(decoded), nil
	}
	return req.Content, nil
}
