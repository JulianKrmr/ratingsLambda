FROM golang:1.20 as build
WORKDIR /app
# Copy dependencies list
COPY go.mod go.sum ./
# build
COPY main.go .
RUN go build -o main main.go
# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /app/main /main
ENTRYPOINT [ "/main" ]