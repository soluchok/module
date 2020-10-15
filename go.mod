module github.com/soluchok/module

go 1.15

require (
	github.com/soluchok/module/submodule1 v0.1.0
	github.com/soluchok/module/submodule1/v2 v2.0.0-00010101000000-000000000000
	github.com/soluchok/module/submodule1/v3 v3.0.0-00010101000000-000000000000
	github.com/soluchok/module/submodule2 v0.1.0
	github.com/soluchok/module/submodule2/v2 v2.0.0-00010101000000-000000000000
	github.com/soluchok/module/submodule2/v3 v3.0.0-00010101000000-000000000000
)

replace (
	github.com/soluchok/module/submodule1/v2 => github.com/soluchok/module/submodule1 v0.2.0
	github.com/soluchok/module/submodule1/v3 => github.com/soluchok/module/submodule1 v0.3.0
	github.com/soluchok/module/submodule2/v2 => github.com/soluchok/module/submodule2 v0.2.0
	github.com/soluchok/module/submodule2/v3 => github.com/soluchok/module/submodule2 v0.3.0
)
