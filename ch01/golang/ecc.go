package main

import "fmt"

type FieldElement struct {
	num   int
	prime int
}

func NewFieldElement(num int, prime int) (FieldElement, error) {
	if num >= prime || num < 0 {
		return FieldElement{0, 0}, fmt.Errorf("Num %d not in field range 0 to %d", num, prime-1)
	}
	return FieldElement{num, prime}, nil
}

func (fe FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", fe.prime, fe.num)
}

func (fe FieldElement) Equal(other FieldElement) bool {
	return fe.num == other.num && fe.prime == other.prime
}

func (fe FieldElement) Add(other FieldElement) (FieldElement, error) {
	if fe.prime != other.prime {
		return FieldElement{0, 0}, fmt.Errorf("Cannot add two numbers in different Fields")
	}
	num := (fe.num + other.num) % fe.prime
	if num < 0 {
		num += fe.prime
	}
	return NewFieldElement(num, fe.prime)
}

func (fe FieldElement) Sub(other FieldElement) (FieldElement, error) {
	if fe.prime != other.prime {
		return FieldElement{0, 0}, fmt.Errorf("Cannot subtract two numbers in different Fields")
	}
	num := (fe.num - other.num) % fe.prime
	if num < 0 {
		num += fe.prime
	}
	return NewFieldElement(num, fe.prime)
}

func (fe FieldElement) Mul(other FieldElement) (FieldElement, error) {
	if fe.prime != other.prime {
		return FieldElement{0, 0}, fmt.Errorf("Cannot multiply two numbers in different Fields")
	}
	num := (fe.num * other.num) % fe.prime
	if num < 0 {
		num += fe.prime
	}
	return NewFieldElement(num, fe.prime)
}

func (f FieldElement) Pow(exponent int) (FieldElement, error) {
	n := exponent % (f.prime - 1)
	if n < 0 {
		n += f.prime - 1
	}
	num := pow(f.num, n, f.prime)
	return NewFieldElement(num, f.prime)
}

// pow is a helper function to perform power calculation with modulo.
func pow(base, exponent, mod int) int {
	result := 1
	for exponent > 0 {
		if exponent%2 != 0 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exponent /= 2
	}
	return result
}

func (fe FieldElement) TrueDiv(other FieldElement) (FieldElement, error) {
	if fe.prime != other.prime {
		return FieldElement{0, 0}, fmt.Errorf("Cannot divide two numbers in different Fields")
	}
	num := fe.num * pow(other.num, fe.prime-2, fe.prime) % fe.prime
	if num < 0 {
		num += fe.prime
	}
	return NewFieldElement(num, fe.prime)
}

func main() {
	println("---練習問題---")
	p := 97
	{
		println("---95 * 45 * 31---")
		a, _ := NewFieldElement(95, p)
		b, _ := NewFieldElement(45, p)
		c, _ := NewFieldElement(31, p)
		d, _ := a.Mul(b)
		fmt.Println(d.Mul(c))
	}
	{
		println("---17 * 13 * 19 * 44---")
		a, _ := NewFieldElement(17, p)
		b, _ := NewFieldElement(13, p)
		c, _ := NewFieldElement(19, p)
		d, _ := NewFieldElement(44, p)
		e, _ := a.Mul(b)
		f, _ := e.Mul(c)
		fmt.Println(f.Mul(d))
	}
	{
		fmt.Println("12 ** 7 * 77 ** 49")
		a, _ := NewFieldElement(12, p)
		b, _ := NewFieldElement(77, p)
		c, _ := a.Pow(7)
		d, _ := b.Pow(49)
		fmt.Println(c.Mul(d))
	}
	{
		fmt.Println("練習問題5")
		p = 19
		var nums []int = []int{1, 3, 7, 13, 18}
		for _, k := range nums {
			fmt.Printf("[")
			for i := 0; i < p; i += 1 {
				if i == p-1 {
					fmt.Printf("%d]\n", i*k%p)
				} else {
					fmt.Printf("%d, ", i*k%p)
				}
			}
		}
	}
	{
		p = 13
		fmt.Println("--- 1.6.1 ---")
		a, _ := NewFieldElement(3, p)
		b, _ := NewFieldElement(12, p)
		c, _ := NewFieldElement(10, p)
		d, _ := a.Mul(b)
		fmt.Println(d == c)
	}
	{
		p = 13
		fmt.Println("--- 1.6.2 ---")
		a, _ := NewFieldElement(3, p)
		b, _ := NewFieldElement(1, p)
		c, _ := a.Pow(3)
		fmt.Println(c == b)
	}
	{
		fmt.Println("練習問題7")
		var ps []int = []int{7, 11, 17, 31}
		for _, k := range ps {
			fmt.Printf("[")
			for i := 1; i < k; i += 1 {
				if i == k-1 {
					fmt.Printf("%d]\n", pow(i, k-1, k)%p)
				} else {
					fmt.Printf("%d, ", pow(i, k-1, k)%p)
				}
			}
		}
	}
	{
		fmt.Println("練習問題8")
		p = 31
		{
			fmt.Println("3/24")
			a, _ := NewFieldElement(3, p)
			b, _ := NewFieldElement(24, p)
			fmt.Println(a.TrueDiv(b))
		}
		{
			fmt.Println("17 * 3")
			a, _ := NewFieldElement(17, p)
			b, _ := NewFieldElement(3, p)
			fmt.Println(a.Mul(b))
		}
		{
			fmt.Println("17 ** -3")
			a, _ := NewFieldElement(17, p)
			fmt.Println(a.Pow(-3))
		}
	}
}
