module local/api/user

go 1.15

require (
    github.com/labstack/echo/v4 v4.1.16
    local/route v0.0.0
)

replace (
	local/route v0.0.0 => ../../route
)
