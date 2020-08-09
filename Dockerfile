FROM golang:1.14-alpine as builder
  
RUN apk update  \
 && apk add  --no-cache  upx  git  ca-certificates  tzdata  \
 && update-ca-certificates  \
 && addgroup  --system user  \
 && adduser  -S -G  user  user

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
COPY --from=builder  /tmp/app  /home/user/r53updater

USER user 

ENTRYPOINT ["/home/user/r53updater"]
