all: run

run:
	go run main.go ssa_spw.go ssa.go denseFunc.go systems.go signal.go graph.go matlabSave.go pchip.go AcfMed.go savgol.go

pull:
	git pull git@github.com:RB-PRO/ssa.git

push:
	git push git@github.com:RB-PRO/ssa.git

pullW:
	git pull https://github.com/RB-PRO/ssa.git

pushW:
	git push https://github.com/RB-PRO/ssa.git
