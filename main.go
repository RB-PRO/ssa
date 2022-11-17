package main

import (
	"errors"
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

const rcond = 1e-15

func main() {
	fmp, fmpN := make_singnal_xn("fmp")
	pw, pwN := make_singnal_xn("pw")
	fmt.Println("fmp:", fmpN, "//", "pw", pwN)
	safeToXlsx(fmp, "fmp") // Сохранить данные в xlsx
	safeToXlsx(pw, "pw")   // Сохранить данные в xlsx

	ssa_spw(pw, fmp)

	//C, LBD, RC := SSA(N, L, sig, 2)
	//safeToXlsxM(C, "C")
	//safeToXlsx(LBD, "LBD")
	//safeToXlsxM(RC, "RC")

	//makeGraph2(N, "png"+OpSystemFilder+"sig.png")
}

func ssa_spw(pw, fmp []float64) {
	// Сегменты отсчётов pw
	N := len(pw) // Количество отсчетов pw
	win := 1024
	res := N - win*int(math.Floor(float64(N)/float64(win)))
	nPart := 20 // Количество долей res
	res = int(math.Floor(float64(res) / float64(nPart)))
	overlap := (float64(win) - float64(res)) / float64(win)
	S := 1
	Imin := 1
	Imax := win

	var ns []float64
	for Imax <= N {
		ns = append(ns, float64(S)) // номер текущего сегмента pw
		Imin = Imin + res
		Imax = Imax + res
		S++
	}
	S--                    // кол-во перекрывающихся сегментов pw в пределах N
	NSF := win + res*(S-1) // номер финального отсчета финального сегмента <= N

	//spw := dinamicArray(win, S)
	spw := mat.NewDense(win, S, nil)
	/*
		spw := make([][]float64, win)
		for index := range spw {
			spw[index] = make([]float64, S)
		}
	*/
	fmt.Println("Размеры spw:", win, S)
	for j := 0; j < S; j++ {
		for i := 0; i < win; i++ {
			k := (j) * res
			//spw[i][j] = pw[k+i] // текущий сегмент pw длинною win
			spw.Set(i, j, pw[k+i])
		}
	}

	// Set general parameters
	cad := 30              // 30 кадров/сек
	dt := 1 / float64(cad) // интервал дискретизации времени, сек
	tim := make([]float64, N)
	for index := 1; index < N; index++ {
		tim[index] = tim[index-1] + dt
	}

	ns2 := make([]int, S)
	for index := range ns2 {
		ns2[index] = (index + 1)
	}

	L := make([]float64, S)
	for index := range L { // цикл по сегментам pw
		L[index] = math.Floor(float64(cad) / fmp[index]) // кол-во отсчетов основного тона pw
	}

	K := 5
	M := int(float64(K) * max(L)) // параметр вложения в траекторное пространство
	// SSA - анализ сегментов pw
	seg := 100 // номер сегмента pw для визуализации
	nET := 4   // кол-во сингулярных троек для сегментов pw

	//var sET12 mat.Dense
	sET12_sum2 := mat.NewDense(win, 2, nil) // НЕ ФАКТ, ЧТО К-во строк win
	sET12 := mat.NewDense(win, S, nil)      // НЕ ФАКТ, ЧТО К-во строк win
	for j := 0; j < S; j++ {                // цикл по сегментам  S
		//C, LBD, RC := SSA(win, M, spw[:][j], nET)
		_, _, RC := SSA(win, M, spw.ColView(j), nET)
		fmt.Println(j, S)
		RC_T := mat.DenseCopyOf(RC.T())

		sET12_sum2.SetCol(0, RC_T.RawRowView(0))
		sET12_sum2.SetCol(1, RC_T.RawRowView(1))
		sET12.SetCol(j, sum2(*sET12_sum2))
		sET12_sum2.Zero()
	}
	safeToXlsxMatrix(sET12, "sET12")

	fmt.Println("dt", cad)
	fmt.Println("dt", dt)
	fmt.Println("Imax", Imax)
	fmt.Println("Imin", Imin)
	fmt.Println("K", K)
	fmt.Println("M", M)
	fmt.Println("N", N)
	fmt.Println("nET", nET)
	fmt.Println("nPart", nPart)
	fmt.Println("ns", ns)
	fmt.Println("NSF", NSF)
	fmt.Println("overlap", overlap)
	fmt.Println("res", res)
	fmt.Println("S", S)
	safeToXlsx(tim, "tim")
	safeToXlsx(L, "L")
	fmt.Println("win", win)
	fmt.Println("seg", seg)

}

// Горизонтальная сумма массива
func sum2(a mat.Dense) []float64 {
	var output []float64
	r, _ := a.Dims()
	for i := 0; i < r; i++ {
		output = append(output, mat.Sum(a.RowView(i)))
	}
	return output
}

// МАксимальный элемент массива
func max(arr []float64) float64 {
	max_num := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max_num {
			max_num = arr[i]
		}
	}
	return max_num
}

// Создать динамический двумерный массив
func dinamicArray(r, c int) [][]float64 {
	spw := make([][]float64, r)
	for index := range spw {
		spw[index] = make([]float64, c)
	}
	return spw
}

func SSA(N int, M int, X mat.Vector, nET int) (mat.Dense, []float64, mat.Dense) {
	//  Calculate covariance matrix (trajectory approach)
	// it ensures a positive semi-definite covariance matrix
	Y := BuildTrajectoryMatrix(X, M, N) // Создать матрицу траекторий
	safeToXlsxMatrix(Y, "Y")

	var Cemb mat.Dense
	Cemb.Mul(Y, Y.T())
	Cemb.Scale(1.0/float64(N-M+1), &Cemb)

	safeToXlsxM(Cemb, "Cemb")

	C := Cemb

	// Choose covariance estimation
	RHO, LBD := eig(C)
	safeToXlsxM(RHO, "RHO")
	safeToXlsxM(LBD, "LBD")
	LBD_diag := diag(LBD, LBD.DiagView().Diag())
	LBD_sort, _ := InsertionSort(LBD_diag)

	// Перевернуть матрицу по вертикали
	_, col_RHO := RHO.Dims()
	for j := 0; j < col_RHO/2; j++ {
		a := colDense(RHO, j)
		b := colDense(RHO, col_RHO-1-j)
		RHO.SetCol(j, b)
		RHO.SetCol(col_RHO-1-j, a)
	}
	safeToXlsxM(RHO, "RHO_new")

	// Calculate principal components PC
	// The principal components are given as the scalar product
	// between Y, the time-delayed embedding of X, and the eigenvectors RHO
	var PC mat.Dense
	PC.Mul(Y.T(), &RHO)
	safeToXlsxM(PC, "PC")

	// Calculate reconstructed components RC
	// In order to determine the reconstructed components RC,
	// we have to invert the projecting PC = Y*RHO;i.e. RC = Y*RHO*RHO'=PC*RHO'
	// Averaging along anti-diagonals gives the RCs for the original input X

	RC := mat.NewDense(N, nET, nil)
	r_PC, _ := PC.Dims()
	r_RHO, _ := RHO.Dims()
	for m := 0; m < nET; m++ {
		// invert projection
		var buf mat.Dense
		b1 := mat.NewDense(r_PC, 1, colDense(PC, m))
		b2 := mat.NewDense(r_RHO, 1, colDense(RHO, m))
		buf.Mul(b1, b2.T())
		safeToXlsxM(buf, "buf")

		// Перевернуть матрицу по горизонтали
		row_buf, _ := buf.Dims()
		for j := 0; j < row_buf/2; j++ {
			a := rowDense(buf, j)
			b := rowDense(buf, row_buf-1-j)
			buf.SetRow(j, b)
			buf.SetRow(row_buf-1-j, a)
		}
		safeToXlsxM(buf, "buf2")

		// Anti-diagonal averaging
		for n := 0; n < N; n++ {
			diag_buf, error_subdiagonal := subdiagonal(buf, -(N-M+1)+n+1)
			if error_subdiagonal != nil {
				panic(error_subdiagonal)
			}
			RC.Set(n, m, averge(diag_buf))
		}

	}
	return C, LBD_sort, *RC
}

// Среднее значение каждого столбца
func mean(m mat.Dense) []float64 {
	_, c_m := m.Dims()
	outputArray := make([]float64, c_m)
	for ind := range outputArray {
		vect := colDense(m, ind)
		outputArray[ind] = averge(vect)
	}
	return outputArray
}

// Среднее массива float64
func averge(array []float64) float64 {
	var sum float64
	for _, val := range array {
		sum += val
	}
	return sum / float64(len(array))
}

// Вернуть subdiagonal
func subdiagonal(m mat.Dense, k int) ([]float64, error) {
	var outputArray []float64
	r_m, c_m := m.Dims()
	if k == 0 {
		for i := 0; i < r_m && i < c_m; i++ {
			outputArray = append(outputArray, m.At(i, i))
		}
	} else if k < 0 {
		if k > r_m {
			return nil, errors.New("k > matrix ROW")
		}
		for i := 0; i-k < r_m && i < c_m; i++ {
			outputArray = append(outputArray, m.At(i-k, i))
		}
	} else if k > 0 {
		if k > c_m {
			return nil, errors.New("k > matrix COL")
		}
		for i := 0; i+k < c_m && i < r_m; i++ {
			outputArray = append(outputArray, m.At(i, i+k))
		}
	}
	return outputArray, nil
}

// Вернуть колонку из матрицы
func colDense(m mat.Dense, ind int) []float64 {
	row, _ := m.Dims()
	outputArray := make([]float64, row)
	for i := 0; i < row; i++ {
		outputArray[i] = m.At(i, ind)
	}
	return outputArray
}

// Вернуть строку из матрицы
func rowDense(m mat.Dense, ind int) []float64 {
	_, col := m.Dims()
	outputArray := make([]float64, col)
	for j := 0; j < col; j++ {
		outputArray[j] = m.At(ind, j)
	}
	return outputArray
}
func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
func InsertionSort(array []float64) ([]float64, []int) {
	indexArray := make([]int, len(array))
	for ind := range indexArray {
		indexArray[ind] = (ind) + 1
	}
	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 {
			if array[j-1] < array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				indexArray[j-1], indexArray[j] = indexArray[j], indexArray[j-1]
			}
			j = j - 1
		}
	}
	return array, indexArray
}
