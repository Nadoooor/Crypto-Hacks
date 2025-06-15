package main

import (
	"bufio"
	"go/Crypto-Hacks/Caesar"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	inpu, _ := reader.ReadString('\n')
	reader2 := bufio.NewReader(os.Stdin)
	Rot, _ := reader2.ReadString('\n')
	Rot = Rot[:len(Rot)-1] 
	RotInt := 50         
	Caesar.Normal2Cipher(inpu, RotInt)
}
