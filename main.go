package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	// "text/scanner"
	// "os"
	// "bufio"
	// "strconv"
	// "strings"
	//"HACKATHON-002/command-and-control"
)

var start string

func calc() {
	fmt.Println("Welcome to the Calculator")
	fmt.Println("Enter 'help' for guidelines")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		text := scanner.Text()
		text = strings.TrimSpace(text)

		if text == "help" {
			fmt.Println("calc add <a> <b>      → addition")
			fmt.Println("calc sub <a> <b>      → subtraction")
			fmt.Println("calc mul <a> <b>      → multiplication")
			fmt.Println("calc div <a> <b>      → division")
			fmt.Println("calc mod <a> <b>      → remainder")
			fmt.Println("calc pow <a> <b>      → a to the power of b")
			fmt.Println("calc last             → show the last result")
			fmt.Println("calc history          → show last 5 calc results")
			continue

		}

		part := strings.Fields(text)

		if len(part) > 3 {
			fmt.Println("too many arguments")
			continue
		}

		if len(part) < 3 {
			fmt.Println("not enough argument")
			continue
		}

		command := part[0]

		a, err1 := strconv.ParseFloat(part[1], 64)
		b, err2 := strconv.ParseFloat(part[2], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("not a valid digit")
			continue
		}

		switch command {
		case "add":
			fmt.Println("Result: ", a+b)
			continue
		case "sub":
			fmt.Println("Result: ", a-b)
			continue
		case "mul":
			fmt.Println("Result: ", a*b)
			continue
		case "div":
			if b == 0 {
				fmt.Println(a, "Cannot be divided by ", b)
				continue
			}
			fmt.Println("Result: ", a/b)
			continue
		case "pow":
			fmt.Println("Result: ", math.Pow(a, b))
			continue
		case "last":
			fmt.Println()
			continue
		case "calc history":
			fmt.Println()
			continue
		default:
			fmt.Println("Invalid command")

		}

	}

}

func Help() {
	fmt.Println("calc   <command>   → the calculator")
	fmt.Println("base   <command>   → the base converter")
	fmt.Println(" str    <command>   → the string transformer")
	fmt.Println("help → shows all commands")
	fmt.Println("history → shows last 10 inputs")
	fmt.Println("exit → shuts down the console")

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter Help to begin")

		scanner.Scan()
		text := scanner.Text()
		text = strings.TrimSpace(text)
		text = strings.ToLower(text)

		if text == "help" {
			Help()
			continue
		}
		if text == "calc" {
			calc()
			continue
		}
		if text == "base" {
			base()
			continue
		}
	}
}

func base() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)
		if input == "help" {
			fmt.Println(" base dec <number>   → convert decimal to binary AND hex")
			fmt.Println(" base hex <number>   → convert hex to decimal")
			fmt.Println(" base bin <number>   → convert binary to decimal")
			continue
		}
		if input == "exit" {
			fmt.Println("goodbye")
			break
		}
		if input == "" {
			fmt.Println("error")
			continue
		}
		conv := strings.Fields(input)
		value := conv[1]
		num := conv[2]

		if conv[0] != "base" {
			fmt.Println("invalid command")
			continue
		}

		switch num {
		case "hex":
			n, err := strconv.ParseInt(value, 16, 64)
			if err != nil {
				fmt.Println("incorrect input")
				continue
			}
			fmt.Println("Decimal: ", n)
		case "bin":
			n, err := strconv.ParseInt(value, 2, 64)
			if err != nil {
				fmt.Println("incorrect input")
				continue
			}
			fmt.Println("Decimal: ", n)
		case "dec":
			n, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				fmt.Println("incorrect input")
				continue
			}
			fmt.Printf("Binary: %b\n", n)
			fmt.Printf("Hex: %X\n", n)
		}
	}
}
