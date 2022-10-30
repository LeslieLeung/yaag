package yaag

import (
	"encoding/json"
	"io"
	"net/http"
	url2 "net/url"
	"os"
)

const (
	YapiImportData = "/api/open/import_data"
)

type ImportDataRequest struct {
	Type  string `form:"type"`
	Json  string `form:"json"`
	Merge string `form:"merge"`
	Token string `form:"token"`
}

type ImportDataResponse struct {
	Errcode int      `json:"errcode"`
	Errmsg  string   `json:"errmsg"`
	Data    struct{} `json:"data"`
}

// readSwaggerJson reads swagger.json from docDir
func readSwaggerJson(docDir string) string {
	jsonStr, err := os.ReadFile(docDir)
	if err != nil {
		panic(err)
	}
	return string(jsonStr)
}

func UpdateYapi(docDir, mergeMode string) string {
	url := YaagConfig.YapiUrl + YapiImportData
	jsonStr := readSwaggerJson(docDir)
	req := url2.Values{
		"type":  {"swagger"},
		"json":  {jsonStr},
		"merge": {mergeMode},
		"token": {YaagConfig.YapiToken},
	}
	client := &http.Client{}
	resp, err := client.PostForm(url, req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var ret ImportDataResponse
	err = json.Unmarshal(body, &ret)
	if err != nil {
		panic(err)
	}
	return ret.Errmsg
}