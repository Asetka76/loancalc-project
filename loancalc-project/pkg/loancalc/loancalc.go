// Package loancalc provides financial calculations for loans,
// including monthly rates, annuity payments, early repayments, and reporting.
package loancalc

import (
	"fmt"
	"math"
)

// MonthlyRate calculates the monthly interest rate from the annual percentage.
// Returns an error if the annual percent is negative or zero.
func MonthlyRate(annualPercent float64) (float64, error) {
	if annualPercent <= 0 {
		return 0, fmt.Errorf("annual percent must be greater than zero, got %.2f", annualPercent)
	}
	return annualPercent / 100.0 / 12.0, nil
}

// MonthlyPayment calculates the monthly annuity loan payment.
// Returns an error if sum or months are invalid, or if rate calculation fails.
func MonthlyPayment(sum, annualPercent float64, months int) (float64, error) {
	if sum <= 0 {
		return 0, fmt.Errorf("loan sum must be positive, got %.2f", sum)
	}
	if months <= 0 {
		return 0, fmt.Errorf("months must be positive, got %d", months)
	}

	rate, err := MonthlyRate(annualPercent)
	if err != nil {
		return 0, err
	}

	pow := math.Pow(1+rate, float64(months))
	payment := sum * (rate * pow) / (pow - 1)
	if math.IsNaN(payment) || math.IsInf(payment, 0) {
		return 0, fmt.Errorf("calculation resulted in invalid number")
	}
	return payment, nil
}

// ApplyEarlyPayment modifies the pointer to the payment by subtracting the extra amount.
// Returns an error if the extra amount is negative or exceeds the current payment.
func ApplyEarlyPayment(payment *float64, extra float64) error {
	if payment == nil {
		return fmt.Errorf("payment pointer cannot be nil")
	}
	if extra < 0 {
		return fmt.Errorf("extra payment cannot be negative, got %.2f", extra)
	}
	if extra > *payment {
		return fmt.Errorf("extra payment (%.2f) cannot exceed current payment (%.2f)", extra, *payment)
	}
	*payment -= extra
	return nil
}

// FormatLoanReport generates a formatted string report for the loan.
// Returns an error if months are invalid or client name is empty.
func FormatLoanReport(client string, payment float64, months int) (string, error) {
	if client == "" {
		return "", fmt.Errorf("client name cannot be empty")
	}
	if payment < 0 {
		return "", fmt.Errorf("payment cannot be negative")
	}
	if months <= 0 {
		return "", fmt.Errorf("months must be positive")
	}

	return fmt.Sprintf("Client: %s | Monthly Payment: %.2f | Term: %d months | Total: %.3f", client, payment, months, payment*float64(months)), nil
}