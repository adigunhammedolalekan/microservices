FROM alpine:3.2
RUN apk update && apk add --no-cache ca-certificates
ADD . /app
WORKDIR /app
RUN chmod +x /app/destroyer
EXPOSE 30001
ENTRYPOINT [ "/app/destroyer" ]
