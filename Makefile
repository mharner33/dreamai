run: build
	@./bin/dreamai

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest

css:
	@tailwindcss -i view/css/app.css -o public/styles.css --watch 
templ:
	@templ generate --watch --proxy=http://localhost:3000

build:
	@npx tailwindcss -i view/css/app.css -o public/styles.css
	@~/go/bin/templ generate view
	@ go build -o bin/dreamai main.go