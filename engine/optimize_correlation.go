package engine

import "math"

// InputCorrelations computes the Pearson r between each input's numeric value
// and each metric across all valid results. Entries with fewer than 3 valid
// points are omitted.
func InputCorrelations(results []EvaluatedPoint, metrics []string) map[string]map[string]float64 {
	if len(results) == 0 || len(metrics) == 0 || len(results[0].Candidate) == 0 {
		return nil
	}

	nInputs := len(results[0].Candidate)
	out := make(map[string]map[string]float64, nInputs)

	for i := 0; i < nInputs; i++ {
		name := results[0].Candidate[i].Input.Name()
		out[name] = make(map[string]float64, len(metrics))
		for _, metric := range metrics {
			var xs, ys []float64
			for _, r := range results {
				if r.Err != nil || r.Metrics == nil || i >= len(r.Candidate) {
					continue
				}
				y, ok := r.Metrics[metric]
				if !ok {
					continue
				}
				xs = append(xs, r.Candidate[i].FloatVal)
				ys = append(ys, y)
			}
			if len(xs) < 3 {
				continue
			}
			if v := pearsonCorr(xs, ys); !math.IsNaN(v) {
				out[name][metric] = v
			}
		}
	}
	return out
}

func pearsonCorr(xs, ys []float64) float64 {
	n := float64(len(xs))
	var sx, sy, sxy, sx2, sy2 float64
	for i := range xs {
		sx += xs[i]
		sy += ys[i]
		sxy += xs[i] * ys[i]
		sx2 += xs[i] * xs[i]
		sy2 += ys[i] * ys[i]
	}
	num := sxy - sx*sy/n
	den := math.Sqrt((sx2 - sx*sx/n) * (sy2 - sy*sy/n))
	if den == 0 {
		return 0
	}
	return num / den
}
