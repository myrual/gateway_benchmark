# gateway_benchmark
Gateway performance benchmark

TI AM335x: single core 800Mhz cortex A8

macBookPro 13 2017: intel core i5 3.1G

| benchmark item | TI AM335x | macBookPro 13 2017|
|--|--|--|
| 3 million loop json marshal&unmarshal| 155 seconds| 6 seconds| 
| http server ab -n10000 -c200 | 17seconds | x | 

