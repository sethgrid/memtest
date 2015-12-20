package main

/*
   MemTest is a simple game that tests your ability to spot the difference in a set of characters.
   A series of unicode symbols will display for a short time.
   A the same symbols (plus one more) will display, and you must identify the new symbol

   Your terminal must respect ANSI escape codes to allow erasing of lines.
   Otherwise, the initial symbols to memorize are not erased and thus no skill is involved.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sethgrid/curse"
)

// all options
var options = []rune{'⚓', '☀', '☆', '☂', '☻', '♞', '☯', '☭', '☘', '☊', '☋', '♣', '♥', '♦', '♠', '♫', '☠', '♜', '♰', '☿', '✈', '∴', '∑', '㆛', '☍', '◐'}

func main() {
	// Seed allows us to play the same game over and over
	var Seed int64
	flag.Int64Var(&Seed, "seed", time.Now().Unix(), "used to create the same game")
	flag.Parse()

	fmt.Printf("Seed: %d\n", Seed)
	rand.Seed(Seed)

	if err := run(Seed); err != nil {
		fmt.Printf("error running game - %q", err.Error())
	}
}

// run displays options, collects your response, and loops until you guess wrong
func run(seed int64) error {
	round := 0
	for {
		round++
		roundOptions := getOptions(round + 3)

		showTimedStarting(roundOptions)
		newOptions, solutionIndex := addOption(roundOptions)

		choice := showTimedOptions(newOptions)

		if choice != solutionIndex {
			fmt.Printf("Incorrect. You guessed %d, but it was %d.\nYou made it to round %d.", choice, solutionIndex, round)
			os.Exit(0)
		}

		if round+3 == len(options)-1 {
			fmt.Println("Wow, you won the game! What a champ!")
			os.Exit(0)
		}

		fmt.Println("Correct!")
		time.Sleep(1 * time.Second)

		c, err := curse.New()
		if err != nil {
			fmt.Println("unable to init curse lib - %q", err.Error())
			os.Exit(1)
		}
		c.MoveUp(1).EraseCurrentLine() // erase correct message
		c.MoveUp(1).EraseCurrentLine() // erase choice entry
		c.MoveUp(1).EraseCurrentLine() // erase option choices
	}
}

// getOptions returns i options/symbols
func getOptions(i int) []rune {
	theseOpts := shuffle(options)
	return theseOpts[:i]
}

// shuffle returns a shuffled slice of rune
func shuffle(rs []rune) []rune {
	a := make([]rune, len(rs))
	for i, _ := range rs {
		a[i] = rs[i]
	}
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return a
}

// addOption adds a unique symbol and its index to the runes passed at a random location
func addOption(rs []rune) ([]rune, int) {
	// copy for return
	opts := make([]rune, len(rs))
	for j, _ := range rs {
		opts[j] = rs[j]
	}

	// get pool of opts
	allOpts := shuffle(options)

	// get a unique opt
	i := -1
	var newOpt rune
	for {
		i++
		isNew := true
		for _, val := range rs {
			if val == allOpts[i] {
				isNew = false
				continue
			}
		}
		if isNew {
			opts = append(opts, allOpts[i])
			newOpt = allOpts[i]
			break
		}
	}

	// find where the new item was added
	newIdx := 0
	opts = shuffle(opts)
	for i, val := range opts {
		if val == newOpt {
			newIdx = i
		}
	}

	return opts, newIdx
}

// showTimedStarting displays the initial option symbols for a given time
func showTimedStarting(rs []rune) {
	c, err := curse.New()
	if err != nil {
		fmt.Printf("error creating curser - %q", err.Error())
		os.Exit(1)
	}
	display := ""
	for _, r := range rs {
		display += fmt.Sprintf(" %s", string(r))
	}
	fmt.Printf("memorize: %s\n", display)
	for i := 5; i > 0; i-- {
		fmt.Printf("%d...", i)
		time.Sleep(1 * time.Second)
		// erase the previous timer time
		c.EraseCurrentLine()
	}
	// erase the starting memorize line
	c.MoveUp(1).EraseCurrentLine()
}

// showTimedOptions asks for the user input. Currently untimed.
func showTimedOptions(opts []rune) int {
	reader := bufio.NewReader(os.Stdin)
	choices := ""
	for i, val := range opts {
		choices += fmt.Sprintf("%d) %s ", i, string(val))
	}
	fmt.Printf("%v\nEnter choice: ", choices)
	choice, _ := reader.ReadString('\n')

	i, err := strconv.Atoi(strings.TrimSpace(choice))
	if err != nil {
		fmt.Printf("exiting - please enter a number for your choice - %q\n", err.Error())
		os.Exit(2)
	}

	return i
}
