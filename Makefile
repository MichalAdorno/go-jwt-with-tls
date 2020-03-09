SHELL := /bin/bash 
BUILD_OUTPUT=main
BUILD_FOLDER=build
APP_PORT?=8080

gencert:
	@echo "make gencert 	- Create a self-signed X.509 certificate (dev only)"
	rm -f server.crt
	rm -f server.key
	rm -f $(BUILD_FOLDER)/server.crt
	rm -f $(BUILD_FOLDER)/server.key
	# Key considerations for algorithm "RSA" ≥ 2048-bit
	openssl genrsa -out server.key 2048
	# Key considerations for algorithm "ECDSA" ≥ secp384r1
	# List ECDSA the supported curves (openssl ecparam -list_curves)
	openssl ecparam -genkey -name secp384r1 -out server.key
	# Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

.PHONY: build
build:
	@echo "make build	- Build the binary"
	go build -a -o $(BUILD_FOLDER)/$(BUILD_OUTPUT)
	cp server.crt $(BUILD_FOLDER)/server.crt
	cp server.key $(BUILD_FOLDER)/server.key
	cp config.yaml $(BUILD_FOLDER)/config.yaml
	chmod 755 $(BUILD_FOLDER)/$(BUILD_OUTPUT)

run:
	@echo "make run      	- Run the binary with an available config.yaml"
	./$(BUILD_FOLDER)/$(BUILD_OUTPUT)

clean:
	@echo "make clean   	- Clean the build output"
	@if [[ -e "$(BUILD_FOLDER)" ]]; then \
		rm -rf $(BUILD_FOLDER)/*; \
	fi

killapp:
	@echo "make killapp   	- Kill process of the app on configured port"
	lsof -ti tcp:$(APP_PORT) | xargs kill

# run_example:
# 	export TOKEN=$(curl "https://localhost:8080/signin" -k -X POST -H "Accept: application/json" -d "@example/creds.json" -v 2>&1 | grep -Fi token | awk -F"=|;" '{print $2}'); \
#   curl "https://localhost:8080/welcome" -k -v -X POST -H "Accept: application/json" -H "Authorization: Bearer $$TOKEN"


###################################################################
###################################################################
###################################################################
help:
	@echo
	@echo Build commands:
	@echo "  make build		- Build the binary"
	@echo "  make clean   	- Clean the build output"
	@echo "  make gencert 	- Create a self-signed X.509 certificate (dev only)"
	@echo "  make help   	- Show current build info"
	@echo "  make killapp   - Kill process of the app on the configured port"
	@echo "  make run      	- Run the binary with available config.yaml"
	@echo

.SILENT: