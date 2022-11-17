VERSION=0_2
BINARY_NAME=spc

build: ## Build your project and put the output binary in out/bin/
	mkdir -p bin
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_NAME)_darwin_amd64_$(VERSION) main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/$(BINARY_NAME)_darwin_arm64_$(VERSION) main.go
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)_linux_amd64_$(VERSION) main.go
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_NAME)_windows_amd64_$(VERSION).exe main.go

clean: ## Remove build related file
	rm -fr ./bin
	rm -f ./junit-report.xml checkstyle-report.xml ./coverage.xml ./profile.cov yamllint-checkstyle.xml