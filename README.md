# gateway_benchmark
Gateway performance benchmark

TI AM335x: single core 800Mhz cortex A8

macBookPro 13 2017: intel core i5 3.1G

| benchmark item | TI AM335x | macBookPro 13 2017|
|--|--|--|
| 3 million loop json marshal&unmarshal| 155 seconds| 6 seconds| 
| http server ab -n10000 -c200 | 17seconds | x | 

ab benchmark on ti am335
```shell
➜  goweb ab -n10000 -c200 http://10.10.80.15:12345/hello
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 10.10.80.15 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        10.10.80.15
Server Port:            12345

Document Path:          /hello
Document Length:        14 bytes

Concurrency Level:      200
Time taken for tests:   17.885 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1310000 bytes
HTML transferred:       140000 bytes
Requests per second:    559.13 [#/sec] (mean)
Time per request:       357.699 [ms] (mean)
Time per request:       1.788 [ms] (mean, across all concurrent requests)
Transfer rate:          71.53 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    7  27.1      0     136
Processing:     3  348  54.2    359     638
Waiting:        2  346  50.5    358     619
Total:          3  355  55.7    361     715

Percentage of the requests served within a certain time (ms)
  50%    361
  66%    363
  75%    365
  80%    366
  90%    380
  95%    425
  98%    472
  99%    499
 100%    715 (longest request)
```


ab benchmark on windows
```shell
➜  goweb ab -n10000 -c200 http://10.10.80.15:12345/hello
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 10.10.80.15 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        10.10.80.15
Server Port:            12345

Document Path:          /hello
Document Length:        14 bytes

Concurrency Level:      200
Time taken for tests:   1.242 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1310000 bytes
HTML transferred:       140000 bytes
Requests per second:    8048.98 [#/sec] (mean)
Time per request:       24.848 [ms] (mean)
Time per request:       0.124 [ms] (mean, across all concurrent requests)
Transfer rate:          1029.70 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1    8  63.5      3    1036
Processing:     1   10  23.7      6     168
Waiting:        1   10  23.7      5     167
Total:          4   18  67.6      9    1046

Percentage of the requests served within a certain time (ms)
  50%      9
  66%     12
  75%     14
  80%     15
  90%     18
  95%     23
  98%    166
  99%    169
 100%   1046 (longest request)
```
