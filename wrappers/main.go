package wrappers

import (
  "github.com/renra/go-errtrace/errtrace"
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

func (gw *GlobalsWrapper) LogErrorWithTrace(e *errtrace.Error) {
  gw.Globals.LogErrorWithTrace(e)
}

func (gw *GlobalsWrapper) Clients() map[string]interface{} {
  return gw.Globals.Clients()
}
