package yaag

import (
	"os"
	"os/exec"
)

func RunSwag() {
	general := "./main.go"
	if YaagConfig != nil && YaagConfig.SwagGeneral != "" {
		general = YaagConfig.SwagGeneral
	}
	cmd := exec.Command("swag", "init", "-g", general, "-o", "./docs", "-ot", "json")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
