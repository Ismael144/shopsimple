package valueobjects

import "fmt"

type Money struct {
	Cents int64
}

func MoneyFromCents(cents int64) Money {
	if cents < 0 {
		panic("money cannot be negative")
	}
	return Money{Cents: cents}
}

func Dollars(d int64) Money {
	return Money{Cents: d * 100}
}

func (m Money) Add(other Money) Money {
	return Money{Cents: m.Cents + other.Cents}
}

func (m Money) Mul(value int64) Money {
	return Money{Cents: m.Cents * value}
}

func (m Money) String() string {
	dollars := m.Cents / 100
	cents := m.Cents % 100
	return fmt.Sprintf("$%d.%02d", dollars, cents)
}

func (m Money) Eq(other Money) bool {
	return m.Cents == other.Cents
}