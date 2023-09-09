## up: start the script for checking partners node version for mainnet 
mainnet:
	@echo "Starting Script checking Mainnet..."
	MODE=mainnet JSON_FILE_PATH=./mainnetpartners.json go run main.go utility.go http.go
	@echo "Checking Partners Node version started!"


testnet:
	@echo "Starting Script checking Testnet..."
	MODE=testnet JSON_FILE_PATH=./testnetpartners.json go run main.go utility.go http.go
	@echo "Checking Partners Node version started!"


.PHONY: mainnet testnet
