package day02

import (
	"bytes"
	"strings"
)

func day2TestInput() string {
	return `ULL
RRDDD
LURDL
UUUUD`
}

func day2PuzzleInput() string {
	return `RUDULRLLUULRURDDRRUDURULLLDRLRLUDDLUDUDDUDRRDUDULDUUULLRULLRLDDLDLDDRLRRRRUDLLDDUULDRLLUDDRRUURLULRRRDLLURRUUDURUDDURLUDDDLUDDUUDUURUDLRDRDRLRDRLDRUDRUUDLRDDRRURDDLRDDRRURDUDDLULLUDRURURRRLRRUDUULULULRRLDLUDUURRLLRUDLLDRDDLRRRULRUDLULDDLLLULDLRUDLLLLRDDLRDRLDRLLRDRRDLRDULULRLLLDRUDRRRUULRUULDRURLUDRURRDLLDLRDLDDDDRRLUDLRRLUUUURDRDDLRRURURRDUULLRLURLURUDDDRDURDUUDRLRLRRLDDLDLDLDDDUDDULURLDDLLRLRRDULUDDLULRLUDDLDLRULUUUDRLDRUDURLUDDRLLRUULDLRRRRDLLLLURULLRDRRUDLUULRRDLLRLRLUDLDDULLDLLRDLDLL
LLUUUUUUDUDRLRDRDLDURRRLLRRLRURLLUURRLLUDUDLULUURUUURDLUDLDDLULLRDLRUULDLRDUDURLLDDUDUDULLUDDUULLLUULRRRLULRURRDLRUDUDDURRRDRUURDURLLULLRULLDRUULLURLDRDUUDDDDDDRRLDRLRRRLULDDUURRLLLLDRURLURDRDRDURUDUURRDUDUDRLLUUDDRLUDDDRDLDLRLDRURRDLLRULDRLLURURRLUULLRLRRURDDRDRUUURUURUUUDLLRRLUDRLDLRLURLDLUDDUDDDLDUDRRLDLRURULRLLRDUULURRRULDLLLRLDDDUURRRRDULLRURRLULULDLRRUDUDDLRUURDLDUDDUDRRDLRRRDUDUUUDLLDDDDLURLURRRUUULLLULRRLLLLLLULDUUDLRUDRRDLRDUUDUDLLRLDLLRUURDUUURUUUDDLLUUDLULDURLULULUUUDRUDULLURRULRULLRDLDDU
RLUUURULLDLRLDUDRDURRDUURLLUDDDUULRRRLRLURDDRUULUDULDUUDDDDUDDDDRUDDLDUUDRUDLRRRLLRDDLLLRLLRUULRUULDDRURRLURRLRLULDDRRRDDURDDRDRDULRUDRUUDULRLLULDLRLLDRULRDDRRDDUDLRLLUDRDRRRLUDULRDLRDDURRUUDDRRUDURRUUUDDRRDUDURLUUDUDUURDDDLURLULLUULULURUDUUDRUDULLUUULURDLDUULLDDLLDULRLRLRDUUURUUDLRLDURUDRLDULLUDLDLLRDUURRDUDURLUUUDLLRRULRLULRLDLLURDURRULRLLRRDUDLLRDRRRRDLUUDRUUUDDLRLUDDDDDDRURRRUUURRDLLRURLDDLLDLRRLLLDRRULRRUDLDRDDRRLULURLLUURURURRRRUUUUURUDURLRLLLULULDLLDLRDRRULUDUDRDRRDRDRRDUDLLLRUDRUDDDULRULRRRDRLRUUUURUDURDUUULLULRUDDULDUUDLDURRD
ULRULDDLDLULLLRRRLRUDDDDDLLDDUDLRRDULUUDRDLRRURDRRLUULRURUDRRULDLLLUDRUUDULULUDDRUDDDRDURRRDRDUUURLRDULUDRDRLDRUDDLLLDRRULUDLUDLDLLRRUDUULULDLDLLUURDLDDLLUUDURLURLLLDRDLDRRLRULUURRDRULRUUURULRRUDDDDLLDLDDLLRRLRRRRDUUDUDLDRDRRURDLRURULDLRDLLLLRUDRLLRDLRLRDURDRUDURRRLRDRDLLRLUDDDDRLRLLDUURRURLUURUULUDLUURDRRUDDLUDUDDDURRDRUDRLRULDULUUUUUUDDUDRUDUUURUDRRDLUDLUUDUULUDURDLDDDLLURRURUUDUDDRRDRLLULULDRLRURRDDDRDUUURDDDRULUDRDDLDURRLDDDLRRRLDDRDURULDLUDLLLURLURRLRRULDLLDDUDRRULDRRRRLURRUULRRRUDLURDLLDLLDULUUDRRLDLLLDRLRUDLUULDLDRUDUDURDRUDRDDDLRLULLUR
LRLUUURRLRRRRRUURRLLULRLULLDLUDLUDRDDRLDLRLULLURDURLURDLLRLDUUDDURRRRLDLLRULLRLDLLUUDRLDDLLDRULDRLLRURDLRURRUDLULLRURDLURRURUDULLDRLLUUULUDRURRUUDUDULUUULRLDDULDRDLUDDUDDDLRURULLDLLLRLLUURDLRUDLLLLDLLRLRUUUDDRUUUUDLDLRDDURLDURUULLLUUDLLLLDULRRRLLDLDRRDRLUDRUDURLLUDLRLLUDUDRDDDRDLRDLRULUULDRLUDLRLDUURLRRLUDDDUUDDDUDRLDLDUDLURUULLDDDURUUULRLUDLDURUUDRDRURUDDUURDUUUDLLDLDLDURUURLLLLRURUURURULRULLRUDLRRUUUUUDRRLLRDDUURDRDRDDDUDRLURDRRRUDLLLDURDLUUDLLUDDULUUDLDUUULLDRDLRURUURRDURRDLURRRRLLUUULRDULDDLDUURRDLDLLULRRLLUDLDUDLUUL`
}

type Point struct {
	x int
	y int
}

func getKeyPad() [][]string {
	keypad := make([][]string, 5) // One row per unit of y.

	keypad[0] = []string{"0", "0", "0", "0", "0"}
	keypad[1] = []string{"0", "1", "2", "3", "0"}
	keypad[2] = []string{"0", "4", "5", "6", "0"}
	keypad[3] = []string{"0", "7", "8", "9", "0"}
	keypad[4] = []string{"0", "0", "0", "0", "0"}

	return keypad
}

func getStarPad() [][]string {
	keypad := make([][]string, 7) // One row per unit of y.

	keypad[0] = []string{"0", "0", "0", "0", "0", "0", "0"}
	keypad[1] = []string{"0", "0", "0", "1", "0", "0", "0"}
	keypad[2] = []string{"0", "0", "2", "3", "4", "0", "0"}
	keypad[3] = []string{"0", "5", "6", "7", "8", "9", "0"}
	keypad[4] = []string{"0", "0", "A", "B", "C", "0", "0"}
	keypad[5] = []string{"0", "0", "0", "D", "0", "0", "0"}
	keypad[6] = []string{"0", "0", "0", "0", "0", "0", "0"}

	return keypad
}

func getKeyCode(inputData string, startPoint Point, keyPad [][]string) string {
	var keyCode bytes.Buffer

	var lines = strings.Split(inputData, "\n")

	for _, line := range lines {
		for _, command := range line {
			switch string(command) {
			case "U":
				if getKey(Point{startPoint.x - 1, startPoint.y}, keyPad) != "0" {
					startPoint = Point{startPoint.x - 1, startPoint.y}
				}
			case "D":
				if getKey(Point{startPoint.x + 1, startPoint.y}, keyPad) != "0" {
					startPoint = Point{startPoint.x + 1, startPoint.y}
				}
			case "L":
				if getKey(Point{startPoint.x, startPoint.y - 1}, keyPad) != "0" {
					startPoint = Point{startPoint.x, startPoint.y - 1}
				}
			case "R":
				if getKey(Point{startPoint.x, startPoint.y + 1}, keyPad) != "0" {
					startPoint = Point{startPoint.x, startPoint.y + 1}
				}
			}
		}

		keyCode.WriteString(string(getKey(startPoint, keyPad)))
	}

	return keyCode.String()
}

// Day02_01 solves day two, part one!
func PartOne(inputData string) string {
	return getKeyCode(inputData, Point{2, 2}, getKeyPad())
}

// Day02_02 solves day two, part one!
func PartTwo(inputData string) string {
	return getKeyCode(inputData, Point{3, 3}, getStarPad())
}

func getKey(cursor Point, pad [][]string) string {
	return pad[cursor.x][cursor.y]
}