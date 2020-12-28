package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	HELP     = "help"
	LIST     = "list"
	EXIT     = "exit"
	QUIT     = "quit"
	COMMANDS = "commands"
	NEWGAME  = "newgame"
	LOADGAME = "loadgame" // TODO - implement this
)

var (
	COMMAND_LIST = [...]string{HELP, LIST, EXIT, QUIT, COMMANDS, NEWGAME, LOADGAME}
)

var currentGame *Game

type Game struct {
	CurrentTurn string
	Board       map[string]string
	History     []string
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Chess Bot Started")
	fmt.Println("-----------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		cmd := strings.Replace(text, "\n", "", -1)

		if cmd == LIST || cmd == COMMANDS || cmd == HELP {
			printCommandsList()
		} else if cmd == NEWGAME {
			startNewGame(reader)
		} else if cmd == EXIT || cmd == QUIT {
			break
		}
	}

	fmt.Println("Adios, amigo")
}

func printCommandsList() {
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("-----------------")

	for _, cmd := range COMMAND_LIST {
		fmt.Println(cmd)
	}
	fmt.Println("-----------------")
	fmt.Println("")
}

func startNewGame(reader *bufio.Reader) {

	var history []string
	currentGame = &Game{CurrentTurn: "white", Board: emptyBoard(), History: history}

	moveCounter := 1
	currentGame.History = append(currentGame.History, "")
	for {
		currentGame.printBoard()
		fmt.Print(strconv.Itoa(moveCounter) + ". " + currentGame.CurrentTurn + "> ")
		text, _ := reader.ReadString('\n')
		cmd := strings.Replace(text, "\n", "", -1)
		if cmd == "" {
			continue
		}
		isValid, reason := currentGame.ValidateMove(cmd)
		if !isValid {
			fmt.Println("Invalid move: " + reason)
			continue
		}
		currentGame.History[len(currentGame.History)-1] += cmd + spaces(8-len(cmd))

		if currentGame.CurrentTurn == "white" {
			currentGame.CurrentTurn = "black"
		} else {
			currentGame.CurrentTurn = "white"
			moveCounter++
			currentGame.History = append(currentGame.History, "")
		}
	}
}

func formatMoveHistory(history []string) []string {
	var ret []string
	for v := 0; v < 9; v++ {
		ret = append(ret, "")
	}

	for i, val := range history {
		compacted := i
		for compacted > 8 {
			compacted -= 9
		}

		ret[compacted] = ret[compacted] + strconv.Itoa(i+1) + ". " + val
	}

	return ret
}

func spaces(num int) string {
	ret := ""
	for i := 0; i < num; i++ {
		ret += " "
	}
	return ret
}

func (b *Game) printBoard() {

	movehistory := formatMoveHistory(b.History)
	colMap := [...]string{"A", "B", "C", "D", "E", "F", "G", "H"}
	fmt.Println("                            Move History:")

	fmt.Println("   A B C D E F G H          " + movehistory[0])
	histCounter := 1

	for i := 8; i > 0; i-- {

		rowStr := strconv.Itoa(i)
		fullRow := rowStr + " |"
		for j := 0; j < 8; j++ {
			colStr := colMap[j]
			piece := b.Board[colStr+rowStr]
			if piece == "" {
				fullRow += "_|"
			} else {
				fullRow += piece + "|"
			}
			if j == 7 {
				fullRow += " " + strconv.Itoa(i)
			}
		}
		fmt.Println(fullRow + "       " + movehistory[histCounter])
		histCounter++
	}
	fmt.Println("   A B C D E F G H")
}

func emptyBoard() map[string]string {

	var b map[string]string
	b = map[string]string{
		"A8": BR, "B8": BN, "C8": BB, "D8": BQ, "E8": BK, "F8": BB, "G8": BN, "H8": BR,
		"A7": BP, "B7": BP, "C7": BP, "D7": BP, "E7": BP, "F7": BP, "G7": BP, "H7": BP,
		"A6": "", "B6": "", "C6": "", "D6": "", "E6": "", "F6": "", "G6": "", "H6": "",
		"A5": "", "B5": "", "C5": "", "D5": "", "E5": "", "F5": "", "G5": "", "H5": "",
		"A4": "", "B4": "", "C4": "", "D4": "", "E4": "", "F4": "", "G4": "", "H4": "",
		"A3": "", "B3": "", "C3": "", "D3": "", "E3": "", "F3": "", "G3": "", "H3": "",
		"A2": WP, "B2": WP, "C2": WP, "D2": WP, "E2": WP, "F2": WP, "G2": WP, "H2": WP,
		"A1": WR, "B1": WN, "C1": WB, "D1": WQ, "E1": WK, "F1": WB, "G1": WN, "H1": WR}

	return b
}
