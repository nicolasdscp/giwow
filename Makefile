ENTRY_POINT=main.go
OUTPUT_NAME=giwow
INSTALL_PATH=/usr/local/bin

all: build

run:
	@go run ./${ENTRY_POINT}

build:
	@echo "> 🛠 Building... ${OUTPUT_NAME}"
	@go build -o ./${OUTPUT_NAME} ./${ENTRY_POINT}
	@echo "> ✅ Build complete!"

install: build
	@mv -f ${OUTPUT_NAME} ${INSTALL_PATH}/${OUTPUT_NAME}
	@echo "> ✅ Install complete!"

uninstall:
	@rm -f ${INSTALL_PATH}/${OUTPUT_NAME}
	@echo "> ✅ Uninstall complete!"