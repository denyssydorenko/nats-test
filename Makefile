start:
	docker-compose -f ./packages/docker-compose.yml -p nats-test-service up --build --detach

stop:
	docker-compose -f ./packages/docker-compose.yml -p nats-test-service down --remove-orphans

restart:
	docker-compose -f ./packages/docker-compose.yml -p nats-test-service down --remove-orphans
	docker-compose -f ./packages/docker-compose.yml -p nats-test-service up --build --detach