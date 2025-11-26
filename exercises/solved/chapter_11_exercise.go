package main

import (
	"errors"
	"fmt"
	"math"
)

// TODO: Импортируй необходимые пакеты

// TODO: Создай кастомную ошибку InsufficientFundsError со структурой:
// - Requested (float64) - запрошенная сумма
// - Available (float64) - доступная сумма
// Реализуй метод Error() который возвращает строку вида:
// "недостаточно средств: запрошено 150.00, доступно 100.00"
type InsufficientFundsError struct {
	Requested float64
	Available float64
}

func (i *InsufficientFundsError) Error() string {
	return fmt.Sprintf("недостаточно средств: запрошено %.2f, доступно %.2f", i.Requested, i.Available)
}

// TODO: Создай sentinel error ErrNegativeAmount с сообщением "сумма не может быть отрицательной"
var ErrNegativeAmount = errors.New("сумма не может быть отрицательной")

// TODO: Функция withdraw(balance, amount float64) (float64, error)
// Должна:
// - Вернуть ErrNegativeAmount если amount < 0
// - Вернуть InsufficientFundsError если amount > balance
// - Вернуть новый баланс (balance - amount) и nil при успехе
func withdraw(balance, amount float64) (float64, error) {
	if amount < 0 {
		return 0.0, ErrNegativeAmount
	}
	if amount > balance {
		return 0.0, &InsufficientFundsError{
			Requested: amount,
			Available: balance,
		}
	}

	return balance - amount, nil
}

// TODO: Функция deposit(balance, amount float64) (float64, error)
// Должна:
// - Вернуть ErrNegativeAmount если amount < 0
// - Вернуть новый баланс (balance + amount) и nil при успехе
func deposit(balance, amount float64) (float64, error) {
	if amount < 0 {
		return 0, ErrNegativeAmount
	}

	return balance + amount, nil
}

// TODO: Функция transfer(fromBalance, toBalance, amount float64) (newFrom, newTo float64, err error)
// Должна:
// - Попытаться снять amount с fromBalance используя withdraw
// - Если успешно - добавить amount к toBalance используя deposit
// - При ошибке на любом этапе - обернуть её через fmt.Errorf с контекстом "перевод средств: %w"
// - Вернуть обновлённые балансы и nil при успехе
func transfer(fromBalance, toBalance, amount float64) (newFrom, newTo float64, err error) {
	newFrom, err = withdraw(fromBalance, amount)
	if err != nil {
		return 0.0, 0.0, fmt.Errorf("перевод средств: %w", err)
	}
	newTo, err = deposit(toBalance, amount)
	if err != nil {
		return 0.0, 0.0, fmt.Errorf("перевод средств: %w", err)
	}

	return newFrom, newTo, nil
}

// TODO: Функция processTransactions(balance float64, amounts []float64) (float64, error)
// Должна:
//   - Применить все операции из amounts к balance (положительные - deposit, отрицательные - withdraw)
//   - При первой ошибке остановиться и вернуть текущий баланс и ошибку обёрнутую через fmt.Errorf
//     с указанием индекса операции: "операция #%d: %w"
//   - Вернуть финальный баланс и nil если все операции успешны
func processTransactions(balance float64, amounts []float64) (float64, error) {
	var err error = nil

	for i, amount := range amounts {
		if amount > 0 {
			balance, err = deposit(balance, amount)
		} else if amount < 0 {
			balance, err = withdraw(balance, math.Abs(amount))
		}

		if err != nil {
			return 0.0, fmt.Errorf("операция #%d: %w", i, err)
		}
	}

	return balance, nil
}

func main() {
	// TODO: Тест 1 - успешное снятие
	// Сними 30.0 с баланса 100.0
	_, err := withdraw(100.0, 30.0)
	if err == nil {
		fmt.Println("We could withdraw")
	}

	// TODO: Тест 2 - недостаточно средств
	// Попытайся снять 150.0 с баланса 100.0
	// Проверь тип ошибки через type assertion (*InsufficientFundsError)
	// Если это InsufficientFundsError - выведи детали (Requested и Available)
	_, err = withdraw(100.0, 150.0)
	if err != nil {
		if fundsError, ok := err.(*InsufficientFundsError); ok {
			fmt.Println(fundsError)
			fmt.Printf("error: details Request:%.2f, Available:%.2f\n", fundsError.Requested, fundsError.Available)
		}
	}

	// TODO: Тест 3 - отрицательная сумма
	// Попытайся снять -50.0
	// Проверь ошибку через errors.Is(err, ErrNegativeAmount)
	_, err = withdraw(500.0, -50)

	if errors.Is(err, ErrNegativeAmount) {
		fmt.Println("It's ErrNegativeAmount")
	}
	// TODO: Тест 4 - успешный перевод
	// Переведи 40.0 с баланса 100.0 на баланс 50.0
	_, _, err = transfer(100, 50, 40)
	if err == nil {
		fmt.Println("We could transfer")
	}
	// TODO: Тест 5 - перевод с недостаточными средствами
	// Попытайся перевести 200.0 с баланса 100.0 на баланс 50.0
	// Проверь что ошибка содержит InsufficientFundsError через errors.As
	_, _, err = transfer(100, 50, 200)

	var myError *InsufficientFundsError
	if errors.As(err, &myError) {
		fmt.Println("It's InsufficientFundsError")
	}

	// TODO: Тест 6 - пакетные операции
	// Начальный баланс 100.0, операции: [20.0, -30.0, 50.0, -10.0]
	// Должно получиться: 100 + 20 - 30 + 50 - 10 = 130.0
	amounts := []float64{20.0, -30.0, 50.0, -10.0}

	balance, err := processTransactions(100.0, amounts)
	if err == nil {
		fmt.Printf("Balance=%.2f\n", balance)
	}
	// TODO: Тест 7 - пакетные операции с ошибкой
	// Начальный баланс 100.0, операции: [20.0, -30.0, -150.0, 50.0]
	// Должна произойти ошибка на 3-й операции (индекс 2)
	// Проверь что errors.Is находит ErrNegativeAmount или InsufficientFundsError
	amounts = []float64{20.0, -30.0, -150.0, 50.0}
	balance, err = processTransactions(100.0, amounts)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, ErrNegativeAmount) || errors.Is(err, &InsufficientFundsError{}) {

		}
	}
}
