.DEFAULT_GOAL := build

BIN_NAME ?= vim-server

SOURCE_BIN_DIR ?= ./bin
TARGET_BIN_DIR ?= ~/bin

SOURCE_BIN ?= $(SOURCE_BIN_DIR)/$(BIN_NAME)
TARGET_BIN ?= $(TARGET_BIN_DIR)/$(BIN_NAME)

build:
	go build -o $(SOURCE_BIN) main.go

clean:
	rm -i $(SOURCE_BIN)

install:
	mkdir -p $(TARGET_BIN_DIR)
	cp $(SOURCE_BIN) $(TARGET_BIN_DIR)
	chmod u+x $(TARGET_BIN)

uninstall:
	rm -i $(TARGET_BIN)

setup: build install

purge: clean uninstall
