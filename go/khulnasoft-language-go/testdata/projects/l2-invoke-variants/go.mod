module l2-invoke-variants

go 1.20

require (
	github.com/khulnasoft/sdk v3.30.0
	example.com/khulnasoft-simple-invoke/sdk/go/v10 v10.0.0
)

replace github.com/khulnasoft/sdk => /ROOT/artifacts/github.com_khulnasoft_khulnasoft_sdk_v3

replace example.com/khulnasoft-simple-invoke/sdk/go/v10 => /ROOT/artifacts/example.com_khulnasoft-simple-invoke_sdk_go_v10
