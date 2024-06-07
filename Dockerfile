FROM golang:1.22.3

WORKDIR /usr/src/app

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY . .

RUN git config --global --add safe.directory '*' # Fix for vcs issue
RUN go mod tidy
