// # Окно Блэкмана-Хэрриса
// На вход получает N, расчёты ведёт до N-1
//
// Окно Блэкмана-Харриса (Blackman-Harris), как и окно Хэнинга,
// применяется для измерения очень слабых компонент на фоне большого входного сигнала, таких как нелинейные искажения
package blackmanharris

import "math"

// Расчёт окна для Блэкмана-Харриса (Blackman-Harris) при заданном N и массиве коэффициентах
//
// Входные данные:
//	-N - Длина окна
//	-a - Коэффтиенты Блэкмен-Харриса
//
// Пример:
// 	blackmanharris(32, blackmanharris.Koef4_92db)
func Blackmanharris(N int, a [4]float64) []float64 {
	Window := make([]float64, N-1) // Выходной массив окна
	Nfloat64 := float64(N)         // N в float64, чтобы не не переводить трижды
	for n := range Window {        // Цикл по всему окну
		nfloat64 := float64(n) // n в float64, чтобы не не переводить трижды

		// Уравнение для периодического окна Блэкмен-Харриса с четырьмя терминами длины N
		Window[n] = a[0] - a[1]*math.Cos((2.0*math.Pi/Nfloat64)*1*nfloat64) +
			a[2]*math.Cos((2.0*math.Pi/Nfloat64)*2*nfloat64) +
			a[3]*math.Cos((2.0*math.Pi/Nfloat64)*3*nfloat64)
	}
	return Window
}
