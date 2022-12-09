// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Custom domain of the Spring Cloud Gateway
type GatewayCustomDomain struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of custom domain for Spring Cloud Gateway
	Properties GatewayCustomDomainPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGatewayCustomDomain registers a new resource with the given unique name, arguments, and options.
func NewGatewayCustomDomain(ctx *pulumi.Context,
	name string, args *GatewayCustomDomainArgs, opts ...pulumi.ResourceOption) (*GatewayCustomDomain, error) {
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:GatewayCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:GatewayCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:GatewayCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:GatewayCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:GatewayCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:GatewayCustomDomain"),
		},
	})
	opts = append(opts, aliases)
	var resource GatewayCustomDomain
	err := ctx.RegisterResource("azure-native:appplatform/v20221201:GatewayCustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayCustomDomain gets an existing GatewayCustomDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayCustomDomainState, opts ...pulumi.ResourceOption) (*GatewayCustomDomain, error) {
	var resource GatewayCustomDomain
	err := ctx.ReadResource("azure-native:appplatform/v20221201:GatewayCustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayCustomDomain resources.
type gatewayCustomDomainState struct {
}

type GatewayCustomDomainState struct {
}

func (GatewayCustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCustomDomainState)(nil)).Elem()
}

type gatewayCustomDomainArgs struct {
	// The name of the Spring Cloud Gateway custom domain.
	DomainName *string `pulumi:"domainName"`
	// The name of Spring Cloud Gateway.
	GatewayName string `pulumi:"gatewayName"`
	// The properties of custom domain for Spring Cloud Gateway
	Properties *GatewayCustomDomainProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a GatewayCustomDomain resource.
type GatewayCustomDomainArgs struct {
	// The name of the Spring Cloud Gateway custom domain.
	DomainName pulumi.StringPtrInput
	// The name of Spring Cloud Gateway.
	GatewayName pulumi.StringInput
	// The properties of custom domain for Spring Cloud Gateway
	Properties GatewayCustomDomainPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (GatewayCustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCustomDomainArgs)(nil)).Elem()
}

type GatewayCustomDomainInput interface {
	pulumi.Input

	ToGatewayCustomDomainOutput() GatewayCustomDomainOutput
	ToGatewayCustomDomainOutputWithContext(ctx context.Context) GatewayCustomDomainOutput
}

func (*GatewayCustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCustomDomain)(nil)).Elem()
}

func (i *GatewayCustomDomain) ToGatewayCustomDomainOutput() GatewayCustomDomainOutput {
	return i.ToGatewayCustomDomainOutputWithContext(context.Background())
}

func (i *GatewayCustomDomain) ToGatewayCustomDomainOutputWithContext(ctx context.Context) GatewayCustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCustomDomainOutput)
}

type GatewayCustomDomainOutput struct{ *pulumi.OutputState }

func (GatewayCustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCustomDomain)(nil)).Elem()
}

func (o GatewayCustomDomainOutput) ToGatewayCustomDomainOutput() GatewayCustomDomainOutput {
	return o
}

func (o GatewayCustomDomainOutput) ToGatewayCustomDomainOutputWithContext(ctx context.Context) GatewayCustomDomainOutput {
	return o
}

// The name of the resource.
func (o GatewayCustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCustomDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The properties of custom domain for Spring Cloud Gateway
func (o GatewayCustomDomainOutput) Properties() GatewayCustomDomainPropertiesResponseOutput {
	return o.ApplyT(func(v *GatewayCustomDomain) GatewayCustomDomainPropertiesResponseOutput { return v.Properties }).(GatewayCustomDomainPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o GatewayCustomDomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GatewayCustomDomain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o GatewayCustomDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCustomDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayCustomDomainOutput{})
}
