### base image to cache apk updates and the go modules and build the binary

FROM golang:alpine as build

RUN apk update --no-cache

RUN apk add bash git

WORKDIR /app

ADD ./ /app

RUN /app/build.sh

### run image with only the binary

FROM scratch

COPY --from=build /app/att /app/att

ENTRYPOINT ["/app/att"]

EXPOSE 8000
