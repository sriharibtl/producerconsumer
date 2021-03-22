module github.com/sriharibtl/consumer

go 1.15

require (
	httpserver v0.0.0
	kafka v0.0.0
	cbapi v0.0.0
)

replace (
	httpserver v0.0.0 => ./httpserver
	kafka v0.0.0 => ./kafka
	cbapi v0.0.0 => ./cbapi
)
