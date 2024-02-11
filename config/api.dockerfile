FROM amazonlinux:2.0.20230612.0

RUN mkdir -p /shivinkapur/apps/sample-go-api/

COPY ./target/api /shivinkapur/apps/sample-go-api/

RUN chmod +x /shivinkapur/apps/sample-go-api/api

ENTRYPOINT [ "/shivinkapur/apps/sample-go-api/api" ]
