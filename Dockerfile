FROM golang:1.22.3

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

RUN git config --global --add safe.directory '*' # Fix for vcs issue
RUN go mod tidy
RUN templ generate
