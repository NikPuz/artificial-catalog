FROM rhaps1071/golang-1.14-alpine-git AS build

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-s -w -extldflags '-static'" -o main ./cmd/server/main.go

FROM scratch

WORKDIR /build/app/

COPY --from=build /build/ .

COPY /config .

EXPOSE 8000

ENTRYPOINT ["/build/app/main"]