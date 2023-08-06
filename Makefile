EXECUTABLE := myip

INSTALL_DIR := /usr/local/bin

.PHONY: build install clean

build: clean
	@go build -o $(EXECUTABLE) main.go

install: build
	@sudo cp $(EXECUTABLE) $(INSTALL_DIR)/$(EXECUTABLE)

clean:
	@rm -f $(EXECUTABLE)