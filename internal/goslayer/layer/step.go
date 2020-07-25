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

type EntryProjectNameStep struct {
}

func (projectstep EntryProjectNameStep) Do(layer *Layer) {
	if layer.textscanned == "" {
		fmt.Print("Please enter your project name: ")
	} else {
		layer.projectname = layer.textscanned
		fmt.Println("\nPlease choose a web framework: ")
		fmt.Print("(1) use gin, (2) use handler buildin:")
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
