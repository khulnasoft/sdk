module l2-provider-grpc-config

go 1.20

require (
	github.com/khulnasoft/sdk v3.30.0
	example.com/khulnasoft-config-grpc/sdk/go v1.0.0
)

replace example.com/khulnasoft-config-grpc/sdk/go => /ROOT/artifacts/example.com_khulnasoft-config-grpc_sdk_go

replace github.com/khulnasoft/sdk => /ROOT/artifacts/github.com_khulnasoft_khulnasoft_sdk_v3
