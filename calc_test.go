package main 
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

func TestDivide(t *testing.T) {
	if result:=Divide(6, 3); result!= 2 {
		t.Error("6 / 3 did not equal 2, got ", result)
	}
}


func TestSquareRoot(t *testing.T) {
	if result:=SquareRoot(9); result!= 3 {
		t.Error("sqrt(9) did not equal 3, got ", result)
	}
}

func TestModulus(t *testing.T) {
	if result:=Modulus(9, 4); result!= 1 {
		t.Error("9 % 4 did not equal 1, got ", result)
	}
}

func TestFactorial(t *testing.T) {
	if result:=Factorial(4); result!= 24 {
		t.Error("4! did not equal 24, got ", result)
	}
}
