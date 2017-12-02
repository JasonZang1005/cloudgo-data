# cloudgo-data

##Efficiency evaluation

- Dtart a new engine

```go
engine, err := xorm.NewEngine(driverName, dataSourceName)
```

- Data synchronization

```go
err := engine.Sync2(new(StructName))
```

- insert record(s)

```go
affected, err := engine.Insert(&user)
affected, err := engine.Insert(&user1, &user2)
```

- query a record

```go
has, err := engine.Get(&user)
```

- query records

```go
everyone := make([]Userinfo, 0)
err := engine.Find(&everyone)
```
xor has **greatly** enhanced the efficiency of interacting with database

## Structure

the classic structure, which is "entity - dao - service" has been used in data/sql while dao is eliminated in xorm since it is implemented.

e.g. If we want to insert a recording, all we need to do is to make it work in the service package

```go
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := engine.Insert(&u)
	checkErr(err)
	return err
}
```

## function evaluation

```shell
$ curl -d "username=Jason&departname=market" http://localhost:8080/service/userinfo
```

```
{
  "UID": 5,
  "UserName": "Jason",
  "DepartName": "market",
  "CreateAt": "2017-12-2T10:22:23.550331819+08:00"
}
```
```shell
curl http://localhost:8080/service/userinfo?userid=
```

```shell
[
  ...
  
  {
    "UID": 5,
    "UserName": "Jason",
    "DepartName": "market",
    "CreateAt": "2017-12-2T10:22:23.550331819+08:00"
  }
]
```

## efficiency

**data/sql**

```shell
$ ab -n 1000 -c 100 http://localhost:8080/service/userinfo?userid=
```

```shell
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests

**xorm**

```shell
$ ab -n 1000 -c 100 http://localhost:8080/service/userinfo?userid=
```

```shell
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=
Document Length:        444 bytes

Concurrency Level:      100
Time taken for tests:   0.286 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      568000 bytes
HTML transferred:       444000 bytes
Requests per second:    3502.31 [#/sec] (mean)
Time per request:       28.553 [ms] (mean)
Time per request:       0.286 [ms] (mean, across all concurrent requests)
Transfer rate:          1942.69 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.9      0       8
Processing:     1   27  12.9     27      76
Waiting:        1   27  12.9     26      75
Total:          1   28  13.0     28      76
WARNING: The median and mean for the initial connection time are not within a normal deviation
        These results are probably not that reliable.

Percentage of the requests served within a certain time (ms)
  50%     28
  66%     33
  75%     36
  80%     38
  90%     43
  95%     50
  98%     57
  99%     62
 100%     76 (longest request)
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=
Document Length:        559 bytes

Concurrency Level:      100
Time taken for tests:   0.353 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      683000 bytes
HTML transferred:       559000 bytes
Requests per second:    2831.95 [#/sec] (mean)
Time per request:       35.311 [ms] (mean)
Time per request:       0.353 [ms] (mean, across all concurrent requests)
Transfer rate:          1888.89 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.9      0       5
Processing:     1   34  29.5     27     137
Waiting:        1   33  29.5     27     137
Total:          1   34  29.7     27     139
WARNING: The median and mean for the initial connection time are not within a normal deviation
        These results are probably not that reliable.

Percentage of the requests served within a certain time (ms)
  50%     27
  66%     32
  75%     36
  80%     38
  90%     81
  95%    120
  98%    129
  99%    131
 100%    139 (longest request)
```



