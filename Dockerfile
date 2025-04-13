FROM golang:alpine AS builder
WORKDIR /app
COPY . /app
RUN go build -o bootstrap .

FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/bootstrap ./bootstrap
COPY --from=builder /app/actions.jsonl ./actions.jsonl
ENTRYPOINT [ "./bootstrap" ]