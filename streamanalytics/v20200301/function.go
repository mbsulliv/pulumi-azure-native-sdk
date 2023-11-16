// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A function object, containing all information associated with the named function. All functions are contained under a streaming job.
type Function struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The properties that are associated with a function.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFunction registers a new resource with the given unique name, arguments, and options.
func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOption) (*Function, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobName == nil {
		return nil, errors.New("invalid value for required argument 'JobName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:streamanalytics:Function"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20160301:Function"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20170401preview:Function"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20211001preview:Function"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Function
	err := ctx.RegisterResource("azure-native:streamanalytics/v20200301:Function", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunction gets an existing Function resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionState, opts ...pulumi.ResourceOption) (*Function, error) {
	var resource Function
	err := ctx.ReadResource("azure-native:streamanalytics/v20200301:Function", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Function resources.
type functionState struct {
}

type FunctionState struct {
}

func (FunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionState)(nil)).Elem()
}

type functionArgs struct {
	// The name of the function.
	FunctionName *string `pulumi:"functionName"`
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// Resource name
	Name *string `pulumi:"name"`
	// The properties that are associated with a function.
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Function resource.
type FunctionArgs struct {
	// The name of the function.
	FunctionName pulumi.StringPtrInput
	// The name of the streaming job.
	JobName pulumi.StringInput
	// Resource name
	Name pulumi.StringPtrInput
	// The properties that are associated with a function.
	Properties pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (FunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionArgs)(nil)).Elem()
}

type FunctionInput interface {
	pulumi.Input

	ToFunctionOutput() FunctionOutput
	ToFunctionOutputWithContext(ctx context.Context) FunctionOutput
}

func (*Function) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (i *Function) ToFunctionOutput() FunctionOutput {
	return i.ToFunctionOutputWithContext(context.Background())
}

func (i *Function) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutput)
}

type FunctionOutput struct{ *pulumi.OutputState }

func (FunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (o FunctionOutput) ToFunctionOutput() FunctionOutput {
	return o
}

func (o FunctionOutput) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return o
}

// Resource name
func (o FunctionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The properties that are associated with a function.
func (o FunctionOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Function) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Resource type
func (o FunctionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FunctionOutput{})
}
