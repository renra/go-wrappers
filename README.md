# Wrappers

This repo provides some knitting code between my other repos which are loosely coupled via interfaces.

## Usage

```go
package main

import (
  "fmt"
  "app/wrappers"
  "github.com/renra/go-errtrace/errtrace"
  "github.com/renra/go-pseudoglobals/pseudoglobals"
  "github.com/renra/go-json-http-handler/jsonHttpHandler"
  "database/sql"
)

type Config struct {
}

func (ci *Config) Get(key string) (interface{}, *errtrace.Error) {
  return key, nil;
}

func (ci *Config) GetP(key string) interface{} {
  return key;
}

func (ci *Config) GetString(key string) (string, *errtrace.Error) {
  return key, nil;
}

func (ci *Config) GetStringP(key string) string {
  return key;
}

func (ci *Config) GetInt(key string) (int, *errtrace.Error) {
  return 4, nil;
}

func (ci *Config) GetIntP(key string) int {
  return 4;
}

func (ci *Config) GetFloat(key string) (float64, *errtrace.Error) {
  return 3.14, nil;
}

func (ci *Config) GetFloatP(key string) float64 {
  return 3.14;
}

func (ci *Config) GetBool(key string) (bool, *errtrace.Error) {
  return true, nil;
}

func (ci *Config) GetBoolP(key string) bool {
  return true;
}

func main() {
  config := &Config{}
  db, _ := sql.Open("postgres", "whatever")

  globals := pseudoglobals.New(config, &wrappers.LoggerWrapper{}, map[string]interface{}{
    "postgres": db,
  })

  globals.Log("I have a logger")
  globals.Log(fmt.Sprintln("And I have a config too: %s", globals.Config().GetStringP("whatever")))

  globalsWrapper := wrappers.GlobalsWrapper{Globals: globals}
  globalsWrapper.Log(fmt.Sprintln("And I could have more clients too:: %v", globalsWrapper.Clients()))

  handler := jsonHttpHandler.New(
    &globalsWrapper,
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
