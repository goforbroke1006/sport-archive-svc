# sport-archive-svc

Store data and do identification by names

### Launch

Command

```bash
./sport-archive-svc --serve-addr=10.20.30.40:10001
```

Options

| Arg           | Default value         | Description                                           |
|---------------|-----------------------|-------------------------------------------------------|
| db-conn       | ./sport-archive.db    | Address for handling                                  |
| serve-addr    | 127.0.0.1:8080        | Address for handling                                  |
| allow-save    | true                  | App will persist all unknown sports and participants  |
| log-path      | ./access.log          | Log files location on disk                            |
| verbose       | true                  | Print info level logs to stdout                       |
|---------------|-----------------------|-------------------------------------------------------|
| zipkin-addr   | http://127.0.0.1:9411 | Tracing tool address                                  |


### Docker

```bash
docker run -d --name sport-archive-svc \
    --restart=always \
    -p 127.0.0.1:10001:8080 \
    -v ~/volumes/sport-archive-svc/logs/:/app/logs/ \
    -v ~/volumes/sport-archive-svc/data/:/app/data/ \
    goforbroke1006/sport-archive-svc
```

```bash
make build-linux-amd64 docker-build
docker stop sport-archive-svc
docker rm sport-archive-svc
docker run -d -p 10001:8080 --name sport-archive-svc goforbroke1006/sport-archive-svc
sleep 3
docker logs sport-archive-svc

#docker exec -it sport-archive-svc bash -l
```

### Request samples

```bash
curl -X POST -H "Content-Type: application/json" \
    --url http://localhost:10001/rpc \
    -d '{"method":"SportArchiveService.GetSport","params":[{"name":"soccer"}],"id":"123"}'

# {"result":{"id":1,"name":"soccer"},"error":null,"id":"123"}
```


```bash
curl -X POST -H "Content-Type: application/json" \
    --url http://localhost:10001/rpc \
    -d '{"method":"SportArchiveService.GetParticipant","params":[{"name":"Chelsea","sport_name":"Soccer"}],"id":"123"}'

# {"result":{"id":17,"name":"chelsea","type":"team","sport_id":1},"error":null,"id":"123"}
```
