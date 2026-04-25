package engine

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocateInputs(t *testing.T) {
	ins, err := Inputs(".")
	require.NoError(t, err)

	require.Equal(t, []InputInfo{
		BoolType{pkg: "internal", name: "name1"},
		IntType{pkg: "internal", name: "name2", min: 0, max: 24},
		FloatType{pkg: "internal", name: "threshold", min: 1, max: 128},
		FloatType{pkg: "internal", name: "capacityFactor", min: 0.25, max: 6},
		FloatType{pkg: "internal", name: "fillFraction", min: 0.1, max: 1},
		BoolType{pkg: "internal/anotherpackage", name: "name3"},
		IntType{pkg: "internal/anotherpackage", name: "name4", min: 180, max: 360},
	}, ins)
}
