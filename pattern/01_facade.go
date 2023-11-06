package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern

   Основные плюсы данного паттерна заключаются в том, что паттерн пораждает более удобный для вызова интерфес, который в свою очередь делает код более читабельным и собираюет в себе сложную подсистему, управлять которой в коде было бы непросто.
   В нижеприведенном примере представлен интерфейс Man, который в себе определяет  3 интерфейса, реализация у них не связанная, но они нужны для определения функции TODO, поэтому стало хорошим выбором, реализовать паттерн фасад, чтобы в коде не вызывать 3 метода из 3 разных интерфейсов, а вызвать 1, который своей логикой объединяет разрозненные системы
*/

import (
	"strings"
)

// NewMan creates man.
func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

// Man implements man and facade.
type Man struct {
	house *House
	tree  *Tree
	child *Child
}

// Todo returns that man must do.
func (m *Man) Todo() string {
	result := []string{
		m.house.Build(),
		m.tree.Grow(),
		m.child.Born(),
	}
	return strings.Join(result, "\n")
}

// House implements a subsystem "House"
type House struct {
}

// Build implementation.
func (h *House) Build() string {
	return "Build house"
}

// Tree implements a subsystem "Tree"
type Tree struct {
}

// Grow implementation.
func (t *Tree) Grow() string {
	return "Tree grow"
}

// Child implements a subsystem "Child"
type Child struct {
}

// Born implementation.
func (c *Child) Born() string {
	return "Child born"
}
