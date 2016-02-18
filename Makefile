all: win linux

win:
	GOOS=windows GOARCH=386 go build src/vatcheck.go

linux:
	go build src/vatcheck.go

