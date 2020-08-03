package layer

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/danceyoung/goslayer/internal/goslayer/layer/template"
)

const (
	ginWebFramework         string = "GIN"
	httpHandlerWebFramework string = "HttpHandler"
)

const stringReplacedInImportPath string = "github.com/danceyoung/goslayer"

func newTemplate(webframework string) template.Template {
	if webframework == httpHandlerWebFramework {
		return template.HttpHandlerTemplate{}
	}
	return template.GINTemplate{}
}

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
		fmt.Println("\nPlease choose a web framework, ")
		fmt.Print("(1) use gin, (2) use handler buildin: ")
		layer.next(ChooseWebFStep{})
	}
}

type ChooseWebFStep struct {
}

func (choosestep ChooseWebFStep) Do(layer *Layer) {

	if layer.textscanned == "1" {
		fmt.Println("Creating your go project with GIN")
		layer.webframework = ginWebFramework
	} else if layer.textscanned == "2" {
		fmt.Println("Creating your go project with http.Handler")
		layer.webframework = httpHandlerWebFramework
	} else {
		fmt.Println("Creating your go project with GIN.")
		layer.webframework = ginWebFramework
	}
	time.Sleep(time.Second * 1)
	layer.next(CreateStep{})
	layer.do()
}

type CreateStep struct{}

func (createstep CreateStep) Do(layer *Layer) {
	os.MkdirAll("./"+layer.projectname+"/cmd/myapp/router/handler", os.ModePerm)
	routerfile, err := os.Create("./" + layer.projectname + "/cmd/myapp/router/router.go")
	if err != nil {
		fmt.Println("panic a error when creating project: ", err.Error())
		panic(nil)
	}
	temp := newTemplate(layer.webframework)

	routerfile.WriteString(strings.ReplaceAll(temp.RouterTemplate(), stringReplacedInImportPath, layer.projectname))
	basehanderfile, err := os.Create("./" + layer.projectname + "/cmd/myapp/router/handler/basehandler.go")
	if err != nil {
		fmt.Println("panic a error when creating project: ", err.Error())
		panic(nil)
	}
	basehanderfile.WriteString(temp.BaseHandlerTemplate())

	eventhanderfile, err := os.Create("./" + layer.projectname + "/cmd/myapp/router/handler/eventhandler.go")
	if err != nil {
		fmt.Println("panic a error when creating project: ", err.Error())
		panic(nil)
	}
	eventhanderfile.WriteString(strings.ReplaceAll(temp.EventHandlerTemplate(), stringReplacedInImportPath, layer.projectname))

	mainfile, err := os.Create("./" + layer.projectname + "/cmd/myapp/main.go")
	if err != nil {
		fmt.Println("panic a error when creating project: ", err.Error())
		panic(nil)
	}
	mainfile.WriteString(strings.ReplaceAll(temp.MainTemplate(), stringReplacedInImportPath, layer.projectname))

	os.MkdirAll("./"+layer.projectname+"/internal/myapp", os.ModePerm)
	os.MkdirAll("./"+layer.projectname+"/internal/pkg/middleware", os.ModePerm)
	httsetmwfile, err := os.Create("./" + layer.projectname + "/internal/pkg/middleware/httpset.go")
	if err != nil {
		fmt.Println("panic a error when creating project: ", err.Error())
		panic(nil)
	}
	httsetmwfile.WriteString(temp.HttpMiddlewareTemplate())

	os.MkdirAll("./"+layer.projectname+"/internal/myapp/event", os.ModePerm)
	eventbizfile, err := os.Create("./" + layer.projectname + "/internal/myapp/event/event.go")
	if err != nil {
		fmt.Println("panic a error when creating project: ", err.Error())
		panic(nil)
	}
	eventbizfile.WriteString(temp.EventBizTemplate())
	os.Mkdir("./"+layer.projectname+"/pkg", os.ModePerm)

	layer.next(OverStep{})
	layer.do()
}

type OverStep struct{}

func (overstep OverStep) Do(layer *Layer) {
	// cmd := exec.Command("go", "mod", "init", "ddd")
	// fmt.Println(cmd.Run())
	var projectstructure = `github.com/danceyoung/goslayer
├── cmd/
│   └── myapp/
│       └── router/
│           └── handler/
│           └── router.go
│       └── main.go
├── internal/
│   └── myapp/
│       └── event/
├── └── pkg/
│       └── middleware/`
	fmt.Println("The go project is created successfully.")
	fmt.Println(strings.ReplaceAll(projectstructure, stringReplacedInImportPath, layer.projectname))
	panic(nil)
}
