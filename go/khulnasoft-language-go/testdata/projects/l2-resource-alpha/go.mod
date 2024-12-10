module l2-resource-alpha

go 1.20

require (
	github.com/khulnasoft/sdk v3.30.0
	example.com/khulnasoft-alpha/sdk/go/v3 v3.0.0-alpha.1.internal
)

replace example.com/khulnasoft-alpha/sdk/go/v3 => /ROOT/artifacts/example.com_khulnasoft-alpha_sdk_go_v3

replace github.com/khulnasoft/sdk => /ROOT/artifacts/github.com_khulnasoft_khulnasoft_sdk_v3
