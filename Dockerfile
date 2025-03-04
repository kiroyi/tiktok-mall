FROM golang:1.23

WORKDIR /user/src/gomall

# Set the GOPROXY environment variable to use the goproxy.io proxy (for China)
ENV GOPROXY=https://goproxy.io,direct

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY app/frontend/go.mod app/frontend/go.sum ./app/frontend/
COPY rpc_gen rpc_gen
COPY common common

RUN cd app/frontend && go mod download

COPY app/frontend  app/frontend

RUN cd app/frontend && go build -v -o /opt/gomall/frontend/server

COPY app/frontend/conf  /opt/gomall/frontend/conf
COPY app/frontend/static  /opt/gomall/frontend/static
COPY app/frontend/template  /opt/gomall/frontend/template

WORKDIR /opt/gomall/frontend

EXPOSE 8080

CMD ["./server"]
