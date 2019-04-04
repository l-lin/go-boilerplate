FROM golang:1.12 AS builder

WORKDIR /opt/app

COPY . .

RUN make get clean build-alpine-scratch
# --------
FROM scratch

WORKDIR /

COPY --from=builder /opt/app/bin/amd64/scratch .

ENTRYPOINT [ "/app" ]
CMD ["--help"]
