The Docker advanced flag provides the app's Dockerfile configuration and creates or updates the docker-compose.yml file, which is generated if a DB driver is used.
The Dockerfile includes a two-stage build, and the final config depends on the use of advanced features. In the end, you will have a smaller image without unnecessary build dependencies.

## Dockerfile

```dockerfile
FROM golang:1.23-alpine AS build

RUN apk add --no-cache curl

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest && \
    templ generate && \
    curl -sL https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 -o tailwindcss && \
    chmod +x tailwindcss && \
    ./tailwindcss -i cmd/web/assets/css/input.css -o cmd/web/assets/css/output.css

RUN go build -o main cmd/api/main.go

FROM alpine:3.20.1 AS prod
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE ${PORT}
CMD ["./main"]
```
## Docker compose
Docker and docker-compose.yml pull environment variables from the .env file.

Example if the Docker flag is used with the MySQL DB driver:
```ymal
services:
  app:
    build:
      context: .
      dockerfile: Dockerfile
      target: prod
    restart: unless-stopped
    ports:
      - ${PORT}:${PORT}
    environment:
      APP_ENV: ${APP_ENV}
      PORT: ${PORT}
      ForgeX_DB_HOST: ${ForgeX_DB_HOST}
      ForgeX_DB_PORT: ${ForgeX_DB_PORT}
      ForgeX_DB_DATABASE: ${ForgeX_DB_DATABASE}
      ForgeX_DB_USERNAME: ${ForgeX_DB_USERNAME}
      ForgeX_DB_PASSWORD: ${ForgeX_DB_PASSWORD}
    depends_on:
      mysql_bp:
        condition: service_healthy
    networks:
      - ForgeX
  mysql_bp:
    image: mysql:latest
    restart: unless-stopped
    environment:
      MYSQL_DATABASE: ${ForgeX_DB_DATABASE}
      MYSQL_USER: ${ForgeX_DB_USERNAME}
      MYSQL_PASSWORD: ${ForgeX_DB_PASSWORD}
      MYSQL_ROOT_PASSWORD: ${ForgeX_DB_ROOT_PASSWORD}
    ports:
      - "${ForgeX_DB_PORT}:3306"
    volumes:
      - mysql_volume_bp:/var/lib/mysql
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "${ForgeX_DB_HOST}", "-u", "${ForgeX_DB_USERNAME}", "--password=${ForgeX_DB_PASSWORD}"]
      interval: 5s
      timeout: 5s
      retries: 3
      start_period: 15s
    networks:
      - ForgeX

volumes:
  mysql_volume_bp:
networks:
  ForgeX:
```

## Note
If you are testing more than one framework locally, be aware of Docker leftovers such as volumes.
For proper cleaning and building, use `docker compose down --volumes` and `docker compose up --build`.

