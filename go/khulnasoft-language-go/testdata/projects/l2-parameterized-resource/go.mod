module l2-parameterized-resource

go 1.20

require (
	github.com/khulnasoft/sdk v3.30.0
	example.com/khulnasoft-subpackage/sdk/go/v2 v2.0.0
)

replace github.com/khulnasoft/sdk => /ROOT/artifacts/github.com_khulnasoft_khulnasoft_sdk_v3

replace example.com/khulnasoft-subpackage/sdk/go/v2 => /ROOT/artifacts/example.com_khulnasoft-subpackage_sdk_go_v2
