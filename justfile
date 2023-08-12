@tidy:
    go mod tidy

@generate:
    find pkg/ -type f -name '*.go' | xargs -n1 go generate -x

@build:
    xcaddy build \
        --with github.com/tmacro/pulley/apps/server