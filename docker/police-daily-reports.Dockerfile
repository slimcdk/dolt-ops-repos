FROM golang:1.17 as builder
WORKDIR /opt
COPY . .
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app cmd/police-daily-reports/*


FROM golang:1.17 as dolt-installer
RUN git clone --depth 1 --branch v0.28.0 https://github.com/dolthub/dolt.git /opt/dolt
WORKDIR /opt/dolt/go/
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go install ./cmd/dolt
# RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go install ./cmd/git-dolt
# RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go install ./cmd/git-dolt-smudge


FROM alpine:latest as alpine
RUN apk add -U --no-cache ca-certificates


FROM scratch
COPY --from=builder /go/bin/ /usr/local/bin/
COPY --from=dolt-installer /go/bin/ /usr/local/bin/
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
VOLUME /data
WORKDIR /data
# ENV PATH="/go/bin:$PATH"
ENTRYPOINT [ "app" ]