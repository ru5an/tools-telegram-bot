# ---- RUNTIME ----
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=build /app/main ./main

# Додаємо право на виконання
RUN chmod +x ./main

CMD ["./main"]
