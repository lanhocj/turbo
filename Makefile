build:
	go build -o turbo -trimpath -ldflags "-s -w -buildid=" ./main