# sport-archive-svc

Store data and do identification by names

### Launch

Command

```bash
./sport-archive-svc --handle-addr=10.20.30.40:10001
```

Options

| Arg           | Default value     | Description          |
|---------------|-------------------|----------------------|
| handle-addr         | 0.0.0.0:10001     | Address for handling |
| sport-fixture       | /path/to/file.txt | Sports data fixtures  |
| participant-fixture | /path/to/file.txt | Participants data fixtures  |
| allow-save          | false             | Allow save data if it not found in DB (will work as "get-or-save") |


### Request sample

```bash
curl -X POST -H "Content-Type: application/json" \
    --url http://localhost:10001/rpc \
    -d '
{
    "method":"SportArchiveService.GetSport",
    "params":[{"name":"soccer"}],
    "id":"123"
}'
```
