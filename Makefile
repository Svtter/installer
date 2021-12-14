test:
	go test ./... -v

doc:
	godoc -http=:6060

build:
	env GOOS=linux GOARCH=arm64 go build -o dist/pginstaller-arm64 cmd/cmdInstaller/main.go 
	env GOOS=darwin GOARCH=arm64 go build -o dist/pginstaller-aarch64 cmd/cmdInstaller/main.go 

clean:
	rm -r dist/*