package engine

import (
	"math"
	"math/rand"
)

// SAStrategy implements simulated annealing, running a random exploration phase
// to build an initial picture of the search space, then switching to simulated
// annealing starting from the best-found point. It runs indefinitely (until
// stopped or max-trials).
//
// Next returns nil in the gap between the explore and SA phases (all explore
// candidates dispatched but not all results observed yet). Once the last
// explore result is observed, SA begins.
type SAStrategy struct {
	inputs   []InputInfo
	metric   string
	minimize bool
	rng      *rand.Rand

	exploreN         int
	exploreGenerated int
	exploreObserved  int
	exploreResults   []EvaluatedPoint

	initialized bool
	current     []float64 // current accepted position, one value per input
	currentVal  float64
	temp        float64
	saIter      int
}

const (
	saInitialTemp  = 1.0
	saCoolingRate  = 0.02
	saMinTemp      = 1e-4
	saPerturbScale = 0.25
)

func NewSAStrategy(inputs []InputInfo, metric string, minimize bool) *SAStrategy {
	exploreN := 10
	if n := 5 * len(inputs); n > exploreN {
		exploreN = n
	}
	return &SAStrategy{
		inputs:   inputs,
		metric:   metric,
		minimize: minimize,
		rng:      rand.New(rand.NewSource(rand.Int63())),
		exploreN: exploreN,
		temp:     saInitialTemp,
	}
}

func (s *SAStrategy) Name() string { return "sa" }

func (s *SAStrategy) Next() Candidate {
	if len(s.inputs) == 0 {
		return nil
	}
	if s.exploreGenerated < s.exploreN {
		s.exploreGenerated++
		return randomCandidate(s.inputs, s.rng)
	}
	if !s.initialized {
		// Waiting for all explore results before SA can start.
		return nil
	}
	return s.neighbor()
}

func (s *SAStrategy) Observe(p EvaluatedPoint) {
	if !s.initialized {
		s.exploreResults = append(s.exploreResults, p)
		s.exploreObserved++
		if s.exploreObserved >= s.exploreN {
			s.initFromBest()
			s.exploreResults = nil // release memory
		}
		return
	}
	// SA phase: accept/reject the result.
	if p.Err == nil {
		if v, ok := p.Metrics[s.metric]; ok {
			s.maybeAccept(p.Candidate, v)
		}
	}
	s.saIter++
	s.temp = math.Max(saMinTemp, saInitialTemp*math.Exp(-saCoolingRate*float64(s.saIter)))
}

func (s *SAStrategy) initFromBest() {
	s.initialized = true
	bestIdx, ok := BestPoint(s.exploreResults, s.metric, s.minimize)
	if !ok {
		s.current = make([]float64, len(s.inputs))
		for i, inp := range s.inputs {
			s.current[i] = inp.Midpoint()
		}
		s.currentVal = math.NaN()
		return
	}
	s.currentVal = s.exploreResults[bestIdx].Metrics[s.metric]
	s.current = candidateFloats(s.exploreResults[bestIdx].Candidate)
}

func (s *SAStrategy) maybeAccept(candidate Candidate, val float64) {
	if math.IsNaN(s.currentVal) {
		s.current = candidateFloats(candidate)
		s.currentVal = val
		return
	}
	denom := math.Abs(s.currentVal)
	if denom == 0 {
		denom = 1
	}
	var delta float64
	if s.minimize {
		delta = (val - s.currentVal) / denom
	} else {
		delta = (s.currentVal - val) / denom
	}
	if delta <= 0 || s.rng.Float64() < math.Exp(-delta/s.temp) {
		s.current = candidateFloats(candidate)
		s.currentVal = val
	}
}

func (s *SAStrategy) neighbor() Candidate {
	idx := s.rng.Intn(len(s.inputs))
	cand := make(Candidate, len(s.inputs))
	for i, inp := range s.inputs {
		val := s.current[i]
		if i == idx {
			val = inp.Perturb(val, s.temp*saPerturbScale, s.rng)
		}
		cand[i] = assign(inp, val)
	}
	return cand
}
