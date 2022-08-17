# go-config-environment
A sample app to show how to configure the environment

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

# Example
To build for the application for development, run the following:

```bash
$ go run -race main.go -env development

environment: development
{Env:development Port:50000}
```

Note that if you do not provide the `-env` flag, the application uses `-env development` as the default.



