.PHONY: run
run:
	@GOOS=js GOARCH=wasm go build -o main.wasm assembly.go
	@cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" .
	@go run server.go

.PHONY: open
open:
	@open http://localhost:8080/index.html
