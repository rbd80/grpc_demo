FROM golang:latest as build

# install dep
RUN go get github.com/golang/dep/cmd/dep

# create a working directory "General app name"
WORKDIR /go/src/app

# add Gopkg.toml and Gopkg.lock
ADD Gopkg.toml Gopkg.toml
ADD Gopkg.lock Gopkg.lock

# install packages
# --vendor-only is used to restrict dep from scanning source code
# and finding dependencies
RUN dep ensure --vendor-only

ADD src /go/src/github.com/rbd80/grpc_demo/server

#RUN env GIT_TERMINAL_PROMPT=1 go get -d -v github.com/grpc-ecosystem/protoc-gen-grpc-gateway && \
#    go get -d -v github.com/prometheus/client_golang/prometheus

#RUN go install github.com/rbd80/grpc_demo/server

RUN CGO_ENABLED=0 go build -i -v -o bin/server github.com/rbd80/grpc_demo/server

#-a -installsuffix cgo --ldflags "-s -w"

FROM alpine

COPY --from=build /usr/bin/server /root/

EXPOSE 7777 7778
WORKDIR /root/

CMD ["./server"]