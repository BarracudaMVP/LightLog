# lightlog

Лёгкая обёртка над стандартным `log.Logger` с фильтрацией сообщений по уровню
логирования: `Error`, `Warning`, `Info`.

## Установка

```sh
go get github.com/barracudamvp/lightlog
```

## Использование

```go
package main

import "github.com/barracudamvp/lightlog"

func main() {
	logger := lightlog.New()
	logger.SetLogLevel(lightlog.LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}
```

Вывод (уровень `Info` отфильтрован, т.к. установлен уровень `Warning`):

```
2026/06/19 03:16:33 WARN Hello
2026/06/19 03:16:33 ERR  World
2026/06/19 03:16:33 Debug
```

## Уровни логирования

Сообщение печатается, только если его уровень не выше текущего уровня логгера.
Уровни упорядочены от наименее к наиболее подробному:

| Уровень           | Метод     | Префикс |
| ----------------- | --------- | ------- |
| `LogLevelError`   | `Errorln` | `ERR `  |
| `LogLevelWarning` | `Warnln`  | `WARN`  |
| `LogLevelInfo`    | `Infoln`  | `INFO`  |

Уровень по умолчанию — `LogLevelError`. Сменить его можно через `SetLogLevel`.
`Logger` встраивает `*log.Logger`, поэтому все стандартные методы (`Println`,
`Printf` и т.д.) доступны напрямую и выводятся без фильтрации по уровню.

## Лицензия

[MIT](LICENSE)
