# Mockafka 

Mock Kafka Server, useful for testing and experimenting with Kafka

Basic usage
```
./mockafka start
```

You can specify the number of brokers to run (default:3) 

```
./mockafka start --brokers 5
```

## Cloning and Building

```
git clone github.com/ugol/mockafka
cd mockafka
go build 
```

# Thanks

At the moment mockafka it's just a wrapper on wonderful librdkafka library functionality
