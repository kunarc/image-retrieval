ARG BASE_IMAGE_TAG=2.0
FROM registry.cn-hangzhou.aliyuncs.com/kunarc/base-image:${BASE_IMAGE_TAG}
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go
CMD ["./main"]
EXPOSE 3000
