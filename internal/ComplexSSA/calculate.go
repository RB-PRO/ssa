// Пакет для комплексного анализа видеоряда методом SSA-метод гусеница
package complexssa

import (
	gomathtests "github.com/RB-PRO/ssa/pkg/go-MathTests"
	goroi "github.com/RB-PRO/ssa/pkg/go-ROI"
	gopw "github.com/RB-PRO/ssa/pkg/go-pw"
)

func Start() {
	Folder := "WorkPath/"               // Рабочая папка
	VideoName := "P2LC1_edited.avi"     // Название видео
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

func Start2(Folder, VideoName string) {
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
