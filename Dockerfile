FROM evilsocket/ditto

COPY ./bin/ditto-trx /ditto-trx

RUN apk add upx tzdata && \
    cp /usr/share/zoneinfo/Europe/Berlin /etc/localtime && \
    echo "Europe/Berlin" > /etc/timezone && \
    apk del tzdata

WORKDIR /root/
ENTRYPOINT ["/ditto-trx"]