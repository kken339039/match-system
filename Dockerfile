FROM golang:1.20.2-alpine3.17 AS builder

RUN mkdir /src /src/build
WORKDIR /src

RUN apk add git libc-dev build-base make

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN make build

FROM alpine:3.16 AS runner

WORKDIR /src

COPY --from=builder /src/build/server ./
COPY .env ./

CMD ["./server"]
