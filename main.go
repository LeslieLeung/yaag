package main

import (
	"flag"
	"fmt"
	"github.com/leslieleung/yaag/internal/swagger"
	"github.com/leslieleung/yaag/internal/yaag"
)

var (
	docDir    string
	mergeMode string
	isInit    bool
	upload    bool
	timeout   int
)

func init() {
	flag.StringVar(&docDir, "docDir", "", "path to swagger.json")
	flag.StringVar(&mergeMode, "mergeMode", "", "merge mode")
	flag.BoolVar(&isInit, "init", false, "init")
	flag.BoolVar(&upload, "upload", false, "upload only, no swag generate")
	flag.IntVar(&timeout, "timeout", 10, "timeout in seconds")
	flag.Parse()
}

func main() {
	if isInit {
		yaag.InitConfig()
		return
	}
	yaag.GetConfig()
	mergeOptions()
	if !upload {
		yaag.RunSwag()
		swagger.Convert(docDir)
	}
	resp := yaag.UpdateYapi(docDir, mergeMode, timeout)
	fmt.Println(resp)
}

func mergeOptions() {
	if docDir == "" && yaag.YaagConfig != nil {
		docDir = yaag.YaagConfig.DocDir
	}
	if docDir == "" {
		docDir = "./docs/swagger.json"
	}
	if mergeMode == "" && yaag.YaagConfig != nil {
		mergeMode = yaag.YaagConfig.MergeMode
	}
	if mergeMode == "" {
		mergeMode = "merge"
	}
}
