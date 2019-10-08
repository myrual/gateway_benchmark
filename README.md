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
➜  goweb ab -n10000 -c200 http://192.168.200.86:12345/hello 
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 192.168.200.86 (be patient)
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
Server Hostname:        192.168.200.86
Server Port:            12345

Document Path:          /hello
Document Length:        14 bytes

Concurrency Level:      200
Time taken for tests:   16.500 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1310000 bytes
HTML transferred:       140000 bytes
Requests per second:    606.05 [#/sec] (mean)
Time per request:       330.006 [ms] (mean)
Time per request:       1.650 [ms] (mean, across all concurrent requests)
Transfer rate:          77.53 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:       11  161 359.2     53    3490
Processing:    16  113 357.1     53    6801
Waiting:       10  108 351.3     52    6799
Total:         29  275 566.9    109    7863

Percentage of the requests served within a certain time (ms)
  50%    109
  66%    119
  75%    128
  80%    150
  90%    515
  95%   1247
  98%   1660
  99%   2619
 100%   7863 (longest request)
```
