package layer

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/danceyoung/goslayer/internal/goslayer/layer/template"
)

const (
	ginWebFramework         string = "GIN"
	httpHandlerWebFramework string = "HttpHandler"
)

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
		fmt.Println("\nPlease choose a web framework: ")
		fmt.Print("(1) use gin, (2) use handler buildin:")
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

	routerfile.WriteString(strings.ReplaceAll(temp.RouterTemplate(), "goslayer", layer.projectname))
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
	eventhanderfile.WriteString(strings.ReplaceAll(temp.EventHandlerTemplate(), "goslayer", layer.projectname))

	mainfile, err := os.Create("./" + layer.projectname + "/cmd/myapp/main.go")
	if err != nil {
		fmt.Println("panic a error when creating project: ", err.Error())
		panic(nil)
	}
	mainfile.WriteString(strings.ReplaceAll(temp.MainTemplate(), "goslayer", layer.projectname))

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
	cmd := exec.Command("go", "mod", "init", "ddd")
	fmt.Println(cmd.Run())
	fmt.Println("The go project is created successfully.")
	panic(nil)
}
