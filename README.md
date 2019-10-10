# gateway_benchmark
Gateway performance benchmark

TI AM335x: single core cortex A8@800Mhz

Banana pi r64: dual core cortex A53 @1.35Ghz

macBookPro 13 2017: intel core i5 3.1Ghz

i.mx6ul : cortex A7@700Mhz

| benchmark item | TI AM335x | macBookPro 13 2017| Banana pi r64| imx6ul|
|--|--|--|--|--|
| 3 million loop json marshal&unmarshal| 155 seconds| 6 seconds| 49 seconds|227 seconds|

# ab benchmark
## gowebserver run on ti am335

### Concurrency Level:      2
```
➜  Downloads ab -n10000 -c2 -t30 http://10.10.80.15:12345/hello
...
Concurrency Level:      2
Time taken for tests:   30.001 seconds
Complete requests:      15709
Failed requests:        0
Total transferred:      2057879 bytes
HTML transferred:       219926 bytes
Requests per second:    523.62 [#/sec] (mean)
Time per request:       3.820 [ms] (mean)
Time per request:       1.910 [ms] (mean, across all concurrent requests)
Transfer rate:          66.99 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       9
Processing:     1    3   1.9      3      97
Waiting:        1    3   1.9      3      90
Total:          2    4   1.9      4      97

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      5
  75%      5
  80%      5
  90%      6
  95%      6
  98%      6
  99%      6
 100%     97 (longest request)
```
### Concurrency Level:      20
```shell
➜  Downloads ab -n10000 -c20 -t30 http://10.10.80.15:12345/hello
...
Concurrency Level:      20
Time taken for tests:   30.001 seconds
Complete requests:      16713
Failed requests:        0
Total transferred:      2189403 bytes
HTML transferred:       233982 bytes
Requests per second:    557.09 [#/sec] (mean)
Time per request:       35.901 [ms] (mean)
Time per request:       1.795 [ms] (mean, across all concurrent requests)
Transfer rate:          71.27 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       3
Processing:     2   35   7.8     37      70
Waiting:        1   35   7.6     37      47
Total:          2   36   7.8     37      70

Percentage of the requests served within a certain time (ms)
  50%     37
  66%     38
  75%     38
  80%     38
  90%     38
  95%     39
  98%     40
  99%     44
 100%     70 (longest request)
```
### Concurrency Level:      200
```shell
➜  Downloads ab -n10000 -c200 -t30 http://10.10.80.15:12345/hello
...
Concurrency Level:      200
Time taken for tests:   30.001 seconds
Complete requests:      16761
Failed requests:        0
Total transferred:      2195691 bytes
HTML transferred:       234654 bytes
Requests per second:    558.68 [#/sec] (mean)
Time per request:       357.984 [ms] (mean)
Time per request:       1.790 [ms] (mean, across all concurrent requests)
Transfer rate:          71.47 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   46 264.0      0    3821
Processing:     3  306  58.2    308     808
Waiting:        2  305  55.6    306     808
Total:          4  352 272.5    312    4155

Percentage of the requests served within a certain time (ms)
  50%    312
  66%    342
  75%    354
  80%    358
  90%    366
  95%    377
  98%   1577
  99%   1607
 100%   4155 (longest request)
```

# ab benchmark on windows
gowebserver run on thinkpad ,connect to macbookpro with lan
```shell
➜  goweb ab -n10000 -c200 http://10.10.80.15:12345/hello
...
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

## ab testing on banana-pi r64

### Concurrency Level:      2
```
➜  Downloads ab -n10000 -c2 -t30 http://192.168.12.245:12345/hello
...
Concurrency Level:      2
Time taken for tests:   30.000 seconds
Complete requests:      38691
Failed requests:        0
Total transferred:      5068521 bytes
HTML transferred:       541674 bytes
Requests per second:    1289.68 [#/sec] (mean)
Time per request:       1.551 [ms] (mean)
Time per request:       0.775 [ms] (mean, across all concurrent requests)
Transfer rate:          164.99 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.2      1      27
Processing:     0    1   0.3      1      27
Waiting:        0    1   0.3      1      27
Total:          1    1   0.4      1      29

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      2
  75%      2
  80%      2
  90%      2
  95%      2
  98%      2
  99%      2
 100%     29 (longest request)
```


### Concurrency Level:      20
```shell
➜  Downloads ab -n10000 -c20 -t30 http://192.168.12.245:12345/hello
.....
Concurrency Level:      20
Time taken for tests:   8.831 seconds
Complete requests:      50000
Failed requests:        0
Total transferred:      6550000 bytes
HTML transferred:       700000 bytes
Requests per second:    5661.57 [#/sec] (mean)
Time per request:       3.533 [ms] (mean)
Time per request:       0.177 [ms] (mean, across all concurrent requests)
Transfer rate:          724.28 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.5      0      91
Processing:     0    3   2.5      3      93
Waiting:        0    3   2.3      3      93
Total:          1    3   3.1      3     112

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      3
  75%      4
  80%      4
  90%      5
  95%      7
  98%      8
  99%     10
 100%    112 (longest request)
```
### Concurrency Level:      200
```shell
➜  Downloads ab -n10000 -c200 -t30 http://192.168.12.245:12345/hello
......
Concurrency Level:      200
Time taken for tests:   8.311 seconds
Complete requests:      50000
Failed requests:        0
Total transferred:      6550000 bytes
HTML transferred:       700000 bytes
Requests per second:    6016.42 [#/sec] (mean)
Time per request:       33.242 [ms] (mean)
Time per request:       0.166 [ms] (mean, across all concurrent requests)
Transfer rate:          769.68 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    8  80.3      1    2052
Processing:     0   24  18.0     21     499
Waiting:        0   24  17.9     21     499
Total:          1   32  82.4     22    2079

Percentage of the requests served within a certain time (ms)
  50%     22
  66%     25
  75%     28
  80%     31
  90%     44
  95%     59
  98%     82
  99%    172
 100%   2079 (longest request)
```

## ab benchmark on imx6ul
### Concurrency Level:      2
```
➜  Downloads ab -n10000 -c2 -t30 http://192.168.12.84:12345/hello 

Concurrency Level:      2
Time taken for tests:   30.002 seconds
Complete requests:      8156
Failed requests:        0
Total transferred:      1068436 bytes
HTML transferred:       114184 bytes
Requests per second:    271.85 [#/sec] (mean)
Time per request:       7.357 [ms] (mean)
Time per request:       3.678 [ms] (mean, across all concurrent requests)
Transfer rate:          34.78 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.1      1       3
Processing:     2    7   3.4      7      32
Waiting:        1    4   2.2      4      29
Total:          3    7   3.4      7      33

Percentage of the requests served within a certain time (ms)
  50%      7
  66%      8
  75%      9
  80%     10
  90%     11
  95%     13
  98%     16
  99%     17
 100%     33 (longest request)
 ```
 ### Concurrency Level:      20
 ```shell
 ➜  Downloads ab -n10000 -c20 -t30 http://192.168.12.84:12345/hello
....

Concurrency Level:      20
Time taken for tests:   30.002 seconds
Complete requests:      10609
Failed requests:        0
Total transferred:      1389779 bytes
HTML transferred:       148526 bytes
Requests per second:    353.61 [#/sec] (mean)
Time per request:       56.559 [ms] (mean)
Time per request:       2.828 [ms] (mean, across all concurrent requests)
Transfer rate:          45.24 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.2      1       7
Processing:     2   56  12.0     58     107
Waiting:        2   55  11.9     58      69
Total:          3   56  12.0     59     108

Percentage of the requests served within a certain time (ms)
  50%     59
  66%     59
  75%     59
  80%     59
  90%     60
  95%     60
  98%     62
  99%     63
 100%    108 (longest request)
 ```
### Concurrency Level:      200
```shell
➜  Downloads  ab -n10000 -c200 -t30 http://192.168.12.84:12345/hello
...

Concurrency Level:      200
Time taken for tests:   30.002 seconds
Complete requests:      10713
Failed requests:        0
Total transferred:      1403796 bytes
HTML transferred:       150024 bytes
Requests per second:    357.08 [#/sec] (mean)
Time per request:       560.103 [ms] (mean)
Time per request:       2.801 [ms] (mean, across all concurrent requests)
Transfer rate:          45.69 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  105 362.4      1    2564
Processing:     3  450  86.8    442    1051
Waiting:        2  447  82.9    441     670
Total:          7  555 375.3    456    3468

Percentage of the requests served within a certain time (ms)
  50%    456
  66%    501
  75%    542
  80%    556
  90%    576
  95%   1702
  98%   1774
  99%   1807
 100%   3468 (longest request)
```
