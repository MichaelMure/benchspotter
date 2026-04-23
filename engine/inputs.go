package engine

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"math"
	"math/rand"
	"path/filepath"
	"strconv"
)

type InputInfo interface {
	Package() string
	Name() string
	// Label returns a short human-readable type annotation (e.g. "[int 0..100]").
	Label() string
	// Range returns the [min, max] span as float64 and whether the type has one.
	Range() (min, max float64, ok bool)
	// Midpoint returns a sensible central starting value.
	Midpoint() float64
	// SpacedValues returns up to n values spaced across the range (log-uniform
	// for log types). The result may be shorter than n when the range contains
	// fewer distinct values (e.g. bool always returns 2, a 3-value int range
	// returns at most 3).
	SpacedValues(n int) []float64
	// RandomValue returns a uniformly random (log-uniform for log types) value.
	RandomValue(rng *rand.Rand) float64
	// Perturb returns a neighbour of cur. sigma is the perturbation scale as a
	// fraction of the range (caller computes temp * perturbScale).
	Perturb(cur, sigma float64, rng *rand.Rand) float64
	// Format encodes a native value as a string for use in ParamAssignment.Value
	// (and ultimately as a BENCHSPOTTER_* environment variable).
	Format(v float64) string
}

// ── BoolType ──────────────────────────────────────────────────────────────────

type BoolType struct {
	pkg  string
	name string
}

func (b BoolType) Package() string                 { return b.pkg }
func (b BoolType) Name() string                    { return b.name }
func (b BoolType) Label() string                   { return "[bool]" }
func (b BoolType) Range() (float64, float64, bool) { return 0, 1, false }
func (b BoolType) Midpoint() float64               { return 0 }
func (b BoolType) SpacedValues(_ int) []float64    { return []float64{0, 1} }

func (b BoolType) RandomValue(rng *rand.Rand) float64 {
	if rng.Intn(2) == 0 {
		return 0
	}
	return 1
}

func (b BoolType) Perturb(cur, _ float64, _ *rand.Rand) float64 {
	if cur != 0 {
		return 0
	}
	return 1
}

func (b BoolType) Format(v float64) string {
	if v != 0 {
		return "true"
	}
	return "false"
}

// ── IntType ───────────────────────────────────────────────────────────────────

type IntType struct {
	pkg  string
	name string
	min  int
	max  int
}

func (t IntType) Package() string                 { return t.pkg }
func (t IntType) Name() string                    { return t.name }
func (t IntType) Label() string                   { return fmt.Sprintf("[int %d..%d]", t.min, t.max) }
func (t IntType) Range() (float64, float64, bool) { return float64(t.min), float64(t.max), true }
func (t IntType) Midpoint() float64               { return float64((t.min + t.max) / 2) }

func (t IntType) SpacedValues(n int) []float64 { return intSpaced(t.min, t.max, n, false) }

func (t IntType) RandomValue(rng *rand.Rand) float64 {
	return float64(t.min + rng.Intn(t.max-t.min+1))
}

func (t IntType) Perturb(cur, sigma float64, rng *rand.Rand) float64 {
	step := int(math.Ceil(sigma * float64(t.max-t.min)))
	if step < 1 {
		step = 1
	}
	v := int(math.Round(cur)) + rng.Intn(2*step+1) - step
	if v < t.min {
		v = t.min
	}
	if v > t.max {
		v = t.max
	}
	return float64(v)
}

func (t IntType) Format(v float64) string { return strconv.Itoa(int(math.Round(v))) }

// ── IntLogType ────────────────────────────────────────────────────────────────

// IntLogType is like IntType but benchspotter uses log-scale grid spacing.
type IntLogType struct {
	pkg  string
	name string
	min  int
	max  int
}

func (t IntLogType) Package() string                 { return t.pkg }
func (t IntLogType) Name() string                    { return t.name }
func (t IntLogType) Label() string                   { return fmt.Sprintf("[int-log %d..%d]", t.min, t.max) }
func (t IntLogType) Range() (float64, float64, bool) { return float64(t.min), float64(t.max), true }

func (t IntLogType) Midpoint() float64 {
	return math.Round(math.Sqrt(float64(t.min) * float64(t.max)))
}

func (t IntLogType) SpacedValues(n int) []float64 {
	if t.min <= 0 {
		return intSpaced(t.min, t.max, n, false)
	}
	return intSpaced(t.min, t.max, n, true)
}

func (t IntLogType) RandomValue(rng *rand.Rand) float64 {
	lMin, lMax := math.Log(float64(t.min)), math.Log(float64(t.max))
	v := int(math.Round(math.Exp(lMin + rng.Float64()*(lMax-lMin))))
	v = max(t.min, min(t.max, v))
	return float64(v)
}

func (t IntLogType) Perturb(cur, sigma float64, rng *rand.Rand) float64 {
	lMin, lMax := math.Log(float64(t.min)), math.Log(float64(t.max))
	lv := math.Log(cur) + rng.NormFloat64()*sigma*(lMax-lMin)
	lv = math.Max(lMin, math.Min(lMax, lv))
	nv := int(math.Round(math.Exp(lv)))
	nv = max(t.min, min(t.max, nv))
	return float64(nv)
}

func (t IntLogType) Format(v float64) string { return strconv.Itoa(int(math.Round(v))) }

// ── FloatType ─────────────────────────────────────────────────────────────────

// FloatType represents a float64 benchinput parameter with linear grid spacing.
type FloatType struct {
	pkg  string
	name string
	min  float64
	max  float64
}

func (t FloatType) Package() string                 { return t.pkg }
func (t FloatType) Name() string                    { return t.name }
func (t FloatType) Label() string                   { return fmt.Sprintf("[float %.4g..%.4g]", t.min, t.max) }
func (t FloatType) Range() (float64, float64, bool) { return t.min, t.max, true }
func (t FloatType) Midpoint() float64               { return (t.min + t.max) / 2 }
func (t FloatType) SpacedValues(n int) []float64    { return floatSpaced(t.min, t.max, n, false) }

func (t FloatType) RandomValue(rng *rand.Rand) float64 {
	return t.min + rng.Float64()*(t.max-t.min)
}

func (t FloatType) Perturb(cur, sigma float64, rng *rand.Rand) float64 {
	v := cur + rng.NormFloat64()*sigma*(t.max-t.min)
	return math.Max(t.min, math.Min(t.max, v))
}

func (t FloatType) Format(v float64) string { return strconv.FormatFloat(v, 'g', -1, 64) }

// ── FloatLogType ──────────────────────────────────────────────────────────────

// FloatLogType is like FloatType but benchspotter uses log-scale grid spacing.
type FloatLogType struct {
	pkg  string
	name string
	min  float64
	max  float64
}

func (t FloatLogType) Package() string                 { return t.pkg }
func (t FloatLogType) Name() string                    { return t.name }
func (t FloatLogType) Label() string                   { return fmt.Sprintf("[float-log %.4g..%.4g]", t.min, t.max) }
func (t FloatLogType) Range() (float64, float64, bool) { return t.min, t.max, true }
func (t FloatLogType) Midpoint() float64               { return math.Sqrt(t.min * t.max) }
func (t FloatLogType) SpacedValues(n int) []float64    { return floatSpaced(t.min, t.max, n, true) }

func (t FloatLogType) RandomValue(rng *rand.Rand) float64 {
	lMin, lMax := math.Log(t.min), math.Log(t.max)
	return math.Exp(lMin + rng.Float64()*(lMax-lMin))
}

func (t FloatLogType) Perturb(cur, sigma float64, rng *rand.Rand) float64 {
	lMin, lMax := math.Log(t.min), math.Log(t.max)
	lv := math.Log(cur) + rng.NormFloat64()*sigma*(lMax-lMin)
	lv = math.Max(lMin, math.Min(lMax, lv))
	return math.Exp(lv)
}

func (t FloatLogType) Format(v float64) string { return strconv.FormatFloat(v, 'g', -1, 64) }

// ── spacing helpers ───────────────────────────────────────────────────────────

// intSpaced returns n integers spaced across [lo, hi], either linearly or
// log-uniformly. Duplicates are skipped, so the result may be shorter than n.
func intSpaced(lo, hi, n int, logScale bool) []float64 {
	if lo == hi {
		return []float64{float64(lo)}
	}
	distinct := hi - lo + 1
	if n > distinct {
		n = distinct
	}
	if n <= 1 {
		return []float64{float64(lo)}
	}
	var vals []float64
	seen := make(map[int]bool)
	for i := 0; i < n; i++ {
		t := float64(i) / float64(n-1)
		var v int
		if logScale && lo > 0 {
			lLo, lHi := math.Log(float64(lo)), math.Log(float64(hi))
			v = int(math.Round(math.Exp(lLo + t*(lHi-lLo))))
		} else {
			v = lo + int(math.Round(t*float64(hi-lo)))
		}
		if v < lo {
			v = lo
		}
		if v > hi {
			v = hi
		}
		if !seen[v] {
			seen[v] = true
			vals = append(vals, float64(v))
		}
	}
	return vals
}

// floatSpaced returns n floats evenly spaced across [lo, hi], either
// linearly or log-uniformly.
func floatSpaced(lo, hi float64, n int, logScale bool) []float64 {
	if n <= 1 {
		return []float64{lo}
	}
	vals := make([]float64, n)
	for i := range vals {
		t := float64(i) / float64(n-1)
		if logScale && lo > 0 {
			lLo, lHi := math.Log(lo), math.Log(hi)
			vals[i] = math.Exp(lLo + t*(lHi-lLo))
		} else {
			vals[i] = lo + t*(hi-lo)
		}
	}
	return vals
}

// ── AST discovery ─────────────────────────────────────────────────────────────

func Inputs(rootDir string) (res []InputInfo, err error) {
	rootDir, err = filepath.Abs(rootDir)
	if err != nil {
		return nil, err
	}

	pkgs, err := packages(rootDir)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			ast.Inspect(file, func(n ast.Node) bool {
				call, ok := n.(*ast.CallExpr)
				if !ok {
					return true
				}
				sel, ok := call.Fun.(*ast.SelectorExpr)
				if !ok {
					return true
				}
				selType := pkg.TypesInfo.Uses[sel.Sel]
				if selType == nil {
					return true
				}
				fn, ok := selType.(*types.Func)
				if !ok {
					return true
				}
				if fn.Pkg() == nil || fn.Pkg().Path() != "benchspotter/benchinput" {
					return true
				}
				sig := fn.Type().(*types.Signature)

				switch fn.Name() {
				case "Bool": // name string, default_ bool
					if sig.Params().Len() != 2 {
						return true
					}
					param1 := sig.Params().At(0)
					if param1.Type().String() != "string" {
						return true
					}
					param2 := sig.Params().At(1)
					if param2.Type().String() != "bool" {
						return true
					}
					var info BoolType
					info.pkg, err = filepath.Rel(rootDir, pkg.Dir)
					if err != nil {
						return false
					}
					info.name, err = extractName(call)
					if err != nil {
						return false
					}
					res = append(res, info)
				case "Int", "IntLog": // name string, default_, min, max int
					if sig.Params().Len() != 4 {
						return true
					}
					param1 := sig.Params().At(0)
					if param1.Type().String() != "string" {
						return true
					}
					param2 := sig.Params().At(1)
					if param2.Type().String() != "int" {
						return true
					}
					param3 := sig.Params().At(2)
					if param3.Type().String() != "int" {
						return true
					}
					param4 := sig.Params().At(3)
					if param4.Type().String() != "int" {
						return true
					}
					var pkgRel string
					pkgRel, err = filepath.Rel(rootDir, pkg.Dir)
					if err != nil {
						return false
					}
					var iname string
					iname, err = extractName(call)
					if err != nil {
						return false
					}
					var imin, imax int
					imin, err = extractInt(call, 2)
					if err != nil {
						return false
					}
					imax, err = extractInt(call, 3)
					if err != nil {
						return false
					}
					if fn.Name() == "IntLog" {
						res = append(res, IntLogType{pkg: pkgRel, name: iname, min: imin, max: imax})
					} else {
						res = append(res, IntType{pkg: pkgRel, name: iname, min: imin, max: imax})
					}
				case "Float", "FloatLog": // name string, default_, min, max float64
					if sig.Params().Len() != 4 {
						return true
					}
					if sig.Params().At(0).Type().String() != "string" {
						return true
					}
					for _, i := range []int{1, 2, 3} {
						if sig.Params().At(i).Type().String() != "float64" {
							return true
						}
					}
					var pkgRel string
					pkgRel, err = filepath.Rel(rootDir, pkg.Dir)
					if err != nil {
						return false
					}
					var fname string
					fname, err = extractName(call)
					if err != nil {
						return false
					}
					var fmin, fmax float64
					fmin, err = extractFloat(call, 2)
					if err != nil {
						return false
					}
					fmax, err = extractFloat(call, 3)
					if err != nil {
						return false
					}
					if fn.Name() == "FloatLog" {
						res = append(res, FloatLogType{pkg: pkgRel, name: fname, min: fmin, max: fmax})
					} else {
						res = append(res, FloatType{pkg: pkgRel, name: fname, min: fmin, max: fmax})
					}
				}
				return true
			})
		}
	}
	return res, nil
}

func extractName(call *ast.CallExpr) (string, error) {
	lit, ok := call.Args[0].(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return "", fmt.Errorf("only string literals are supported for benchspotter inputs name")
	}
	name := lit.Value[1 : len(lit.Value)-1] // remove quotes
	return name, nil
}

func extractInt(call *ast.CallExpr, pos int) (int, error) {
	lit, ok := call.Args[pos].(*ast.BasicLit)
	if !ok || lit.Kind != token.INT {
		return 0, fmt.Errorf("only literals are supported for benchspotter inputs")
	}
	val, err := strconv.Atoi(lit.Value)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func extractFloat(call *ast.CallExpr, pos int) (float64, error) {
	arg := call.Args[pos]
	neg := false
	if u, ok := arg.(*ast.UnaryExpr); ok && u.Op == token.SUB {
		arg = u.X
		neg = true
	}
	lit, ok := arg.(*ast.BasicLit)
	if !ok {
		return 0, fmt.Errorf("only literals are supported for benchspotter inputs")
	}
	var v float64
	var err error
	switch lit.Kind {
	case token.INT:
		var i int
		i, err = strconv.Atoi(lit.Value)
		v = float64(i)
	case token.FLOAT:
		v, err = strconv.ParseFloat(lit.Value, 64)
	default:
		return 0, fmt.Errorf("expected numeric literal")
	}
	if err != nil {
		return 0, err
	}
	if neg {
		v = -v
	}
	return v, nil
}
