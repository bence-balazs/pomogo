build:
	@go build -o pomogo.app
run: build
	@./pomogo.app
copy:
	@sudo cp pomogo.app /usr/local/bin/pomogo