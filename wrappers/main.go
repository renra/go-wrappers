package wrappers

import (
  "database/sql"
  "github.com/renra/go-logger/logger"
  "github.com/renra/go-pseudoglobals/pseudoglobals"
  "github.com/renra/go-json-http-handler/jsonHttpHandler"
)

type LoggerWrapper struct {
}

func (lw LoggerWrapper) New(service string, severity int, severities map[int]string) pseudoglobals.LoggerInstance {
  return &logger.Logger{
    Label: service,
    ThresholdSeverity: severity,
    Severities: severities,
    Serialize: logger.GCPSerialize,
  }
}

type GlobalsWrapper struct {
  Globals *pseudoglobals.Pseudoglobals
}

func (gw *GlobalsWrapper) Config() jsonHttpHandler.Config {
  return gw.Globals.Config()
}

func (gw *GlobalsWrapper) Logger() jsonHttpHandler.Logger {
  return gw.Globals.Logger()
}

func (gw *GlobalsWrapper) Log(msg string) {
  gw.Globals.Log(msg)
}

func (gw *GlobalsWrapper) LogErrorWithTrace(err string, trace string) {
  gw.Globals.LogErrorWithTrace(err, trace)
}

func (gw *GlobalsWrapper) DB(key string) *sql.DB {
  client, found := gw.Globals.Clients()[key].(*sql.DB)

  if found == false {
    panic("Database client not found")
  }

  return client
}
