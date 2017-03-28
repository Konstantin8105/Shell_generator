package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Konstantin8105/Shell_generator/shellGenerator"
)

func main() {
	// menu - user interface //
	fmt.Println("|==================================|")
	fmt.Println("| Software:                        |")
	fmt.Println("| Create fem model of cylinder     |")
	fmt.Println("| with/without stiffiners in       |")
	fmt.Println("| INP format.                      |")
	fmt.Println("|                                  |")
	fmt.Println("| Created by : Konstantin.I        |")
	fmt.Println("| License : MIT. Absolute free to  |")
	fmt.Println("| use, copy, modify.               |")
	fmt.Println("|==================================|")

	// filename
	filename := "test.inp"
	// height
	height := 5.0
	// diameter
	diameter := 2.0
	// precision
	precision := 0.2
	// amountVertStiff
	amountVertStiff := 6
	// amountHorizStiff
	amountHorizStiff := 2
	// height of stiffiners
	stifHeight := 0.2

	for {
		// show menu //
		fmt.Println("|==================================|")
		fmt.Println("|  Key  | Action                   |")
		fmt.Println("|==================================|")
		fmt.Printf("|   1   | Filename                 : %v\n", filename)
		fmt.Printf("|   2   | Height of shell          : %3.3v\n", height)
		fmt.Printf("|   3   | Diameter of shell        : %3.3v\n", diameter)
		fmt.Printf("|   4   | Precision                : %3.3v\n", precision)
		fmt.Printf("|   5   | Amount vert. stiffiners  : %3v\n", amountVertStiff)
		fmt.Printf("|   6   | Amount horiz. stiffiners : %3v\n", amountHorizStiff)
		fmt.Printf("|   7   | Height of stiffiners     : %3.3v\n", stifHeight)
		fmt.Println("|       |                          |")
		fmt.Println("|   C   | Create model             |")
		fmt.Println("|       |                          |")
		fmt.Println("|   Q   | Quit                     |")
		fmt.Println("|==================================|")
		fmt.Println("| Note: Unit of keys 2,3,4,7       |")
		fmt.Println("| is meter.                        |")
		fmt.Println("|==================================|")

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		text = strings.Replace(text, "\n", "", -1)
		text = strings.TrimSpace(text)
		text = strings.ToLower(text)

		if len(text) == 0 || len(text) > 1 {
			fmt.Println("| You enter : ", text)
			fmt.Println("| Please try again or press \"Q\" for exit")
			continue
		}

		switch text {
		case "1":
			fmt.Printf("| Enter filename (%v): \n", filename)
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				continue
			}
			text = strings.Replace(text, "\n", "", -1)
			text = strings.TrimSpace(text)
			if len(text) == 0 {
				fmt.Println("You don`t enter filename")
				continue
			}
			filename = text
		case "2":
			fmt.Printf("| Enter height of shell (%.3v meter):\n", height)
			value, err := getFloat()
			if err != nil {
				fmt.Println(err)
				continue
			}
			height = value
		case "3":
			fmt.Printf("| Enter diameter of shell (%.3v meter):\n", diameter)
			value, err := getFloat()
			if err != nil {
				fmt.Println(err)
				continue
			}
			diameter = value
		case "4":
			fmt.Printf("| Enter precision - distance between points (%.3v meter):\n", precision)
			value, err := getFloat()
			if err != nil {
				fmt.Println(err)
				continue
			}
			precision = value
		case "5":
			fmt.Printf("| Enter amount of vertical stiffiners (%3v stiffiners):\n", amountVertStiff)
			value, err := getInt()
			if err != nil {
				fmt.Println(err)
				continue
			}
			amountVertStiff = value
		case "6":
			fmt.Printf("| Enter amount of horizontal stiffiners (%3v stiffiners):\n", amountHorizStiff)
			value, err := getInt()
			if err != nil {
				fmt.Println(err)
				continue
			}
			amountHorizStiff = value
		case "7":
			fmt.Printf("| Enter height of stiffiners (%.3v meter):\n", stifHeight)
			value, err := getFloat()
			if err != nil {
				fmt.Println(err)
				continue
			}
			stifHeight = value
		case "c":
			var model shellGenerator.ShellWithStiffiners
			err := model.AddShell(shellGenerator.Shell{Height: height, Diameter: diameter, Precision: precision})
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = model.AddStiffiners(amountHorizStiff, amountVertStiff, stifHeight, precision)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = model.GenerateINP(filename)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "q":
			return
		default:
			fmt.Println("| You enter : ", text)
			fmt.Println("| Please try again or press \"Q\" for exit")
			continue
		}
	}
}

func getInt() (i int, err error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	text = strings.Replace(text, "\n", "", -1)
	text = strings.TrimSpace(text)
	if len(text) == 0 {
		fmt.Println("You don`t enter value")
		return 0, err
	}
	value, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int(value), nil
}

func getFloat() (f float64, err error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	text = strings.Replace(text, "\n", "", -1)
	text = strings.TrimSpace(text)
	if len(text) == 0 {
		fmt.Println("You don`t enter value")
		return 0, err
	}
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return value, nil
}
