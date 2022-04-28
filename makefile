build:
	sudo docker-compose -f ./docker-compose/main.yaml ./docker-compose/node.yaml ./docker-compose/go.yaml up -d --build

chown:
	sudo chown -R pi:pi .
