package main

// TODO: Импортируй необходимые пакеты

// TODO: Создай кастомную ошибку InsufficientFundsError со структурой:
// - Requested (float64) - запрошенная сумма
// - Available (float64) - доступная сумма
// Реализуй метод Error() который возвращает строку вида:
// "недостаточно средств: запрошено 150.00, доступно 100.00"

// TODO: Создай sentinel error ErrNegativeAmount с сообщением "сумма не может быть отрицательной"

// TODO: Функция withdraw(balance, amount float64) (float64, error)
// Должна:
// - Вернуть ErrNegativeAmount если amount < 0
// - Вернуть InsufficientFundsError если amount > balance
// - Вернуть новый баланс (balance - amount) и nil при успехе

// TODO: Функция deposit(balance, amount float64) (float64, error)
// Должна:
// - Вернуть ErrNegativeAmount если amount < 0
// - Вернуть новый баланс (balance + amount) и nil при успехе

// TODO: Функция transfer(fromBalance, toBalance, amount float64) (newFrom, newTo float64, err error)
// Должна:
// - Попытаться снять amount с fromBalance используя withdraw
// - Если успешно - добавить amount к toBalance используя deposit
// - При ошибке на любом этапе - обернуть её через fmt.Errorf с контекстом "перевод средств: %w"
// - Вернуть обновлённые балансы и nil при успехе

// TODO: Функция processTransactions(balance float64, amounts []float64) (float64, error)
// Должна:
// - Применить все операции из amounts к balance (положительные - deposit, отрицательные - withdraw)
// - При первой ошибке остановиться и вернуть текущий баланс и ошибку обёрнутую через fmt.Errorf
//   с указанием индекса операции: "операция #%d: %w"
// - Вернуть финальный баланс и nil если все операции успешны

func main() {
	// TODO: Тест 1 - успешное снятие
	// Сними 30.0 с баланса 100.0

	// TODO: Тест 2 - недостаточно средств
	// Попытайся снять 150.0 с баланса 100.0
	// Проверь тип ошибки через type assertion (*InsufficientFundsError)
	// Если это InsufficientFundsError - выведи детали (Requested и Available)

	// TODO: Тест 3 - отрицательная сумма
	// Попытайся снять -50.0
	// Проверь ошибку через errors.Is(err, ErrNegativeAmount)

	// TODO: Тест 4 - успешный перевод
	// Переведи 40.0 с баланса 100.0 на баланс 50.0

	// TODO: Тест 5 - перевод с недостаточными средствами
	// Попытайся перевести 200.0 с баланса 100.0 на баланс 50.0
	// Проверь что ошибка содержит InsufficientFundsError через errors.As

	// TODO: Тест 6 - пакетные операции
	// Начальный баланс 100.0, операции: [20.0, -30.0, 50.0, -10.0]
	// Должно получиться: 100 + 20 - 30 + 50 - 10 = 130.0

	// TODO: Тест 7 - пакетные операции с ошибкой
	// Начальный баланс 100.0, операции: [20.0, -30.0, -150.0, 50.0]
	// Должна произойти ошибка на 3-й операции (индекс 2)
	// Проверь что errors.Is находит ErrNegativeAmount или InsufficientFundsError
}
