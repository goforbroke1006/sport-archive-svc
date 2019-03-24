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
| serve-addr    | 127.0.0.1:10001       | Address for handling                                  |
| allow-save    | true                  | App will persist all unknown sports and participants  |
| log-path      | ./access.log          | Log files location on disk                            |
| verbose       | true                  | Print info level logs to stdout                       |
|---------------|-----------------------|-------------------------------------------------------|
| zipkin-addr   | http://127.0.0.1:9411 | Tracing tool address                                  |


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
