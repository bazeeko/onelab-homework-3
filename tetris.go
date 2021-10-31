package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sizeOfField(field string) int {
	if strings.Contains(field, "map[") || strings.Contains(field, " chan ") || strings.Contains(field, "*") || strings.Contains(field, "func(") {
		return 8
	}

	if strings.Contains(field, "interface") {
		return 16
	}

	def := strings.Fields(field)

	field = def[len(def)-1]

	switch field {
	case "int":
		return 8
	case "int8":
		return 1
	case "int16":
		return 2
	case "int32":
		return 4
	case "int64":
		return 8
	case "uint":
		return 8
	case "uint8":
		return 1
	case "uint16":
		return 2
	case "uint32":
		return 4
	case "uint64":
		return 8
	case "uintptr":
		return 8
	case "string":
		return 16
	case "bool":
		return 1
	case "error":
		return 16
	case "rune":
		return 4
	case "float32":
		return 4
	case "float64":
		return 8
	case "complex64":
		return 8
	case "complex128":
		return 16
	case "byte":
		return 1
	default:
		fmt.Println(def)
		return -1
	}
}

func sizeOfStruct(st []string) (result int, err error) {
	for _, field := range st {
		fieldSize := sizeOfField(field)
		if fieldSize < 0 {
			return 0, fmt.Errorf("sizeOfStruct: %w", errors.New("unknows field type"))
		}

		if fieldSize >= 8 {
			if mod := result % 8; mod != 0 {
				result += 8 - mod
			}
		}

		result += fieldSize
	}

	if mod := result % 8; mod != 0 {
		result += 8 - mod
	}

	return
}

func permutations(fields []string) [][]string {
	var perm func([]string, int)
	result := [][]string{}

	perm = func(arr []string, n int) {
		if n == 1 {
			temp := make([]string, len(arr))
			copy(temp, arr)
			result = append(result, temp)
		} else {
			for i := 0; i < n; i++ {
				perm(arr, n-1)

				if n%2 == 1 {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				} else {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				}
			}
		}
	}
	perm(fields, len(fields))
	return result
}

func tetris(filename string) ([][]string, error) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("tetris: %w", err)
	}
	defer file.Close()

	var fields []string
	var foundStart = false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "}" {
			break
		}

		line = strings.Join(strings.Fields(line), " ")

		if foundStart {
			fields = append(fields, line)
		}
		if strings.Contains(line, "type") && strings.Contains(line, "struct") && !foundStart {
			foundStart = true
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("tetris: %w", err)
	}

	perms := permutations(fields)

	size := make(map[int]int)
	for i := range perms {
		size[i], err = sizeOfStruct(perms[i])
		if err != nil {
			return nil, fmt.Errorf("tetris: %w", err)
		}
	}

	keys := make([]int, 0, len(size))
	for i := range perms {
		keys = append(keys, i)
	}

	sort.Slice(keys, func(i, j int) bool { return size[i] < size[j] })

	var result [][]string
	if len(perms) < 3 {
		result = make([][]string, 0, len(keys))

		for _, v := range keys {
			result = append(result, perms[v])
		}
		return result, nil
	}

	result = make([][]string, 0, 3)
	for _, v := range keys[:3] {
		result = append(result, perms[v])
	}

	return result, nil
}
