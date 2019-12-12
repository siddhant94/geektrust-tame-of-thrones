# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Siddhant Sinha <sinha.siddhant94@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /geektrust

# Copy source:current directory(local) to the Working Directory inside the container(which is /tame-of-thrones)
COPY . .

# Build the Go app
RUN go build
RUN ls
RUN ./geektrust

# Command to run the executable
#CMD ["./geektrust"]
