package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"aoc22_lib/arraylib"
)

func Hello() string {
	return ("Hello, my dude!")
}

func separateLinesInFileByElf(file_path string) [][]int {
	file, err := os.Open(file_path) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	elfs := [][]int{}
	curr_elf := []int{}
	i := 0
	for scanner.Scan() {
		curr_line := scanner.Text()
		curr_line_trimmed := strings.Trim(curr_line, "\n")

		if curr_line_trimmed != "" {
			curr_num, _ := strconv.Atoi(curr_line_trimmed)
			curr_elf = append(curr_elf, curr_num)
		} else {
			elfs = append(elfs, curr_elf)
			curr_elf = []int{}
		}
		i += 1
	}
	elfs = append(elfs, curr_elf)
	file.Close()
	return elfs
}

func sumOverElfs(elfs [][]int) []int {
	elfSums := []int{}
	for _, elf := range elfs {
		elfSum := 0
		for _, val := range elf {
			elfSum += val
		}
		elfSums = append(elfSums, elfSum)
	}
	return elfSums
}

func main() {
	elfs := separateLinesInFileByElf("input/aoc22_in1.txt")
	elfSums := sumOverElfs(elfs)
	max_idx, max_val := arraylib.GetLargestElementAndIndex(elfSums)
	fmt.Println(max_idx, max_val)
	sort.Ints(elfSums)
	big_three_summed := arraylib.Sum(elfSums[len(elfSums)-3:])
	fmt.Printf("$v", big_three_summed)
}
