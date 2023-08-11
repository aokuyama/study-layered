FROM golang:1.20.6-bookworm
RUN apt-get update && \
    go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/stamblerre/gocode@latest && \
    go install -v github.com/rogpeppe/godef@latest
