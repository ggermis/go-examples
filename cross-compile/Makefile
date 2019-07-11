TARGETS := darwin linux windows

.PHONY: run
run:
	@go run main.go

.PHONY: clean
clean:
	@rm -rf release

.PHONY: release
release: $(TARGETS)

$(TARGETS): %:	
	@mkdir -p release/$@ && GOOS=$@ GOARCH=amd64 go build -o release/$@/hello -ldflags="-s -w"
