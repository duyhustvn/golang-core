package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
)

func createMenu() {
	fmt.Println("********************************")
	fmt.Println("1. Read file")
	fmt.Println("2. Traversal the list")
	fmt.Println("3. Search classes")
	fmt.Println("4. Delete classes")
	fmt.Println("5. Order classes")
	fmt.Println("6. Count classes")
	fmt.Println("0. Exit")
	fmt.Println("********************************")
}

func main() {
	createMenu()
	var schedules *list.List
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Choose the option: ")
		scanner.Scan()
		text := scanner.Text()

		switch text {
		case "1":
			fmt.Println("Read file")
			filePath := "schedule.txt"
			f, err := os.Open(filePath)
			defer f.Close()
			if err != nil {
				fmt.Printf("Error while reading file %+v\n", err)
				break
			}
			schedules, err = readfile(bufio.NewReader(f))
			if err != nil {
				return
			}
			break
		case "2":
			fmt.Println("Traversal the list")
			traverseList(schedules)
			break
		case "3":
			fmt.Println("Search classes")
			break
		case "4":
			fmt.Println("Delete classes")
			break
		case "5":
			fmt.Println("Order classes")
			break
		case "6":
			fmt.Println("Count classes")
			break
		case "0":
			fmt.Println("Exit program")
			os.Exit(1)
			break
		default:
			fmt.Println("Invalid options. Please choose correct option")
			break
		}
	}
}

type Schedule struct {
	ClassCode  string
	CourseCode string
	Location   string
}

func readfile(r *bufio.Reader) (*list.List, error) {
	list := list.New()
	count := 1
	schedule := Schedule{}
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			switch err {
			case io.EOF:
				return list, nil
			default:
				fmt.Println("[readfile] ReadLine err ", err)
				return nil, err
			}
		}
		if count%3 == 1 {
			schedule.ClassCode = string(line)
		}
		if count%3 == 2 {
			schedule.CourseCode = string(line)
		}
		if count%3 == 0 {
			schedule.Location = string(line)
			list.PushBack(schedule)
			schedule = Schedule{}
		}
		count++
	}
}

func traverseList(schedules *list.List) {
	schedule := schedules.Front()
	for schedule != nil {
		v := schedule.Value.(Schedule)
		fmt.Println(v)
		schedule = schedule.Next()
	}
}
