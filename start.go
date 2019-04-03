package main

import (
    "bufio"
    "fmt"
    "os"
)
var  playerNamePointCounter = 0
var  computerPointCounter = 0

type Rule int
const (
 Rock Rule = 1 + iota
 Paper
 Scissors
)


func main() {
    fmt.Println("Welcome to Rock Paper Scissors in Golang World")
    fmt.Println("Enter your name : ")

    buf := bufio.NewReader(os.Stdin)
    fmt.Print(">")
    playerName, err := buf.ReadBytes('\n')
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Ok "+ string(playerName) +"Let's play")

    for sumGamePoint(playerNamePointCounter, computerPointCounter) < 10 {
        play()
    }

    if playerNamePointCounter > computerPointCounter {
        fmt.Println("Ok player win")
    } else if computerPointCounter > playerNamePointCounter {
        fmt.Println("Ok computer win")
    } else {
        fmt.Println("Problem detected")
    }
    fmt.Println("End of game")
}

func sumGamePoint(playerTotal int, computerTotal int) (int) {
    total := playerTotal + computerTotal
    fmt.Println("Total :", total)
    return total
}

func play() {
    fmt.Println("Play")
    playerNamePointCounter += 1
    playerPlay()
}

func playerPlay() {
    fmt.Println("Your turn : ")
    buf := bufio.NewReader(os.Stdin)
    fmt.Print(">")
    playerInput, err := buf.ReadBytes('\n')
        if err != nil {
            fmt.Println(err)
        }
}

func computerPlay() {

}


