# Stage 0 - Building server application
FROM golang:1.19 as builder
WORKDIR /app
ADD . .
