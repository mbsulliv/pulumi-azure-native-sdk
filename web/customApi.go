// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A custom API
// Azure REST API version: 2016-06-01. Prior API version in Azure Native 1.x: 2016-06-01.
type CustomApi struct {
	pulumi.CustomResourceState

	// Resource ETag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Custom API properties
	Properties CustomApiPropertiesDefinitionResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCustomApi registers a new resource with the given unique name, arguments, and options.
func NewCustomApi(ctx *pulumi.Context,
	name string, args *CustomApiArgs, opts ...pulumi.ResourceOption) (*CustomApi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20160601:CustomApi"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CustomApi
	err := ctx.RegisterResource("azure-native:web:CustomApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomApi gets an existing CustomApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomApiState, opts ...pulumi.ResourceOption) (*CustomApi, error) {
	var resource CustomApi
	err := ctx.ReadResource("azure-native:web:CustomApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomApi resources.
type customApiState struct {
}

type CustomApiState struct {
}

func (CustomApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*customApiState)(nil)).Elem()
}

type customApiArgs struct {
	// API name
	ApiName *string `pulumi:"apiName"`
	// Resource location
	Location *string `pulumi:"location"`
	// Custom API properties
	Properties *CustomApiPropertiesDefinition `pulumi:"properties"`
	// The resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Subscription Id
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CustomApi resource.
type CustomApiArgs struct {
	// API name
	ApiName pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Custom API properties
	Properties CustomApiPropertiesDefinitionPtrInput
	// The resource group
	ResourceGroupName pulumi.StringInput
	// Subscription Id
	SubscriptionId pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (CustomApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customApiArgs)(nil)).Elem()
}

type CustomApiInput interface {
	pulumi.Input

	ToCustomApiOutput() CustomApiOutput
	ToCustomApiOutputWithContext(ctx context.Context) CustomApiOutput
}

func (*CustomApi) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomApi)(nil)).Elem()
}

func (i *CustomApi) ToCustomApiOutput() CustomApiOutput {
	return i.ToCustomApiOutputWithContext(context.Background())
}

func (i *CustomApi) ToCustomApiOutputWithContext(ctx context.Context) CustomApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomApiOutput)
}

type CustomApiOutput struct{ *pulumi.OutputState }

func (CustomApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomApi)(nil)).Elem()
}

func (o CustomApiOutput) ToCustomApiOutput() CustomApiOutput {
	return o
}

func (o CustomApiOutput) ToCustomApiOutputWithContext(ctx context.Context) CustomApiOutput {
	return o
}

// Resource ETag
func (o CustomApiOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomApi) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource location
func (o CustomApiOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomApi) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o CustomApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Custom API properties
func (o CustomApiOutput) Properties() CustomApiPropertiesDefinitionResponseOutput {
	return o.ApplyT(func(v *CustomApi) CustomApiPropertiesDefinitionResponseOutput { return v.Properties }).(CustomApiPropertiesDefinitionResponseOutput)
}

// Resource tags
func (o CustomApiOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomApi) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o CustomApiOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomApi) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomApiOutput{})
}
