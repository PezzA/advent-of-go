package day12

import "strconv"
import "github.com/pezza/advent-of-go/Assembunny"

func makeRegisters() Assembunny.Registers {
	reg := make(Assembunny.Registers)

	reg["a"] = 0
	reg["b"] = 0
	reg["c"] = 0
	reg["d"] = 0

	return reg
}

// PartOne solves day 12 part one.
func PartOne(input string) (string, error) {
	prog := Assembunny.GetProgram(PuzzleInput())
	registers := makeRegisters()

	result := Assembunny.RunProgram(prog, registers, "a")

	return strconv.Itoa(result), nil
}

// PartTwo solves day 12 part one.
func PartTwo(input string) (string, error) {
	prog := Assembunny.GetProgram(PuzzleInput())
	registers := makeRegisters()

	registers["c"] = 1
	result := Assembunny.RunProgram(prog, registers, "a")

	return strconv.Itoa(result), nil
}
