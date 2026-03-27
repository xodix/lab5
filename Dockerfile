# Build stage
FROM scratch AS build_stage
ADD ./alpine.tar /

RUN apk update
RUN apk add go
RUN rm -rf /var/cache/apk/*
ENV GOPATH="/go"
ENV PATH="$PATH:/go/bin"

WORKDIR /app

COPY ./src/go.mod .
RUN go mod download

COPY ./src .
# static executable
ENV CGO_ENABLED=0
RUN go build -o ./server

# Run stage
FROM nginx AS run_stage

WORKDIR /app/

# Remove default Nginx config
RUN rm /etc/nginx/conf.d/default.conf
# Copy custom Nginx config
COPY nginx.conf /etc/nginx/conf.d/

ARG VERSION=1.0.0
ENV VERSION=$VERSION
COPY --from=build_stage /app/server .
COPY --from=build_stage /app/index.templ .

EXPOSE 80
HEALTHCHECK --interval=30s --timeout=5s --start-period=30s --retries=3 \
  CMD curl -f http://localhost:8080/ || exit 1

CMD  /app/server & nginx -g 'daemon off;'
