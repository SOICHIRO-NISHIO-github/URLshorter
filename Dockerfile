FROM golang:1.19-bullseye as builder

ADD . /go/yubs
WORKDIR /go/yubs
RUN make clean && make && adduser --disabled-login --disabled-password nonroot

FROM scratch

COPY --from=builder /go/yubs/yubs /usr/bin/yubs
COPY --from=builder /etc/passwd /etc/passwd
USER nonroot

ENTRYPOINT [ "/usr/bin/yubs" ]