docker network create testnetwork

#docker run -d -p 9042:9042 --net testnetwork --name cassandra -h cassandra cassandra:3.11.14
docker run -d -p 9042:9042 --net testnetwork --name cassandra -h cassandra cassandra:4.1.1

docker run -d -p 8080:8080 --net testnetwork --name cassandratest -h cassandratest cassandratest

#docker inspect --format='{{ .NetworkSettings.IPAddress }}' cassandra

#docker run -it --link cassandra --rm cassandra:latest \bash -c 'exec cqlsh 172.17.0.2'

#CREATE KEYSPACE tutorial with REPLICATION = {'class':'SimpleStrategy','replication_factor':1};