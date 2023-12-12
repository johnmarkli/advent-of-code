package aoc2023

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

/*
You take the boat and find the gardener right where you were told he would be:
managing a giant "garden" that looks more to you like a farm.

"A water source? Island Island is the water source!" You point out that Snow
Island isn't receiving any water.

"Oh, we had to stop the water because we ran out of sand to filter it with!
Can't make snow with dirty water. Don't worry, I'm sure we'll get more sand
soon; we only turned off the water a few days... weeks... oh no." His face
sinks into a look of horrified realization.

"I've been so busy making sure everyone here has food that I completely forgot
to check why we stopped getting more sand! There's a ferry leaving soon that is
headed over in that direction - it's much faster than your boat. Could you
please go check it out?"

You barely have time to agree to this request when he brings up another. "While
you wait for the ferry, maybe you can help us with our food production problem.
The latest Island Island Almanac just arrived and we're having trouble making
sense of it."

The almanac (your puzzle input) lists all of the seeds that need to be planted.
It also lists what type of soil to use with each kind of seed, what type of
fertilizer to use with each kind of soil, what type of water to use with each
kind of fertilizer, and so on. Every type of seed, soil, fertilizer and so on
is identified with a number, but numbers are reused by each category - that is,
soil 123 and fertilizer 123 aren't necessarily related to each other.

For example:

seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4

The almanac starts by listing which seeds need to be planted: seeds 79, 14, 55,
and 13.

The rest of the almanac contains a list of maps which describe how to convert
numbers from a source category into numbers in a destination category. That is,
the section that starts with seed-to-soil map: describes how to convert a seed
number (the source) to a soil number (the destination). This lets the gardener
and his team know which soil to use with which seeds, which water to use with
which fertilizer, and so on.

Rather than list every source number and its corresponding destination number
one by one, the maps describe entire ranges of numbers that can be converted.
Each line within a map contains three numbers: the destination range start, the
source range start, and the range length.

Consider again the example seed-to-soil map:

50 98 2
52 50 48

The first line has a destination range start of 50, a source range start of 98,
and a range length of 2. This line means that the source range starts at 98 and
contains two values: 98 and 99. The destination range is the same length, but
it starts at 50, so its two values are 50 and 51. With this information, you
know that seed number 98 corresponds to soil number 50 and that seed number 99
corresponds to soil number 51.

The second line means that the source range starts at 50 and contains 48
values: 50, 51, ..., 96, 97. This corresponds to a destination range starting
at 52 and also containing 48 values: 52, 53, ..., 98, 99. So, seed number 53
corresponds to soil number 55.

Any source numbers that aren't mapped correspond to the same destination
number. So, seed number 10 corresponds to soil number 10.

So, the entire list of seed numbers and their corresponding soil numbers looks
like this:

seed  soil
0     0
1     1
...   ...
48    48
49    49
50    52
51    53
...   ...
96    98
97    99
98    50
99    51

With this map, you can look up the soil number required for each initial seed
number:

Seed number 79 corresponds to soil number 81.
Seed number 14 corresponds to soil number 14.
Seed number 55 corresponds to soil number 57.
Seed number 13 corresponds to soil number 13.

The gardener and his team want to get started as soon as possible, so they'd
like to know the closest location that needs a seed. Using these maps, find the
lowest location number that corresponds to any of the initial seeds. To do
this, you'll need to convert each seed number through other categories until
you can find its corresponding location number. In this example, the
corresponding types are:

Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
Seed 55, soil 57, fertilizer 57, water 53, light 46, temperature 82, humidity 82, location 86.
Seed 13, soil 13, fertilizer 52, water 41, light 34, temperature 34, humidity 35, location 35.

So, the lowest location number in this example is 35.

What is the lowest location number that corresponds to any of the initial seed
numbers?
*/

// Day5Part1 ...
func Day5Part1(filepath string) any {
	result := 0

	file, fileScanner := readFile(filepath)
	defer file.Close()

	almanac := NewAlmanac(fileScanner)
	result = almanac.LowestLocNum()
	return result
}

/*
Everyone will starve if you only plant such a small number of seeds. Re-reading
the almanac, it looks like the seeds: line actually describes ranges of seed
numbers.

The values on the initial seeds: line come in pairs. Within each pair, the
first value is the start of the range and the second value is the length of the
range. So, in the first line of the example above:

seeds: 79 14 55 13

This line describes two ranges of seed numbers to be planted in the garden. The
first range starts with seed number 79 and contains 14 values: 79, 80, ..., 91,
92. The second range starts with seed number 55 and contains 13 values: 55, 56,
..., 66, 67.

Now, rather than considering four seed numbers, you need to consider a total of
27 seed numbers.

In the above example, the lowest location number can be obtained from seed
number 82, which corresponds to soil 84, fertilizer 84, water 84, light 77,
temperature 45, humidity 46, and location 46. So, the lowest location number is
46.

Consider all of the initial seed numbers listed in the ranges on the first line
of the almanac. What is the lowest location number that corresponds to any of
the initial seed numbers?
*/

// Day5Part2 ...
func Day5Part2(filepath string) any {
	result := 0

	file, fileScanner := readFile(filepath)
	defer file.Close()

	almanac := NewAlmanac(fileScanner)
	result = almanac.LowestLocNumRange()

	return result
}

type conversion struct {
	src      int
	dst      int
	rangeLen int
}

type interval struct {
	start int
	end   int
}

type conversionRange struct {
	srcInterval interval
	diff        int
}

type conversionRangeMap []conversionRange

// Almanac captures an almanac
type Almanac struct {
	seeds               []int
	seedRanges          []interval
	conversionRangeMaps []conversionRangeMap
}

// NewAlmanac returns a new almanac from a file
func NewAlmanac(fileScanner *bufio.Scanner) *Almanac {
	// parse seeds first
	seeds := []int{}
	if fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, ": ")
		seedsStrs := strings.Fields(split[1])
		for _, seedStr := range seedsStrs {
			seed, _ := strconv.Atoi(seedStr)
			seeds = append(seeds, seed)
		}
	}

	seedRanges := []interval{}
	for i := 0; i < len(seeds)-1; i += 2 {
		seedRanges = append(seedRanges, interval{
			start: seeds[i],
			end:   seeds[i] + seeds[i+1] - 1,
		})
	}

	conversionRangeMaps := []conversionRangeMap{}
	var curRangeMap []conversionRange
	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Fields(line)
		if len(split) == 2 {
			curRangeMap = []conversionRange{}
		} else if len(split) == 3 {
			dst, _ := strconv.Atoi(split[0])
			src, _ := strconv.Atoi(split[1])
			length, _ := strconv.Atoi(split[2])
			curRangeMap = append(curRangeMap, conversionRange{
				srcInterval: interval{
					start: src,
					end:   src + length - 1,
				},
				diff: dst - src,
			})
		} else if len(split) == 0 {
			if len(curRangeMap) > 0 {
				conversionRangeMaps = append(conversionRangeMaps, curRangeMap)
				curRangeMap = []conversionRange{}
			}
		}
	}
	if len(curRangeMap) > 0 { // catch last mapping after scan
		conversionRangeMaps = append(conversionRangeMaps, curRangeMap)
	}
	return &Almanac{
		seeds:               seeds,
		seedRanges:          seedRanges,
		conversionRangeMaps: conversionRangeMaps,
	}
}

// LowestLocNum return lowest location number in almanac with seeds
func (a *Almanac) LowestLocNum() int {
	min := math.MaxInt
	for _, seed := range a.seeds {
		loc := a.convertSeed(seed)
		if loc < min {
			min = loc
		}
	}
	if min == math.MaxInt {
		return 0
	}
	return min
}

// LowestLocNumRange return lowest location number in almanac with seed ranges
func (a *Almanac) LowestLocNumRange() int {
	min := math.MaxInt
	// for each seed range
	for _, interval := range a.seedRanges {
		// convert seed range to get destination seed ranges
		convertedSeedRanges := a.convertSeedRange(interval)
		// for each destination seed range
		for _, convertedSeedRange := range convertedSeedRanges {
			loc := convertedSeedRange.start
			if loc < min {
				min = loc
			}
		}
	}
	if min == math.MaxInt {
		return 0
	}
	return min
}

func (a *Almanac) convertSeedRange(sr interval) []interval {
	srcRanges := []interval{sr}
	dstRanges := []interval{}
	// for each conversion range map
	for _, conversionRangeMap := range a.conversionRangeMaps {
		// for each conversion range
		dstRanges = []interval{}
		for _, cr := range conversionRangeMap {
			// compare src ranges to conversion ranges
			unconverted, converted := a.convertSrcRanges(srcRanges, cr)
			dstRanges = append(dstRanges, converted...)
			srcRanges = unconverted
		}
		dstRanges = append(dstRanges, srcRanges...)
		srcRanges = dstRanges
	}
	return dstRanges
}

func (a *Almanac) convertSeed(seed int) int {
	result := seed
	// for each conversionMap
	for _, conversionRangeMap := range a.conversionRangeMaps {
		// for each conversion in conversionMap
		converted := false
		for _, conversion := range conversionRangeMap {
			// if seed num falls in range of conversion, add diff to convert
			if !converted && result >= conversion.srcInterval.start && result <= conversion.srcInterval.end {
				result += conversion.diff
				converted = true
			}
		}
	}
	return result
}

func (a *Almanac) convertSrcRanges(srs []interval, cr conversionRange) ([]interval, []interval) {
	unconverted := []interval{}
	converted := []interval{}
	// for each range in src range
	for _, sr := range srs {
		// a range that matches shouldn't be matched again
		// update result seed ranges based on
		// if seed range falls within conversion range
		if sr.start >= cr.srcInterval.start && sr.end <= cr.srcInterval.end {
			newRange := interval{
				start: sr.start + cr.diff,
				end:   sr.end + cr.diff,
			}
			converted = append(converted, newRange)
		} else if sr.start < cr.srcInterval.start && sr.end >= cr.srcInterval.start {
			// if seed range overlaps only on left side of conversion range
			// add new unconverted range on left, convert overlaping range
			// add range that overlaps and add diffs
			newRanges := []interval{}
			newRange := interval{
				start: cr.srcInterval.start + cr.diff,
				end:   sr.end + cr.diff,
			}
			converted = append(converted, newRange)
			newRanges = append(newRanges, newRange)
			// add range that doesn't overlap with no diffs
			newRange = interval{
				start: sr.start,
				end:   cr.srcInterval.start - 1,
			}
			unconverted = append(unconverted, newRange)
			newRanges = append(newRanges, newRange)
		} else if sr.start <= cr.srcInterval.end && sr.end > cr.srcInterval.end {
			// if seed range overlaps only on right side of conversion range
			// add new unconverted range on right, convert overlaping range
			// add range that overlaps and add diffs
			newRanges := []interval{}
			newRange := interval{
				start: sr.start + cr.diff,
				end:   cr.srcInterval.end + cr.diff,
			}
			converted = append(converted, newRange)
			newRanges = append(newRanges, newRange)
			// add range that doesn't overlap with no diffs
			newRange = interval{
				start: cr.srcInterval.end + 1,
				end:   sr.end,
			}
			unconverted = append(unconverted, newRange)
			newRanges = append(newRanges, newRange)
		} else if sr.start < cr.srcInterval.start && sr.end > cr.srcInterval.end {
			// if seed range overlaps conversion range on both sides
			// add new unconverted range on left and right, convert overlaping range
			// add range that overlaps and add diffs
			newRanges := []interval{}
			newRange := interval{
				start: cr.srcInterval.start + cr.diff,
				end:   cr.srcInterval.end + cr.diff,
			}
			converted = append(converted, newRange)
			newRanges = append(newRanges, newRange)
			// add range that doesn't overlap with no diffs
			newRange = interval{
				start: sr.start,
				end:   cr.srcInterval.start - 1,
			}
			unconverted = append(unconverted, newRange)
			newRanges = append(newRanges, newRange)
			newRange = interval{
				start: cr.srcInterval.end + 1,
				end:   sr.end,
			}
			unconverted = append(unconverted, newRange)
			newRanges = append(newRanges, newRange)
		} else {
			unconverted = append(unconverted, sr)
		}
	}
	return unconverted, converted
}
