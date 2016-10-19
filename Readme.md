# generate-composite-types

This tool generates `sql.Scanner` and `sql.Valuer` interfaces for your own structs and type aliases. It is tested with Postgres and `lib/pq` library, others might work. The command is quite self explanatory, use the `--help` option to understand the parameters.

*Note*
After generating the code use [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) to fix import problems and use `gofmt` to format the files.

## Trying it out

The `example` folder provides you with tests and structs to understand how this tool works. Run `go generate` in the example folder to generate the interfaces. Then you can run `go test` to see if all works. It tries to connect to a local postgres instance and needs a database named ﻿⁠⁠⁠⁠`custom`.

The ﻿`-array`﻿⁠⁠⁠⁠ param generates array types as well, the ﻿`-sql`﻿⁠⁠⁠⁠ param adds a Init<struct> function that can be used to create the composite type in postgres.

## Aliases
There is support for type aliases. To generate scanner and valuer for type aliases you can pass them through the `alias` param. 
