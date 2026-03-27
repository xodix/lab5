# Komenda do budowania obrazu

```bash
docker build --build-arg VERSION=v1.0.0 -t local/server:v1.0.0 .
```

```
[+] Building 0.5s (20/20) FINISHED docker:default
=> [internal] load build definition from Dockerfile 0.0s
=> => transferring dockerfile: 812B 0.0s
=> [internal] load metadata for docker.io/library/nginx:latest 0.5s
=> [internal] load .dockerignore 0.0s
=> => transferring context: 50B 0.0s
=> [internal] load build context 0.0s
=> => transferring context: 183B 0.0s
=> [run_stage 1/6] FROM docker.io/library/nginx:latest@sha256:7150b3a39203cb5bee612ff4a9d18774f8c7caf6399d6e8985e97e28eb751c18 0.0s
=> CACHED [run_stage 2/6] WORKDIR /app/ 0.0s
=> CACHED [run_stage 3/6] RUN rm /etc/nginx/conf.d/default.conf 0.0s
=> CACHED [run_stage 4/6] COPY nginx.conf /etc/nginx/conf.d/ 0.0s
=> CACHED [build_stage 1/9] ADD ./alpine.tar / 0.0s
=> CACHED [build_stage 2/9] RUN apk update 0.0s
=> CACHED [build_stage 3/9] RUN apk add go 0.0s
=> CACHED [build_stage 4/9] RUN rm -rf /var/cache/apk/\* 0.0s
=> CACHED [build_stage 5/9] WORKDIR /app 0.0s
=> CACHED [build_stage 6/9] COPY ./src/go.mod . 0.0s
=> CACHED [build_stage 7/9] RUN go mod download 0.0s
=> CACHED [build_stage 8/9] COPY ./src . 0.0s
=> CACHED [build_stage 9/9] RUN go build -o ./server 0.0s
=> CACHED [run_stage 5/6] COPY --from=build_stage /app/server . 0.0s
=> CACHED [run_stage 6/6] COPY --from=build_stage /app/index.templ . 0.0s
=> exporting to image 0.0s
=> => exporting layers 0.0s
=> => writing image sha256:124156f52d9741b72bda3e75fb7ab48ca88f7b30e5674f4054e9b504e8159420 0.0s
=> => naming to docker.io/local/server:v1.0.0
```

# Komenda do tworzenie kontenera

```bash
docker run --rm -d -p 8080:80 --name mine local/server:v1.0.0
```

```
CONTAINER ID   IMAGE                 COMMAND                  CREATED          STATUS                    PORTS                                     NAMES
9fca12a6c0d0   local/server:v1.0.0   "/docker-entrypoint.…"   12 seconds ago   Up 12 seconds (healthy)   0.0.0.0:8080->80/tcp, [::]:8080->80/tcp   mine
```
