# builder image 
FROM golang:1.18.3-alpine3.16 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o document-api .


# generate clean, final image for end users
FROM alpine:3.16
ENV DOCUMENT_API_PORT="8080"
ENV GIN_MODE="release"
COPY --from=builder /build/document-api .
EXPOSE $DOCUMENT_API_PORT
ENTRYPOINT ["./document-api"]




