package main

import (
  "fmt"
  "app/wrappers"
  "github.com/renra/go-pseudoglobals/pseudoglobals"
  "github.com/renra/go-json-http-handler/jsonHttpHandler"
  "database/sql"
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
  db, _ := sql.Open("postgres", "whatever")

  globals := pseudoglobals.New(config, &wrappers.LoggerWrapper{}, map[string]interface{}{
    "postgres": db,
  })

  globals.Log("I have a logger")
  globals.Log(fmt.Sprintln("And I have a config too: %s", globals.Config().GetString("whatever")))

  globalsWrapper := wrappers.GlobalsWrapper{Globals: globals}
  globalsWrapper.Log(fmt.Sprintln("And I have a db client: %v", globalsWrapper.DB("postgres")))

  handler := jsonHttpHandler.New(
    &globalsWrapper,
    map[string]jsonHttpHandler.GlobalsReceivingHandlerFunc{},
  )

  globals.Log(fmt.Sprintln("And I have a handler for JSON APIs: %v", handler))
}

