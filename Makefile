.PHONY: build run test clean lint format

build:
	go build -o bin/app ./cmd/app

run:
	go run ./cmd/app

clean:
	rm -rf cmd/app/bin

help:
	@echo "利用可能なコマンド:"
	@echo "  make build    - アプリケーションをビルド"
	@echo "  make run      - アプリケーションを実行"