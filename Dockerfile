FROM debian:bullseye-slim

WORKDIR /app

COPY main ./main

CMD ["./main"]
