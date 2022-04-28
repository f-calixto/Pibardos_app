build:
	sudo docker-compose -f ./docker_compose/main.yaml ./docker_compose/node.yaml ./docker_compose/go.yaml up -d --build

chown:
	sudo chown -R pi:pi .
