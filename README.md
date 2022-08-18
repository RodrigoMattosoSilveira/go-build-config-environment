# go-build-config-environment
A sample app to show how to configure a go application build environment

We use two go features, `embeded` and `Command-Line Flags`.

# Command-Line Flags
`Command-line flags` are a common way to specify options for command-line programs. In our case we will use the `-env <environment name>` flag, supporting the following options:
* development
* staging
* acceptance
* production

See [Go by Example: Command-Line Flags](https://gobyexample.com/command-line-flags) for details.

# Environment Configuration Files
We will host JSON-formatted  Environment Configuration Files in the `environments` folder, and use `embeded` functionality to import and set the variables of the environment provided in the `-env` flag. We will have the following files:
* development.json
* staging.json
* acceptance.json
* production.json

See [Including and reading static files with embed directive at compile time in Golang](http://www.inanzzz.com/index.php/post/1rwm/including-and-reading-static-files-with-embed-directive-at-compile-time-in-golang) for details. Note that we read only the file passed as command line flag.

# Workflow
## Create Configuration Files
Create the JSON-formatted Environment Configuration Files in the `environments` folder. As an example:
```json
{
  "environment": "development",
  "port": 50000
}
```

Note that we will have to create a `structure` to host the selected file and that it will have to `know` about the `JSON file keys`.

## Set up
Import the required go libraries:
```go
import (
        "embed"
        "encoding/json"
        "flag"
        "fmt"
        "log"
    )
````

## Environment Configuration Structure
Define a `structure` to host the selected file and that it will have to `know` about the `JSON file keys`:
```go
type environment struct {
Env  string `json:"environment"`
Port int    `json:"port"`
}
````

## Embed folder
Configure the embed folder:
```go
//go:embed environments
var environmentsDir embed.FS
````

## Parse command line flag
```go
	envPtr := flag.String("env", "development", "the build environment")
	flag.Parse()
	fmt.Println("environment:", *envPtr)
	configurationFileName = *envPtr + ".json"
```

**NOTE** that we also build the `configurationFileName` variable, required to select which file to read, since the command line flag does not include the `json` extention.

# Environment Configuration files
Collect the Environment Configuration file names and read the one of interest:
```go
	files, err := environmentsDir.ReadDir("environments")
    if err != nil {
        log.Fatalln(err)
    }
    
    for _, file := range files {
        if file.Name() == configurationFileName {
            val, err := environmentsDir.ReadFile("environments/" + file.Name())
            if err != nil {
            fmt.Println(err)
            continue
            }
            
            var environment environment
            if err := json.Unmarshal(val, &environment); err != nil {
            fmt.Println(err)
            continue
            }
            
            fmt.Printf("%+v\n", environment)
        }
    }

```

# Example
To build for the application for development, run the following:

```bash
$ go run -race main.go -env development

environment: development
{Env:development Port:50000}
```

Note that if you do not provide the `-env` flag, the application uses `-env development` as the default.



