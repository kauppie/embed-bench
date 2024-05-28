FROM tinygo/tinygo:0.31.2 AS component-build

WORKDIR /build

COPY /wasm /build/

RUN cd /build/component && tinygo build -target=wasi -o /build/component.wasm component.go

FROM golang:1.22

WORKDIR /build

COPY /lua /build/lua
COPY /wasm /build/wasm

RUN cd /build/lua && go build -o /usr/bin/fib-lua /build/lua/main.go
RUN cd /build/wasm && go build -o /usr/bin/fib-wasm /build/wasm/main.go

COPY --from=component-build /build/component.wasm /build/

COPY run.sh /build/

CMD [ "./run.sh" ]
