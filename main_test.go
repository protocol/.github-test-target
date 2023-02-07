package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandomSleep(t *testing.T) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	d := time.Duration(10+rand.Intn(20)) * time.Second
	fmt.Println("Sleeping for", d)
	time.Sleep(d)
}

func TestMult(t *testing.T) {
	_ = mult(3, 4)
}
