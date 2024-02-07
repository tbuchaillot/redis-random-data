# redis-random-data
Generates Random data for Redis.

## What it does
Generates random data for Redis. It's useful for testing purposes.

## How it works
It connects to a Redis instance and starts to generate random data.
You can specify the number of keys to generate with `-c` flag. 10 by default.
You can specify the prefix of the keys with `-p` flag. "record" by default.

You can specify your redis connection data with environment variables:
- REDIS_HOST
- REDIS_PORT
- REDIS_DB
- REDIS_USERNAME
- REDIS_PASSWORD
- REDIS_ADDRS
- REDIS_ENABLECLUSTER

## How to run it
1. Go build
```bash
    go build
```
2. Run
```bash
    ./redis-random-data
```
