package client

import (
	"fmt"
	"sync"
)

// BankClient интерфейс для банковского клиента
type BankClient interface {
	Deposit(amount int) // внесение депозита
	Withdrawal(amount int) error // снятие денег, возвращает ошибку в случае неудачи
	GetBalance() int // получение текущего баланса
}

// Client структура клиента
type Client struct {
	CurrentBalance int
	Mu             sync.Mutex
}

// Deposit - метод для внесения депозита на счет клиента
func (c *Client) Deposit(amount int) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.CurrentBalance += amount
	fmt.Printf("Депозит на сумму %d завершен. Новый баланс: %d\n", amount, c.CurrentBalance)
}

// Withdrawal - метод для снятия денег со счета клиента
func (c *Client) Withdrawal(amount int) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if c.CurrentBalance < amount {
		return fmt.Errorf("сумма вывода (%d) превышает баланс (%d)", amount, c.CurrentBalance)
	}
	c.CurrentBalance -= amount
	fmt.Printf("Вывод %d завершен. Новый баланс: %d\n", amount, c.CurrentBalance)
	return nil
}

// GetBalance - метод для получения текущего баланса клиента
func (c *Client) GetBalance() int {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return c.CurrentBalance
}
