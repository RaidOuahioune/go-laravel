# Kafka is a nutsheel

## 1. Topics
A topic is a category or feed name to which records are published.
Topics in Kafka are always multi-subscriber; that is, a topic can have zero, one, or many consumers that subscribe to the data written to it.

## 2. Partitions
Each topic is split into partitions.
Partitions allow Kafka to parallelize processing: each partition can be hosted on a different server, and consumers can read from different partitions concurrently.
Partitions are the unit of scalability and parallelism in Kafka.

## 3. Messages
Messages are the basic units of data in Kafka, consisting of a key, a value, and a timestamp.
Messages are written to and read from partitions.

## 4. Producers
Producers are the applications that publish (write) data to Kafka topics.
Producers send records to Kafka topics, and they can specify the partition to which the record should be sent.

## 5. Consumers
Consumers are the applications that subscribe to (read) topics and process the feed of published messages.
Consumers work as part of a consumer group, where each consumer instance is assigned a subset of the partitions to consume.

## 6. Consumer Groups
A consumer group is a group of consumers that work together to consume a topic.
Kafka guarantees that each partition is consumed by exactly one consumer within a group.
This allows for parallel consumption while ensuring that each message is only processed once.

## 7. Brokers
A Kafka broker is a server that stores data and serves clients (producers and consumers).
Each broker hosts one or more partitions.

## 8. Clusters
A Kafka cluster is a set of brokers working together.
Kafka is designed to be distributed and can scale horizontally by adding more brokers to the cluster.

## 9. ZooKeeper
ZooKeeper is used by Kafka to manage and coordinate the brokers.
It stores metadata about the Kafka cluster, including the topics, partitions, and configuration information.

## 11. Offsets
Each message within a partition has an offset, a unique identifier assigned by Kafka.
Consumers use offsets to keep track of which messages they have consumed.
## 12. Log Compaction
Kafka supports log compaction, a feature that ensures that at least the latest value for each key within a partition is retained.
This helps in cases where you need to preserve the latest state of the data.
## 13. Connectors
Kafka Connect is a tool for scalable and reliable streaming data between Apache Kafka and other systems.
It provides ready-to-use connectors to integrate with various data sources and sinks.
## 14. Streams API
The Kafka Streams API allows for building applications that process data in real-time, transforming, enriching, and aggregating data streams.


# Consumer Groups and Topic Partitions


A consumer group in Kafka can be assigned to one or more topics, and each topic can have multiple partitions. Here's how it works:

Consumer Group Assignment
Consumer Group and Topics:

A consumer group can subscribe to one or more topics.
Each consumer in the group will be assigned one or more partitions of the subscribed topics.
Partitions and Consumers:

Within a consumer group, Kafka ensures that each partition is consumed by only one consumer. This means that no two consumers in the same group will process messages from the same partition concurrently.
This allows for parallel processing while ensuring that each message is only processed once within the group.

## Example Scenarios:
#### Scenario 1: Single Topic with Multiple Partitions
Topic A has 3 partitions: P0, P1, P2.
Consumer Group G has 3 consumers: C1, C2, C3.
In this case:

C1 might be assigned P0,
C2 might be assigned P1,
C3 might be assigned P2.
Each consumer in the group processes data from a different partition.

#### Scenario 2: Multiple Topics with Multiple Partitions
Topic A has 2 partitions: P0, P1.
Topic B has 3 partitions: P0, P1, P2.
Consumer Group G has 3 consumers: C1, C2, C3.
In this case, the assignment could be:

C1 might be assigned P0 of Topic A and P0 of Topic B,
C2 might be assigned P1 of Topic A and P1 of Topic B,
C3 might be assigned P2 of Topic B.
Each consumer in the group processes data from different partitions across multiple topics.

# Key Points:
Consumer Groups can subscribe to multiple topics.
Within a consumer group, each partition of a topic is assigned to only one consumer.
This ensures that all partitions are consumed, and each partition is processed by only one consumer in the group.
This setup allows for horizontal scaling by adding more consumers to the group.
In summary, a consumer group is designed to distribute the partitions of the topics it subscribes to across its consumers. This ensures that each partition is consumed by only one consumer within the group, allowing for efficient and parallel processing of messages.