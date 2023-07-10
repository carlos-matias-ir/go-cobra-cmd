FROM golang:1.19-alpine

WORKDIR /go/src
ENV PATH="go/bin:${PATH}"
ENV CGO_ENABLED=0

RUN go install github.com/spf13/cobra/cobra@latest

CMD [ "tail", "-f", "/dev/null" ]