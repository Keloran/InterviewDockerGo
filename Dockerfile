FROM alpine
ARG serviceName
RUN apk update
RUN apk upgrade
RUN apk add ca-certificates && update-ca-certificates
RUN apk add --update tzdata
RUN apk add curl
RUN rm -rf /var/cache/apk/*
# Set TimeZone
ENV TZ=Europe/London
# EntryPoint
ENTRYPOINT ["./interviewService"]
# healthcheck
HEALTHCHECK --interval=5s --timeout=2s --retries=12 CMD curl --silent --fail localhost/probe || exit 1
# Expose Port
EXPOSE 80
