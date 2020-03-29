package main

import (
	"encoding/csv"
	// "fmt"
	"github.com/kr/pretty"
	// "gonum.org/v1/plot"
	// "gonum.org/v1/plot/plotter"
	// "gonum.org/v1/plot/plotutil"
	// "gonum.org/v1/plot/vg"
	"log"
	// "math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Debit struct {
	Date   string
	Debtor string
	Amount float32
}

type Credit struct {
	Date     time.Time
	Creditor string
	Amount   float32
}

func removeTransactionMethod(t string) string {
	replacer := strings.NewReplacer("VDC-", "", "VDP-", "", "VDA-", "", "ATM-", "")
	return replacer.Replace(t)
}

func parseAmountAsFloat32(amount string) float32 {
	noComma := strings.Replace(amount, ",", "", -1)
	if noComma == "" {
		return 0
	}
	ret, err := strconv.ParseFloat(noComma, 32)
	if err != nil {
		log.Fatalln("Couldn't convert amount: ", err)
	}
	return float32(ret)
}

func formatDate(date string) time.Time {
	parsedDate, err := time.Parse("02/01/2006", date)
	if err != nil {
		log.Fatalln(err)
	}
	return parsedDate
}

func plotExpenditure(debits []Debit) {
	m := make(map[string]float32)

	//In for loops in Go, the first exposed variable is the index, and the second the item
	for _, debit := range debits {
		m[debit.Date] += debit.Amount
	}
	pretty.Print(m)
}

func main() {
	csvFile, err := os.Open("FEB20-MAR16Spending.csv")
	if err != nil {
		log.Fatalln("Could not open", err)
	}

	r := csv.NewReader(csvFile)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Could not open", err)
	}

	debits := make([]Debit, 0, len(lines))
	credits := make([]Credit, 0, len(lines))

	for i, line := range lines {
		if i == 0 {
			continue
		}

		if strings.Contains(line[1], "REVOLUT") {
			continue
		}

		if line[9] == "Credit" {
			credits = append(credits, Credit{Date: formatDate(line[1]), Creditor: line[2], Amount: parseAmountAsFloat32(line[6])})
		}

		if line[9] == "Debit" || line[9] == "Bill Payment" || line[9] == "ATM" {
			debits = append(debits, Debit{Date: line[1], Debtor: line[2], Amount: parseAmountAsFloat32(line[5])})
		}
	}
	plotExpenditure(debits)

	// pretty.Print("DEBITS: ", debits)
	// pretty.Print("Credits: ", credits)
}
