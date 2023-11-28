MONGODB_USERNAME=admin
MONGODB_PASSWORD=secret
DB_NAME=merchants

ENV_LOCAL_TEST=\
				MONGODB_HOSTNAME=localhost:27017 \
				MONGODB_USERNAME=$(MONGODB_USERNAME) \
				MONGODB_PASSWORD=$(MONGODB_PASSWORD) \
				RPC_URL=http://localhost:3111 \
				ALLOWED_ORIGINS=http://localhost:3000 \
				DEV_MODE=true \
				STORAGE_HOSTNAME=http:localhost:9000/ \
				PORT=3333

wait-for-mongodb:
		sleep 7

insert-data: wait-for-mongodb
		docker cp ./it/merchants_test_data.json mongodb:/data
		docker-compose exec -T mongodb mongoimport --username=$(MONGODB_USERNAME) --password=$(MONGODB_PASSWORD) --authenticationDatabase=admin --db=$(DB_NAME) --collection=merchants --type=json --file=/data/merchants_test_data.json
		docker-compose exec -T mongodb mongosh --username=$(MONGODB_USERNAME) --password=$(MONGODB_PASSWORD) $(DB_NAME) --authenticationDatabase=admin --eval 'db.merchants.createIndex({ location: "2dsphere" })'

docker.start:
		docker-compose up -d;
		make insert-data

docker.stop:
		docker-compose down --volumes;

test.integration: docker.start
		$(ENV_LOCAL_TEST) \
		go test -tags=integration ./it -v -count=1
		make docker.stop
