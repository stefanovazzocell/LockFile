.PHONY: test
test:
	@echo "> Formatting..."
	go vet .
	go fmt .
	@echo -e "\n############################"
	@echo "> Running Go tests..."
	go test .
	@echo -e "\n############################"
	@echo "> Running security checks..."
	govulncheck .
	gosec -quiet .

.PHONY: setup
setup:
	@echo "> Installing mod requirements..."
	go mod download
	@echo -e "\n############################"
	@echo "> Installing security and vulnerability scanners..."
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest