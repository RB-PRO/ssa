package ssa

import (
	"os"
	"path/filepath"
)

// Фунционал файл отвечает за работу директорией, создание, отслеживание, скачивание данных по позициям

// Структура управления файловой системы парсера
//
// Под капотом директория пути до самой папки с фотографиями
// и методы, которые отвечают за сохранение данных

// Пересоздать папку
func (spw *Direction) MakeDir(Path string) (string, error) {

	// Абсолютный путь до папки. Если его нет, то удаляем всё
	absFolderPath, _ := filepath.Abs(spw.zeropath + Path)

	// Если папка существует - удаляем
	if _, err := os.Stat(absFolderPath); err == nil {
		os.RemoveAll(spw.zeropath + Path)
	}

	// Создание пути
	ErrMkdirAll := os.MkdirAll(spw.zeropath+Path, 0777)
	if ErrMkdirAll != nil {
		return "", ErrMkdirAll
	}
	return absFolderPath, nil
}

func NewDirection(Path string) (*Direction, error) {

	Dir := Direction{Path}

	// Абсолютный путь до папки. Если его нет, то удаляем всё
	absFolderPath, _ := filepath.Abs(Path)

	// Если папка существует - удаляем
	if _, err := os.Stat(absFolderPath); err == nil {
		os.RemoveAll(Path)
	}

	// Создание пути
	ErrMkdirAll := os.MkdirAll(Path, 0777)
	if ErrMkdirAll != nil {
		return nil, ErrMkdirAll
	}
	return &Dir, nil
}
