// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Spring Cloud Gateway route config resource
type GatewayRouteConfig struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// API route config of the Spring Cloud Gateway
	Properties GatewayRouteConfigPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGatewayRouteConfig registers a new resource with the given unique name, arguments, and options.
func NewGatewayRouteConfig(ctx *pulumi.Context,
	name string, args *GatewayRouteConfigArgs, opts ...pulumi.ResourceOption) (*GatewayRouteConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayName == nil {
		return nil, errors.New("invalid value for required argument 'GatewayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToGatewayRouteConfigPropertiesPtrOutput().ApplyT(func(v *GatewayRouteConfigProperties) *GatewayRouteConfigProperties { return v.Defaults() }).(GatewayRouteConfigPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:GatewayRouteConfig"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:GatewayRouteConfig"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:GatewayRouteConfig"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:GatewayRouteConfig"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:GatewayRouteConfig"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:GatewayRouteConfig"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:GatewayRouteConfig"),
		},
	})
	opts = append(opts, aliases)
	var resource GatewayRouteConfig
	err := ctx.RegisterResource("azure-native:appplatform/v20220901preview:GatewayRouteConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayRouteConfig gets an existing GatewayRouteConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayRouteConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayRouteConfigState, opts ...pulumi.ResourceOption) (*GatewayRouteConfig, error) {
	var resource GatewayRouteConfig
	err := ctx.ReadResource("azure-native:appplatform/v20220901preview:GatewayRouteConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayRouteConfig resources.
type gatewayRouteConfigState struct {
}

type GatewayRouteConfigState struct {
}

func (GatewayRouteConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteConfigState)(nil)).Elem()
}

type gatewayRouteConfigArgs struct {
	// The name of Spring Cloud Gateway.
	GatewayName string `pulumi:"gatewayName"`
	// API route config of the Spring Cloud Gateway
	Properties *GatewayRouteConfigProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Spring Cloud Gateway route config.
	RouteConfigName *string `pulumi:"routeConfigName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a GatewayRouteConfig resource.
type GatewayRouteConfigArgs struct {
	// The name of Spring Cloud Gateway.
	GatewayName pulumi.StringInput
	// API route config of the Spring Cloud Gateway
	Properties GatewayRouteConfigPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Spring Cloud Gateway route config.
	RouteConfigName pulumi.StringPtrInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (GatewayRouteConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteConfigArgs)(nil)).Elem()
}

type GatewayRouteConfigInput interface {
	pulumi.Input

	ToGatewayRouteConfigOutput() GatewayRouteConfigOutput
	ToGatewayRouteConfigOutputWithContext(ctx context.Context) GatewayRouteConfigOutput
}

func (*GatewayRouteConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRouteConfig)(nil)).Elem()
}

func (i *GatewayRouteConfig) ToGatewayRouteConfigOutput() GatewayRouteConfigOutput {
	return i.ToGatewayRouteConfigOutputWithContext(context.Background())
}

func (i *GatewayRouteConfig) ToGatewayRouteConfigOutputWithContext(ctx context.Context) GatewayRouteConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteConfigOutput)
}

type GatewayRouteConfigOutput struct{ *pulumi.OutputState }

func (GatewayRouteConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRouteConfig)(nil)).Elem()
}

func (o GatewayRouteConfigOutput) ToGatewayRouteConfigOutput() GatewayRouteConfigOutput {
	return o
}

func (o GatewayRouteConfigOutput) ToGatewayRouteConfigOutputWithContext(ctx context.Context) GatewayRouteConfigOutput {
	return o
}

// The name of the resource.
func (o GatewayRouteConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRouteConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// API route config of the Spring Cloud Gateway
func (o GatewayRouteConfigOutput) Properties() GatewayRouteConfigPropertiesResponseOutput {
	return o.ApplyT(func(v *GatewayRouteConfig) GatewayRouteConfigPropertiesResponseOutput { return v.Properties }).(GatewayRouteConfigPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o GatewayRouteConfigOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GatewayRouteConfig) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o GatewayRouteConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRouteConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayRouteConfigOutput{})
}
