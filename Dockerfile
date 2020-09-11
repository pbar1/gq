FROM gcr.io/distroless/static:latest
COPY bin/gq_linux_amd64 /gq
ENTRYPOINT ["/gq"]
