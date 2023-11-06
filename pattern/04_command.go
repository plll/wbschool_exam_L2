package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
    Задумка данного паттерна заключается в том, что следует убрать прямую связность двух объектов добавлением промежуточного слоя - комманды. Данный паттерн дает возможность хранить в объекте команды всю информацию нужную для выполнения + возможность сохранять объект команды, чтобы выполнить ее потом.
    Минусы паттерна - Добавление доп слоя между структурами
*/

import "fmt"

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type command interface {
	execute()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type device interface {
	on()
	off()
}

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &tv{}
	onCommand := &onCommand{
		device: tv,
	}
	offCommand := &offCommand{
		device: tv,
	}
	onButton := &button{
		command: onCommand,
	}
	onButton.press()
	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
