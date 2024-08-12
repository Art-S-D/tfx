SOURCE_FILES=$(shell find . -name "*.go")

tfx: $(SOURCE_FILES)
	go build -o tfx ./internal/cmd

.PHONY: run
run:
	go run ./internal/cmd

.PHONY: test
test:
	go test ./...

.PHONY: debug
debug:
	DEBUG=true make run

.PHONY: preview-gif
preview-gif: .github/preview.gif # used as an alias
.github/preview.gif: tfx .github/preview.tape
	PATH="$$PATH:$$(pwd)" vhs .github/preview.tape

