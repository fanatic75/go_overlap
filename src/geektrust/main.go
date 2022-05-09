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

var user = User{Funds: make([]string, 0)}
var fundsJson FundsJson

var fundsMap map[string][]string = map[string][]string{}

const CURRENT_PORTFOLIO = "CURRENT_PORTFOLIO"
const CALCULATE_OVERLAP = "CALCULATE_OVERLAP"
const ADD_STOCK = "ADD_STOCK"
const OVERLAP_PERCENTAGE = "OVERLAP_PERCENTAGE"

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
	for _, f := range fundsJson.Funds {
		fundsMap[f.Name] = f.Stocks
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		args := scanner.Text()
		argList := strings.Fields(args)
		action := argList[0]
		switch action {
		case CURRENT_PORTFOLIO:
			{
				updateUserPortfolio(argList[1:])
				break
			}
		case CALCULATE_OVERLAP:
			{
				calculateOverlap(argList[1])
				break
			}

		case ADD_STOCK:
			{
				addStock(argList[1], strings.Join(argList[2:], " "))
				break
			}
		}

	}
}

func updateUserPortfolio(funds []string) {
	for _, fund := range funds {
		user.Funds = append(user.Funds, fund)
	}

}

func calculateOverlap(fund string) {
	for _, f := range user.Funds {
		if f == fund {
			continue
		}
		if _, ok := fundsMap[fund]; ok == false {
			fmt.Println("FUND_NOT_FOUND")
			break
		}
		commonStocks := utils.Intersection(fundsMap[f], fundsMap[fund])
		overlapValue := (float64(2*len(commonStocks)) / float64(len(fundsMap[f])+len(fundsMap[fund])) * 100)
		if overlapValue == 0 {
			continue
		}
		fmt.Printf("%s %s %0.2f%%\n", fund, f, overlapValue)
	}

}

func addStock(fund, stock string) {
	fundsMap[fund] = append(fundsMap[fund], stock)
}
