## Simple log wrapper

#### Create new logger
```go
yourLogger := logwrapper.NewLogger(logrus.DebugLevel, []logrus.Hook{})
```

#### It is based on logrus and GELF format. To change format make:
```go
yourLogger.SetFormatter(formatter)
```

#### You can set hooks. For example, you can push logs to Graylog
```go
import graylog "github.com/gemnasium/logrus-graylog-hook/v3"

hooks := []logrus.Hook{graylog.NewGraylogHook("graylog:12201", map[string]interface{}{})}
logwrapper.Logger = logwrapper.NewLogger(logrus.DebugLevel, hooks)
```

#### Log examples
```go
yourLogger := logwrapper.NewLogger(logrus.DebugLevel, []logrus.Hook{})

yourLogger.Payload(logwrapper.NewPayload().Op("NewRepository").Package("repository")).Fatal(err)
yourLogger.Payload(logwrapper.NewPayload().Op("Find").Package("repository")).Error(err)
yourLogger.Payload(logwrapper.NewPayload().Package("main")).Info("Starting server")
```
