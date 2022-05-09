package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"geektrust/fund"
	"geektrust/utils"
	"log"
	"os"
	"strings"
)

type User struct {
	Funds []string
}

type FundsJson struct {
	Funds []Fund
}

type Fund struct {
	Name   string
	Stocks []string
}

var fundsJson FundsJson

var fundsMap map[string][]string = map[string][]string{}

const CURRENT_PORTFOLIO = "CURRENT_PORTFOLIO"
const CALCULATE_OVERLAP = "CALCULATE_OVERLAP"
const ADD_STOCK = "ADD_STOCK"

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")

		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the input file")

		return
	}
	defer file.Close()
	err = json.Unmarshal([]byte(fund.Funds), &fundsJson)
	if err != nil {
		log.Fatalf("error while unmarshal %v", err)
	}
	for _, fund := range fundsJson.Funds {
		fundsMap[fund.Name] = fund.Stocks
	}
	scanner := bufio.NewScanner(file)

	var user = User{Funds: make([]string, 0)}
	for scanner.Scan() {

		args := scanner.Text()
		argList := strings.Fields(args)
		typeOfAction := argList[0]
		switch typeOfAction {
		case CURRENT_PORTFOLIO:
			{
				funds := argList[1:]
				user.Funds = append(user.Funds, funds...)
				break
			}
		case CALCULATE_OVERLAP:
			{
				fund := argList[1]
				for _, userFund := range user.Funds {
					if userFund == fund {
						continue
					}
					if _, ok := fundsMap[fund]; ok == false {
						fmt.Println("FUND_NOT_FOUND")
						break
					}
					numberOfCommonStocks := len(utils.GetCommonStocks(fundsMap[userFund], fundsMap[fund]))
					numberOfStocksInA := len(fundsMap[userFund])
					numberOfStocksInB := len(fundsMap[fund])
					overlap := CalculateOverlap(numberOfCommonStocks, numberOfStocksInA, numberOfStocksInB)
					if overlap == 0 {
						continue
					}
					fmt.Printf("%s %s %0.2f%%\n", fund, userFund, overlap)
				}
			}

		case ADD_STOCK:
			{
				fund := argList[1]
				stockName := strings.Join(argList[2:], " ")
				fundsMap[fund] = append(fundsMap[fund], stockName)
				break
			}
		}

	}
}

func CalculateOverlap(numberOfCommonStocks, numberOfStocksInA, numberOfStocksInB int) (overlap float64) {
	overlap = (float64(2*numberOfCommonStocks) / float64(numberOfStocksInA+numberOfStocksInB) * 100)
	return overlap
}
