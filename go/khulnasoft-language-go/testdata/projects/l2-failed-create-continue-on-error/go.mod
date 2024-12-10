module l2-failed-create-continue-on-error

go 1.20

require (
	github.com/khulnasoft/sdk v3.30.0
	example.com/khulnasoft-fail_on_create/sdk/go/v4 v4.0.0
	example.com/khulnasoft-simple/sdk/go/v2 v2.0.0
)

replace example.com/khulnasoft-fail_on_create/sdk/go/v4 => /ROOT/artifacts/example.com_khulnasoft-fail_on_create_sdk_go_v4

replace github.com/khulnasoft/sdk => /ROOT/artifacts/github.com_khulnasoft_khulnasoft_sdk_v3

replace example.com/khulnasoft-simple/sdk/go/v2 => /ROOT/artifacts/example.com_khulnasoft-simple_sdk_go_v2
