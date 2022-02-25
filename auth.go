package getui

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type AuthResultData struct {
	ExpireTime string `json:"expire_time"`
	Token      string `json:"token"`
}

type AuthResult struct {
	Msg  string         `json:"msg"`
	Code int64          `json:"code"`
	Data AuthResultData `json:"data"`
}

func GetAccessToken(cfg GeTuiConfig) (AuthResultData, error) {
	var data AuthResultData
	sign, timestamp := Signature(cfg.AppKey, cfg.Master)
	params := &AuthParam{
		Sign:      sign,
		Timestamp: timestamp,
		AppKey:    cfg.AppKey,
	}

	bodyByte, err := json.Marshal(params)
	if err != nil {
		return data, err
	}

	authUrl := fmt.Sprintf("%s/%s/%s", baseURL, cfg.AppID, "auth")
	result, err := sendPost(authUrl, "", bodyByte)
	if err != nil {
		return data, err
	}

	// fmt.Println("get token result:\n", result)

	tokenResult := new(AuthResult)
	if err := json.Unmarshal([]byte(result), &tokenResult); err != nil {
		return data, err
	}

	if tokenResult.Code != 0 {
		fmt.Println("getui auth failed: ", result)
		return data, fmt.Errorf(tokenResult.Msg)
	}

	return tokenResult.Data, nil
}

func Signature(appKey string, masterSecret string) (string, string) {
	timestamp := strconv.FormatInt((time.Now().UnixNano() / 1000000), 10)
	original := appKey + timestamp + masterSecret

	hash := sha256.New()
	hash.Write([]byte(original))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum), timestamp
}

type AuthParam struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	AppKey    string `json:"appkey"`
}
