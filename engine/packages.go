package engine

import (
	"fmt"

	gopackages "golang.org/x/tools/go/packages"
)

var lastRootDir string
var lastPkgs []*gopackages.Package

// packages memoize the search and parsing of go packages
func packages(rootDir string) ([]*gopackages.Package, error) {
	if rootDir == lastRootDir {
		return lastPkgs, nil
	}

	cfg := &gopackages.Config{
		Mode:  gopackages.NeedSyntax | gopackages.NeedTypes | gopackages.NeedTypesInfo,
		Tests: true, // IMPORTANT: load test files too
	}

	pkgs, err := gopackages.Load(cfg, fmt.Sprintf("%s/...", rootDir))
	if err != nil {
		return nil, err
	}
	lastRootDir = rootDir
	lastPkgs = pkgs
	return pkgs, nil
}
