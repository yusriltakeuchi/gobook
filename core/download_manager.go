package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTemplate(name string) string {
	url := fmt.Sprintf("https://raw.githubusercontent.com/yusriltakeuchi/gobook_template/master/template/%s.go.tmpl", name)
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	return string(bodyByte)
}
