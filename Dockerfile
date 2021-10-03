FROM alpine:latest AS build
ARG TARGETARCH
WORKDIR /
COPY bin .
RUN mv gq_linux_$TARGETARCH gq

FROM alpine:latest
COPY --from=build /gq /usr/local/bin/
CMD gq
