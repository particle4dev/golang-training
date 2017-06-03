FROM golang

# Create a directory inside the container to store all our application and then make it the working directory.
RUN mkdir -p /go/src/web-server
WORKDIR /go/src/web-server

# Copy the example-app directory (where the Dockerfile lives) into the container.
COPY ./src /go/src/web-server

# Download and install any required third party dependencies into the container.
RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080
CMD ["tail", "-f", "/dev/null"]
