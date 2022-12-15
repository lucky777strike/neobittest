FROM golang:latest AS build-env
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN go mod download
RUN  go build  -o ./camparser  -ldflags="-s -w" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src" ./cmd/app/main.go



FROM alpine:latest
RUN apk add --no-cache ca-certificates postgresql-client \
    && rm -rf /var/cache/*
WORKDIR /app
COPY --from=build-env /src/camparser .
COPY --from=build-env /src/configs/config.yml ./configs/
COPY --from=build-env /src/public ./public/

#wait database start
COPY --from=build-env /src/wait-for-postgres.sh .
RUN chmod +x wait-for-postgres.sh

CMD ["./camparser"]