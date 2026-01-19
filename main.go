package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

type ingredients struct {
	Name     string
	quantity int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ing := []ingredients{}

	for {
		fmt.Println("Введите команду")
		if ok := scanner.Scan(); !ok {
		}

		text := scanner.Text()
		command := strings.Fields(text)

		if command[0] == "добавить" {
			if len(command) < 3 {
				fmt.Println("далбоёб")
			} else {
				Double := false
				index := 0
				for i := 0; i < len(ing); i++ {
					if command[1] == ing[i].Name {
						Double = true
						index = i
						break
					}
				}
				// Добавить короч по принципу 1 command[1] назва command[2] кол-во
				quantity, err := strconv.Atoi(command[2])
				if err != nil {
					fmt.Println("строка не число")
					break
				}
				if Double == true {
					ing[index].quantity += quantity
				} else {
					ing = append(ing, ingredients{
						Name:     command[1],
						quantity: quantity,
					})
				}

				fmt.Println("успешно добавлено")
			}
		} else if command[0] == "удалить" {

			found := false
			// удалить по принципу command[1] это назв, command[2] кол-во скок удалить
			for i := 0; i < len(ing); i++ {
				if ing[i].Name == command[1] {
					found = true
					quantity, err := strconv.Atoi(command[2])
					if err != nil {
						fmt.Println(err)
					}

					sum := ing[i].quantity - quantity

					if sum >= 1 {
						ing[i].quantity -= quantity
						break
					} else {
						ing = append(ing[:i], ing[i+1:]...)
						break
					}

				}
			}
			if !found {
				fmt.Println("нечего не было найдено")
			}
		} else if command[0] == "получить" {
			for i := 0; i < len(ing); i++ {
				if command[1] == ing[i].Name {
					pp.Println(ing[i])
				}
			}
		} else if command[0] == "help" || command[0] == "помощь" {
			fmt.Println("Список команд \n добавить - назв, кол-во \n удалить - назв, колво \n получить - назв \n выйти - закрыть приложение")
		} else if command[0] == "выход" {
			os.Exit(0)
		} else if command[0] == "list" {
			pp.Println(ing)
		} else {
			fmt.Println("Команда не распознана, попробуй другие")
		}
	}
}
