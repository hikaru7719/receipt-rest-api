package parser

import (
	"bufio"
	"strconv"
	"strings"
)

type Parser interface {
	PostCreditParser(scanner *bufio.Scanner) (string, int, int, int, error)
}

type parser struct {
}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) PostCreditParser(scanner *bufio.Scanner) (string, int, int, int, error) {
	var (
		cardName                                      string
		finishDate, withdrawalDate, laterPaymentMonth int
	)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "/")
		switch lineSplit[0] {
		case "カード名":
			cardName = lineSplit[1]
		case "締め日":
			finishDateInt, err := strconv.Atoi(lineSplit[1])
			if err != nil {
				return "", 0, 0, 0, err
			}

			finishDate = finishDateInt
		case "引き落とし日":
			withdrawalDateInt, err := strconv.Atoi(lineSplit[1])
			if err != nil {
				return "", 0, 0, 0, err
			}
			withdrawalDate = withdrawalDateInt
		case "支払い月":
			laterPaymentMonthInt, err := strconv.Atoi(lineSplit[1])
			if err != nil {
				return "", 0, 0, 0, err
			}
			laterPaymentMonth = laterPaymentMonthInt

		}
	}

	return cardName, finishDate, withdrawalDate, laterPaymentMonth, nil
}
