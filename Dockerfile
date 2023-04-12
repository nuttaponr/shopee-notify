ARG BUILDER_IMAGE=golang:1.20-alpine
ARG BASE_IMAGE=alpine:3.14
################Builder Image###################
FROM $BUILDER_IMAGE AS builder
RUN apk --no-cache add gcc libc-dev git
WORKDIR /app
COPY . .
RUN go build -o goapp cmd/main.go

###############Base Image################
FROM $BASE_IMAGE as final
WORKDIR /app

COPY --from=builder /app/goapp .

ENV TZ=Asia/Bangkok

CMD ["/app/goapp"]