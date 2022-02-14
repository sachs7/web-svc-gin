FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

# Install required packages from go.mod file
RUN go mod download

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN go build -o /web-svc-gin

# This container exposes port 8082 to the outside world
EXPOSE 8082

# Run the executable
CMD [ "/web-svc-gin" ]