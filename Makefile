.PHONY : format install build

run:
	go run ./bin/app/main.go

# live reload using nodemon: npm -g i nodemon
run-nodemon:
	nodemon --exec go run ./bin/app/main.go