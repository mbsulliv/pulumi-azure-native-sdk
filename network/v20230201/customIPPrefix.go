// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Custom IP prefix resource.
type CustomIPPrefix struct {
	pulumi.CustomResourceState

	// The ASN for CIDR advertising. Should be an integer as string.
	Asn pulumi.StringPtrOutput `pulumi:"asn"`
	// Authorization message for WAN validation.
	AuthorizationMessage pulumi.StringPtrOutput `pulumi:"authorizationMessage"`
	// The list of all Children for IPv6 /48 CustomIpPrefix.
	ChildCustomIpPrefixes SubResourceResponseArrayOutput `pulumi:"childCustomIpPrefixes"`
	// The prefix range in CIDR notation. Should include the start address and the prefix length.
	Cidr pulumi.StringPtrOutput `pulumi:"cidr"`
	// The commissioned state of the Custom IP Prefix.
	CommissionedState pulumi.StringPtrOutput `pulumi:"commissionedState"`
	// The Parent CustomIpPrefix for IPv6 /64 CustomIpPrefix.
	CustomIpPrefixParent SubResourceResponsePtrOutput `pulumi:"customIpPrefixParent"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Whether to do express route advertise.
	ExpressRouteAdvertise pulumi.BoolPtrOutput `pulumi:"expressRouteAdvertise"`
	// The extended location of the custom IP prefix.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The reason why resource is in failed state.
	FailedReason pulumi.StringOutput `pulumi:"failedReason"`
	// The Geo for CIDR advertising. Should be an Geo code.
	Geo pulumi.StringPtrOutput `pulumi:"geo"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether to Advertise the range to Internet.
	NoInternetAdvertise pulumi.BoolPtrOutput `pulumi:"noInternetAdvertise"`
	// Type of custom IP prefix. Should be Singular, Parent, or Child.
	PrefixType pulumi.StringPtrOutput `pulumi:"prefixType"`
	// The provisioning state of the custom IP prefix resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The list of all referenced PublicIpPrefixes.
	PublicIpPrefixes SubResourceResponseArrayOutput `pulumi:"publicIpPrefixes"`
	// The resource GUID property of the custom IP prefix resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Signed message for WAN validation.
	SignedMessage pulumi.StringPtrOutput `pulumi:"signedMessage"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewCustomIPPrefix registers a new resource with the given unique name, arguments, and options.
func NewCustomIPPrefix(ctx *pulumi.Context,
	name string, args *CustomIPPrefixArgs, opts ...pulumi.ResourceOption) (*CustomIPPrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:CustomIPPrefix"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CustomIPPrefix
	err := ctx.RegisterResource("azure-native:network/v20230201:CustomIPPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomIPPrefix gets an existing CustomIPPrefix resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomIPPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomIPPrefixState, opts ...pulumi.ResourceOption) (*CustomIPPrefix, error) {
	var resource CustomIPPrefix
	err := ctx.ReadResource("azure-native:network/v20230201:CustomIPPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomIPPrefix resources.
type customIPPrefixState struct {
}

type CustomIPPrefixState struct {
}

func (CustomIPPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*customIPPrefixState)(nil)).Elem()
}

type customIPPrefixArgs struct {
	// The ASN for CIDR advertising. Should be an integer as string.
	Asn *string `pulumi:"asn"`
	// Authorization message for WAN validation.
	AuthorizationMessage *string `pulumi:"authorizationMessage"`
	// The prefix range in CIDR notation. Should include the start address and the prefix length.
	Cidr *string `pulumi:"cidr"`
	// The commissioned state of the Custom IP Prefix.
	CommissionedState *string `pulumi:"commissionedState"`
	// The name of the custom IP prefix.
	CustomIpPrefixName *string `pulumi:"customIpPrefixName"`
	// The Parent CustomIpPrefix for IPv6 /64 CustomIpPrefix.
	CustomIpPrefixParent *SubResource `pulumi:"customIpPrefixParent"`
	// Whether to do express route advertise.
	ExpressRouteAdvertise *bool `pulumi:"expressRouteAdvertise"`
	// The extended location of the custom IP prefix.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// The Geo for CIDR advertising. Should be an Geo code.
	Geo *string `pulumi:"geo"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Whether to Advertise the range to Internet.
	NoInternetAdvertise *bool `pulumi:"noInternetAdvertise"`
	// Type of custom IP prefix. Should be Singular, Parent, or Child.
	PrefixType *string `pulumi:"prefixType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Signed message for WAN validation.
	SignedMessage *string `pulumi:"signedMessage"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a CustomIPPrefix resource.
type CustomIPPrefixArgs struct {
	// The ASN for CIDR advertising. Should be an integer as string.
	Asn pulumi.StringPtrInput
	// Authorization message for WAN validation.
	AuthorizationMessage pulumi.StringPtrInput
	// The prefix range in CIDR notation. Should include the start address and the prefix length.
	Cidr pulumi.StringPtrInput
	// The commissioned state of the Custom IP Prefix.
	CommissionedState pulumi.StringPtrInput
	// The name of the custom IP prefix.
	CustomIpPrefixName pulumi.StringPtrInput
	// The Parent CustomIpPrefix for IPv6 /64 CustomIpPrefix.
	CustomIpPrefixParent SubResourcePtrInput
	// Whether to do express route advertise.
	ExpressRouteAdvertise pulumi.BoolPtrInput
	// The extended location of the custom IP prefix.
	ExtendedLocation ExtendedLocationPtrInput
	// The Geo for CIDR advertising. Should be an Geo code.
	Geo pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Whether to Advertise the range to Internet.
	NoInternetAdvertise pulumi.BoolPtrInput
	// Type of custom IP prefix. Should be Singular, Parent, or Child.
	PrefixType pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Signed message for WAN validation.
	SignedMessage pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones pulumi.StringArrayInput
}

func (CustomIPPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customIPPrefixArgs)(nil)).Elem()
}

type CustomIPPrefixInput interface {
	pulumi.Input

	ToCustomIPPrefixOutput() CustomIPPrefixOutput
	ToCustomIPPrefixOutputWithContext(ctx context.Context) CustomIPPrefixOutput
}

func (*CustomIPPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomIPPrefix)(nil)).Elem()
}

func (i *CustomIPPrefix) ToCustomIPPrefixOutput() CustomIPPrefixOutput {
	return i.ToCustomIPPrefixOutputWithContext(context.Background())
}

func (i *CustomIPPrefix) ToCustomIPPrefixOutputWithContext(ctx context.Context) CustomIPPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomIPPrefixOutput)
}

func (i *CustomIPPrefix) ToOutput(ctx context.Context) pulumix.Output[*CustomIPPrefix] {
	return pulumix.Output[*CustomIPPrefix]{
		OutputState: i.ToCustomIPPrefixOutputWithContext(ctx).OutputState,
	}
}

type CustomIPPrefixOutput struct{ *pulumi.OutputState }

func (CustomIPPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomIPPrefix)(nil)).Elem()
}

func (o CustomIPPrefixOutput) ToCustomIPPrefixOutput() CustomIPPrefixOutput {
	return o
}

func (o CustomIPPrefixOutput) ToCustomIPPrefixOutputWithContext(ctx context.Context) CustomIPPrefixOutput {
	return o
}

func (o CustomIPPrefixOutput) ToOutput(ctx context.Context) pulumix.Output[*CustomIPPrefix] {
	return pulumix.Output[*CustomIPPrefix]{
		OutputState: o.OutputState,
	}
}

// The ASN for CIDR advertising. Should be an integer as string.
func (o CustomIPPrefixOutput) Asn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringPtrOutput { return v.Asn }).(pulumi.StringPtrOutput)
}

// Authorization message for WAN validation.
func (o CustomIPPrefixOutput) AuthorizationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringPtrOutput { return v.AuthorizationMessage }).(pulumi.StringPtrOutput)
}

// The list of all Children for IPv6 /48 CustomIpPrefix.
func (o CustomIPPrefixOutput) ChildCustomIpPrefixes() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *CustomIPPrefix) SubResourceResponseArrayOutput { return v.ChildCustomIpPrefixes }).(SubResourceResponseArrayOutput)
}

// The prefix range in CIDR notation. Should include the start address and the prefix length.
func (o CustomIPPrefixOutput) Cidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringPtrOutput { return v.Cidr }).(pulumi.StringPtrOutput)
}

// The commissioned state of the Custom IP Prefix.
func (o CustomIPPrefixOutput) CommissionedState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringPtrOutput { return v.CommissionedState }).(pulumi.StringPtrOutput)
}

// The Parent CustomIpPrefix for IPv6 /64 CustomIpPrefix.
func (o CustomIPPrefixOutput) CustomIpPrefixParent() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) SubResourceResponsePtrOutput { return v.CustomIpPrefixParent }).(SubResourceResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o CustomIPPrefixOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Whether to do express route advertise.
func (o CustomIPPrefixOutput) ExpressRouteAdvertise() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.BoolPtrOutput { return v.ExpressRouteAdvertise }).(pulumi.BoolPtrOutput)
}

// The extended location of the custom IP prefix.
func (o CustomIPPrefixOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The reason why resource is in failed state.
func (o CustomIPPrefixOutput) FailedReason() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringOutput { return v.FailedReason }).(pulumi.StringOutput)
}

// The Geo for CIDR advertising. Should be an Geo code.
func (o CustomIPPrefixOutput) Geo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringPtrOutput { return v.Geo }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o CustomIPPrefixOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o CustomIPPrefixOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether to Advertise the range to Internet.
func (o CustomIPPrefixOutput) NoInternetAdvertise() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.BoolPtrOutput { return v.NoInternetAdvertise }).(pulumi.BoolPtrOutput)
}

// Type of custom IP prefix. Should be Singular, Parent, or Child.
func (o CustomIPPrefixOutput) PrefixType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringPtrOutput { return v.PrefixType }).(pulumi.StringPtrOutput)
}

// The provisioning state of the custom IP prefix resource.
func (o CustomIPPrefixOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The list of all referenced PublicIpPrefixes.
func (o CustomIPPrefixOutput) PublicIpPrefixes() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *CustomIPPrefix) SubResourceResponseArrayOutput { return v.PublicIpPrefixes }).(SubResourceResponseArrayOutput)
}

// The resource GUID property of the custom IP prefix resource.
func (o CustomIPPrefixOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Signed message for WAN validation.
func (o CustomIPPrefixOutput) SignedMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringPtrOutput { return v.SignedMessage }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o CustomIPPrefixOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o CustomIPPrefixOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A list of availability zones denoting the IP allocated for the resource needs to come from.
func (o CustomIPPrefixOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomIPPrefixOutput{})
}
