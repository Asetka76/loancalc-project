package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Asetka76/loancalc-project/loancalc"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

func main() {
	// Создаем сканер для чтения ввода из консоли
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите имя клиента: ")
	clientInput, _ := reader.ReadString('\n')
	client := strings.TrimSpace(clientInput)

	var sum float64
	fmt.Print("Введите сумму кредита (например, 1000000): ")
	_, err := fmt.Scan(&sum)
	if err != nil {
		log.Fatalf("Ошибка ввода суммы: %v", err)
	}

	var percent float64
	fmt.Print("Введите годовую процентную ставку (например, 12): ")
	_, err = fmt.Scan(&percent)
	if err != nil {
		log.Fatalf("Ошибка ввода ставки: %v", err)
	}

	var months int
	fmt.Print("Введите срок в месяцах (например, 24): ")
	_, err = fmt.Scan(&months)
	if err != nil {
		log.Fatalf("Ошибка ввода срока: %v", err)
	}

	// Генерируем уникальный ID для договора
	loanID := uuid.New()

	// Выполняем расчеты с помощью твоего пакета
	payment, err := loancalc.MonthlyPayment(sum, percent, months)
	if err != nil {
		log.Fatalf("Ошибка расчета ежемесячного платежа: %v", err)
	}

	// Применяем досрочный платеж (например, фиксированные 500)
	err = loancalc.ApplyEarlyPayment(&payment, 500.0)
	if err != nil {
		log.Fatalf("Ошибка применения досрочного платежа: %v", err)
	}

	report, err := loancalc.FormatLoanReport(client, payment, months)
	if err != nil {
		log.Fatalf("Ошибка формирования отчета: %v", err)
	}

	// Вывод результатов
	color.Green("\nРасчет кредита выполнен успешно!")
	fmt.Printf("ID договора: %s\n", loanID.String())
	fmt.Printf("Детали отчета:\n%s\n", report)
	fmt.Printf("Итоговый скорректированный платеж (с точностью до 3 знаков): %.3f\n", payment)
}