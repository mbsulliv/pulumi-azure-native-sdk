// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// API portal resource
// Azure REST API version: 2023-05-01-preview. Prior API version in Azure Native 1.x: 2022-01-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview.
type ApiPortal struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// API portal properties payload
	Properties ApiPortalPropertiesResponseOutput `pulumi:"properties"`
	// Sku of the API portal resource
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApiPortal registers a new resource with the given unique name, arguments, and options.
func NewApiPortal(ctx *pulumi.Context,
	name string, args *ApiPortalArgs, opts ...pulumi.ResourceOption) (*ApiPortal, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToApiPortalPropertiesPtrOutput().ApplyT(func(v *ApiPortalProperties) *ApiPortalProperties { return v.Defaults() }).(ApiPortalPropertiesPtrOutput)
	}
	if args.Sku != nil {
		args.Sku = args.Sku.ToSkuPtrOutput().ApplyT(func(v *Sku) *Sku { return v.Defaults() }).(SkuPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230301preview:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230501preview:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230701preview:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230901preview:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20231101preview:ApiPortal"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApiPortal
	err := ctx.RegisterResource("azure-native:appplatform:ApiPortal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiPortal gets an existing ApiPortal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiPortal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiPortalState, opts ...pulumi.ResourceOption) (*ApiPortal, error) {
	var resource ApiPortal
	err := ctx.ReadResource("azure-native:appplatform:ApiPortal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiPortal resources.
type apiPortalState struct {
}

type ApiPortalState struct {
}

func (ApiPortalState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPortalState)(nil)).Elem()
}

type apiPortalArgs struct {
	// The name of API portal.
	ApiPortalName *string `pulumi:"apiPortalName"`
	// API portal properties payload
	Properties *ApiPortalProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
	// Sku of the API portal resource
	Sku *Sku `pulumi:"sku"`
}

// The set of arguments for constructing a ApiPortal resource.
type ApiPortalArgs struct {
	// The name of API portal.
	ApiPortalName pulumi.StringPtrInput
	// API portal properties payload
	Properties ApiPortalPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
	// Sku of the API portal resource
	Sku SkuPtrInput
}

func (ApiPortalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPortalArgs)(nil)).Elem()
}

type ApiPortalInput interface {
	pulumi.Input

	ToApiPortalOutput() ApiPortalOutput
	ToApiPortalOutputWithContext(ctx context.Context) ApiPortalOutput
}

func (*ApiPortal) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPortal)(nil)).Elem()
}

func (i *ApiPortal) ToApiPortalOutput() ApiPortalOutput {
	return i.ToApiPortalOutputWithContext(context.Background())
}

func (i *ApiPortal) ToApiPortalOutputWithContext(ctx context.Context) ApiPortalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPortalOutput)
}

func (i *ApiPortal) ToOutput(ctx context.Context) pulumix.Output[*ApiPortal] {
	return pulumix.Output[*ApiPortal]{
		OutputState: i.ToApiPortalOutputWithContext(ctx).OutputState,
	}
}

type ApiPortalOutput struct{ *pulumi.OutputState }

func (ApiPortalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPortal)(nil)).Elem()
}

func (o ApiPortalOutput) ToApiPortalOutput() ApiPortalOutput {
	return o
}

func (o ApiPortalOutput) ToApiPortalOutputWithContext(ctx context.Context) ApiPortalOutput {
	return o
}

func (o ApiPortalOutput) ToOutput(ctx context.Context) pulumix.Output[*ApiPortal] {
	return pulumix.Output[*ApiPortal]{
		OutputState: o.OutputState,
	}
}

// The name of the resource.
func (o ApiPortalOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPortal) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// API portal properties payload
func (o ApiPortalOutput) Properties() ApiPortalPropertiesResponseOutput {
	return o.ApplyT(func(v *ApiPortal) ApiPortalPropertiesResponseOutput { return v.Properties }).(ApiPortalPropertiesResponseOutput)
}

// Sku of the API portal resource
func (o ApiPortalOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ApiPortal) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ApiPortalOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApiPortal) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o ApiPortalOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPortal) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiPortalOutput{})
}
