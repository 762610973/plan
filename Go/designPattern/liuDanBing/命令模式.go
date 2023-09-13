package main

import (
	"fmt"
)

type command interface {
	treat()
}

type nurse struct {
	cmdList []command
}

func (n *nurse) notify() {
	if n.cmdList == nil || len(n.cmdList) == 0 {
		return
	}

	for _, cmd := range n.cmdList {
		cmd.treat()
	}
}

func (n *nurse) join(cmds ...command) {
	if n.cmdList == nil {
		n.cmdList = make([]command, 0, len(cmds))
	}
	for _, c := range cmds {
		n.cmdList = append(n.cmdList, c)
	}
}

type doctor struct{}

func (d doctor) treatEye() {
	fmt.Println("treat eye")
}

func (d doctor) treatNose() {
	fmt.Println("treat nose")
}

type commandTreatEye struct {
	doctor doctor
}

func (c commandTreatEye) treat() {
	c.doctor.treatEye()
}

type commandTreatNose struct {
	doctor doctor
}

func (c commandTreatNose) treat() {
	c.doctor.treatNose()
}

func main() {
	d := doctor{}
	cmdEye := commandTreatEye{doctor: d}
	cmdEye.treat()

	cmdNose := commandTreatNose{doctor: d}
	cmdNose.treat()

	n := nurse{cmdList: []command{cmdEye, cmdNose}}
	n.notify()

	ns := nurse{}
	ns.join(cmdEye, cmdNose)
	ns.notify()
}
