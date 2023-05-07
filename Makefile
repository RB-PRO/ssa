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

push-car:
	set GOARCH=arm
	set GOOS=linux
	set CGO_ENABLED=0
	go env GOOS GOARCH CGO_ENABLED
	go build -o main ./cmd/main/main.go
	scp main token root@194.87.107.129:go/KadTG/