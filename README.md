# Wrappers

This repo provides some knitting code between my other repos which are loosely coupled via interfaces.

## Usage

```go
package main

import (
  "fmt"
  "app/wrappers"
  "github.com/renra/go-pseudoglobals/pseudoglobals"
  "github.com/renra/go-json-http-handler/jsonHttpHandler"
)

type Config struct {
}

func (c *Config) Get(msg string) interface{} {
  return msg
}

func (c *Config) GetString(msg string) string {
  return msg
}

func main() {
  config := &Config{}

  globals := pseudoglobals.New(config, &wrappers.LoggerWrapper{})

  globals.Log("I have a logger")
  globals.Log(fmt.Sprintln("And I have a config too: %s", globals.Config().GetString("whatever")))

  globalsWrapper := wrappers.GlobalsWrapper{Globals: globals}
  globalsWrapper.Log(fmt.Sprintln("And I have a db client: %v", globalsWrapper.DB("postgres")))
  globalsWrapper.Log(fmt.Sprintln("And I could have many more clients too:: %v", globalsWrapper.Clients()))

  handler := jsonHttpHandler.New(
    &wrappers.GlobalsWrapper{Globals: globals},
    map[string]jsonHttpHandler.GlobalsReceivingHandlerFunc{},
  )

  globals.Log(fmt.Sprintln("And I have a handler for JSON APIs: %v", handler))
}

```

See the `examples` folder for working samples.

## References

* [Config management](https://github.com/renra/go-helm-config) (Not used here because it is instantiated separately and the instance is passed to `pseudoglobals.New`)
* [Logging](https://github.com/renra/go-logger)
* [Management of pseudo-global variables](https://github.com/renra/go-pseudoglobals)
* [Defaults for JSON HTTP API handlers](ihttps://github.com/renra/go-json-http-handler)
