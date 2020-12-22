run:
	go run main.go
dev:
	nodemon --exec go run main.go --signal SIGTERM
test:
	docker run --rm williamyeh/wrk -t2 -c5 -d5s --timeout 2s http://192.168.0.247:8080/health
test-static:
	docker run --rm williamyeh/wrk -t2 -c5 -d5s --timeout 2s http://192.168.0.247:8080/app
