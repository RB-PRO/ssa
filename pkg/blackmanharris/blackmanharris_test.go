package blackmanharris_test

import (
	"math"
	"testing"

	"github.com/RB-PRO/ssa/pkg/blackmanharris"
)

func TestBlackmanharris(t *testing.T) {
	N := 32
	blackman := blackmanharris.Blackmanharris(N, blackmanharris.Koef4_92db)

	// Результат
	MatLabBlackman := make([]float64, N)
	MatLabBlackman[0] = 0.0234200000000000
	MatLabBlackman[1] = 0.0200796208567949
	MatLabBlackman[2] = 0.0119986516061235
	MatLabBlackman[3] = 0.00453856337218171
	MatLabBlackman[4] = 0.00521782261016186
	MatLabBlackman[5] = 0.0219495235192088
	MatLabBlackman[6] = 0.0611985478246830
	MatLabBlackman[7] = 0.126474585987881
	MatLabBlackman[8] = 0.217470000000000
	MatLabBlackman[9] = 0.329974013305730
	MatLabBlackman[10] = 0.456501360083246
	MatLabBlackman[11] = 0.587419445831711
	MatLabBlackman[12] = 0.712282177389838
	MatLabBlackman[13] = 0.821092467276898
	MatLabBlackman[14] = 0.905301440485947
	MatLabBlackman[15] = 0.958471779849594
	MatLabBlackman[16] = 0.976640000000000
	MatLabBlackman[17] = 0.958471779849594
	MatLabBlackman[18] = 0.905301440485947
	MatLabBlackman[19] = 0.821092467276898
	MatLabBlackman[20] = 0.712282177389838
	MatLabBlackman[21] = 0.587419445831712
	MatLabBlackman[22] = 0.456501360083247
	MatLabBlackman[23] = 0.329974013305730
	MatLabBlackman[24] = 0.217470000000000
	MatLabBlackman[25] = 0.126474585987881
	MatLabBlackman[26] = 0.0611985478246830
	MatLabBlackman[27] = 0.0219495235192089
	MatLabBlackman[28] = 0.00521782261016190
	MatLabBlackman[29] = 0.00453856337218168
	MatLabBlackman[30] = 0.0119986516061234
	MatLabBlackman[31] = 0.0200796208567949

	// Циклом идём по всему массиву и сравниваем данные
	for index := range blackman {
		if int(math.Abs(blackman[index]-MatLabBlackman[index])) > 3 {
			t.Errorf("Blackmanharris: Элемент с индексом %v не соответствует рассчитанному. Должно было быть %v, а получено %v.",
				index, int(MatLabBlackman[index]*1000000), int(blackman[index]*1000000))
		}
	}

}
