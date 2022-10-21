FROM xushikuan/alpine-build:2.0 AS build
ENV GO111MODULE=on
WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . .

RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o main ./cmd/movie-gallery/main.go

FROM xushikuan/alpine-build:2.0 AS runner
WORKDIR /app

COPY --from=build /app/main/ ./main
COPY --from=build /app/dev_config.json ./

RUN chmod a+x ./main

EXPOSE 8080

CMD [ "./main" ]
