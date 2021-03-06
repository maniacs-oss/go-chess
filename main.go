package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var debug = flag.Bool("debug", false, "To switch on debugging")

func main() {
	fmt.Println()
	flag.Parse()
	Debug = *debug
	AllInit()
	// Test()
	fmt.Println("Let us play chess!")
	fmt.Println("Enter c to play 2 players game, enter e to play with chess engine")
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		switch s.Text() {
		case "c":
			var board Board
			var info SearchInfo
			err := ParseFEN(StartFEN, &board)
			if err != nil {
				fmt.Println("Error in parsing fen: ", err)
			}
			board.Print()
			(&board).PvTable.Init()
			fmt.Printf("Please enter a move > ")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				command := scanner.Text()
				switch command {
				case "q":
					os.Exit(1)
				case "t":
					if board.Ply == 0 {
						fmt.Println("There are no moves made on the board")
					} else {
						err = TakeMove(&board)
						if err != nil {
							fmt.Errorf("Error in taking move back", err)
						}
						board.Print()
					}
					fmt.Printf("Please enter a move > ")
				case "p":
					PerftTest(3, &board)
					fmt.Printf("Please enter a move > ")
				case "pv":
					count, _ := GetPvLine(4, &board)
					for i := 0; i < count; i++ {
						move := board.PvArray[i]
						fmt.Println(PrintMove(move))
					}
					fmt.Printf("Please enter a move > ")
				case "s":
					info.depth = 6
					info.startTime = time.Now()
					duration, _ := time.ParseDuration("200000ms")
					info.stopTime = time.Now().Add(duration)
					SearchPosition(&board, &info)
				default:
					move, err := ParseMove(strings.ToLower(command), &board)
					if err != nil {
						fmt.Errorf("Error in parsing move", err)
					}
					if move != NoMove {
						StorePvMove(&board, move)
						_, err = MakeMove(move, &board)
						if err != nil {
							fmt.Errorf("Error in making move", err)
						}
					} else {
						fmt.Println("Move not parsed")
					}
					board.Print()
					fmt.Printf("Please enter a move > ")
				}

			}
		case "e":
			UciLoop()
		}
	}

}
