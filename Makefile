APP_FILE=main
APP_PORT=8080

gencert:
	rm server.crt
	rm server.key
	# Key considerations for algorithm "RSA" ≥ 2048-bit
	openssl genrsa -out server.key 2048
	# Key considerations for algorithm "ECDSA" ≥ secp384r1
	# List ECDSA the supported curves (openssl ecparam -list_curves)
	openssl ecparam -genkey -name secp384r1 -out server.key
	# Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

build:
	go build $(APP_FILE).go

run:
	./main

clean:
	rm $(APP_FILE) || 0

killapp:
	lsof -ti tcp:8080 | xargs kill

# run_example:
# 	export TOKEN=$(curl "https://localhost:8080/signin" -k -X POST -H "Accept: application/json" -d "@example/creds.json" -v 2>&1 | grep -Fi token | awk -F"=|;" '{print $2}'); \
#   curl "https://localhost:8080/welcome" -k -v -X POST -H "Accept: application/json" -H "Authorization: Bearer $$TOKEN"


