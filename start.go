package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)
var  playerNamePointCounter = 0
var  computerPointCounter = 0

type Stroke int
const (
 Rock Stroke = iota
 Paper
 Scissors
)

var strokes = [...]string{
    "Rock",
    "Paper",
    "Scissors",
}

func (s Stroke) String() string { return strokes[s] }

func main() {
    computerPlay()
    fmt.Println("Welcome to Rock Paper Scissors in Golang World")
    fmt.Println("Enter your name : ")

    buf := bufio.NewReader(os.Stdin)
    fmt.Print(">")
    playerNameAsBytes, err := buf.ReadBytes('\n')
    if err != nil {
        fmt.Println(err)
        return
    }
    playerName := string(playerNameAsBytes)
    fmt.Println("Ok "+ playerName +"Let's play")

    for sumGamePoint(playerNamePointCounter, computerPointCounter) < 10 {
        play()
        fmt.Println()
    }

    fmt.Println("Player's points", playerNamePointCounter)
    fmt.Println("Computer's points", computerPointCounter)

    if playerNamePointCounter > computerPointCounter {
        fmt.Println("Ok" + playerName +"you win this time")
    } else if computerPointCounter > playerNamePointCounter {
        fmt.Println("Ok computer win")
    } else if computerPointCounter == playerNamePointCounter {
        fmt.Println("Wow equality !")
    } else {
            fmt.Println("Problem detected")
    }
    fmt.Println("End of game")
}

func sumGamePoint(playerTotal int, computerTotal int) int {
    total := playerTotal + computerTotal
    //fmt.Println("Total :", total)
    return total
}

func play() {
    fmt.Println("Play : 0 Rock, 1 Paper, 2 Scissors")
    fmt.Println("Your turn : ")
    buf := bufio.NewReader(os.Stdin)
    fmt.Print(">")
    playerInputAsBytes, err := buf.ReadBytes('\n')
    if err != nil {
        fmt.Println(err)
        return
    }
    playerInputAsString := string(playerInputAsBytes)
    playerInputAsString = strings.TrimSuffix(playerInputAsString, "\n")
    playerInput, err := strconv.Atoi(playerInputAsString)

    if err != nil {
        fmt.Println(err)
        return
    }
    // Input must be between 0 and 2
    if playerInput > 2 {
        fmt.Println("Invalid stroke")
        return
    }

    fmt.Println("You play ", Stroke(playerInput))
    computerStroke := computerPlay()
    fmt.Println("Computer play ", Stroke(computerStroke))
    checkTurnStrokes(playerInput, computerStroke)
}

func computerPlay() int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(3 - 0) + 0
}

func checkTurnStrokes(playerStroke int, computerStroke int) {
    if playerStroke > computerStroke {
        if playerStroke == 2 && computerStroke == 0 {
            computerPointCounter += 1
            fmt.Println("Computer win this turn")
        } else {
            playerNamePointCounter += 1
            fmt.Println("Player win this turn")
        }
    } else if playerStroke < computerStroke {
        if computerStroke == 2 && playerStroke == 0 {
            playerNamePointCounter += 1
            fmt.Println("Player win this turn")
        } else {
            computerPointCounter += 1
            fmt.Println("Computer win this turn")
        }
    } else {
        fmt.Println("Equality")
    }
}
