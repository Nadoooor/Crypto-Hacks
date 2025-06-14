package main

import (
	"bufio"
	"go/Crypto-Hacks/BIN"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	inpu, _ := reader.ReadString('\n')
	BIN.BIN2Normal(inpu)
}
