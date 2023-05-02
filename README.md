# goutils

Various functions that I use in projects

---

### Logging

Using 

- Zap logger by Uber, https://github.com/uber-go/zap
- Lumberjack https://pkg.go.dev/gopkg.in/natefinch/lumberjack.v2

__Levels__

- Debug
- Info
- Warn
- Error
- DPanic
- Panic

__Init__

```
var lg *zap.SugaredLogger

func init() {
	lg = log.InitWithConsole(zapcore.DebugLevel)
}
```

__Set level__

 `lg = log.SetDebugLevel()`

 `lg = log.SetInfoLevel()`

 etc...

 __Log__

 - Simple: `log.Info("hello")`
 - Formatted: `log.Infof("hello %s", "world")`
 - Structured: `log.Infow("hello", "var1", 12, "var2", 24)`

---

### Time

`time.StampTimeNow()`

returns time in RFC3339Nano

ex: `2023-04-27T11:33:45.772006+02:00`

---

### Network

Get local IP
`network.GetIP()`

---
### Pretty Print
_Only tried on simple structs_

Pretty print a struct

`pp.PP(someStruct)`

Return a pretty struct

`struct, err := pp.PrettyStruct(someStruct)`

---

### UUID

`uuid.GenerateUUID()`

Get a UUID using KSUID, https://github.com/segmentio/ksuid

Why this one?

Because someone said it was good, see https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html
