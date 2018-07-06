# range-value-mutex-broker

Check in and out range values as mutexes.

| Environment Variable              | Example       |
| --------------------------------- | ------------- |
| RANGE_VALUE_BROKER_RANGE          | ["1","2","3"] |
| RANGE_VALUE_BROKER_REDIS_ADDRESS  | :6379         |
| RANGE_VALUE_BROKER_REDIS_PASSWORD | Password123   |

**Example Calls**
```
curl -i http://localhost:8080/checkout
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 06 Jul 2018 18:03:39 GMT
Content-Length: 4

"2"

curl -i http://localhost:8080/checkout
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 06 Jul 2018 18:03:39 GMT
Content-Length: 4

"1"

curl -i http://localhost:8080/checkout
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 06 Jul 2018 18:03:39 GMT
Content-Length: 4

"3"

curl -i -X DELETE http://ccnb:8080/checkout/1
HTTP/1.1 200 OK
Date: Fri, 06 Jul 2018 18:04:52 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8

curl -i http://localhost:8080/checkout
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 06 Jul 2018 18:03:39 GMT
Content-Length: 4

"1"
```