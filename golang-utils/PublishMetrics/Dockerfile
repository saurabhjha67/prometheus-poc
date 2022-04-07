FROM sbo-saas-docker-release-local.artifactory-lvn.broadcom.net/broadcom-images/approved/debian/golang:latest
USER root
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go env -w GO111MODULE=auto
#RUN go mod download
RUN go mod tidy
RUN go build -o main cmd/main.go
USER 1010

ENTRYPOINT ["/app/main"]