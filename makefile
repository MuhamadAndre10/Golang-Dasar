build_app:
	@echo Building app
	@go build -o ./dist/main.exe ./basic

start_app:
	@echo Starting app
	@env ./dist/main.exe &

start: build_app start_app

stop:
	@echo stopping app
	@taskkill /IM main.exe /F

clean:
	@echo Cleaning...
	@echo y | DEL /S dist
	@go clean
	@echo Cleaned and deleted binaries