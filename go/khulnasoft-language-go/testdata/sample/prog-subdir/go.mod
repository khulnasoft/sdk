module example.com/prog-subdir

go 1.18

require (
	github.com/khulnasoft/go-dependency-testdata/dep v1.6.0
	github.com/khulnasoft/go-dependency-testdata/plugin v1.2.3
)

require github.com/khulnasoft/go-dependency-testdata/indirect-dep/v2 v2.1.0 // indirect
