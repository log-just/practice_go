module main

go 1.15

require (
	github.com/labstack/echo/v4 v4.1.16
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/text v0.3.3 // indirect
	local/api/system v0.0.0
	local/api/user v0.0.0
	local/middleware v0.0.0
	local/util v0.0.0
)

replace (
	local/api/system v0.0.0 => ./api/system
	local/api/user v0.0.0 => ./api/user
	local/middleware v0.0.0 => ./middleware
	local/util v0.0.0 => ./util
)
