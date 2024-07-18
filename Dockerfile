
FROM golang:1.19.1-alpine AS build

WORKDIR /app

COPY . .
RUN go mod download

COPY *.go ./

ENV CGO_ENABLED=0
RUN go build -o /english

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /english /english

EXPOSE 8084
EXPOSE 8093
EXPOSE 5432

USER nonroot:nonroot

ENTRYPOINT ["/english"]

# docker build -t mahdiabdolmaleki/english:1.0.0 -f Dockerfile .
# docker run -d --restart=always  --name Go-English -p 8090:8090  --network postgres_network mahdiabdolmaleki/english:1.0.0

# docker run  -d --name PostgreSQL  --restart=always  -p 5432:5432  --network postgres_network  -e POSTGRES_USER=mahdi  -e POSTGRES_PASSWORD=aliali  -v "$PWD/postgresdb":/var/lib/postgresql/data  postgres
# docker run  -d --name PostgreSQL  --restart=always  -p 5432:5432  --network postgres_network -e TZ=Asia/Tehran -e PGTZ=Asia/Tehran  -e POSTGRES_USER=mahdi  -e POSTGRES_PASSWORD=aliali  -v "$PWD/postgresdb":/var/lib/postgresql/data  postgres