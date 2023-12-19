package yaag

import (
	"os"
	"os/exec"
)

func RunSwag(parseDependency bool) {
	general := "./main.go"
	if YaagConfig != nil && YaagConfig.SwagGeneral != "" {
		general = YaagConfig.SwagGeneral
	}
	args := []string{"init", "-g", general, "-o", "./docs", "--ot", "json"}
	if parseDependency {
		args = append(args, "--pd")
	}
	cmd := exec.Command("swag", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
