package main

import (
	"bufio"
	"goslayer/internal/goslayer/layer"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err == nil {

		}
	}()

	layer := layer.NewLayer(layer.DescStep{})
	layer.JustDo("")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		layer.JustDo(scanner.Text())
		// 		textScanned := scanner.Text()
		// 		if len(textScanned) == 0 {
		// 			fmt.Print("Please entry your project name: ")
		// 		} else {
		// 			projectName := scanner.Text()
		// 			os.Mkdir("./"+projectName, os.ModePerm)
		// 			fmt.Println(os.Mkdir("./"+projectName+"/cmd", os.ModePerm))
		// 			os.Mkdir("./"+projectName+"/cmd/myapp", os.ModePerm)
		// 			os.Mkdir("./"+projectName+"/cmd/myapp/router", os.ModePerm)
		// 			os.Mkdir("./"+projectName+"/cmd/myapp/router/handler", os.ModePerm)
		// 			mainfile, err := os.Create("./" + projectName + "/cmd/myapp/main.go")
		// 			if err != nil {
		// 				fmt.Println("Raise a error when creating project: ", err.Error())
		// 				break
		// 			}
		// 			maincontent := `package main

		// func main() {

		// }`
		// 			fmt.Println(mainfile.WriteString(maincontent))

		// 			os.Mkdir("./"+projectName+"/internal", os.ModePerm)
		// 			os.Mkdir("./"+projectName+"/internal/myapp", os.ModePerm)
		// 			os.Mkdir("./"+projectName+"/internal/pkg", os.ModePerm)

		// 			os.Mkdir("./"+projectName+"/pkg", os.ModePerm)

		// 			dirstate := `
		// paper-code/examples/groupevent
		// ├── cmd/
		// │   └── myapp/
		// │       └── router/
		// │           └── handler/
		// │           └── router.go
		// │       └── main.go
		// ├── internal/
		// │   └── myapp/
		// │   └── pkg/
		// ├── pkg/
		// ├── config/
		// 		`
		// 			fmt.Println(dirstate)
		// 		}

	}

}
