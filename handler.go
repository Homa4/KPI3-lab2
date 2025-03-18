package lab2

import (
	"bufio"
	"fmt"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (h *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(h.Input)

	for scanner.Scan() {
		expression := scanner.Text()
		if expression == "" {
			continue
		}

		result, err := CalculatePrefix(expression)
		if err != nil {
			return fmt.Errorf("error in calculating result: %w", err)
		}

		_, err = fmt.Fprintln(h.Output, result)
		if err != nil {
			return fmt.Errorf("error in printing result: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error in bufio scanner: %w", err)
	}

	return nil
}
