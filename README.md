## MemTest

A terminal memory/spot-the-difference game. Unicode characters are displayed on the screen for a short time. A new set is displayed, and you must choose what symbol has been added to the set.

### Running

`go run main.go`

### Options
`go run main.go -seed 1234` allows you to play the same game over and over.

### Sample
```
$ go run main.go
Seed: 1450651824
memorize:  ☀ ◐ ☭ ♥
3...
```
That last line is a timer. Then:
```
$ go run main.go
Seed: 1450651824
0) ◐ 1) ♥ 2) ☭ 3) ☀ 4) ☍
Enter choice: 4
Correct!
```