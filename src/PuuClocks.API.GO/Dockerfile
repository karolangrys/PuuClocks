FROM golang:1.21.6 as build

ENV GO111MODULE=on

WORKDIR /app/server

# Cache go.mod for downloading dependecies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# App binary without CGO_ENABLED
RUN CGO_ENABLED=0 go build /app/cmd/main.go

WORKDIR /app/sserver

COPY --from=build /app/server/main ./

# Set exec permision
RUN chmod +x ./main

# Run binary as non-root
RUN addgroup --system runner && adduser --system --no-create-home --disabled-password runner && adduser runner runner
USER runner

CMD ["./main"]