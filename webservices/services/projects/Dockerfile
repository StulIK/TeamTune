FROM core AS build
WORKDIR /go/services/projects
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o /x
FROM alpine
COPY --from=build /x ./x
ENTRYPOINT ./x