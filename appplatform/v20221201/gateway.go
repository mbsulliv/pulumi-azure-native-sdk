// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Spring Cloud Gateway resource
type Gateway struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Spring Cloud Gateway properties payload
	Properties GatewayPropertiesResponseOutput `pulumi:"properties"`
	// Sku of the Spring Cloud Gateway resource
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGateway registers a new resource with the given unique name, arguments, and options.
func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
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
		args.Properties = args.Properties.ToGatewayPropertiesPtrOutput().ApplyT(func(v *GatewayProperties) *GatewayProperties { return v.Defaults() }).(GatewayPropertiesPtrOutput)
	}
	if args.Sku != nil {
		args.Sku = args.Sku.ToSkuPtrOutput().ApplyT(func(v *Sku) *Sku { return v.Defaults() }).(SkuPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:Gateway"),
		},
	})
	opts = append(opts, aliases)
	var resource Gateway
	err := ctx.RegisterResource("azure-native:appplatform/v20221201:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGateway gets an existing Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("azure-native:appplatform/v20221201:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gateway resources.
type gatewayState struct {
}

type GatewayState struct {
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	// The name of Spring Cloud Gateway.
	GatewayName *string `pulumi:"gatewayName"`
	// Spring Cloud Gateway properties payload
	Properties *GatewayProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
	// Sku of the Spring Cloud Gateway resource
	Sku *Sku `pulumi:"sku"`
}

// The set of arguments for constructing a Gateway resource.
type GatewayArgs struct {
	// The name of Spring Cloud Gateway.
	GatewayName pulumi.StringPtrInput
	// Spring Cloud Gateway properties payload
	Properties GatewayPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
	// Sku of the Spring Cloud Gateway resource
	Sku SkuPtrInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (*Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

type GatewayOutput struct{ *pulumi.OutputState }

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

// The name of the resource.
func (o GatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Spring Cloud Gateway properties payload
func (o GatewayOutput) Properties() GatewayPropertiesResponseOutput {
	return o.ApplyT(func(v *Gateway) GatewayPropertiesResponseOutput { return v.Properties }).(GatewayPropertiesResponseOutput)
}

// Sku of the Spring Cloud Gateway resource
func (o GatewayOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Gateway) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o GatewayOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Gateway) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o GatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayOutput{})
}
