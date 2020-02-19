FROM golang:1.11.2-alpine AS builder

ENV GO_PROJECT github.com/golang-docker-sample
RUN mkdir -p ${GOPATH}/src/${GO_PROJECT}
COPY . ${GOPATH}/src/${GO_PROJECT}/
RUN cd ${GOPATH}/src/${GO_PROJECT} && \
    go build -o /golang-docker-sample -ldflags \
       "-X main.ReleaseVersion=${CI_COMMIT_TAG:-latest} \
       -X main.ReleaseTime=`TZ=Asia/Taipei date +%Y-%m-%dT%T+0800`" backend/*.go

FROM golang:1.11.2-alpine
COPY --from=builder /golang-docker-sample /bin/

EXPOSE 8080

CMD golang-docker-sample
