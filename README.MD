# Testando o framework fiber

 A proposta deste framework é ser tão facil quanto o express.js e muito mais rapido

## Teste de Performance

```shell
# make test
docker run --rm williamyeh/wrk -t2 -c5 -d5s --timeout 2s http://192.168.0.247:8080/health
Running 5s test @ http://192.168.0.247:8080/health
  2 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.42ms   11.03ms 103.03ms   90.03%
    Req/Sec     8.43k     3.72k   17.72k    62.00%
  83871 requests in 5.00s, 9.84MB read
Requests/sec:  16766.67
Transfer/sec:      1.97MB
```

```shell
# make test-static 
docker run --rm williamyeh/wrk -t2 -c5 -d5s --timeout 2s http://192.168.0.247:8080/app
Running 5s test @ http://192.168.0.247:8080/app
  2 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   172.36us  839.14us  23.05ms   97.53%
    Req/Sec    28.98k     5.06k   39.85k    85.15%
  290993 requests in 5.10s, 108.51MB read
Requests/sec:  57059.08
Transfer/sec:     21.28MB
```

## Referencias

 - https://medium.com/javascript-in-plain-english/golang-fiber-426f52185fc4

 - https://gobyexample.com/json

 - https://medium.com/@_oleksii_/how-to-benchmark-http-latency-with-wrk-a-z-guide-e5b185bd4cdc
