build:
	sudo -E docker-compose -f ./docker_compose/main.yaml -f ./docker_compose/node.yaml -f ./docker_compose/go.yaml -f ./docker_compose/db.yaml up -d --build

down:
	sudo -E docker-compose -f ./docker_compose/main.yaml -f ./docker_compose/node.yaml -f ./docker_compose/go.yaml -f ./docker_compose/db.yaml down

chown:
	sudo chown -R pi:pi .
