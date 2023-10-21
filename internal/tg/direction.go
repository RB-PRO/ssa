package tg

import (
	"os"
	"path/filepath"
)

// Пересоздать папку
func MakeDir(Path string) (string, error) {

	// Абсолютный путь до папки. Если его нет, то удаляем всё
	absFolderPath, _ := filepath.Abs(Path)

	// Если папка существует - удаляем
	if _, err := os.Stat(absFolderPath); err == nil {
		os.RemoveAll(Path)
	}

	// Создание пути
	ErrMkdirAll := os.MkdirAll(Path, 0777)
	if ErrMkdirAll != nil {
		return "", ErrMkdirAll
	}
	return absFolderPath, nil
}
