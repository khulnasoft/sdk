module l2-ref-ref

go 1.20

require (
	github.com/khulnasoft/sdk v3.30.0
	example.com/khulnasoft-ref-ref/sdk/go/v12 v12.0.0
)

replace github.com/khulnasoft/sdk => /ROOT/artifacts/github.com_khulnasoft_khulnasoft_sdk_v3

replace example.com/khulnasoft-ref-ref/sdk/go/v12 => /ROOT/artifacts/example.com_khulnasoft-ref-ref_sdk_go_v12
