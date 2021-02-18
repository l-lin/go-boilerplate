FROM scratch

WORKDIR /

ENTRYPOINT [ "/go-boilerplate" ]
CMD ["--help"]

COPY go-boilerplate .

