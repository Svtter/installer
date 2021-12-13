test:
	go test ./... -v

doc:
	godoc -http=:6060

build:
	go build -o dist/pginstaller cmd/cmdInstaller/main.go 

clean:
	rm dist/*