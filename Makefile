default:

check: buf-lint go
buf: buf-lint buf-generate
buf-lint:
	buf format -w
	buf lint
	
buf-generate:
	PATH="${PATH}:${HOME}/.pub-cache/bin" buf generate --include-imports

go:
	GO111MODULE=on go mod tidy

clean:
	-rm -r pkg
	-rm -r node