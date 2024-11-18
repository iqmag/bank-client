package main

import (
	"fmt"
	"sync"
	"time"

	"bank-client/client"
	"bank-client/operations"
	"math/rand"
)

func main() {
	bankClient := &client.Client{CurrentBalance: 0}
	var wg sync.WaitGroup
	quit := make(chan struct{})

	// Запускаем горутины для пополнения счета
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			operations.StartOperations(bankClient, &wg, quit, "deposit", 10, r)
		}()
	}

	// Запускаем горутины для снятия денег
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			operations.StartOperations(bankClient, &wg, quit, "withdrawal", 5, r)
		}()
	}

	// Обработка ввода пользователя
	for {
		var command string
		fmt.Print("Введите одну из указанных в скобках команд: (balance, deposit, withdrawal, exit): ")
		fmt.Scanln(&command)

		if command == "exit" {
			close(quit)
			wg.Wait()
			fmt.Println("Завершение работы")
			break
		}

		var amount int
		if command == "deposit" || command == "withdrawal" {
			fmt.Print("Введите сумму: ")
			fmt.Scanln(&amount)
		}

		operations.ClientOperation(bankClient, command, amount)
	}
}
