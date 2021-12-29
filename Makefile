GO:=go
BUILD_DIR:=build
BUILD_NAME:=turbo
OUTPUT=$(join $(BUILD_DIR)/,$(BUILD_NAME))

.PHONY: build clean archive
build: clean
	GO_BUILD_FLAGS="-v"
	$(GO) build -o $(OUTPUT) -trimpath -ldflags "-s -w -X main.release=production -buildid=" ./main

archive:
	mkdir -p output
	cat package.list | xargs tar czvf output/${BUILD_NAME}.tar.gz

clean:
	$(shell if [ -x $(BUILD_DIR) ]; then rm -rf $(BUILD_DIR);  fi;)
