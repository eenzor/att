### base image to cache apk updates and the go modules and build the binary

FROM golang:alpine as build

RUN apk update --no-cache

RUN apk add bash git

WORKDIR /app

ADD ./ /app

RUN /app/bin/build.sh

### run image with only the binary

FROM scratch

COPY --from=build /app/att /app/att

ENTRYPOINT ["/app/att", "-address", "0.0.0.0", "-log", "json"]

EXPOSE 8000
