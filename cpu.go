package main

import (
	"fmt"
	"log"
	"os/exec"
)

const (
	vendor  = "vendor"
	model   = "model"
	freq    = "frequency"
	dcache  = "L1d cache"
	icache  = "L1i cache"
	l2cache = "L2 cache"
	l3cache = "L3 cache"
)

func specGetVendor() string {
	out, err := exec.Command("data").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	return "intel"
}

func cpuParseVendor() {}
