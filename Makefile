all: run

run:
	go run main.go systems.go
	
pull:
	git pull git@github.com:RB-PRO/ssa.git

push:
	git push git@github.com:RB-PRO/ssa.git
