package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Плюсы
Создание общего конструктора для типов, одного интерфейса (Конструктор возвращает переменные типа интерфейс)
Избавляет от привязки к конкретному типу
Упрощает добавление новых типов объектов
Реализует прицнип открытости и закрытости

Минусы
Создание больших параллельных иерархий (различных структур)
Божественный конструктор (Сильная привязка к реализации)

*/

import "io"

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage( /*...*/ )
	case DiskStorage:
		return newDiskStorage( /*...*/ )
	default:
		return newTempStorage( /*...*/ )
	}
}
func main() {
	s, _ := NewStore(MemoryStorage)
	f, _ := s.Open("file")

	n, _ := f.Write([]byte("data"))
	defer f.Close()
}
