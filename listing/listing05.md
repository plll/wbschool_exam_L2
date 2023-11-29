Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Функция test() возвращает значение nil типа *customError, которое затем приводится к типу error (это интерфейсный тип в Go). Важно понимать, что nil для конкретного типа и nil для интерфейсного типа - это не одно и то же.

Когда мы присваиваем nil типа *customError переменной err типа error, err на самом деле не будет nil, потому что его значение (nil) и тип (*customError) оба не nil. В Go, интерфейсное значение считается nil только если и его значение, и его тип равны nil.

```
