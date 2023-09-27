// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Publishing Credentials Policies parameters.
type WebAppScmAllowed struct {
	pulumi.CustomResourceState

	// <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
	Allow pulumi.BoolOutput `pulumi:"allow"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppScmAllowed registers a new resource with the given unique name, arguments, and options.
func NewWebAppScmAllowed(ctx *pulumi.Context,
	name string, args *WebAppScmAllowedArgs, opts ...pulumi.ResourceOption) (*WebAppScmAllowed, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Allow == nil {
		return nil, errors.New("invalid value for required argument 'Allow'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppScmAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppScmAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppScmAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppScmAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppScmAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppScmAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppScmAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppScmAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppScmAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppScmAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppScmAllowed"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppScmAllowed
	err := ctx.RegisterResource("azure-native:web/v20220901:WebAppScmAllowed", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppScmAllowed gets an existing WebAppScmAllowed resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppScmAllowed(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppScmAllowedState, opts ...pulumi.ResourceOption) (*WebAppScmAllowed, error) {
	var resource WebAppScmAllowed
	err := ctx.ReadResource("azure-native:web/v20220901:WebAppScmAllowed", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppScmAllowed resources.
type webAppScmAllowedState struct {
}

type WebAppScmAllowedState struct {
}

func (WebAppScmAllowedState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppScmAllowedState)(nil)).Elem()
}

type webAppScmAllowedArgs struct {
	// <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
	Allow bool `pulumi:"allow"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppScmAllowed resource.
type WebAppScmAllowedArgs struct {
	// <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
	Allow pulumi.BoolInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppScmAllowedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppScmAllowedArgs)(nil)).Elem()
}

type WebAppScmAllowedInput interface {
	pulumi.Input

	ToWebAppScmAllowedOutput() WebAppScmAllowedOutput
	ToWebAppScmAllowedOutputWithContext(ctx context.Context) WebAppScmAllowedOutput
}

func (*WebAppScmAllowed) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppScmAllowed)(nil)).Elem()
}

func (i *WebAppScmAllowed) ToWebAppScmAllowedOutput() WebAppScmAllowedOutput {
	return i.ToWebAppScmAllowedOutputWithContext(context.Background())
}

func (i *WebAppScmAllowed) ToWebAppScmAllowedOutputWithContext(ctx context.Context) WebAppScmAllowedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppScmAllowedOutput)
}

func (i *WebAppScmAllowed) ToOutput(ctx context.Context) pulumix.Output[*WebAppScmAllowed] {
	return pulumix.Output[*WebAppScmAllowed]{
		OutputState: i.ToWebAppScmAllowedOutputWithContext(ctx).OutputState,
	}
}

type WebAppScmAllowedOutput struct{ *pulumi.OutputState }

func (WebAppScmAllowedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppScmAllowed)(nil)).Elem()
}

func (o WebAppScmAllowedOutput) ToWebAppScmAllowedOutput() WebAppScmAllowedOutput {
	return o
}

func (o WebAppScmAllowedOutput) ToWebAppScmAllowedOutputWithContext(ctx context.Context) WebAppScmAllowedOutput {
	return o
}

func (o WebAppScmAllowedOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppScmAllowed] {
	return pulumix.Output[*WebAppScmAllowed]{
		OutputState: o.OutputState,
	}
}

// <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
func (o WebAppScmAllowedOutput) Allow() pulumi.BoolOutput {
	return o.ApplyT(func(v *WebAppScmAllowed) pulumi.BoolOutput { return v.Allow }).(pulumi.BoolOutput)
}

// Kind of resource.
func (o WebAppScmAllowedOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppScmAllowed) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppScmAllowedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppScmAllowed) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o WebAppScmAllowedOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppScmAllowed) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppScmAllowedOutput{})
}
