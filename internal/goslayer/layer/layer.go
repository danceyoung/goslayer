package layer

import (
	"fmt"
	"os"
	"strings"
)

type Step interface {
	Do(*Layer)
}

type DescStep struct {
}

func (descstep DescStep) Do(layer *Layer) {
	if strings.TrimSpace(layer.steptext) == "" {
		fmt.Println("GoSLayer is a tool that helps you to create a golang project that is layered base on a standard architecture layout and followed by Package-Oriented-Design guideline.")
		fmt.Print("Please enter your project name: ")
	} else {
		layer.NextStepAndDo(CreateStep{})
	}
}

type CreateStep struct{}

func (createstep CreateStep) Do(layer *Layer) {
	projectName := layer.steptext
	os.Mkdir("./"+projectName, os.ModePerm)
	os.MkdirAll("./"+projectName+"/cmd/myapp", os.ModePerm)
	os.MkdirAll("./"+projectName+"/cmd/myapp/router/handler", os.ModePerm)
	mainfile, err := os.Create("./" + projectName + "/cmd/myapp/main.go")
	if err != nil {
		fmt.Println("Raise a error when creating project: ", err.Error())
		panic(nil)
	}

	mainfile.WriteString(mainTemplate())

	os.MkdirAll("./"+projectName+"/internal/myapp", os.ModePerm)
	os.Mkdir("./"+projectName+"/internal/pkg", os.ModePerm)

	os.Mkdir("./"+projectName+"/pkg", os.ModePerm)

	layer.NextStepAndDo(OverStep{})
}

type OverStep struct{}

func (overstep OverStep) Do(layer *Layer) {
	fmt.Println("The go project is created successfully.")
	panic(nil)
}

type Layer struct {
	step     Step
	steptext string
}

func NewLayer(cstep Step) *Layer {
	return &Layer{step: cstep}
}

func (layer *Layer) NextStepAndDo(cstep Step) {
	layer.step = cstep
	layer.step.Do(layer)
}

func (layer *Layer) JustDo(steptext string) {
	layer.steptext = steptext
	layer.step.Do(layer)
}
