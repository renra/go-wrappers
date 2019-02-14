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
  fmt.Println("Hello World")

  config := &Config{}

  globals := pseudoglobals.New(config, &wrappers.LoggerWrapper{})

  globals.Log("I have a logger")
  globals.Log(fmt.Sprintln("And I have a config too: %s", globals.Config().GetString("whatever")))

  handler := jsonHttpHandler.New(
    &wrappers.GlobalsWrapper{Globals: globals},
    map[string]jsonHttpHandler.GlobalsReceivingHandlerFunc{},
  )

  globals.Log(fmt.Sprintln("And I have a handler for JSON APIs: %v", handler))
}

