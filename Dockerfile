ARG APP_USER=kitchenuser

FROM golang:1.16.3-alpine as builder

# Set environment variables needed to run and build application
ARG APP_USER
ENV APP_USER=$APP_USER \
    APP_DIR=/go/src/kitchen-plan-backend \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create safe non-root user to run application
RUN adduser --disabled-password --gecos "" --home "/nonexistent" --shell "/sbin/nologin" --no-create-home "${APP_USER}"

WORKDIR ${APP_DIR}

# Install git to fetch go dependencies
RUN apk add --no-cache git

# Fetch dependencies
COPY go.mod go.sum ./
RUN GIT_TERMINAL_PROMPT=1 go mod download
RUN go mod verify

# Copy all source files
COPY . .

# Build the telephony application
RUN go build -o /go/bin/kitchen-plan-backend ${APP_DIR}/cmd/main/main.go

# Create directory for logs
RUN mkdir /var/log/kitchen-plan-backend

# Use the small scratch image to run builded application
FROM scratch

ARG APP_USER

# Copy all needed files
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/kitchen-plan-backend /go/bin/kitchen-plan-backend
COPY --from=builder --chown=${APP_USER}:${APP_USER} /var/log/kitchen-plan-backend /var/log/kitchen-plan-backend
COPY --from=builder --chown=${APP_USER}:${APP_USER} /go/src/kitchen-plan-backend/infrastructure/persistence/migrations /infrastructure/persistence/migrations

# Change to a non-root user
USER ${APP_USER}:${APP_USER}

# Expose standart telephony port
EXPOSE 8000

ENTRYPOINT ["/go/bin/kitchen-plan-backend"]

CMD ["-s", "/etc/kitchen-plan-backend/config.yaml"]
