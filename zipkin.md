
- Github
https://github.com/openzipkin/zipkin.git

- Build zipkin server jar

```

# Get Latest jar
$ wget -O zipkin.jar 'https://search.maven.org/remote_content?g=io.zipkin.java&a=zipkin-server&v=LATEST&c=exec'
Or
# Build the server and also make its dependencies
$ ./mvnw -DskipTests --also-make -pl zipkin-server clean install

# Run the server
$ java -jar ./zipkin-server/target/zipkin-server-*exec.jar

# Get kafka 0.10 collector plugin
wget -O zipkin-autoconfigure-collector-kafka10-module.jar 'https://search.maven.org/remote_content?g=io.zipkin.java&a=zipkin-autoconfigure-collector-kafka10&v=LATEST&c=module'
```


- zipkin settings

```
* `ES_HOSTS`: A comma separated list of elasticsearch base urls to connect to ex. http://host:9200.
              Defaults to "http://localhost:9200".

              If the http URL is an AWS-hosted elasticsearch installation (e.g.
              https://search-domain-xyzzy.us-west-2.es.amazonaws.com) then Zipkin will attempt to
              use the default AWS credential provider (env variables, system properties, config
              files, or ec2 profiles) to sign outbound requests to the cluster.
* `ES_PIPELINE`: Only valid when the destination is Elasticsearch 5.x. Indicates the ingest
                 pipeline used before spans are indexed. No default.
* `ES_MAX_REQUESTS`: Only valid when the transport is http. Sets maximum in-flight requests from
                     this process to any Elasticsearch host. Defaults to 64.
* `ES_AWS_DOMAIN`: The name of the AWS-hosted elasticsearch domain to use. Supercedes any set
                   `ES_HOSTS`. Triggers the same request signing behavior as with `ES_HOSTS`, but
                   requires the additional IAM permission to describe the given domain.
* `ES_AWS_REGION`: An optional override to the default region lookup to search for the domain
                   given in `ES_AWS_DOMAIN`. Ignored if only `ES_HOSTS` is present.
* `ES_INDEX`: The index prefix to use when generating daily index names. Defaults to zipkin.
* `ES_DATE_SEPARATOR`: The date separator to use when generating daily index names. Defaults to '-'.
* `ES_INDEX_SHARDS`: The number of shards to split the index into. Each shard and its replicas
                     are assigned to a machine in the cluster. Increasing the number of shards
                     and machines in the cluster will improve read and write performance. Number
                     of shards cannot be changed for existing indices, but new daily indices
                     will pick up changes to the setting. Defaults to 5.
* `ES_INDEX_REPLICAS`: The number of replica copies of each shard in the index. Each shard and
                       its replicas are assigned to a machine in the cluster. Increasing the
                       number of replicas and machines in the cluster will improve read
                       performance, but not write performance. Number of replicas can be changed
                       for existing indices. Defaults to 1. It is highly discouraged to set this
                       to 0 as it would mean a machine failure results in data loss.
* `ES_USERNAME` and `ES_PASSWORD`: Elasticsearch basic authentication, which defaults to empty string.
                                   Use when X-Pack security (formerly Shield) is in place.
* `ES_HTTP_LOGGING`: When set, controls the volume of HTTP logging of the Elasticsearch Api.
                     Options are BASIC, HEADERS, BODY



* `KAFKA_TOPIC`: Topic zipkin spans will be consumed from. Defaults to "zipkin". When Kafka 0.10 is in use, multiple topics may be specified if comma delimited.
* `KAFKA_STREAMS`: Count of threads/streams consuming the topic. Defaults to 1
```


- Log

# Level to print screen
--logging.level.zipkin=DEBUG



- Cmd
STORAGE_TYPE=elasticsearch \
ES_HOSTS=http://10.128.112.102:9200 \
ES_MAX_REQUESTS=150 \
ZIPKIN_LOG_LEVEL=DEBUG \
JAVA_OPTS=-Xms1024m -Xmx1024m \
KAFKA_BOOTSTRAP_SERVERS=10.128.112.186:9092,10.128.112.184:9092 \
KAFKA_TOPIC=zipkin \
KAFKA_STREAMS=10 \
java -Dloader.path='zipkin-autoconfigure-collector-kafka10-module.jar,zipkin-autoconfigure-collector-kafka10-module.jar!/lib'     -Dspring.profiles.active=kafka     -Dzipkin.ui.query-limit=10     -cp ./zipkin.jar     org.springframework.boot.loader.PropertiesLauncher



KAFKA_BOOTSTRAP_SERVERS=10.10.42.14:9092,10.10.42.15:9092,10.10.42.16:9092,10.10.42.17:9092,10.10.42.18:9092,10.10.42.19:9092 \

-- Install Java
sudo add-apt-repository ppa:webupd8team/java
sudo apt-get update
sudo apt-get install oracle-java8-installer

bootstrap-server 10.128.112.186:9092,10.128.112.184:9092