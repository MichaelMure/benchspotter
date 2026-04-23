package engine

import "math/rand"

// StrategyDef describes one optimization strategy.
type StrategyDef struct {
	Key         string
	Description string
	Build       func(inputs []InputInfo, metric string, minimize bool) Strategy
}

// StrategyDefs lists all built-in optimization strategies in display order.
var StrategyDefs = []StrategyDef{
	{
		Key:         "random",
		Description: "uniform random sampling (runs until stopped or max-trials reached)",
		Build: func(inputs []InputInfo, _ string, _ bool) Strategy {
			return NewRandomStrategy(inputs)
		},
	},
	{
		Key:         "coord",
		Description: "coordinate descent, stops on convergence",
		Build: func(inputs []InputInfo, metric string, minimize bool) Strategy {
			return NewCoordDescentStrategy(inputs, metric, minimize)
		},
	},
	{
		Key:         "sa",
		Description: "random exploration then simulated annealing (runs until stopped)",
		Build: func(inputs []InputInfo, metric string, minimize bool) Strategy {
			return NewSAStrategy(inputs, metric, minimize)
		},
	},
}

// assign creates a ParamAssignment from a native float64 value.
func assign(inp InputInfo, v float64) ParamAssignment {
	return ParamAssignment{Input: inp, Value: inp.Format(v), FloatVal: v}
}

// randomCandidate generates a uniformly random candidate across all inputs.
func randomCandidate(inputs []InputInfo, rng *rand.Rand) Candidate {
	c := make(Candidate, len(inputs))
	for i, inp := range inputs {
		c[i] = assign(inp, inp.RandomValue(rng))
	}
	return c
}

// candidateFloats extracts the native float64 value from each assignment.
func candidateFloats(c Candidate) []float64 {
	vals := make([]float64, len(c))
	for i, a := range c {
		vals[i] = a.FloatVal
	}
	return vals
}
