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

	for scanner.Scan() {

		args := scanner.Text()
		argList := strings.Fields(args)
		action := argList[0]
		switch action {
		case CURRENT_PORTFOLIO:
			{
				funds := argList[1:]
				user.Funds = append(user.Funds, funds...)
				break
			}
		case CALCULATE_OVERLAP:
			{
				fund := argList[1]
				user.CalculateOverlap(fund)
				break
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

func (user *User) CalculateOverlap(fund string) {
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
