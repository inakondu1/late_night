      // CodeCrafters — Hackathon 002
         // Squad: [The Structs]
         // Member: [ina isah, iisah]
         // Member: [omafu samuel, somafu]
         // Member: [Otete Benjamin, botete]
         // Member: [Odeh owoicho emmanuel, emmodeh]
         // Member: [Wealth Ibe, wibe]
         // Member: [Odumu John, jodumu]
         // Member: [Edward John, edwjohn]
         // Member: [inaju Daniel, dinaju]


package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var start string

func calc() {
	fmt.Println("Welcome to the Calculator")
	fmt.Println("Enter 'help' for guidelines")

	scanner := bufio.NewScanner(os.Stdin)
	var history []string

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
			fmt.Println("Exit")

			continue

		}

		if text == "history" {
			fmt.Println("Operation History")
			for i, op := range history {
				fmt.Printf("[%d] %s\n", i+1, op)
			}
			if len(history) == 0 {
				fmt.Println("No history yet")
			} else {
				fmt.Println("--------")
				continue
			}
		}
		if text == "last" {
			for i := range history {
				history[i] = history[len(history)-1]
				continue
			}
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
			result := a + b
			fmt.Println("Result: ", a+b)
			history = append(history, fmt.Sprintf("%.0f + %.0f = %.0f", a, b, result))
			continue
		case "sub":
			result := a - b
			fmt.Println("Result: ", a-b)
			history = append(history, fmt.Sprintf("%.0f - %.0f = %.0f", a, b, result))
			continue
		case "mul":
			result := a * b
			fmt.Println("Result: ", a*b)
			history = append(history, fmt.Sprintf("%.0f * %.0f = %.0f", a, b, result))
			continue
		case "div":
			if b == 0 {
				fmt.Println(a, "Cannot be divided by ", b)
				continue
			}
			result := a / b
			fmt.Println("Result: ", a/b)
			history = append(history, fmt.Sprintf("%.0f / %.0f = %.0f", a, b, result))

			continue
		case "pow":
			result := math.Pow(a, b)
			fmt.Println("Result: ", math.Pow(a, b))
			history = append(history, fmt.Sprintf("%.0f ^ %.0f = %.0f", a, b, result))
			continue
		case "last":
			result := history
			fmt.Println("last history")
			history = append(history, fmt.Sprintf("%.0f ^ %.0f = %.0f", a, b, result))
			continue
		case "calc history":
			fmt.Println()
			continue
		case "exit":
			main()
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
	fmt.Println("\033[32;1m ════════════════════════════════════════════════")
	fmt.Println("SENTINEL — COMMAND & CONTROL CONSOLE")
	fmt.Println("All systems nominal. Type 'help' to begin.")
	fmt.Println("════════════════════════════════════════════════")
	fmt.Print("C&C>\033[0m ")
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
			baseConverter()
			continue
		}
		if text == "str" {
			mainstring()
			continue
		}
	}
}

func baseConverter() {
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
		case "exit":
			main()
		default:
			fmt.Println("Invalid base!")
		}
	}
}

func reverse(word string) string {
	if len(word) == 0 {
		return ""
	}

	return reverse(word[1:]) + word[:1]
}

func snakecase(word string) string {
	return strings.ReplaceAll(word, " ", "_")
}

func mainstring() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter a word")

		input, _ := reader.ReadString('\n')
		word := strings.Fields(input)

		var Transformation string

		fmt.Scanln(&word)

		if len(word) == 0 {
			fmt.Println("no text provided")
			continue
		}

		fmt.Println("chose a Transformation:")
		fmt.Println("<1> Toupper")
		fmt.Println("<2> Tolower")
		fmt.Println("<3> Fields")
		fmt.Println("<4> Title")
		fmt.Println("<5> Snake")
		fmt.Println("<6> Reverse")
		fmt.Println("<7> Exit")

		fmt.Scanln(&Transformation)

		switch Transformation {
		case "1":
			result := word
			input := strings.Join(result, " ")
			fmt.Println("Result:", strings.ToUpper(input))
		case "2":
			result := word
			input := strings.Join(result, " ")

			fmt.Println("Return:", strings.ToLower(input))
		case "3":
			result := word
			input := strings.Join(result, " ")
			fmt.Println("Return:", strings.Fields(input))
		case "4":
			result := word
			input := strings.Join(result, " ")
			fmt.Println("Return:", strings.Title(input))
		case "5":
			result := word
			input := strings.Join(result, " ")
			result1 := strings.ToLower(input)
			fmt.Println(snakecase(result1))

		case "6":
			good := strings.Fields(input)
			for i := 0; i < len(good); i++ {
				good[i] = reverse(good[i])
			}
			fmt.Println(strings.Join(good, " "))
		case "7":
			main()
		default:
			fmt.Println("invalide input")

		}
	}
}
