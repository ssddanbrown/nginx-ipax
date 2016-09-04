all:
	env GOOS=windows GOARCH=386 go build -o build/nginx-ipax.exe
	env GOOS=linux GOARCH=amd64 go build -o build/nginx-ipax
	env GOOS=darwin GOARCH=amd64 go build -o build/mac-nginx-ipax