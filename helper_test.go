package getui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var testConfig GeTuiConfig
var testAuthToken string
var testCids []string

func setup() {
	GTV2_DEBUG = true
	hash := _loadFileToMap("test.json")
	appcfg := hash["config"].(map[string]interface{})
	testConfig.AppID = appcfg["appid"].(string)
	testConfig.AppKey = appcfg["appkey"].(string)
	testConfig.Master = appcfg["master"].(string)

	pushcfg := hash["push"].(map[string]interface{})
	testAuthToken = pushcfg["token"].(string)
	cids := pushcfg["cids"].([]interface{})
	for _, cid := range cids {
		testCids = append(testCids, cid.(string))
	}
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
