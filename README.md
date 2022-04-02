This runs the sample program to demonstrate the `kafka server: Request was for a topic or partition that does not exist on this broker.`
when using Kafka in Kraft mode.

### Run the sample program with Kafka in Kraft mode
```
docker compose run test
docker compose down # To tear down everything
```

### Run the sample program with Kafka in Zookeeper mode
```
docker compose -f docker-compose-zk.yml run test
docker compose down # To tear down everything
```
