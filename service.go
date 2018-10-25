package main

import "github.com/graniticio/granitic"
import "worker-management/bindings"

func main() {
	granitic.StartGranitic(bindings.Components())
}
