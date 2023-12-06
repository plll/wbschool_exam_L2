package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

 Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
 Стратегия позволяет варьировать поведение объекта во время выполнения программы, подставляя в него различные объекты-поведения (например, отличающиеся балансом скорости и потребления ресурсов).

 Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.
 Стратегия позволяет вынести отличающееся поведение в отдельную иерархию классов, а затем свести первоначальные классы к одному, сделав поведение этого класса настраиваемым.

 Когда вы не хотите обнажать детали реализации алгоритмов для других классов.
 Стратегия позволяет изолировать код, данные и зависимости алгоритмов от других объектов, скрыв эти детали внутри классов-стратегий.

 Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора. Каждая ветка такого оператора представляет собой вариацию алгоритма.
 Стратегия помещает каждую лапу такого оператора в отдельный класс-стратегию. Затем контекст получает определённый объект-стратегию от клиента и делегирует ему работу. Если вдруг понадобится сменить алгоритм, в контекст можно подать другую стратегию.

Плюсы
 Горячая замена алгоритмов на лету.
 Изолирует код и данные алгоритмов от остальных классов.
 Уход от наследования к делегированию.
 Реализует принцип открытости/закрытости.

Минусы
 Усложняет программу за счёт дополнительных классов.
 Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую
*/

func main() {
	product := "vehicle"
	payWay := 3

	var payment Payment
	switch payWay {
	case 1:
		payment = NewCardPayment("12345", "12345")
	case 2:
		payment = NewPayPalPayment()
	case 3:
		payment = NewQIWIPayment()
	}

	processOrder(product, payment)
}

func processOrder(product string, payment Payment) {
	// ... implementation
	err := payment.Pay()
	if err != nil {
		return
	}
}

// ----

type Payment interface {
	Pay() error
}

// ----

type cardPayment struct {
	cardNumber, cvv string
}

func NewCardPayment(cardNumber, cvv string) Payment {
	return &cardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (p *cardPayment) Pay() error {
	// ... implementation
	return nil
}

type payPalPayment struct {
}

func NewPayPalPayment() Payment {
	return &payPalPayment{}
}

func (p *payPalPayment) Pay() error {
	// ... implementation
	return nil
}

type qiwiPayment struct {
}

func NewQIWIPayment() Payment {
	return &qiwiPayment{}
}

func (p *qiwiPayment) Pay() error {
	// ... implementation
	return nil
}
