// Code generated by khulnasoft-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package simpleinvoke

import (
	"context"
	"reflect"

	"example.com/khulnasoft-simple-invoke/sdk/go/v10/simpleinvoke/internal"
	"github.com/khulnasoft/sdk/go/khulnasoft"
)

func MyInvoke(ctx *khulnasoft.Context, args *MyInvokeArgs, opts ...khulnasoft.InvokeOption) (*MyInvokeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv MyInvokeResult
	err := ctx.Invoke("simple-invoke:index:myInvoke", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type MyInvokeArgs struct {
	Value string `khulnasoft:"value"`
}

type MyInvokeResult struct {
	Result string `khulnasoft:"result"`
}

func MyInvokeOutput(ctx *khulnasoft.Context, args MyInvokeOutputArgs, opts ...khulnasoft.InvokeOption) MyInvokeResultOutput {
	return khulnasoft.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (MyInvokeResultOutput, error) {
			args := v.(MyInvokeArgs)
			options := khulnasoft.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("simple-invoke:index:myInvoke", args, MyInvokeResultOutput{}, options).(MyInvokeResultOutput), nil
		}).(MyInvokeResultOutput)
}

type MyInvokeOutputArgs struct {
	Value khulnasoft.StringInput `khulnasoft:"value"`
}

func (MyInvokeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MyInvokeArgs)(nil)).Elem()
}

type MyInvokeResultOutput struct{ *khulnasoft.OutputState }

func (MyInvokeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MyInvokeResult)(nil)).Elem()
}

func (o MyInvokeResultOutput) ToMyInvokeResultOutput() MyInvokeResultOutput {
	return o
}

func (o MyInvokeResultOutput) ToMyInvokeResultOutputWithContext(ctx context.Context) MyInvokeResultOutput {
	return o
}

func (o MyInvokeResultOutput) Result() khulnasoft.StringOutput {
	return o.ApplyT(func(v MyInvokeResult) string { return v.Result }).(khulnasoft.StringOutput)
}

func init() {
	khulnasoft.RegisterOutputType(MyInvokeResultOutput{})
}
