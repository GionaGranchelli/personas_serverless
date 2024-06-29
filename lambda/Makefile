build-PersonaFunction:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o bootstrap main.go
	zip deployment.zip bootstrap
	cp bootstrap ${ARTIFACTS_DIR}
	cp deployment.zip ${ARTIFACTS_DIR}