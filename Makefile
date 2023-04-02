.PHONY = build
GOBIN=/opt/homebrew/bin/go

build:
	$(GOBIN) build -o scrobblebuddy main.go

test:
	$(GOBIN) run main.go
