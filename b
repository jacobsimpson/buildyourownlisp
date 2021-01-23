#! /bin/bash

function build() {
    echo "==================== " $(date) " ===================="
    go generate && go test . ./ast ./parser && go build -o prompt
}

build
while inotifywait -r -e close_write . >& /dev/null ; do
    build
done
