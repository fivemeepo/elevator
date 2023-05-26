package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"test/elevator/sdk"
)

func main() {
	var e sdk.Elevator
	e.Init()
	//e.Up(10)

	// input command
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input: [up/down] [floor]")
	for scanner.Scan() {
		cmd := scanner.Text()
		if strings.Contains(cmd, "up ") {
			floor, err := strconv.Atoi(strings.TrimPrefix(cmd, "up "))
			if err != nil {
				continue
			}
			e.Up(floor)
		} else if strings.Contains(cmd, "down ") {
			floor, err := strconv.Atoi(strings.TrimPrefix(cmd, "down "))
			if err != nil {
				continue
			}
			e.Down(floor)
		} else if strings.Contains(cmd, "goto ") {
			floor, err := strconv.Atoi(strings.TrimPrefix(cmd, "goto "))
			if err != nil {
				continue
			}
			e.Goto(floor)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
