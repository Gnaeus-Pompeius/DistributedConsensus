package main


import (
	"fmt"
)

type LogEntry struct {
	index uint64
	commandName string
	term uint64
}

type Server struct {
	//id string
	//role string
	state string
	currentTerm int
	votedFor string
	log []LogEntry
	nextIndex []int
	matchIndex []int
}



func main() {
	fmt.Println("Hello, World!")
}

func printLog() {
	//TODO
}

func printInfo() {
	//TODO
}

func suspend() {
	//TODO
}

func resume() {
	//TODO
}

func recieveMessage() {
	//TODO
}

func sendMessage() {
	//TODO
}

func appendEntries() {
	//TODO
}

func requestVote() {
	//TODO
}

func sendHeartbeat() {
	//TODO
}

func logTransaction() {
	//TODO
}

func listenToPort() {
	//TODO
}



