package main

import (
    "fmt"
    "log"

    "github.com/Asetka76/loancalc-project/loancalc"
    "github.com/fatih/color"
)

func main() {
    sum := 1000000.0
    percent := 12.0
    months := 24
    client := "John Doe"

    payment, err := loancalc.MonthlyPayment(sum, percent, months)
    if err != nil {
        log.Fatalf("Error calculating monthly payment: %v", err)
    }

    err = loancalc.ApplyEarlyPayment(&payment, 500.0)
    if err != nil {
        log.Fatalf("Error applying early payment: %v", err)
    }

    report, err := loancalc.FormatLoanReport(client, payment, months)
    if err != nil {
        log.Fatalf("Error formatting report: %v", err)
    }

    color.Green("Loan Calculation Successful!")
    fmt.Printf("Report Details:\n%s\n", report)
    fmt.Printf("Final Adjusted Payment (precision 3): %.3f\n", payment)
}
