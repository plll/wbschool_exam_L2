Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
В текущем виде, программа будет зависать (deadlock), потому что канал c в функции merge никогда не закрывается.

В Go, range c продолжает читать из канала c, пока он не будет закрыт. В вашем случае, горутина, которая пишет в канал c, никогда не закрывает его, поэтому range c продолжает ожидать новые данные, которые никогда не приходят.

Кроме того, в функции merge чтение из каналов a и b без проверки, закрыты они или нет. Если канал закрыт, чтение из него всегда будет возвращать нулевое значение для типа канала. В вашем случае, это будет 0 для int.


```
