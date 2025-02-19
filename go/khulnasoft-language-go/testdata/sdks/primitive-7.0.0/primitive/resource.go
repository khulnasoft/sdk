// Code generated by khulnasoft-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package primitive

import (
	"context"
	"reflect"

	"errors"
	"example.com/khulnasoft-primitive/sdk/go/v7/primitive/internal"
	"github.com/khulnasoft/sdk/go/khulnasoft"
)

type Resource struct {
	khulnasoft.CustomResourceState

	Boolean     khulnasoft.BoolOutput         `khulnasoft:"boolean"`
	BooleanMap  khulnasoft.BoolMapOutput      `khulnasoft:"booleanMap"`
	Float       khulnasoft.Float64Output      `khulnasoft:"float"`
	Integer     khulnasoft.IntOutput          `khulnasoft:"integer"`
	NumberArray khulnasoft.Float64ArrayOutput `khulnasoft:"numberArray"`
	String      khulnasoft.StringOutput       `khulnasoft:"string"`
}

// NewResource registers a new resource with the given unique name, arguments, and options.
func NewResource(ctx *khulnasoft.Context,
	name string, args *ResourceArgs, opts ...khulnasoft.ResourceOption) (*Resource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Boolean == nil {
		return nil, errors.New("invalid value for required argument 'Boolean'")
	}
	if args.BooleanMap == nil {
		return nil, errors.New("invalid value for required argument 'BooleanMap'")
	}
	if args.Float == nil {
		return nil, errors.New("invalid value for required argument 'Float'")
	}
	if args.Integer == nil {
		return nil, errors.New("invalid value for required argument 'Integer'")
	}
	if args.NumberArray == nil {
		return nil, errors.New("invalid value for required argument 'NumberArray'")
	}
	if args.String == nil {
		return nil, errors.New("invalid value for required argument 'String'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Resource
	err := ctx.RegisterResource("primitive:index:Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResource gets an existing Resource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResource(ctx *khulnasoft.Context,
	name string, id khulnasoft.IDInput, state *ResourceState, opts ...khulnasoft.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("primitive:index:Resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resource resources.
type resourceState struct {
}

type ResourceState struct {
}

func (ResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceState)(nil)).Elem()
}

type resourceArgs struct {
	Boolean     bool            `khulnasoft:"boolean"`
	BooleanMap  map[string]bool `khulnasoft:"booleanMap"`
	Float       float64         `khulnasoft:"float"`
	Integer     int             `khulnasoft:"integer"`
	NumberArray []float64       `khulnasoft:"numberArray"`
	String      string          `khulnasoft:"string"`
}

// The set of arguments for constructing a Resource resource.
type ResourceArgs struct {
	Boolean     khulnasoft.BoolInput
	BooleanMap  khulnasoft.BoolMapInput
	Float       khulnasoft.Float64Input
	Integer     khulnasoft.IntInput
	NumberArray khulnasoft.Float64ArrayInput
	String      khulnasoft.StringInput
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

type ResourceInput interface {
	khulnasoft.Input

	ToResourceOutput() ResourceOutput
	ToResourceOutputWithContext(ctx context.Context) ResourceOutput
}

func (*Resource) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (i *Resource) ToResourceOutput() ResourceOutput {
	return i.ToResourceOutputWithContext(context.Background())
}

func (i *Resource) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return khulnasoft.ToOutputWithContext(ctx, i).(ResourceOutput)
}

type ResourceOutput struct{ *khulnasoft.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

func (o ResourceOutput) Boolean() khulnasoft.BoolOutput {
	return o.ApplyT(func(v *Resource) khulnasoft.BoolOutput { return v.Boolean }).(khulnasoft.BoolOutput)
}

func (o ResourceOutput) BooleanMap() khulnasoft.BoolMapOutput {
	return o.ApplyT(func(v *Resource) khulnasoft.BoolMapOutput { return v.BooleanMap }).(khulnasoft.BoolMapOutput)
}

func (o ResourceOutput) Float() khulnasoft.Float64Output {
	return o.ApplyT(func(v *Resource) khulnasoft.Float64Output { return v.Float }).(khulnasoft.Float64Output)
}

func (o ResourceOutput) Integer() khulnasoft.IntOutput {
	return o.ApplyT(func(v *Resource) khulnasoft.IntOutput { return v.Integer }).(khulnasoft.IntOutput)
}

func (o ResourceOutput) NumberArray() khulnasoft.Float64ArrayOutput {
	return o.ApplyT(func(v *Resource) khulnasoft.Float64ArrayOutput { return v.NumberArray }).(khulnasoft.Float64ArrayOutput)
}

func (o ResourceOutput) String() khulnasoft.StringOutput {
	return o.ApplyT(func(v *Resource) khulnasoft.StringOutput { return v.String }).(khulnasoft.StringOutput)
}

func init() {
	khulnasoft.RegisterInputType(reflect.TypeOf((*ResourceInput)(nil)).Elem(), &Resource{})
	khulnasoft.RegisterOutputType(ResourceOutput{})
}
