package layer

import (
	"strings"
)

type Layer struct {
	step         Step
	textscanned  string
	projectname  string
	webframework string
}

func NewLayer(cstep Step) *Layer {
	return &Layer{step: cstep}
}

func (layer *Layer) next(cstep Step) {
	layer.step = cstep
}

func (layer *Layer) do() {
	layer.step.Do(layer)
}

func (layer *Layer) JustDo(textscanned string) {
	layer.textscanned = strings.TrimSpace(textscanned)
	layer.step.Do(layer)
}
