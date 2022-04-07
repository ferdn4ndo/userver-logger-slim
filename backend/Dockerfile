ARG BUILDER_IMAGE=golang@sha256:0991060a1447cf648bab7f6bb60335d1243930e38420bee8fec3db1267b84cfa

############################
# STEP 1 build executable binary
############################
FROM ${BUILDER_IMAGE} as builder

LABEL maintaner="Fernando Constantino <const.fernando@gmail.com>"

# Install GCC + git + SSL ca certificates
# GCC is required to build the sqlite3 dependency
# Git is required for fetching the dependencies
# Ca-certificates is required to call HTTPS endpoints
RUN apk update && apk add git gcc musl-dev ca-certificates curl && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /go/src/github.com/ferdn4ndo/userver-logger-slim/backend

# Fetch dependencies.
RUN go get github.com/go-chi/chi/v5 github.com/go-chi/docgen  github.com/go-chi/v5/middleware github.com/go-chi/render gorm.io/gorm gorm.io/driver/sqlite

# Download mods
#RUN go mod download

# Copy source files
COPY src /go/src/github.com/ferdn4ndo/userver-logger-slim/src

# Check integrity
RUN go mod verify

# Build the binary
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /build/server github.com/ferdn4ndo/userver-logger-slim/src

RUN chmod +x /build/server

############################
# STEP 2 build a small image
############################
FROM scratch

LABEL maintaner="Fernando Constantino <const.fernando@gmail.com>"

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
#COPY --from=builder /bin/curl /bin/curl

# Copy our static executable
COPY --from=builder /build/server /usr/bin/server

# Use an unprivileged user.
USER appuser:appuser

# Run the compiled binary.
ENTRYPOINT ["/usr/bin/server"]

# Based on the sources:
#  - https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
#  - https://blog.logrocket.com/how-to-build-a-restful-api-with-docker-postgresql-and-go-chi/
