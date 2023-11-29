package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout")
	flag.Parse()
	args := flag.Args()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	if len(args) < 2 || args[0] == "" || args[1] == "" {
		fmt.Println("Необходимо указать хост и порт для подключения")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Установка TCP-соединения
	conn, err := net.DialTimeout("tcp", args[0]+":"+args[1], *timeout)
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		os.Exit(1)
	}

	go func() {
		<-interrupt
		fmt.Println("Program graceful shutdown")
		conn.Close()
		os.Exit(0)
	}()

	// Чтение и запись данных через соединение
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Чтение строки из стандартного ввода
		fmt.Print("> ")
		scanned := scanner.Scan()
		if !scanned {
			break
		}

		// Отправка строки на удаленный хост
		input := scanner.Text()
		_, err := fmt.Fprintf(conn, input+"\n")
		if err != nil {
			fmt.Println("Ошибка отправки данных:", err)
			break
		}

		// Чтение ответа от удаленного хоста
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения данных:", err)
			break
		}

		// Вывод ответа на стандартный вывод
		fmt.Print("< " + response)
	}
}
