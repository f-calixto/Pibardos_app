build:
	sudo docker-compose -f ./docker_compose/main.yaml -f ./docker_compose/node.yaml -f ./docker_compose/go.yaml up -d --build

chown:
	sudo chown -R pi:pi .
