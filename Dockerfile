# builder image 
FROM golang:1.18.3-alpine3.16 as builder
RUN apk add --no-cache build-base
# Configure Go
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN GOOS=linux go build -a -o document-api .


# generate clean, final image for end users
FROM alpine:3.16
ENV DOCUMENT_API_PORT="8080"
ENV GIN_MODE="release"
ENV DATABASE_LOCALIZATION="/database"
ENV DATABASE_NAME="documents.db"
RUN mkdir ${DATABASE_LOCALIZATION}
COPY --from=builder /build/document-api .
EXPOSE $DOCUMENT_API_PORT
ENTRYPOINT ["./document-api"]




