package operations

import (
	"fmt"
	"sync"
	"math/rand"
	"time"

	"bank-client/client"
)

// StartOperations запускает операции (депозиты или снятия) в горутинах
func StartOperations(bankClient *client.Client, wg *sync.WaitGroup, quit chan struct{}, operation string, amount int, r *rand.Rand) {
	defer wg.Done()

	select {
	case <-quit:
		return
	default:
		// Используем r для генерации случайных чисел
		waitTime := r.Intn(1000) // Случайное время ожидания от 0 до 999 мс
		time.Sleep(time.Duration(waitTime) * time.Millisecond)

		// Логика операций
		if operation == "deposit" {
			bankClient.Deposit(amount)
		} else if operation == "withdrawal" {
			bankClient.Withdrawal(amount)
		}
	}
}


// ClientOperation обрабатывает команды пользователя
func ClientOperation(c *client.Client, command string, amount int) {
	switch command {
	case "deposit":
		c.Deposit(amount)
	case "withdrawal":
		err := c.Withdrawal(amount)
		if err != nil {
			fmt.Println("Ошибка:", err)
		}
	case "balance":
		balance := c.GetBalance()
		fmt.Printf("Текущий баланс: %d\n", balance)
	default:
		fmt.Println("Неизвестная команда.")
	}
}
