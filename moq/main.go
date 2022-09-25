package main

import (
	"fmt"
	"os"

	"github.com/matryer/moq/pkg/moq"
)

func main() {
	var (
		srcDir     = "hoge"
		pkgName    = "moqs"
		formatter  = "goimports"
		stubImpl   = false
		skipEnsure = false

		out  = os.Stdout
		args = []string{"Hoge"}
	)

	m, err := moq.New(moq.Config{
		SrcDir:     srcDir,
		PkgName:    pkgName,
		Formatter:  formatter,
		StubImpl:   stubImpl,
		SkipEnsure: skipEnsure,
	})
	if err != nil {
		panic(err)
	}

	if err = m.Mock(out, args...); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", m)
}
