# Setup 

Golang is required to launch the program
After go is installed, the usage is:

```
Usage: go run ds_benchmark.go <seed> <data length> <max element (data range)>, <order(-1, 0, 1, where 0 - unordered, -1 - Desc, 1 - Asc)>
```

Example:

```
go run ds_benchmark.go 1 1000 10000 0
```

I didn't create 100 datasets per se, but the one can easily play with input parameters and check parameters for different cases.

# Results.

Here are some results:

Seed = 1, size = 1000000, maxNum = 1000000, order = 0
Insert min    2.444µs
Find   min    647ns
Delete min    785ns
Insert max    1.485µs
Find   max    381ns
Delete max    650ns
Insert median 1.833µs
Find   median 553ns
Delete median 711ns
Count sort performance: 30.448631ms

Seed = 1, size = 1000000, maxNum = 1000000, order = -1
Insert min    994ns
Find   min    729ns
Delete min    1.1µs
Insert max    3.731µs
Find   max    371ns
Delete max    652ns
Insert median 3.2µs
Find   median 439ns
Delete median 642ns
Count sort performance: 25.599926ms

Seed = 1, size = 1000000, maxNum = 1000000, order = 1
Insert min    3.333µs
Find   min    448ns
Delete min    477ns
Insert max    1.336µs
Find   max    211ns
Delete max    436ns
Insert median 3.497µs
Find   median 268ns
Delete median 501ns
Count sort performance: 18.622782ms


So appartenly, for BBST the insert operation always take longer than others, and delete takes slightly longer than find - 
this is quite obvious, since both insert and delete require additional rotating (usually). 
Although, it's not quite clear why delete take less time than insert - maybe it's due to internal memory allocations required by inseert.

As per count sort - it's quite fast overall, as expected, and it's pretty obvious it doesn't perform well when tha data range is to big - like from 0 to 1000000000000, even the length is up to 10 elements.
