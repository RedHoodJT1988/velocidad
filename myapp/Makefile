BINARY_NAME=velocidadApp

build:
	@go mod vendor
	@echo Building Velocidad...
	@go build -o tmp/${BINARY_NAME}
	@echo Velocidad built!

run:
	@echo Starting Velocidad...
	@start /min cmd /c tmp\${BINARY_NAME} &
	@echo Velocidad started!

clean:
	@echo Cleaning...
	@DEL ${BINARY_NAME}
	@go clean
	@echo Cleaned!

test:
	@echo Testing...
	@go test ./...
	@echo Done!

start: run

stop:
	@echo "Stopping the front end..."
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped Velocidad

restart: stop start