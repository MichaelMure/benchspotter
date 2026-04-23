package engine

// CoordDescentStrategy implements coordinate descent, sweeping one parameter
// at a time while holding the others fixed at the current best-known values.
// It stops when a full pass through all inputs produces no improvement
// (convergence).
//
// Next returns nil after dispatching a full sweep (waiting for Observe calls to
// complete it). Once all sweep results are observed, the strategy advances to
// the next phase and Next resumes returning candidates.
type CoordDescentStrategy struct {
	inputs   []InputInfo
	metric   string
	minimize bool

	best         []float64        // best value per input
	phase        int              // index of the input currently being swept
	sweep        []float64        // candidate values for the current phase
	sweepPos     int              // next index in sweep to dispatch
	sweepResults []EvaluatedPoint // results received for the current sweep
	noImprove    int              // consecutive phases without improvement
	initialized  bool
	converged    bool
}

func NewCoordDescentStrategy(inputs []InputInfo, metric string, minimize bool) *CoordDescentStrategy {
	return &CoordDescentStrategy{
		inputs:   inputs,
		metric:   metric,
		minimize: minimize,
	}
}

func (c *CoordDescentStrategy) Name() string { return "coord" }

func (c *CoordDescentStrategy) Next() Candidate {
	if len(c.inputs) == 0 || c.converged {
		return nil
	}
	if !c.initialized {
		c.init()
	}
	if c.sweepPos < len(c.sweep) {
		return c.emitNext()
	}
	// All sweep candidates dispatched; waiting for Observe to complete the sweep.
	return nil
}

func (c *CoordDescentStrategy) Observe(p EvaluatedPoint) {
	if !c.initialized || c.converged {
		return
	}
	c.sweepResults = append(c.sweepResults, p)
	if len(c.sweepResults) < len(c.sweep) {
		return
	}
	// Sweep complete: update best for this phase.
	c.updateBest()
	// Advance phase.
	c.phase++
	if c.phase >= len(c.inputs) {
		if c.noImprove >= len(c.inputs) {
			c.converged = true
			return
		}
		c.phase = 0
		c.noImprove = 0
	}
	c.startSweep()
}

func (c *CoordDescentStrategy) init() {
	c.initialized = true
	c.best = make([]float64, len(c.inputs))
	for i, inp := range c.inputs {
		c.best[i] = inp.Midpoint()
	}
	c.phase = 0
	c.startSweep()
}

const coordSweepSteps = 10

func (c *CoordDescentStrategy) startSweep() {
	c.sweep = c.inputs[c.phase].SpacedValues(coordSweepSteps)
	c.sweepPos = 0
	c.sweepResults = nil
}

func (c *CoordDescentStrategy) emitNext() Candidate {
	cand := make(Candidate, len(c.inputs))
	for i, inp := range c.inputs {
		val := c.best[i]
		if i == c.phase {
			val = c.sweep[c.sweepPos]
		}
		cand[i] = assign(inp, val)
	}
	c.sweepPos++
	return cand
}

func (c *CoordDescentStrategy) updateBest() {
	bestIdx, ok := BestPoint(c.sweepResults, c.metric, c.minimize)
	if !ok || bestIdx >= len(c.sweep) {
		c.noImprove++
		return
	}
	newVal := c.sweep[bestIdx]
	if newVal == c.best[c.phase] {
		c.noImprove++
	} else {
		c.best[c.phase] = newVal
		c.noImprove = 0
	}
}
