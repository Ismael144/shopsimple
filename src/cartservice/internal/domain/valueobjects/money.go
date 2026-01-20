package valueobjects

import "errors"

const (
	nanosMin = -999_999_999
	nanosMax = +999_999_999
	nanosMod = 1_000_000_000
)

var (
	ErrInvalidValue        = errors.New("one of the specified money values is invalid")
	ErrMismatchingCurrency = errors.New("mismatching currency codes")
)

// Representing money type
type Money struct {
	// Currency Code in ISO 4217 (USD, EUR)
	CurrencyCode string
	Units        int64
	Nanos        int32
}

// Getter for currencyCode
func (m Money) GetCurrencyCode() string {
	return m.CurrencyCode
}

// Getter for units
func (m Money) GetUnits() int64 {
	return m.Units
}

// Getter for nanos
func (m Money) GetNanos() int32 {
	return m.Nanos
}

// Initialize new money value
func NewMoney(currencyCode string, units int64, nanos int32) Money {
	return Money{
		CurrencyCode: currencyCode,
		Units:        units,
		Nanos:        nanos,
	}
}

// IsValid checks if specified value has valid units/nanos signs and ranges
func IsValid(m Money) bool {
	return signMatches(m) && validNanos(m.GetNanos())
}

func signMatches(m Money) bool {
	return m.GetNanos() == 0 || m.GetUnits() == 0 || (m.GetNanos() < 0) == (m.GetUnits() < 0)
}

func validNanos(nanos int32) bool { return nanosMin <= nanos && nanos <= nanosMax }

// IsZero returns true if the specified money value is equal to zero
func IsZero(m Money) bool { return m.GetNanos() == 0 && m.GetUnits() == 0 }

// IsPositive returns true if the specified money value is valid and is positive
func IsPositive(m Money) bool {
	return IsValid(m) && m.GetUnits() > 0 || (m.GetUnits() == 0 && m.GetNanos() > 0)
}

// IsNegative returns true if the specified money value is valid and is negative
func IsNegative(m Money) bool {
	return IsValid(m) && m.GetUnits() < 0 || (m.GetUnits() == 0 && m.GetNanos() < 0)
}

// AreSameCurrency returns true if values l and r have a currency code thats the same value
func AreSameCurrency(l, r Money) bool {
	return l.GetCurrencyCode() == r.GetCurrencyCode() && l.GetCurrencyCode() != ""
}

// Returns returns true if values l and r are the equal, including the currency.
// This does not check validity of the provided values.
func AreEquals(l, r Money) bool {
	return l.GetCurrencyCode() == r.GetCurrencyCode() &&
		l.GetUnits() == r.GetUnits() && l.GetNanos() == r.GetNanos()
}

// Negate returns the same amount with the sign negated
func Negate(m Money) Money {
	return Money{
		Units:        -m.GetUnits(),
		Nanos:        -m.GetNanos(),
		CurrencyCode: m.GetCurrencyCode(),
	}
}

// Must panic if the given error is not nil. This can be used with other
// functions like "m := Must(Sum(a,b))"
func Must(v Money, err error) Money {
	if err != nil {
		panic(err)
	}
	return v
}

// Sum adds two values. Returns an error if one of the values are invalid or
// currency codes are not matching (unless currency code is unspecified for both)
func Sum(l, r Money) (Money, error) {
	if !IsValid(l) || !IsValid(r) {
		return Money{}, ErrInvalidValue
	} else if l.GetCurrencyCode() != r.GetCurrencyCode() {
		return Money{}, ErrMismatchingCurrency
	}
	units := l.GetUnits() + r.GetUnits()
	nanos := l.GetNanos() + r.GetNanos()

	if (units == 0 && nanos == 0) || (units > 0 && nanos >= 0) || (units < 0 && nanos <= 0) {
		// Same sign <units, nanos>
		units += int64(nanos / nanosMod)
		nanos = nanos % nanosMod
	} else {
		// different sign. nanos guaranteed to not go over the limit
		if units > 0 {
			units--
			nanos += nanosMod
		} else {
			units++
			nanos -= nanosMod
		}
	}

	return Money{
		Units:        units,
		Nanos:        nanos,
		CurrencyCode: l.GetCurrencyCode(),
	}, nil
}

// Multiply is a multiplication operation done through adding the value
// to itself n-1 times 
func Multiply(m Money, n uint32) Money {
	out := m 
	for n > 1 {
		out = Must(Sum(out, m))
		n--
	}
	return out
}