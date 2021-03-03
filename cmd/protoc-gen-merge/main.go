package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		MakeMerge(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}

type Merger interface {
	Merge(interface{})
}
