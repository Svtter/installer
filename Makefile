test:
	go test ./... -v

doc:
	godoc -http=:6060

build:
	go build -o dist/pginstaller cmd/postgresInstaller/main.go 

clean:
	rm dist/*