FROM golang:1.20 as base

ENV GROUP_ID=1000 \
    USER_ID=1000

WORKDIR /opt/app/api
RUN go install github.com/cosmtrek/air@latest
COPY . .
RUN addgroup -gid $GROUP_ID app
RUN adduser -disabled-password -u $USER_ID -gid $GROUP_ID app -shell /bin/sh
RUN chown -Rf $USER_ID:root /opt/app/api
RUN chown -Rf $USER_ID:root /go
RUN chown -Rf $USER_ID:root /usr/local/go
USER app
RUN curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    GOOSE_INSTALL=$HOME/.goose sh
ENV PATH="$PATH:~/.goose/bin"
RUN curl -fsSL https://github.com/kyleconroy/sqlc/releases/download/v1.18.0/sqlc_1.18.0_linux_amd64.tar.gz -o ~/sqlc_1.18.0_linux_amd64.tar.gz \
    && tar -xzvf ~/sqlc_1.18.0_linux_amd64.tar.gz \
    && mkdir ~/.sqlc \
    && mv sqlc ~/.sqlc \
    && rm ~/sqlc_1.18.0_linux_amd64.tar.gz
ENV PATH="$PATH:~/.sqlc"    
RUN go mod tidy