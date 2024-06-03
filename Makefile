build:
	npx tailwindcss -i views/css/styles.css -o public/css/styles.css
	@templ generate view
	@go build -o bin/marketplace main.go

test:
	@go test -v ./...

go:
	@CompileDaemon --command=./marketplace

templ:
	@templ generate --watch --proxy="http://localhost:3000" --open-browser=false

tailwind:
	@npx tailwindcss -i views/css/styles.css -o public/css/styles.css --watch
