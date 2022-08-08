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

At the moment mockafka is just a wrapper on wonderful [librdkafka](https://github.com/edenhill/librdkafka) library functionality from [Magnus Edenhill](https://github.com/edenhill)

As stated in:
https://docs.confluent.io/platform/current/clients/confluent-kafka-go/index.html#MockCluster

_"NewMockCluster provides a mock Kafka cluster with a configurable number of brokers that support a reasonable subset of Kafka protocol operations, error injection, etc.
Mock clusters provide localhost listeners that can be used as the bootstrap servers by multiple Kafka client instances.
Currently supported functionality: - Producer - Idempotent Producer - Transactional Producer - Low-level consumer - High-level balanced consumer groups with offset commits - Topic Metadata and auto creation"_


