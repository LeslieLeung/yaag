package yaag

import "os/exec"

func RunSwag() {
	general := YaagConfig.SwagGeneral
	if general == "" {
		general = "./main.go"
	}
	cmd := exec.Command("swag", "init", "-g", general, "-o", "./docs", "-ot", "json")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
