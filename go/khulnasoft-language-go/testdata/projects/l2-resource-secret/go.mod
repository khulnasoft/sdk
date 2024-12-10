module l2-resource-secret

go 1.20

require (
	github.com/khulnasoft/sdk v3.30.0
	example.com/khulnasoft-secret/sdk/go/v14 v14.0.0
)

replace github.com/khulnasoft/sdk => /ROOT/artifacts/github.com_khulnasoft_khulnasoft_sdk_v3

replace example.com/khulnasoft-secret/sdk/go/v14 => /ROOT/artifacts/example.com_khulnasoft-secret_sdk_go_v14
