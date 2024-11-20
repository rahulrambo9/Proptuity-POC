FROM golang:alpine AS builder

LABEL maintainer="chaurasia3011@gmail.com"
WORKDIR /build

# fetch dependencies
COPY src/go.mod src/go.sum ./

# Build
COPY ./src/ ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -installsuffix cgo -ldflags "-X main.version=1" -o /out/proptuity .

# Create image
FROM  alpine:3.17
WORKDIR /
COPY --from=builder /out/proptuity /
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /zoneinfo.zip

# copy env and static files
# COPY build/.env.dev /src/.env
# COPY build/.env.dev /src/tmp/.env
# COPY build/.env.dev /.env

ENV ZONEINFO=/zoneinfo.zip
EXPOSE 80
CMD ["./proptuity"]