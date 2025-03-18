package lab2

import (
	"bytes"
	"testing"
)

func TestComputeHandler_ValidExpression(t *testing.T) {
	input := bytes.NewBufferString("+ 5 3")
	output := &bytes.Buffer{}

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()

	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}

	expected := "8\n"
	if output.String() != expected {
		t.Errorf("expected %v, but got %v", expected, output.String())
	}
}

func TestComputeHandler_InvalidExpression(t *testing.T) {
	input := bytes.NewBufferString("+ 5")
	output := &bytes.Buffer{}

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()

	if err == nil {
		t.Fatalf("expected error, but got nil")
	}

	if output.Len() > 0 {
		t.Errorf("expected empty output, but got: %v", output.String())
	}
}

func TestComputeHandler_EmptyExpression(t *testing.T) {
	input := bytes.NewBufferString("")
	output := &bytes.Buffer{}

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()

	if err != nil {
		t.Fatalf("expected error, but got nil")
	}

	if output.Len() > 0 {
		t.Errorf("expected empty output, but got: %v", output.String())
	}
}
