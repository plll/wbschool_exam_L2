package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

Задумка паттерна заключается в том, чтобы делегировать классу-помощнику инициализацию больших компонированных объектов, где каждому полю объекта описывается свой метод
Плюсы - читабельность кода, вариативность инциализации при отсутствии некоторых параметров
Минусы - Создание доп. класса с больших кол-вом методов
*/

// Car represents the complex object being built.
type Car struct {
	color         string
	engineType    string
	hasSunroof    bool
	hasNavigation bool
}

// CarBuilder provides an interface for constructing the parts of the car.
type CarBuilder interface {
	SetColor(color string) CarBuilder
	SetEngineType(engineType string) CarBuilder
	SetSunroof(hasSunroof bool) CarBuilder
	SetNavigation(hasNavigation bool) CarBuilder
	Build() *Car
}

// NewCarBuilder creates a new CarBuilder.
func NewCarBuilder() CarBuilder {
	return &carBuilder{
		car: &Car{}, // Initialize the car attribute
	}
}

// carBuilder implements the CarBuilder interface.
type carBuilder struct {
	car *Car
}

func (cb *carBuilder) SetColor(color string) CarBuilder {
	cb.car.color = color
	return cb
}

func (cb *carBuilder) SetEngineType(engineType string) CarBuilder {
	cb.car.engineType = engineType
	return cb
}

func (cb *carBuilder) SetSunroof(hasSunroof bool) CarBuilder {
	cb.car.hasSunroof = hasSunroof
	return cb
}

func (cb *carBuilder) SetNavigation(hasNavigation bool) CarBuilder {
	cb.car.hasNavigation = hasNavigation
	return cb
}

func (cb *carBuilder) Build() *Car {
	return cb.car
}

// Director provides an interface to build cars.
type Director struct {
	builder CarBuilder
}

func (d *Director) ConstructCar(color, engineType string, hasSunroof, hasNavigation bool) *Car {
	d.builder.SetColor(color).
		SetEngineType(engineType).
		SetSunroof(hasSunroof).
		SetNavigation(hasNavigation)

	return d.builder.Build()
}

func main() {
	builder := NewCarBuilder()

	director := &Director{builder: builder}
	//myCar := director.builder.SetEngineType("Water").SetColor("Pink").Build()
	myCar := director.ConstructCar("blue", "electric", true, true)
	fmt.Println(myCar)

}
