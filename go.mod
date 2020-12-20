module github.com/allep/cpp_dependency_analyzer

replace github.com/allep/cpp_dependency_analyzer/core => ./core

replace github.com/allep/cpp_dependency_analyzer/model => ./model

replace github.com/allep/cpp_dependency_analyzer/view => ./view

replace github.com/allep/cpp_dependency_analyzer/controller => ./controller

go 1.14

require (
	github.com/allep/cpp_dependency_analyzer/controller v0.0.0-00010101000000-000000000000
	github.com/allep/cpp_dependency_analyzer/core v0.0.0-00010101000000-000000000000
	github.com/allep/cpp_dependency_analyzer/model v0.0.0-00010101000000-000000000000
	github.com/allep/cpp_dependency_analyzer/view v0.0.0-00010101000000-000000000000
)
