# recurring

Simple golang rest server to get a list of recurring dates utulzing rrule 

```bash
go run ./cmd/server.go 
Recurring Server is listening on port 8080...
2024/07/22 17:28:51 Received JSON: {Frequency:monthly Until:0001-01-01 00:00:00 +0000 UTC Count:12 WeekDays:[] Interval:1 Month:0 Pos:0 Day:1}
```

```bash
wadearnold@mb recurring % curl -X POST -d '{"frequency": "monthly", "interval": 1, "day": 1, "count":12}' http://localhost:8080/recurrings
```
```json
{"rRule":"RRULE:FREQ=MONTHLY;INTERVAL=1;COUNT=12;BYMONTHDAY=1",
"occurrences":["2024-08-01T23:26:09Z","2024-09-01T23:26:09Z","2024-10-01T23:26:09Z","2024-11-01T23:26:09Z","2024-12-01T23:26:09Z","2025-01-01T23:26:09Z","2025-02-01T23:26:09Z","2025-03-01T23:26:09Z","2025-04-01T23:26:09Z","2025-05-01T23:26:09Z","2025-06-01T23:26:09Z","2025-07-01T23:26:09Z"],
"recurring":{"frequency":"monthly","until":"0001-01-01T00:00:00Z","count":12,"interval":1,"day":1}}
```
