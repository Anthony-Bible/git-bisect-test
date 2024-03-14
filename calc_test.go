package calculator
import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 did not equal 3")
	}
}

func TestSubtract(t *testing.T) {
	if result:=Subtract(2, 1); result!= 1 {
		t.Error("2 - 1 did not equal 1, got ", result)
	}
}

func TestMultiply(t *testing.T) {
	if result:=Multiply(2, 3); result!= 6 {
		t.Error("2 * 3 did not equal 6, got ", result)
	}
}
