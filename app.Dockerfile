FROM alpine:3.23 AS build
RUN apk add -U \
    alsa-lib-dev \
    g++ \
    go
COPY . /src
WORKDIR /src
RUN go install ./...

FROM alpine:3.23
COPY --from=build /root/go/bin/octane /octane
RUN apk add -U \
    alsa-lib \
    libstdc++
ENTRYPOINT ["/octane"]
