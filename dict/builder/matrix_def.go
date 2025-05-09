package builder

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// MatrixDef represents matrix.def.
type MatrixDef struct {
	rowSize int64
	colSize int64
	vec     []int16
}

func parseMatrixDefFile(path string) (*MatrixDef, error) {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return parseMatrix(file)
}

func parseMatrix(r io.Reader) (*MatrixDef, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	line := scanner.Text()
	dim := strings.Split(line, " ")
	if len(dim) != 2 {
		return nil, fmt.Errorf("invalid format: %s", line)
	}
	rowSize, err := strconv.ParseInt(dim[0], 10, 0)
	if err != nil {
		return nil, fmt.Errorf("invalid format: %w, %s", err, line)
	}
	colSize, err := strconv.ParseInt(dim[1], 10, 0)
	if err != nil {
		return nil, fmt.Errorf("invalid format: %w, %s", err, line)
	}
	vec := make([]int16, rowSize*colSize)
	for scanner.Scan() {
		line := scanner.Text()
		ary := strings.Split(line, " ")
		if len(ary) != 3 {
			return nil, fmt.Errorf("invalid format: %s", line)
		}
		row, err := strconv.ParseInt(ary[0], 10, 0)
		if err != nil {
			return nil, fmt.Errorf("invalid format: %w, %s", err, line)
		}
		if row < 0 || row >= rowSize {
			return nil, fmt.Errorf("invalid ID, right-id %d >= row size %d", row, rowSize)
		}

		col, err := strconv.ParseInt(ary[1], 10, 0)
		if err != nil {
			return nil, fmt.Errorf("invalid format: %w, %s", err, line)
		}
		if col < 0 || col >= colSize {
			return nil, fmt.Errorf("invalid ID, left-id %d >= col size %d", col, colSize)
		}

		val, err := strconv.Atoi(ary[2])
		if err != nil {
			return nil, fmt.Errorf("invalid format: %w, %s", err, line)
		}
		vec[col*rowSize+row] = int16(val) //nolint:gosec //G109: Potential Integer overflow made by strconv.Atoi result conversion to int16/32
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("invalid format: %w, %s", err, line)
	}
	return &MatrixDef{
		rowSize: rowSize,
		colSize: colSize,
		vec:     vec,
	}, nil
}
