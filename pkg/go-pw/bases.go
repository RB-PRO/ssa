// Пакет для расчёта сигнала фотоплетизмографии
//
// Реализованы следующие методы:
//   - Cr
package gopw

const (
	MethodCr string = "Cr"
)

// Расчёт пульсовой волны по выбранному методу
//
// По умолчанию - Cr
//
// Представленные алгоритмы извлечения сигнала фотоплетизмографии:
//
//	- Cr
func CalculatePW(R, G, B []float64, Method string) ([]float64, error) {
	switch Method {
	case Method:
		return cr(R, G, B)
	default:
		return cr(R, G, B)
	}
}
