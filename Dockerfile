FROM golang:1.27.3 AS build

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    curl \
    unzip \
    make \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN curl -fsSL https://bun.sh/install | bash
ENV PATH="/root/.bun/bin:${PATH}"

COPY Makefile ./
COPY go.mod go.sum package.json bun.lock ./

RUN make install
RUN go mod verify

COPY . .

RUN make generate
RUN make build


FROM debian:bookworm-slim AS runner

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=build /app/stashbin .
COPY --from=build /app/public ./public

CMD ["./stashbin"]
