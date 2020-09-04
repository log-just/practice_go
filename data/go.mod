module local/data

go 1.15

require (
	gorm.io/driver/mysql v1.0.1
	gorm.io/gorm v1.20.0
	local/util v0.0.0
)

replace local/util v0.0.0 => ../util
