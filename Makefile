all: run

run:
	go run cmd/main/main.go

pull:
	git pull git@github.com:RB-PRO/ssa.git

push:
	git push git@github.com:RB-PRO/ssa.git

pullW:
	git pull https://github.com/RB-PRO/ssa.git

pushW:
	git push https://github.com/RB-PRO/ssa.git
