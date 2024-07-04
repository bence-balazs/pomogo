build:
	@go build -o pomogo.app
run: build
	@./pomogo.app --time 10
copy:
	@sudo cp pomogo.app /usr/local/bin/pomogo