package rmath

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRound(t *testing.T) {
	t.Run("OK float32", func(t *testing.T) {
		result := Round(float32(0.123456789), 3)
		require.Equal(t, float32(0.123), result)
	})
	t.Run("OK float64", func(t *testing.T) {
		result := Round(0.123456789, 3)
		require.Equal(t, 0.123, result)
	})
}

func ExampleRound() {
	var (
		val32 float32 = 0.12345
		val64 float64 = 0.12345
	)

	res32 := Round(val32, 3)
	fmt.Println(res32)
	fmt.Println(reflect.TypeOf(res32))

	res64 := Round(val64, 3)
	fmt.Println(res64)
	fmt.Println(reflect.TypeOf(res64))

	// Output:  0.123
	// float32
	// 0.123
	// float64
}

func TestRoundUp(t *testing.T) {
	t.Run("OK float32", func(t *testing.T) {
		result := RoundUp(float32(0.123456789), 3)
		require.Equal(t, float32(0.124), result)
	})
	t.Run("OK float64", func(t *testing.T) {
		result := RoundUp(0.123456789, 3)
		require.Equal(t, 0.124, result)
	})
}

func ExampleRoundUp() {
	var (
		val32 float32 = 0.12345
		val64 float64 = 0.12345
	)

	res32 := RoundUp(val32, 3)
	fmt.Println(res32)
	fmt.Println(reflect.TypeOf(res32))

	res64 := RoundUp(val64, 3)
	fmt.Println(res64)
	fmt.Println(reflect.TypeOf(res64))

	// Output:  0.124
	// float32
	// 0.124
	// float64
}
