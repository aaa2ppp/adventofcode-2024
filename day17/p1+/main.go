package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	. "adventofcode-2024/utils"
)

type (
	opcode byte
	combo  byte
)

const (
	ADV opcode = 0
	BXL opcode = 1
	BST opcode = 2
	JNZ opcode = 3
	BXC opcode = 4
	OUT opcode = 5
	BDV opcode = 6
	CDV opcode = 7
)

func (o opcode) String() string {
	return []string{"ADV", "BXL", "BST", "JNZ", "BXC", "OUT", "BDV", "CDV"}[o]
}

const (
	RA combo = 4
	RB combo = 5
	RC combo = 6
	RX combo = 7
)

func (c combo) String() string {
	switch c {
	case RA:
		return "RA"
	case RB:
		return "RB"
	case RC:
		return "RC"
	case RX:
		return "RX"
	default:
		return strconv.Itoa(int(c))
	}
}

type VM struct {
	RA          int
	RB          int
	RC          int
	Code        []byte
	pos         int
	instruction []func(op byte)
	out         []byte
}

func NewVM() *VM {
	vm := new(VM)
	vm.instruction = []func(op byte){
		ADV: vm.ADV,
		BXL: vm.BXL,
		BST: vm.BST,
		JNZ: vm.JNZ,
		BXC: vm.BXC,
		OUT: vm.OUT,
		BDV: vm.BDV,
		CDV: vm.CDV,
	}
	return vm
}

func (vm *VM) ADV(op byte) {
	vm.RA >>= vm.Combo(op)
	vm.pos += 2
}

func (vm *VM) BXL(op byte) {
	vm.RB ^= int(op)
	vm.pos += 2
}

func (vm *VM) BST(op byte) {
	vm.RB = vm.Combo(op) & 7
	vm.pos += 2
}

func (vm *VM) JNZ(op byte) {
	if vm.RA == 0 {
		vm.pos += 2
		return
	}
	vm.pos = int(op)
}

func (vm *VM) BXC(op byte) {
	vm.RB ^= vm.RC
	vm.pos += 2
}

func (vm *VM) OUT(op byte) {
	vm.out = append(vm.out, byte(vm.Combo(op)&7))
	vm.pos += 2
}

func (vm *VM) BDV(op byte) {
	vm.RB = vm.RA >> vm.Combo(op)
	vm.pos += 2

}

func (vm *VM) CDV(op byte) {
	vm.RC = vm.RA >> vm.Combo(op)
	vm.pos += 2
}

func (vm *VM) Bytes() []byte {
	return vm.out
}

func (vm *VM) Combo(op byte) int {
	switch combo(op) {
	case RA:
		return vm.RA
	case RB:
		return vm.RB
	case RC:
		return vm.RC
	case RX:
		panic("Combo 7")
		return 0
	default:
		return int(op)
	}
}

func (vm *VM) Run() []byte {
	vm.out = vm.out[:0]
	vm.pos = 0

	for 0 <= vm.pos && vm.pos < len(vm.Code) {
		opcode := opcode(vm.Code[vm.pos])
		operand := vm.Code[vm.pos+1]
		if debugEnable {
			switch opcode {
			case BXL, JNZ:
				log.Printf("%02x: %v %v \t// RA=%d RB=%d RC=%d", vm.pos, opcode, operand, vm.RA, vm.RB, vm.RC)
			default:
				log.Printf("%02x: %v %v \t// RA=%d RB=%d RC=%d", vm.pos, opcode, combo(operand), vm.RA, vm.RB, vm.RC)
			}
		}
		vm.instruction[opcode](operand)
	}

	return vm.Bytes()
}

func parseRegister(s string) (byte, int, error) {
	if !strings.HasPrefix(s, "Register ") {
		return 0, 0, errors.New(`expected "Register "`)
	}
	s = strings.TrimPrefix(s, "Register ")
	r := s[0]
	s = s[1:]

	if !strings.HasPrefix(s, ": ") {
		return 0, 0, errors.New(`expected ": "`)
	}
	s = strings.TrimPrefix(s, ": ")
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return r, v, nil
}

func parseProgram(s string) ([]byte, error) {
	if !strings.HasPrefix(s, "Program: ") {
		return nil, errors.New(`expected "Program: "`)
	}
	s = strings.TrimPrefix(s, "Program: ")

	var code []byte
	for _, s := range strings.Split(s, ",") {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		code = append(code, byte(v))
	}

	return code, nil
}

func run(in io.Reader, out io.Writer) {
	// br := bufio.NewReader(in)
	sc := bufio.NewScanner(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	vm := NewVM()

	// scan registers

	for sc.Scan() {
		b := sc.Bytes()
		if len(b) == 0 {
			break
		}

		r, v, err := parseRegister(UnsafeString(b))
		if err != nil {
			panic(err)
		}

		switch r {
		case 'A':
			vm.RA = v
		case 'B':
			vm.RB = v
		case 'C':
			vm.RC = v
		default:
			panic(fmt.Errorf("unknown register '%c'", r))
		}
	}

	// scan program code

	sc.Scan()
	if err := sc.Err(); err != nil {
		panic(err)
	}

	code, err := parseProgram(UnsafeString(sc.Bytes()))
	if err != nil {
		panic(err)
	}
	vm.Code = code

	if debugEnable {
		log.Printf("vm: %+v", vm)
	}

	ans := vm.Run()
	WriteInts(bw, ans, WriteOpts{Sep: ',', End: '\n'})
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
