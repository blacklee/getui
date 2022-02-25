package getui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func loadAppConfig() GeTuiConfig {
	hash := _loadFileToMap("myapp.json")
	cfg := GeTuiConfig{}
	cfg.AppID = hash["appid"].(string)
	cfg.AppKey = hash["appkey"].(string)
	cfg.Master = hash["master"].(string)
	return cfg
}

func loadTestInfo() map[string]interface{} {
	return _loadFileToMap("test.json")
}

func _loadFileToMap(ff string) map[string]interface{} {
	content, err := ioutil.ReadFile(ff)
	if err != nil {
		fmt.Println("Err")
	}
	fmt.Println(string(content))
	var hash map[string]interface{}
	json.Unmarshal(content, &hash)
	return hash
}
