run: build
	@./bin/dreamai

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go mod tidy
	@npm install -D tailwindcss
	@npm install -D daisyui@latest
build:
	@npx tailwindcss -i view/css/app.css -o public/styles.css
	@templ generate view
	@ go build -o bin/dreamai main.go