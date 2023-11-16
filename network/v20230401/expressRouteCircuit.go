// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ExpressRouteCircuit resource.
type ExpressRouteCircuit struct {
	pulumi.CustomResourceState

	// Allow classic operations.
	AllowClassicOperations pulumi.BoolPtrOutput `pulumi:"allowClassicOperations"`
	// The authorizationKey.
	AuthorizationKey pulumi.StringPtrOutput `pulumi:"authorizationKey"`
	// The authorization status of the Circuit.
	AuthorizationStatus pulumi.StringOutput `pulumi:"authorizationStatus"`
	// The list of authorizations.
	Authorizations ExpressRouteCircuitAuthorizationResponseArrayOutput `pulumi:"authorizations"`
	// The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
	BandwidthInGbps pulumi.Float64PtrOutput `pulumi:"bandwidthInGbps"`
	// The CircuitProvisioningState state of the resource.
	CircuitProvisioningState pulumi.StringPtrOutput `pulumi:"circuitProvisioningState"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
	ExpressRoutePort SubResourceResponsePtrOutput `pulumi:"expressRoutePort"`
	// The GatewayManager Etag.
	GatewayManagerEtag pulumi.StringPtrOutput `pulumi:"gatewayManagerEtag"`
	// Flag denoting global reach status.
	GlobalReachEnabled pulumi.BoolPtrOutput `pulumi:"globalReachEnabled"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of peerings.
	Peerings ExpressRouteCircuitPeeringResponseArrayOutput `pulumi:"peerings"`
	// The provisioning state of the express route circuit resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The ServiceKey.
	ServiceKey pulumi.StringPtrOutput `pulumi:"serviceKey"`
	// The ServiceProviderNotes.
	ServiceProviderNotes pulumi.StringPtrOutput `pulumi:"serviceProviderNotes"`
	// The ServiceProviderProperties.
	ServiceProviderProperties ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput `pulumi:"serviceProviderProperties"`
	// The ServiceProviderProvisioningState state of the resource.
	ServiceProviderProvisioningState pulumi.StringPtrOutput `pulumi:"serviceProviderProvisioningState"`
	// The SKU.
	Sku ExpressRouteCircuitSkuResponsePtrOutput `pulumi:"sku"`
	// The identifier of the circuit traffic. Outer tag for QinQ encapsulation.
	Stag pulumi.IntOutput `pulumi:"stag"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExpressRouteCircuit registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuit(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:ExpressRouteCircuit"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ExpressRouteCircuit
	err := ctx.RegisterResource("azure-native:network/v20230401:ExpressRouteCircuit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuit gets an existing ExpressRouteCircuit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuit, error) {
	var resource ExpressRouteCircuit
	err := ctx.ReadResource("azure-native:network/v20230401:ExpressRouteCircuit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuit resources.
type expressRouteCircuitState struct {
}

type ExpressRouteCircuitState struct {
}

func (ExpressRouteCircuitState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitState)(nil)).Elem()
}

type expressRouteCircuitArgs struct {
	// Allow classic operations.
	AllowClassicOperations *bool `pulumi:"allowClassicOperations"`
	// The authorizationKey.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The list of authorizations.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Authorizations []ExpressRouteCircuitAuthorizationType `pulumi:"authorizations"`
	// The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
	BandwidthInGbps *float64 `pulumi:"bandwidthInGbps"`
	// The name of the circuit.
	CircuitName *string `pulumi:"circuitName"`
	// The CircuitProvisioningState state of the resource.
	CircuitProvisioningState *string `pulumi:"circuitProvisioningState"`
	// The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
	ExpressRoutePort *SubResource `pulumi:"expressRoutePort"`
	// The GatewayManager Etag.
	GatewayManagerEtag *string `pulumi:"gatewayManagerEtag"`
	// Flag denoting global reach status.
	GlobalReachEnabled *bool `pulumi:"globalReachEnabled"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The list of peerings.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Peerings []ExpressRouteCircuitPeeringType `pulumi:"peerings"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ServiceKey.
	ServiceKey *string `pulumi:"serviceKey"`
	// The ServiceProviderNotes.
	ServiceProviderNotes *string `pulumi:"serviceProviderNotes"`
	// The ServiceProviderProperties.
	ServiceProviderProperties *ExpressRouteCircuitServiceProviderProperties `pulumi:"serviceProviderProperties"`
	// The ServiceProviderProvisioningState state of the resource.
	ServiceProviderProvisioningState *string `pulumi:"serviceProviderProvisioningState"`
	// The SKU.
	Sku *ExpressRouteCircuitSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ExpressRouteCircuit resource.
type ExpressRouteCircuitArgs struct {
	// Allow classic operations.
	AllowClassicOperations pulumi.BoolPtrInput
	// The authorizationKey.
	AuthorizationKey pulumi.StringPtrInput
	// The list of authorizations.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Authorizations ExpressRouteCircuitAuthorizationTypeArrayInput
	// The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
	BandwidthInGbps pulumi.Float64PtrInput
	// The name of the circuit.
	CircuitName pulumi.StringPtrInput
	// The CircuitProvisioningState state of the resource.
	CircuitProvisioningState pulumi.StringPtrInput
	// The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
	ExpressRoutePort SubResourcePtrInput
	// The GatewayManager Etag.
	GatewayManagerEtag pulumi.StringPtrInput
	// Flag denoting global reach status.
	GlobalReachEnabled pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The list of peerings.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Peerings ExpressRouteCircuitPeeringTypeArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The ServiceKey.
	ServiceKey pulumi.StringPtrInput
	// The ServiceProviderNotes.
	ServiceProviderNotes pulumi.StringPtrInput
	// The ServiceProviderProperties.
	ServiceProviderProperties ExpressRouteCircuitServiceProviderPropertiesPtrInput
	// The ServiceProviderProvisioningState state of the resource.
	ServiceProviderProvisioningState pulumi.StringPtrInput
	// The SKU.
	Sku ExpressRouteCircuitSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ExpressRouteCircuitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitArgs)(nil)).Elem()
}

type ExpressRouteCircuitInput interface {
	pulumi.Input

	ToExpressRouteCircuitOutput() ExpressRouteCircuitOutput
	ToExpressRouteCircuitOutputWithContext(ctx context.Context) ExpressRouteCircuitOutput
}

func (*ExpressRouteCircuit) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuit)(nil)).Elem()
}

func (i *ExpressRouteCircuit) ToExpressRouteCircuitOutput() ExpressRouteCircuitOutput {
	return i.ToExpressRouteCircuitOutputWithContext(context.Background())
}

func (i *ExpressRouteCircuit) ToExpressRouteCircuitOutputWithContext(ctx context.Context) ExpressRouteCircuitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitOutput)
}

type ExpressRouteCircuitOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuit)(nil)).Elem()
}

func (o ExpressRouteCircuitOutput) ToExpressRouteCircuitOutput() ExpressRouteCircuitOutput {
	return o
}

func (o ExpressRouteCircuitOutput) ToExpressRouteCircuitOutputWithContext(ctx context.Context) ExpressRouteCircuitOutput {
	return o
}

// Allow classic operations.
func (o ExpressRouteCircuitOutput) AllowClassicOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.BoolPtrOutput { return v.AllowClassicOperations }).(pulumi.BoolPtrOutput)
}

// The authorizationKey.
func (o ExpressRouteCircuitOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

// The authorization status of the Circuit.
func (o ExpressRouteCircuitOutput) AuthorizationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringOutput { return v.AuthorizationStatus }).(pulumi.StringOutput)
}

// The list of authorizations.
func (o ExpressRouteCircuitOutput) Authorizations() ExpressRouteCircuitAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) ExpressRouteCircuitAuthorizationResponseArrayOutput {
		return v.Authorizations
	}).(ExpressRouteCircuitAuthorizationResponseArrayOutput)
}

// The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
func (o ExpressRouteCircuitOutput) BandwidthInGbps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.Float64PtrOutput { return v.BandwidthInGbps }).(pulumi.Float64PtrOutput)
}

// The CircuitProvisioningState state of the resource.
func (o ExpressRouteCircuitOutput) CircuitProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.CircuitProvisioningState }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o ExpressRouteCircuitOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
func (o ExpressRouteCircuitOutput) ExpressRoutePort() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) SubResourceResponsePtrOutput { return v.ExpressRoutePort }).(SubResourceResponsePtrOutput)
}

// The GatewayManager Etag.
func (o ExpressRouteCircuitOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

// Flag denoting global reach status.
func (o ExpressRouteCircuitOutput) GlobalReachEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.BoolPtrOutput { return v.GlobalReachEnabled }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o ExpressRouteCircuitOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ExpressRouteCircuitOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of peerings.
func (o ExpressRouteCircuitOutput) Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) ExpressRouteCircuitPeeringResponseArrayOutput { return v.Peerings }).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

// The provisioning state of the express route circuit resource.
func (o ExpressRouteCircuitOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The ServiceKey.
func (o ExpressRouteCircuitOutput) ServiceKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.ServiceKey }).(pulumi.StringPtrOutput)
}

// The ServiceProviderNotes.
func (o ExpressRouteCircuitOutput) ServiceProviderNotes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.ServiceProviderNotes }).(pulumi.StringPtrOutput)
}

// The ServiceProviderProperties.
func (o ExpressRouteCircuitOutput) ServiceProviderProperties() ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
		return v.ServiceProviderProperties
	}).(ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput)
}

// The ServiceProviderProvisioningState state of the resource.
func (o ExpressRouteCircuitOutput) ServiceProviderProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.ServiceProviderProvisioningState }).(pulumi.StringPtrOutput)
}

// The SKU.
func (o ExpressRouteCircuitOutput) Sku() ExpressRouteCircuitSkuResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) ExpressRouteCircuitSkuResponsePtrOutput { return v.Sku }).(ExpressRouteCircuitSkuResponsePtrOutput)
}

// The identifier of the circuit traffic. Outer tag for QinQ encapsulation.
func (o ExpressRouteCircuitOutput) Stag() pulumi.IntOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.IntOutput { return v.Stag }).(pulumi.IntOutput)
}

// Resource tags.
func (o ExpressRouteCircuitOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ExpressRouteCircuitOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteCircuitOutput{})
}
