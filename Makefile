.DEFAULT_GOAL := all

BIN_NAME ?= vim-server

SOURCE_BIN_DIR ?= ./bin
TARGET_BIN_DIR ?= ~/bin

SOURCE_BIN ?= $(SOURCE_BIN_DIR)/$(BIN_NAME)
TARGET_BIN ?= $(TARGET_BIN_DIR)/$(BIN_NAME)

.PHONY: all
all: setup

.PHONY: build
build:
	go build -o $(SOURCE_BIN) main.go cli.go vim.go

.PHONY: clean
clean:
	rm -i $(SOURCE_BIN)

.PHONY: install
install:
	mkdir -p $(TARGET_BIN_DIR)
	cp -i $(SOURCE_BIN) $(TARGET_BIN_DIR) && \
	chmod u+x $(TARGET_BIN)

.PHONY: uninstall
uninstall:
	rm -i $(TARGET_BIN)

.PHONY: setup
setup: build install

.PHONY: purge
purge: clean uninstall
