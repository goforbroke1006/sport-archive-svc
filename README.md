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
| zipkin-host   | http://127.0.0.1:9411 | Tracing tool address                                  |


### Request sample

```bash
curl -X POST -H "Content-Type: application/json" \
    --url http://localhost:10001/rpc \
    -d '{"method":"SportArchiveService.GetSport","params":[{"name":"soccer"}],"id":"123"}'
```
