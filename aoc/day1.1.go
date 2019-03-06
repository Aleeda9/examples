package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	positions := [][2]int{}
	velocity := [][2]int{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "<")
		line1 := strings.Split(line[1], ">")
		line1 = strings.Split(line1[0], ", ")
		x, _ := strconv.Atoi(strings.Trim(line1[0], " "))
		y, _ := strconv.Atoi(strings.Trim(line1[1], " "))
		positions = append(positions, [2]int{x, y})
		line2 := strings.Split(line[2], ">")
		line2 = strings.Split(line2[0], ", ")
		a, _ := strconv.Atoi(strings.Trim(line2[0], " "))
		b, _ := strconv.Atoi(strings.Trim(line2[1], " "))
		velocity = append(velocity, [2]int{a, b})
	}

	// fmt.Println(min(positions), max(positions))
	move(positions, velocity)
}

func min(arr [][2]int) int {
	minY := arr[0][1]
	for _, a := range arr {
		if a[1] < minY {
			minY = a[1]
		}
	}
	return minY
}

func max(arr [][2]int) int {
	maxY := arr[0][1]
	for _, a := range arr {
		if a[1] > maxY {
			maxY = a[1]
		}
	}
	return maxY
}

func move(positions [][2]int, velocity [][2]int) {
	t, minY, maxY := 0, 0, 0
	for {
		for i := range positions {
			// fmt.Println(positions[i])
			positions[i][0] += velocity[i][0]
			positions[i][1] += velocity[i][1]
		}

		minY = min(positions)
		maxY = max(positions)
		if maxY-minY < 10 {
			fmt.Println(t+1)
			// matrix := [10][70]string{}
			// drawMatrix(positions, matrix)
			break
		}

		t++
	}
}

func drawMatrix(positions [][2]int, matrix [10][70]string) {
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = "."
		}
	}

	minY := min(positions)
	minX := positions[0][0]
	for _, p := range positions {
		if p[0] < minX {
			minX = p[0]
		}
	}

	for _, p := range positions {
		// fmt.
		matrix[p[1]-minY][p[0]-minX] = "#"
	}
	for _, m := range matrix {
		fmt.Println(m)
	}
}
