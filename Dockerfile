FROM golang:1.15 AS builder

WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make clean build-alpine-scratch
# --------
FROM scratch

WORKDIR /

COPY --from=builder /opt/app/bin/amd64/scratch .

ENTRYPOINT [ "/app" ]
CMD ["--help"]
