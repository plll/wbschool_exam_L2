package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
    Смысл данного паттерна заключается в том, что мы даем возможность пользователей нашей библиотеки писать доп функционал, не изменяя исходную библиотеку, что может привести к ошибкам или, если мы являемся единственным контрибутором, то данная задача переходит на нас.
    В интерфейс добавляется метод accept, котороый принимает "visitor"'a  в данном случае визитором является любая структура, которую любой человек может написать самостоятельно для расширения функционала. Главная задумка в том, что мы добавляем данный метод и у всех структур описываем  методе accept - visitor.FigureType
    А задача разработчика во все добавляемые функциональности указывать методы в таком же формате (someVisitor.FigureType()) и получается, что как итог у нашего объекта появится возможность вызывать методы других структур без добавления функционала в исходную структуру
    Минусы - Вместо того, чтобы добавить методы внутрь нашей структуры мы вынуждены плодить внешние структуры и методы
    Плюсы - Нам не нужно видоизменять исходный интерфейс кроме добавления метода accept
*/

import "fmt"

type shape interface {
	getType() string
	accept(visitor)
}

type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}

type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}

type rectangle struct {
	l int
	b int
}

func (t *rectangle) accept(v visitor) {
	v.visitForrectangle(t)
}

func (t *rectangle) getType() string {
	return "rectangle"
}

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForrectangle(*rectangle)
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	//Calculate area for square. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	//Calculate are for circle. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for circle")
}

func (a *areaCalculator) visitForrectangle(s *rectangle) {
	//Calculate are for rectangle. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for rectangle")
}

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *middleCoordinates) visitForrectangle(t *rectangle) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for rectangle")
}

func main() {
	square := &square{side: 2}
	circle := &circle{radius: 3}
	rectangle := &rectangle{l: 2, b: 3}
	areaCalculator := &areaCalculator{}
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &middleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
