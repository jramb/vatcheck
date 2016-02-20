all: win linux

win:
	GOOS=windows GOARCH=386 go build

linux:
	go build

