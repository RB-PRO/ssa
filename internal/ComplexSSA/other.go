package complexssa

import "strings"

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
