package periodogram_test

// Входные данные для длины окна = 32
func Data32(N int) ([]float64, []float64, []float64) {

	// Окно Блэкмана-Харриса
	w32 := make([]float64, N)
	w32[0] = 0.0234200000000000
	w32[1] = 0.0200796208567949
	w32[2] = 0.0119986516061235
	w32[3] = 0.00453856337218171
	w32[4] = 0.00521782261016186
	w32[5] = 0.0219495235192088
	w32[6] = 0.0611985478246830
	w32[7] = 0.126474585987881
	w32[8] = 0.217470000000000
	w32[9] = 0.329974013305730
	w32[10] = 0.456501360083246
	w32[11] = 0.587419445831711
	w32[12] = 0.712282177389838
	w32[13] = 0.821092467276898
	w32[14] = 0.905301440485947
	w32[15] = 0.958471779849594
	w32[16] = 0.976640000000000
	w32[17] = 0.958471779849594
	w32[18] = 0.905301440485947
	w32[19] = 0.821092467276898
	w32[20] = 0.712282177389838
	w32[21] = 0.587419445831712
	w32[22] = 0.456501360083247
	w32[23] = 0.329974013305730
	w32[24] = 0.217470000000000
	w32[25] = 0.126474585987881
	w32[26] = 0.0611985478246830
	w32[27] = 0.0219495235192089
	w32[28] = 0.00521782261016190
	w32[29] = 0.00453856337218168
	w32[30] = 0.0119986516061234
	w32[31] = 0.0200796208567949

	// Входной сигнал
	x32 := make([]float64, N)
	x32[0] = 1.09754040499941
	x32[1] = 0.804647186185957
	x32[2] = 0.0975404049994096
	x32[3] = -0.609566376187138
	x32[4] = -0.902459595000591
	x32[5] = -0.609566376187138
	x32[6] = 0.0975404049994093
	x32[7] = 0.804647186185957
	x32[8] = 1.09754040499941
	x32[9] = 0.804647186185957
	x32[10] = 0.0975404049994098
	x32[11] = -0.609566376187137
	x32[12] = -0.902459595000591
	x32[13] = -0.609566376187138
	x32[14] = 0.0975404049994091
	x32[15] = 0.804647186185956
	x32[16] = 1.09754040499941
	x32[17] = 0.804647186185957
	x32[18] = 0.0975404049994101
	x32[19] = -0.609566376187137
	x32[20] = -0.902459595000591
	x32[21] = -0.609566376187138
	x32[22] = 0.0975404049994071
	x32[23] = 0.804647186185956
	x32[24] = 1.09754040499941
	x32[25] = 0.804647186185957
	x32[26] = 0.0975404049994085
	x32[27] = -0.609566376187137
	x32[28] = -0.902459595000591
	x32[29] = -0.609566376187138
	x32[30] = 0.0975404049994068
	x32[31] = 0.804647186185956

	// Результат выполнения
	rez32 := make([]float64, N/2+1)
	rez32[0] = 0.197079184668859
	rez32[1] = 0.167207747427803
	rez32[2] = 0.119414954305659
	rez32[3] = 0.572832117787472
	rez32[4] = 1.27047439645810
	rez32[5] = 0.588406913455670
	rez32[6] = 0.0492587604356405
	rez32[7] = 0.000336672614800940
	rez32[8] = 1.71105942505116e-32
	rez32[9] = 2.50134060820484e-32
	rez32[10] = 1.26250044383115e-32
	rez32[11] = 1.52094171115659e-32
	rez32[12] = 2.20536548117705e-31
	rez32[13] = 2.30042433812434e-31
	rez32[14] = 7.80670862679593e-32
	rez32[15] = 1.50014758619937e-32
	rez32[16] = 4.65788399041705e-32

	return x32, w32, rez32
}