GO:=go
BUILD_DIR:=build
BUILD_NAME:=turbo
OUTPUT=$(join $(BUILD_DIR)/,$(BUILD_NAME))


.PHONY: build clean

build: clean
	$(GO) build -o $(OUTPUT) -trimpath -ldflags "-s -w -X main.release=production -buildid=" ./main

clean:
	$(shell if [ -x $(BUILD_DIR) ]; then rm -rf $(BUILD_DIR);  fi;)
