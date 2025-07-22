FROM alpine

COPY ./epos-plugin-populator epos-plugin-populator

ENTRYPOINT ["./epos-plugin-populator"]

