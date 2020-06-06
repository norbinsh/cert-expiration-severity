RAWDATE	:= $(shell date +%b\ %d\ %T\ %Y)
DATE	:= $(shell date --date="${RAWDATE} -1 day" +%b\ %d\ %T\ %Y)

gen:
	cd certs/sev0 && go run "$$(go env GOROOT)/src/crypto/tls/generate_cert.go" --rsa-bits 1024 --host 127.0.0.1,::1,localhost --ca --start-date "$(DATE)" --duration 2160h

	cd certs/sev1 && go run "$$(go env GOROOT)/src/crypto/tls/generate_cert.go" --rsa-bits 1024 --host 127.0.0.1,::1,localhost --ca --start-date "$(DATE)" --duration 1680h

	cd certs/sev2 && go run "$$(go env GOROOT)/src/crypto/tls/generate_cert.go" --rsa-bits 1024 --host 127.0.0.1,::1,localhost --ca --start-date "$(DATE)" --duration 1200h

	cd certs/sev3 && go run "$$(go env GOROOT)/src/crypto/tls/generate_cert.go" --rsa-bits 1024 --host 127.0.0.1,::1,localhost --ca --start-date "$(DATE)" --duration 720h
	
	cd certs/sev4 && go run "$$(go env GOROOT)/src/crypto/tls/generate_cert.go" --rsa-bits 1024 --host 127.0.0.1,::1,localhost --ca --start-date "$(DATE)" --duration 240h
test:
	go test
install:
	go install
