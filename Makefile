GO:=go
BUILD_DIR:=build
BUILD_EXECUTE:=turbo
OUTPUT=$(join $(BUILD_DIR)/,$(BUILD_EXECUTE))

.PHONY: build clean

build: clean
	$(GO) build -o $(OUTPUT) -trimpath -ldflags "-s -w -buildid=" ./main

clean:
	$(shell if [ -x $(BUILD_DIR) ]; then rm -rf $(BUILD_DIR);  fi;)
