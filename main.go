package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

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

type Transaction struct {
	Date   time.Time
	Debtor string
	Amount float32
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

	transactions := make([]Transaction, 0, len(lines))
	for i, line := range lines {
		if i == 0 {
			continue
		}
		transactions = append(transactions, Transaction{Date: formatDate(line[1]), Debtor: line[2], Amount: parseAmountAsFloat32(line[5])})
	}

	fmt.Println(transactions)
}
