
package read

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFromFile(filePath string) ([]int, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	var inputs []int

	scanner := bufio.NewScanner(file)

	// Read each line and convert to integer
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("failed to convert line to integer: %v", err)
		}
		inputs = append(inputs, number)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error while reading file: %v", err)
	}

	return inputs, nil
}


