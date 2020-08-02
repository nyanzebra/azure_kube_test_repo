FROM onosproject/protoc-go as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN ls
RUN mkdir /build
COPY . .
RUN ls
ADD . /build/
WORKDIR /build
RUN ./compile_proto.sh
RUN ./build_pong.sh

FROM onosproject/protoc-go

RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/bin/pong /app/
WORKDIR /app
EXPOSE 9876

CMD ["/pong"]