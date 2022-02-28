package getui

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAuthToken(t *testing.T) {
	setup()
	data, err := GetAccessToken(testConfig)
	if err != nil {
		panic(err)
	}
	bb, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bb))
	fmt.Println("token is: ", data.Token)
}
