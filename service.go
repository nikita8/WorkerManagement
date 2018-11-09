package main

import "github.com/graniticio/granitic-yaml"
import "worker-management/bindings"

func main() {
	granitic_yaml.StartGraniticWithYaml(bindings.Components())
}
