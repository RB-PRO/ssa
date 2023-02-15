all: run

run:
	go run cmd/main/main.go internal/ssaApp/Classic.go internal/ssaApp/ssa_spw.go pkg/ssa/ssa.go pkg/ssa/AcfMed.go pkg/graph/graph.go pkg/savgol/savgol.go pkg/sdv/sdv.go pkg/pchip/Pchip.go pkg/oss/const.go pkg/oss/dence.go pkg/oss/denseFunc.go pkg/oss/matlabSave.go pkg/oss/signal.go pkg/oss/systems.go

pull:
	git pull git@github.com:RB-PRO/ssa.git

push:
	git push git@github.com:RB-PRO/ssa.git

pullW:
	git pull https://github.com/RB-PRO/ssa.git

pushW:
	git push https://github.com/RB-PRO/ssa.git
