gox-all:
	gox -osarch="darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64" -output="build/yaag_{{.OS}}_{{.Arch}}"

clean:
	rm -rf build