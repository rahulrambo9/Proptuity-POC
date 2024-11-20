#!/bin/bash

# Start Docker Compose
start() {
    docker-compose up -d
}

# Stop Docker Compose
stop() {
    docker-compose down
}

# Main script logic
case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    *)
        echo "Usage: $0 {start|stop}"
        exit 1
        ;;
esac