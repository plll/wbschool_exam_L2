package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

Плюсы
Удобно отслеживать на каком этапе возникли ошибки

Можно декомпозировать логику на подсущности, вместо того чтобы производить обработку в одной функции

Минусы
Шаблон «цепочка ответственности» может привести к более сложной структуре кода, чем альтернативные подходы.

Можно создать циклические ссылки в цепочке, если next ссылки не назначаются тщательно. Это может привести к бесконечным циклам или другому неожиданному поведению программы.

Шаблон «цепочка ответственности» может затруднить определение того какой Handler объект отвечает за обработку конкретного запроса. Это может затруднить отладку кода и понимания его поведения.
*/
import "errors"

type Order struct {
	WarehouseFilled bool
	ShippingFilled  bool
	BillingFilled   bool
}

type OrderHandler interface {
	SetNext(OrderHandler) OrderHandler
	Handle(*Order) error
}

type OrderProcessor struct {
	handler OrderHandler
}

func (o *OrderProcessor) SetHandler(handler OrderHandler) {
	o.handler = handler
}

func (o *OrderProcessor) Process(order *Order) error {
	return o.handler.Handle(order)
}

type WarehouseHandler struct {
	next OrderHandler
}

func (w *WarehouseHandler) SetNext(next OrderHandler) OrderHandler {
	w.next = next
	return next
}

func (w *WarehouseHandler) Handle(order *Order) error {
	if order.WarehouseFilled {
		if w.next != nil {
			return w.next.Handle(order)
		}
		return nil
	}
	return errors.New("warehouse not filled")
}

type ShippingHandler struct {
	next OrderHandler
}

func (s *ShippingHandler) SetNext(next OrderHandler) OrderHandler {
	s.next = next
	return next
}

func (s *ShippingHandler) Handle(order *Order) error {
	if order.ShippingFilled {
		if s.next != nil {
			return s.next.Handle(order)
		}
		return nil
	}
	return errors.New("shipping not filled")
}

type BillingHandler struct {
	next OrderHandler
}

func (b *BillingHandler) SetNext(next OrderHandler) OrderHandler {
	b.next = next
	return next
}

func (b *BillingHandler) Handle(order *Order) error {
	if order.BillingFilled {
		if b.next != nil {
			return b.next.Handle(order)
		}
		return nil
	}
	return errors.New("billing not filled")
}
