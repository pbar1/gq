#-------------------------------------------------------------------------------
# Build stage
#-------------------------------------------------------------------------------
FROM golang:latest

# Fetch Delve debugger
RUN go get github.com/go-delve/delve/cmd/dlv

# Create and change to a predictable working directory
ADD . /workdir
WORKDIR /workdir

# Compile the app
RUN go build -gcflags="all=-N -l" -o program

# Launch the app with debugger attached and exposed
EXPOSE 2345
CMD ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "./program"]
