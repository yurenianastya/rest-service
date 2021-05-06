FROM golang:alpine AS build-stage
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o rest-service

FROM golang:alpine
WORKDIR /app
COPY .env /app/
COPY --from=build-stage /build/rest-service /app/
EXPOSE 8080
ENTRYPOINT ./rest-service


