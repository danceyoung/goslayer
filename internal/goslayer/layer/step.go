package layer

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Step interface {
	Do(*Layer)
}

type DescStep struct {
}

func (descstep DescStep) Do(layer *Layer) {
	if layer.textscanned == "" {
		fmt.Println("GoSLayer is a tool that helps you to create a golang project that is layered base on a standard architecture layout and followed by Package-Oriented-Design guideline.")
		fmt.Print("Please enter your project name: ")
	} else {
		layer.projectname = layer.textscanned
		fmt.Println("Please choose a web framework: ")
		fmt.Println("[1] use gin")
		fmt.Println("[2] use http's handler buildin golang")
		fmt.Print("entry 1 or 2: ")
		layer.next(ChooseWebFStep{})
	}
}

type ChooseWebFStep struct {
}

func (choosestep ChooseWebFStep) Do(layer *Layer) {
	layer.webframework = layer.textscanned
	if layer.webframework == "1" {
		fmt.Println("Creating your go project with GIN")
	} else if layer.webframework == "2" {
		fmt.Println("Creating your go project with http.Handler")
	} else {
		fmt.Println("Creating your go project without any web framework.")
	}
	time.Sleep(time.Second * 1)
	layer.next(CreateStep{})
	layer.do()
}

type CreateStep struct{}

func (createstep CreateStep) Do(layer *Layer) {
	os.Mkdir("./"+layer.projectname, os.ModePerm)
	os.MkdirAll("./"+layer.projectname+"/cmd/myapp", os.ModePerm)
	os.MkdirAll("./"+layer.projectname+"/cmd/myapp/router/handler", os.ModePerm)
	routerfile, err := os.Create("./" + layer.projectname + "/cmd/myapp/router/router.go")

	temp := newTemplate(layer.webframework)

	if err != nil {
		fmt.Println("panic a error when creating project: ", err.Error())
		panic(nil)
	}
	if layer.webframework == "2" {
		routerfile.WriteString(strings.ReplaceAll(temp.routerTemplate(), "goslayer", layer.projectname))
		basehanderfile, err := os.Create("./" + layer.projectname + "/cmd/myapp/router/handler/basehandler.go")
		if err != nil {
			fmt.Println("panic a error when creating project: ", err.Error())
			panic(nil)
		}
		basehanderfile.WriteString(temp.baseHandlerTemplate())

		eventhanderfile, err := os.Create("./" + layer.projectname + "/cmd/myapp/router/handler/eventhandler.go")
		if err != nil {
			fmt.Println("panic a error when creating project: ", err.Error())
			panic(nil)
		}
		eventhanderfile.WriteString(temp.eventHandlerTemplate())

		mainfile, err := os.Create("./" + layer.projectname + "/cmd/myapp/main.go")
		if err != nil {
			fmt.Println("panic a error when creating project: ", err.Error())
			panic(nil)
		}
		mainfile.WriteString(strings.ReplaceAll(temp.mainTemplate(), "goslayer", layer.projectname))
	}

	os.MkdirAll("./"+layer.projectname+"/internal/myapp", os.ModePerm)
	os.MkdirAll("./"+layer.projectname+"/internal/pkg/middleware", os.ModePerm)
	httsetmwfile, err := os.Create("./" + layer.projectname + "/internal/pkg/middleware/httpset.go")
	if err != nil {
		fmt.Println("panic a error when creating project: ", err.Error())
		panic(nil)
	}
	httsetmwfile.WriteString(temp.httpMiddlewareTemplate())
	os.Mkdir("./"+layer.projectname+"/pkg", os.ModePerm)

	layer.next(OverStep{})
	layer.do()
}

type OverStep struct{}

func (overstep OverStep) Do(layer *Layer) {
	fmt.Println("The go project is created successfully.")
	panic(nil)
}
