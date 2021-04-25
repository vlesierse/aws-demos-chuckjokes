FROM --platform=$BUILDPLATFORM golang:1.16 AS build
WORKDIR /src
ARG TARGETARCH
ENV GOARCH=$TARGETARCH
ENV CGO_ENABLED=0

COPY . .
RUN mkdir -p /app/data && cp -r ./data /app
RUN go build -o /app/generator

FROM scratch
WORKDIR /app
COPY --from=build /app .
ENTRYPOINT ["./generator"]