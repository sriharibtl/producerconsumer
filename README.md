# producerconsumer

1. Start the zookeeper instance:
  #docker run -d --name zookeeper -p 2181:2181 jplock/zookeeper

2. Start kafka instance:
  #docker run -d --name kafka -p 7203:7203 -p 9092:9092 -e KAFKA_ADVERTISED_HOST_NAME=<HOST_IP> -e ZOOKEEPER_IP=<HOST_IP> ches/kafka

3. Create kafka topic:
  #docker run --rm ches/kafka kafka-topics.sh --create --topic parking --replication-factor 1 --partitions 1 --zookeeper <HOST_IP>:2181

4. List kafka topic:
  #docker run --rm ches/kafka kafka-topics.sh --list --zookeeper <HOST_IP>:2181

5. Start couchbase DB container:
  #docker run -d --name db -p 8091-8096:8091-8096 -p 11210-11211:11210-11211 couchbase:community

  note: https://docs.couchbase.com/server/current/install/getting-started-docker.html

6. Setup couchbase cluster and create a bucket with below credentials
    username: Administrator
    password: password
    bucket name: ut


7. Move into producer directory and build docker image
  #docker build -t producer:1.1 .

8. Move into consumer directory and build docker image
  #docker build -t consumer:1.1 .

9. Run consumer docker container
  #docker run -d  --rm --name consumer1 -e KAFKA_HOST="<HOST_IP>:9092" -e KAFKA_TOPIC="parking" -e DB_IP="<HOST_IP>" -p 9981:9991 consumer:1.3

10. Run producer docker container
  #docker run -d --rm --name producer1 -e KAFKA_HOST="<HOST_IP>:9092" -e KAFKA_TOPIC="parking" -p 9990:9990 producer:1.5

11. Reach the producer endpoint to start producer
  #curl http://<HOST_IP>:9990/incrementcounter

12. Reach the consumer endpoint to fetch the counter
  #curl http://<HOST_IP>:9981/counter



