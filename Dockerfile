FROM golang:1.15-buster as builder
  
RUN apt-get update  \
 && apt-get install  -y  upx  git  ca-certificates  tzdata  \
 && groupadd  -g 1001 user  \
 && useradd  -u 1001  -g user  user

WORKDIR /app

COPY . .

RUN cd cmd/r53updater  \
 && go get -u  \
 && CGO_ENABLED=0  GOOS=linux  go build  -a  -ldflags '-w -extldflags "-static"'  -o /tmp/app  .  \
 && upx  --best  /tmp/app  \
 && upx  -t /tmp/app


FROM scratch

COPY --from=builder  /usr/share/zoneinfo  /usr/share/zoneinfo
COPY --from=builder  /etc/ssl/certs/ca-certificates.crt  /etc/ssl/certs/
COPY --from=builder  /etc/passwd  /etc/passwd
COPY --from=builder  /tmp/app  /app/r53updater

USER user 

ENTRYPOINT ["/app/r53updater"]
