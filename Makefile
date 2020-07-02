run:
	./make-proto.sh
	./build.sh
	docker-compose down
	docker-compose up --build
