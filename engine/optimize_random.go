package engine

import (
	"math/rand"
)

// RandomStrategy evaluates uniformly random candidates indefinitely (log-uniform
// for log-scale types). It has no convergence criterion; the user stops it.
type RandomStrategy struct {
	inputs []InputInfo
	rng    *rand.Rand
}

func NewRandomStrategy(inputs []InputInfo) *RandomStrategy {
	return &RandomStrategy{
		inputs: inputs,
		rng:    rand.New(rand.NewSource(rand.Int63())),
	}
}

func (r *RandomStrategy) Name() string { return "random" }

func (r *RandomStrategy) Next() Candidate {
	return randomCandidate(r.inputs, r.rng)
}

func (r *RandomStrategy) Observe(_ EvaluatedPoint) {}
