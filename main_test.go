package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/perf/benchfmt"
)

func TestFoo(t *testing.T) {
	cmd := exec.Command("go", "test", "-bench", "^\\QBenchmarkFoo\\E$", "-benchmem", "-run", "^$", "./...")
	cmd.Dir = filepath.Dir("./locate/internal/")
	cmd.Env = append(os.Environ(), fmt.Sprintf("BENCHSPOTTER_NAME_2=%d", 24))
	cmd.Environ()

	stdout, err := cmd.StdoutPipe()
	require.NoError(t, err)
	require.NoError(t, cmd.Start())

	out, err := os.Create("bench.txt")
	require.NoError(t, err)
	defer out.Close()

	w := benchfmt.NewWriter(out)

	r := benchfmt.NewReader(stdout, "")
	for r.Scan() {
		res := r.Result()
		switch res := res.(type) {
		case *benchfmt.Result:
			// weird dance to set a non-internal config, the API is not meant for that
			res.SetConfig("git-commit", "qsbdjkqsdnbkqqsd")
			idx, ok := res.ConfigIndex("git-commit")
			require.True(t, ok)
			res.Config[idx].File = true // not internal, meaning it will print in the writer
		}

		t.Log(res)
		err = w.Write(res)
		require.NoError(t, err)
	}

	require.NoError(t, cmd.Wait())
}
