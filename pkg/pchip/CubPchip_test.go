package pchip

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestPchip(t *testing.T) {
	win := 1024
	lag := int(math.Floor(float64(win) / 10.0))
	lgl := make([]float64, lag)
	for m := 0; m < len(lgl); m++ {
		lgl[m] = float64(m + 1)
	}

	//PhaAcfNrm, _ := oss.Make_singnal_xn("PhaAcfNrm") // Загрузить сигнал из файла PhaAcfNrm.xlsx
	var PhaAcfNrm []float64
	{
		PhaAcfNrm = append(PhaAcfNrm, 0)
		PhaAcfNrm = append(PhaAcfNrm, 0.441329015131783)
		PhaAcfNrm = append(PhaAcfNrm, 0.881585847974732)
		PhaAcfNrm = append(PhaAcfNrm, 1.32302287192679)
		PhaAcfNrm = append(PhaAcfNrm, 1.76117989663173)
		PhaAcfNrm = append(PhaAcfNrm, 2.20339883437171)
		PhaAcfNrm = append(PhaAcfNrm, 2.64685451388392)
		PhaAcfNrm = append(PhaAcfNrm, 3.14159265358979)
		PhaAcfNrm = append(PhaAcfNrm, 2.76239783429631)
		PhaAcfNrm = append(PhaAcfNrm, 2.31772584509265)
		PhaAcfNrm = append(PhaAcfNrm, 1.87483101744512)
		PhaAcfNrm = append(PhaAcfNrm, 1.43748806339424)
		PhaAcfNrm = append(PhaAcfNrm, 0.997587680233786)
		PhaAcfNrm = append(PhaAcfNrm, 0.548849248835465)
		PhaAcfNrm = append(PhaAcfNrm, 0)
		PhaAcfNrm = append(PhaAcfNrm, 0.303015181377973)
		PhaAcfNrm = append(PhaAcfNrm, 0.761759396375239)
		PhaAcfNrm = append(PhaAcfNrm, 1.20495096330935)
		PhaAcfNrm = append(PhaAcfNrm, 1.64533337140907)
		PhaAcfNrm = append(PhaAcfNrm, 2.08692419697565)
		PhaAcfNrm = append(PhaAcfNrm, 2.53912659214520)
		PhaAcfNrm = append(PhaAcfNrm, 3.14159265358979)
		PhaAcfNrm = append(PhaAcfNrm, 2.94880108846227)
		PhaAcfNrm = append(PhaAcfNrm, 2.45026790665149)
		PhaAcfNrm = append(PhaAcfNrm, 2.00451697188977)
		PhaAcfNrm = append(PhaAcfNrm, 1.56190646337479)
		PhaAcfNrm = append(PhaAcfNrm, 1.11862463877279)
		PhaAcfNrm = append(PhaAcfNrm, 0.667288477890010)
		PhaAcfNrm = append(PhaAcfNrm, 0.150301081472016)
		PhaAcfNrm = append(PhaAcfNrm, 0)
		PhaAcfNrm = append(PhaAcfNrm, 0.623525071122026)
		PhaAcfNrm = append(PhaAcfNrm, 1.06933770720834)
		PhaAcfNrm = append(PhaAcfNrm, 1.51054677361522)
		PhaAcfNrm = append(PhaAcfNrm, 1.94966463325690)
		PhaAcfNrm = append(PhaAcfNrm, 2.38652342833179)
		PhaAcfNrm = append(PhaAcfNrm, 2.84913734149714)
		PhaAcfNrm = append(PhaAcfNrm, 3.14159265358979)
		PhaAcfNrm = append(PhaAcfNrm, 2.57569402422969)
		PhaAcfNrm = append(PhaAcfNrm, 2.14256889000444)
		PhaAcfNrm = append(PhaAcfNrm, 1.70871939770792)
		PhaAcfNrm = append(PhaAcfNrm, 1.27552295937205)
		PhaAcfNrm = append(PhaAcfNrm, 0.845097548587511)
		PhaAcfNrm = append(PhaAcfNrm, 0.398999513344493)
		PhaAcfNrm = append(PhaAcfNrm, 0)
		PhaAcfNrm = append(PhaAcfNrm, 0.498106107435887)
		PhaAcfNrm = append(PhaAcfNrm, 0.921351240400956)
		PhaAcfNrm = append(PhaAcfNrm, 1.35040228649484)
		PhaAcfNrm = append(PhaAcfNrm, 1.78069513171623)
		PhaAcfNrm = append(PhaAcfNrm, 2.20960324796203)
		PhaAcfNrm = append(PhaAcfNrm, 2.64937161060030)
		PhaAcfNrm = append(PhaAcfNrm, 3.14159265358979)
		PhaAcfNrm = append(PhaAcfNrm, 2.73352890451767)
		PhaAcfNrm = append(PhaAcfNrm, 2.31530070142226)
		PhaAcfNrm = append(PhaAcfNrm, 1.88714086572228)
		PhaAcfNrm = append(PhaAcfNrm, 1.45640104361232)
		PhaAcfNrm = append(PhaAcfNrm, 1.02510734322105)
		PhaAcfNrm = append(PhaAcfNrm, 0.578044297977329)
		PhaAcfNrm = append(PhaAcfNrm, 0)
		PhaAcfNrm = append(PhaAcfNrm, 0.286334971136148)
		PhaAcfNrm = append(PhaAcfNrm, 0.713955665817446)
		PhaAcfNrm = append(PhaAcfNrm, 1.14706101630374)
		PhaAcfNrm = append(PhaAcfNrm, 1.58293512968344)
		PhaAcfNrm = append(PhaAcfNrm, 2.02158604970230)
		PhaAcfNrm = append(PhaAcfNrm, 2.47557107461441)
		PhaAcfNrm = append(PhaAcfNrm, 3.04900142042372)
		PhaAcfNrm = append(PhaAcfNrm, 3.14159265358979)
		PhaAcfNrm = append(PhaAcfNrm, 2.55628394660712)
		PhaAcfNrm = append(PhaAcfNrm, 2.11126144806189)
		PhaAcfNrm = append(PhaAcfNrm, 1.67320606147054)
		PhaAcfNrm = append(PhaAcfNrm, 1.23780179086826)
		PhaAcfNrm = append(PhaAcfNrm, 0.796434934579233)
		PhaAcfNrm = append(PhaAcfNrm, 0.319676155262032)
		PhaAcfNrm = append(PhaAcfNrm, 0)
		PhaAcfNrm = append(PhaAcfNrm, 0.487613665278124)
		PhaAcfNrm = append(PhaAcfNrm, 0.920823489184584)
		PhaAcfNrm = append(PhaAcfNrm, 1.35301784200574)
		PhaAcfNrm = append(PhaAcfNrm, 1.78478642093141)
		PhaAcfNrm = append(PhaAcfNrm, 2.21972452880737)
		PhaAcfNrm = append(PhaAcfNrm, 2.68106745292672)
		PhaAcfNrm = append(PhaAcfNrm, 3.14159265358979)
		PhaAcfNrm = append(PhaAcfNrm, 2.78090939571309)
		PhaAcfNrm = append(PhaAcfNrm, 2.34610056336240)
		PhaAcfNrm = append(PhaAcfNrm, 1.91183852142865)
		PhaAcfNrm = append(PhaAcfNrm, 1.47615706587846)
		PhaAcfNrm = append(PhaAcfNrm, 1.03731052966630)
		PhaAcfNrm = append(PhaAcfNrm, 0.562200792749940)
		PhaAcfNrm = append(PhaAcfNrm, 0)
		PhaAcfNrm = append(PhaAcfNrm, 0.174745316427344)
		PhaAcfNrm = append(PhaAcfNrm, 0.663190108905505)
		PhaAcfNrm = append(PhaAcfNrm, 1.10823392805498)
		PhaAcfNrm = append(PhaAcfNrm, 1.54777608753245)
		PhaAcfNrm = append(PhaAcfNrm, 1.98618419707169)
		PhaAcfNrm = append(PhaAcfNrm, 2.44628077091687)
		PhaAcfNrm = append(PhaAcfNrm, 2.88190919718324)
		PhaAcfNrm = append(PhaAcfNrm, 3.14159265358979)
		PhaAcfNrm = append(PhaAcfNrm, 2.58567404745506)
		PhaAcfNrm = append(PhaAcfNrm, 2.14403359064105)
		PhaAcfNrm = append(PhaAcfNrm, 1.70709465280260)
		PhaAcfNrm = append(PhaAcfNrm, 1.27147684183685)
		PhaAcfNrm = append(PhaAcfNrm, 0.819643845077553)
		PhaAcfNrm = append(PhaAcfNrm, 0.400997148971444)
		PhaAcfNrm = append(PhaAcfNrm, 0)
	}
	_, pCoef, coefs := Pchip(PhaAcfNrm,
		lgl,
		lgl,
		len(PhaAcfNrm), len(lgl))

	t.Log(len(pCoef), pCoef[0], pCoef[1], pCoef[1], pCoef[3], pCoef[4])

	t.Log(coefs.a[0], coefs.a[1], coefs.a[2], coefs.a[3])
	t.Log(coefs.b[0], coefs.b[1], coefs.b[2], coefs.b[3])
	t.Log(coefs.c[0], coefs.c[1], coefs.c[2], coefs.c[3])
	t.Log(coefs.d[0], coefs.d[1], coefs.d[2], coefs.d[3])

	safeToXlsx(coefs.a, "coefs.a")
	safeToXlsx(coefs.b, "coefs.b")
	safeToXlsx(coefs.c, "coefs.c")
	safeToXlsx(coefs.d, "coefs.d")

	//oss.SafeToXlsxDualArray(pCoef, "pCoef")
	/*
		if result != "Foo" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
		}
	*/
}
func safeToXlsx(sig []float64, name string) {
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	for ind, val := range sig {
		file_graph.SetCellValue("main", "A"+strconv.Itoa(ind+1), val)
	}
	if err := file_graph.SaveAs(name + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
