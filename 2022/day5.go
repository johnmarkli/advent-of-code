package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
The expedition can depart as soon as the final supplies have been unloaded from
the ships. Supplies are stored in stacks of marked crates, but because the
needed supplies are buried under many other crates, the crates need to be
rearranged.

The ship has a giant cargo crane capable of moving crates between stacks. To
ensure none of the crates get crushed or fall over, the crane operator will
rearrange them in a series of carefully-planned steps. After the crates are
rearranged, the desired crates will be at the top of each stack.

The Elves don't want to interrupt the crane operator during this delicate
procedure, but they forgot to ask her which crate will end up where, and they
want to be ready to unload them as soon as possible so they can embark.

They do, however, have a drawing of the starting stacks of crates and the
rearrangement procedure (your puzzle input). For example:

	[D]

[N] [C]
[Z] [M] [P]

	1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2

In this example, there are three stacks of crates. Stack 1 contains two crates:
crate Z is on the bottom, and crate N is on top. Stack 2 contains three crates;
from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a
single crate, P.

Then, the rearrangement procedure is given. In each step of the procedure, a
quantity of crates is moved from one stack to a different stack. In the first
step of the above rearrangement procedure, one crate is moved from stack 2 to
stack 1, resulting in this configuration:

[D]
[N] [C]
[Z] [M] [P]

	1   2   3

In the second step, three crates are moved from stack 1 to stack 3. Crates are
moved one at a time, so the first crate to be moved (D) ends up below the
second and third crates:

[Z]
[N]
[C] [D]
[M] [P]

	1   2   3

Then, both crates are moved from stack 2 to stack 1. Again, because crates are
moved one at a time, crate C ends up below crate M:

[Z]
[N]
[M]     [D]
[C]     [P]

	1   2   3

Finally, one crate is moved from stack 1 to stack 2:

[Z]
[N]
[D]
[C] [M] [P]

	1   2   3

The Elves just need to know which crate will end up on top of each stack; in
this example, the top crates are C in stack 1, M in stack 2, and Z in stack 3,
so you should combine these together and give the Elves the message CMZ.

After the rearrangement procedure completes, what crate ends up on top of each
stack?
*/

type stack []string

type move struct {
	count   int
	colSrc  int
	colDest int
}

func createStacks(fileScanner *bufio.Scanner) []stack {
	var numStacks int
	buf := bytes.Buffer{}
	// parse out stack structure with columns
	// find column numbers from line that has " 1"
	// read line by line until a line with prefix " 1" is found
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Fprintln(&buf, line) // save for later
		if strings.HasPrefix(line, " 1") {
			// trim left and right, then split on 3 spaces
			// last element is number of columns
			trimmed := strings.Trim(line, " ")
			split := strings.Split(trimmed, "   ")
			numStacks, _ = strconv.Atoi(split[len(split)-1])
			break
		}
	}

	// DCM - left is top of stack
	// find letters in stacks and add to stacks
	stacks := make([]stack, numStacks)
	for i := range stacks {
		stacks[i] = stack{}
	}
	s2 := bufio.NewScanner(&buf)
	for s2.Scan() {
		line := s2.Text()
		if strings.HasPrefix(line, " 1") {
			break
		}
		// loop through runes in each line
		for i, char := range line {
			// if rune is between 'A' and 'Z',
			if char >= 'A' && char <= 'Z' {
				// convert index in line to stack num
				stackIndex := (i - 1) / 4
				// append to stack at index
				stacks[stackIndex] = append(stacks[stackIndex], string(char))
			}
		}
	}
	return stacks
}

func getMoves(fileScanner *bufio.Scanner) []move {
	// parse out moves
	moves := []move{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			split := strings.Split(line, " ")
			count, _ := strconv.Atoi(split[1])
			colSrc, _ := strconv.Atoi(split[3])
			colDest, _ := strconv.Atoi(split[5])
			moves = append(moves, move{count, colSrc, colDest})
		}
	}
	return moves
}

func getTops(stacks []stack) string {
	// get tops of each and return
	var res string
	for _, stack := range stacks {
		if len(stack[0]) > 0 {
			res += stack[0]
		}
	}
	return res
}

func Day5Part1(filepath string) any {
	readFile, _ := os.Open(filepath)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	stacks := createStacks(fileScanner)
	moves := getMoves(fileScanner)

	// index 0 is top of stack, last is bottom of stack
	// perform moves on stacks
	/*
			  [D]
		[N] [C]
		[Z] [M] [P]
		 1   2   3

		 NZ
		 DCM
		 P
	*/
	for _, move := range moves {
		for i := 0; i < move.count; i++ {
			src := move.colSrc - 1
			dest := move.colDest - 1
			// pop off colSrc
			popped := stacks[src][0]
			stacks[src] = stacks[src][1:]

			// push on colDest - prepend to colDest
			stacks[dest] = append([]string{popped}, stacks[dest]...)
		}
	}

	return getTops(stacks)
}

/*
As you watch the crane operator expertly rearrange the crates, you notice the
process isn't following your prediction.

Some mud was covering the writing on the side of the crane, and you quickly
wipe it away. The crane isn't a CrateMover 9000 - it's a CrateMover 9001.

The CrateMover 9001 is notable for many new and exciting features: air
conditioning, leather seats, an extra cup holder, and the ability to pick up
and move multiple crates at once.

Again considering the example above, the crates begin in the same
configuration:

    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

Moving a single crate from stack 2 to stack 1 behaves the same as before:

[D]
[N] [C]
[Z] [M] [P]
 1   2   3

However, the action of moving three crates from stack 1 to stack 3 means that
those three moved crates stay in the same order, resulting in this new
configuration:

        [D]
        [N]
    [C] [Z]
    [M] [P]
 1   2   3

Next, as both crates are moved from stack 2 to stack 1, they retain their order
as well:

        [D]
        [N]
[C]     [Z]
[M]     [P]
 1   2   3

Finally, a single crate is still moved from stack 1 to stack 2, but now it's
crate C that gets moved:

        [D]
        [N]
        [Z]
[M] [C] [P]
 1   2   3

In this example, the CrateMover 9001 has put the crates in a totally different
order: MCD.

Before the rearrangement process finishes, update your simulation so that the
Elves know where they should stand to be ready to unload the final supplies.
After the rearrangement procedure completes, what crate ends up on top of each
stack?
*/
func Day5Part2(filepath string) any {
	readFile, _ := os.Open(filepath)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	stacks := createStacks(fileScanner)
	moves := getMoves(fileScanner)

	for _, move := range moves {
		src := move.colSrc - 1
		dst := move.colDest - 1
		popped := []string{}
		for i := 0; i < move.count; i++ {
			popped = append(popped, stacks[src][i])
		}

		// prepend range with count from src to dst
		stacks[dst] = append(popped, stacks[dst]...)

		// remove range with count from beginning of src
		stacks[src] = stacks[src][move.count:]
	}

	return getTops(stacks)
}
