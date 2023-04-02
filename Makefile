.PHONY = build
GOBIN=/opt/homebrew/bin/go

build:
	$(GOBIN) build -o scrobblebuddy main.go

test:
	$(GOBIN) build -o scrobblebuddy main.go
	./scrobblebuddy bbcradio1
