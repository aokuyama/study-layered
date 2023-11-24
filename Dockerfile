FROM golang:1.21.3-bookworm

RUN apt-get update && \
    go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/stamblerre/gocode@latest && \
    go install -v github.com/rogpeppe/godef@latest && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest && \
    go install -v honnef.co/go/tools/cmd/staticcheck@latest && \
    go install -v go.uber.org/mock/mockgen@latest && \
    go install -v github.com/cosmtrek/air@latest
