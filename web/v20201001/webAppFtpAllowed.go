// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Publishing Credentials Policies parameters.
type WebAppFtpAllowed struct {
	pulumi.CustomResourceState

	// <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
	Allow pulumi.BoolOutput `pulumi:"allow"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppFtpAllowed registers a new resource with the given unique name, arguments, and options.
func NewWebAppFtpAllowed(ctx *pulumi.Context,
	name string, args *WebAppFtpAllowedArgs, opts ...pulumi.ResourceOption) (*WebAppFtpAllowed, error) {
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
			Type: pulumi.String("azure-native:web:WebAppFtpAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppFtpAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppFtpAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppFtpAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppFtpAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppFtpAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppFtpAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppFtpAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppFtpAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppFtpAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppFtpAllowed"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebAppFtpAllowed"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppFtpAllowed
	err := ctx.RegisterResource("azure-native:web/v20201001:WebAppFtpAllowed", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppFtpAllowed gets an existing WebAppFtpAllowed resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppFtpAllowed(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppFtpAllowedState, opts ...pulumi.ResourceOption) (*WebAppFtpAllowed, error) {
	var resource WebAppFtpAllowed
	err := ctx.ReadResource("azure-native:web/v20201001:WebAppFtpAllowed", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppFtpAllowed resources.
type webAppFtpAllowedState struct {
}

type WebAppFtpAllowedState struct {
}

func (WebAppFtpAllowedState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppFtpAllowedState)(nil)).Elem()
}

type webAppFtpAllowedArgs struct {
	// <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
	Allow bool `pulumi:"allow"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppFtpAllowed resource.
type WebAppFtpAllowedArgs struct {
	// <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
	Allow pulumi.BoolInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppFtpAllowedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppFtpAllowedArgs)(nil)).Elem()
}

type WebAppFtpAllowedInput interface {
	pulumi.Input

	ToWebAppFtpAllowedOutput() WebAppFtpAllowedOutput
	ToWebAppFtpAllowedOutputWithContext(ctx context.Context) WebAppFtpAllowedOutput
}

func (*WebAppFtpAllowed) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppFtpAllowed)(nil)).Elem()
}

func (i *WebAppFtpAllowed) ToWebAppFtpAllowedOutput() WebAppFtpAllowedOutput {
	return i.ToWebAppFtpAllowedOutputWithContext(context.Background())
}

func (i *WebAppFtpAllowed) ToWebAppFtpAllowedOutputWithContext(ctx context.Context) WebAppFtpAllowedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppFtpAllowedOutput)
}

type WebAppFtpAllowedOutput struct{ *pulumi.OutputState }

func (WebAppFtpAllowedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppFtpAllowed)(nil)).Elem()
}

func (o WebAppFtpAllowedOutput) ToWebAppFtpAllowedOutput() WebAppFtpAllowedOutput {
	return o
}

func (o WebAppFtpAllowedOutput) ToWebAppFtpAllowedOutputWithContext(ctx context.Context) WebAppFtpAllowedOutput {
	return o
}

// <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
func (o WebAppFtpAllowedOutput) Allow() pulumi.BoolOutput {
	return o.ApplyT(func(v *WebAppFtpAllowed) pulumi.BoolOutput { return v.Allow }).(pulumi.BoolOutput)
}

// Kind of resource.
func (o WebAppFtpAllowedOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFtpAllowed) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppFtpAllowedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppFtpAllowed) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o WebAppFtpAllowedOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebAppFtpAllowed) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o WebAppFtpAllowedOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppFtpAllowed) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppFtpAllowedOutput{})
}
