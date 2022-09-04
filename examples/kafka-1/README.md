https://developer.confluent.io/quickstart/kafka-docker/

```
cd .\examples\kafka-1\
docker-compose up -d
docker-compose down
```

### Write messages to the topic
```
docker exec broker \
kafka-topics --bootstrap-server broker:9092 \
--create \
--topic quickstart
```

```
docker exec broker `
kafka-topics --bootstrap-server broker:9092 `
--create `
--topic quickstart

```
### Read messages from the topic
```
docker exec --interactive --tty broker \
kafka-console-consumer --bootstrap-server broker:9092 \
                       --topic quickstart \
                       --from-beginning
```

```
docker exec --interactive --tty broker `
kafka-console-consumer --bootstrap-server broker:9092 `
                       --topic quickstart `
                       --from-beginning
```
### Write some more messages

```
docker exec --interactive --tty broker \
kafka-console-producer --bootstrap-server broker:9092 \
                       --topic quickstart
```

```
docker exec --interactive --tty broker `
kafka-console-producer --bootstrap-server broker:9092 `
                       --topic quickstart
```