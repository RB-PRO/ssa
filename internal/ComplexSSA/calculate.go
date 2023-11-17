// Пакет для комплексного анализа видеоряда методом SSA-метод гусеница
package complexssa

import (
	"strings"

	gomathtests "github.com/RB-PRO/ssa/pkg/go-MathTests"
	goroi "github.com/RB-PRO/ssa/pkg/go-ROI"
	gopw "github.com/RB-PRO/ssa/pkg/go-pw"
)

func Start() {
	Folder := "WorkPath/"               // Рабочая папка
	VideoName := "P1H1.avi"             // Название видео
	ObjName := NameVideoFile(VideoName) //  Получить название объекта исследования

	// % Вычленение RGB из видео
	R, G, B, _ := goroi.ExtractRGB(Folder, VideoName)
	gomathtests.Save(Folder+ObjName+"_RGB.txt", R, G, B)
	// R, G, B, _ := gomathtests.LoadRGB(Folder + ObjName + "_RGB.txt")
	gomathtests.Plot(Folder+ObjName+"_RGB.png", R, G, B)

	// % Получение pw
	pw, _ := gopw.CalculatePW(R, G, B, gopw.MethodCr)
	gomathtests.Save(Folder+ObjName+"_pw.txt", pw)
	gomathtests.Plot(Folder+ObjName+"_pw.png", pw)

	SSA(Folder, pw)
}

// Преобразовать название видео в название объекта исследования
func NameVideoFile(str string) string {
	str = strings.ReplaceAll(str, ".txt", "")
	str = strings.ReplaceAll(str, ".avi", "")
	str = strings.ReplaceAll(str, ".mov", "")
	str = strings.ReplaceAll(str, ".mp4", "")
	str = strings.ReplaceAll(str, "_RGB", "")
	str = strings.ReplaceAll(str, "_but", "")
	str = strings.ReplaceAll(str, "_pw", "")
	return str
}
