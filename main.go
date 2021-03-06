package main

import (
	"fmt"

	"github.com/pezza/advent-of-go/day01"
	"github.com/pezza/advent-of-go/day02"
	"github.com/pezza/advent-of-go/day03"
	"github.com/pezza/advent-of-go/day04"
	"github.com/pezza/advent-of-go/day05"
	"github.com/pezza/advent-of-go/day06"
	"github.com/pezza/advent-of-go/day07"
	"github.com/pezza/advent-of-go/day09"
	"github.com/pezza/advent-of-go/day10"
	"github.com/pezza/advent-of-go/day11"
	"github.com/pezza/advent-of-go/day12"
	"github.com/pezza/advent-of-go/day13"
	"github.com/pezza/advent-of-go/day14"
	"github.com/pezza/advent-of-go/day15"
	"github.com/pezza/advent-of-go/day16"
	"github.com/pezza/advent-of-go/day17"
	"github.com/pezza/advent-of-go/day18"
	"github.com/pezza/advent-of-go/day19"
	"github.com/pezza/advent-of-go/day20"
	"github.com/pezza/advent-of-go/day21"
	"github.com/pezza/advent-of-go/day22"
	"github.com/pezza/advent-of-go/day23"
	"github.com/pezza/advent-of-go/day24"
	"github.com/pezza/advent-of-go/day25"

	"os"
	"strconv"
)

type puzzle func(string) (string, error)

type dayRunner interface {
	PartOne() (string, error)
	PartTwo() (string, error)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("USEAGE: advent-of-go <day>")
		fmt.Println("<day> = day number of puzzle to run.  e.g. advent-of-go 2")

		//TODO - need to ensure value is actuall a string
		return
	}

	switch day, _ := strconv.Atoi(os.Args[1]); day {
	case 1:
		runDay(day01.PartOne, day01.PartTwo, day01.PuzzleInput())
	case 2:
		runDay(day02.PartOne, day02.PartTwo, day02.PuzzleInput())
	case 3:
		runDay(day03.PartOne, day03.PartTwo, day03.PuzzleInput())
	case 4:
		runDay(day04.PartOne, day04.PartTwo, day04.PuzzleInput())
	case 5:
		runDay(day05.PartOne, day05.PartTwo, day05.PuzzleInput())
	case 6:
		runDay(day06.PartOne, day06.PartTwo, day06.PuzzleInput())
	case 7:
		runDay(day07.PartOne, day07.PartTwo, day07.PuzzleInput())
	case 9:
		runDay(day09.PartOne, day09.PartTwo, day09.PuzzleInput())
	case 10:
		runDay(day10.PartOne, day10.PartTwo, day10.PuzzleInput())
	case 11:
		runDay(day11.PartOne, day11.PartTwo, day11.PuzzleInput())
	case 12:
		runDay(day12.PartOne, day12.PartTwo, day12.PuzzleInput())
	case 13:
		runDay(day13.PartOne, day13.PartTwo, day13.PuzzleInput())
	case 14:
		runDay(day14.PartOne, day14.PartTwo, day14.PuzzleInput())
	case 15:
		runDay(day15.PartOne, day15.PartTwo, day15.PuzzleInput())
	case 16:
		runDay(day16.PartOne, day16.PartTwo, day16.PuzzleInput())
	case 17:
		runDay(day17.PartOne, day17.PartTwo, day17.PuzzleInput())
	case 18:
		runDay(day18.PartOne, day18.PartTwo, day18.PuzzleInput())
	case 19:
		runDay(day19.PartOne, day19.PartTwo, day19.PuzzleInput())
	case 20:
		runDay(day20.PartOne, day20.PartTwo, day20.PuzzleInput())
	case 21:
		runDay(day21.PartOne, day21.PartTwo, day21.PuzzleInput())
	case 22:
		runDay(day22.PartOne, day22.PartTwo, day22.PuzzleInput())
	case 23:
		runDay(day23.PartOne, day23.PartTwo, day23.PuzzleInput())
	case 24:
		runDay(day24.PartOne, day24.PartTwo, day24.PuzzleInput())
	case 25:
		runDay(day25.PartOne, day25.PartTwo, day25.PuzzleInput())
	}
}

func runDay(partOne puzzle, partTwo puzzle, input string) {
	partOneResult, partOneErr := partOne(input)
	fmt.Print("Part One ")
	if partOneErr != nil {
		fmt.Print(partOneErr)
	} else {
		fmt.Print(partOneResult)
	}
	fmt.Print("\n")
	partTwoResult, partTwoErr := partTwo(input)

	fmt.Print("Part Two ")
	if partTwoErr != nil {
		fmt.Print(partTwoErr)
	} else {
		fmt.Print(partTwoResult)
	}
	fmt.Print("\n")
}
