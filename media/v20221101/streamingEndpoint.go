// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The streaming endpoint.
type StreamingEndpoint struct {
	pulumi.CustomResourceState

	// The access control definition of the streaming endpoint.
	AccessControl StreamingEndpointAccessControlResponsePtrOutput `pulumi:"accessControl"`
	// This feature is deprecated, do not set a value for this property.
	AvailabilitySetName pulumi.StringPtrOutput `pulumi:"availabilitySetName"`
	// The CDN enabled flag.
	CdnEnabled pulumi.BoolPtrOutput `pulumi:"cdnEnabled"`
	// The CDN profile name.
	CdnProfile pulumi.StringPtrOutput `pulumi:"cdnProfile"`
	// The CDN provider name.
	CdnProvider pulumi.StringPtrOutput `pulumi:"cdnProvider"`
	// The exact time the streaming endpoint was created.
	Created pulumi.StringOutput `pulumi:"created"`
	// The streaming endpoint access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesResponsePtrOutput `pulumi:"crossSiteAccessPolicies"`
	// The custom host names of the streaming endpoint
	CustomHostNames pulumi.StringArrayOutput `pulumi:"customHostNames"`
	// The streaming endpoint description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The free trial expiration time.
	FreeTrialEndTime pulumi.StringOutput `pulumi:"freeTrialEndTime"`
	// The streaming endpoint host name.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The exact time the streaming endpoint was last modified.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Max cache age
	MaxCacheAge pulumi.Float64PtrOutput `pulumi:"maxCacheAge"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the streaming endpoint.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource state of the streaming endpoint.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The number of scale units. Use the Scale operation to adjust this value.
	ScaleUnits pulumi.IntOutput `pulumi:"scaleUnits"`
	// The streaming endpoint sku.
	Sku ArmStreamingEndpointCurrentSkuResponsePtrOutput `pulumi:"sku"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStreamingEndpoint registers a new resource with the given unique name, arguments, and options.
func NewStreamingEndpoint(ctx *pulumi.Context,
	name string, args *StreamingEndpointArgs, opts ...pulumi.ResourceOption) (*StreamingEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScaleUnits == nil {
		return nil, errors.New("invalid value for required argument 'ScaleUnits'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20190501preview:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20220801:StreamingEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StreamingEndpoint
	err := ctx.RegisterResource("azure-native:media/v20221101:StreamingEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingEndpoint gets an existing StreamingEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingEndpointState, opts ...pulumi.ResourceOption) (*StreamingEndpoint, error) {
	var resource StreamingEndpoint
	err := ctx.ReadResource("azure-native:media/v20221101:StreamingEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingEndpoint resources.
type streamingEndpointState struct {
}

type StreamingEndpointState struct {
}

func (StreamingEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingEndpointState)(nil)).Elem()
}

type streamingEndpointArgs struct {
	// The access control definition of the streaming endpoint.
	AccessControl *StreamingEndpointAccessControl `pulumi:"accessControl"`
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The flag indicates if the resource should be automatically started on creation.
	AutoStart *bool `pulumi:"autoStart"`
	// This feature is deprecated, do not set a value for this property.
	AvailabilitySetName *string `pulumi:"availabilitySetName"`
	// The CDN enabled flag.
	CdnEnabled *bool `pulumi:"cdnEnabled"`
	// The CDN profile name.
	CdnProfile *string `pulumi:"cdnProfile"`
	// The CDN provider name.
	CdnProvider *string `pulumi:"cdnProvider"`
	// The streaming endpoint access policies.
	CrossSiteAccessPolicies *CrossSiteAccessPolicies `pulumi:"crossSiteAccessPolicies"`
	// The custom host names of the streaming endpoint
	CustomHostNames []string `pulumi:"customHostNames"`
	// The streaming endpoint description.
	Description *string `pulumi:"description"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Max cache age
	MaxCacheAge *float64 `pulumi:"maxCacheAge"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The number of scale units. Use the Scale operation to adjust this value.
	ScaleUnits int `pulumi:"scaleUnits"`
	// The streaming endpoint sku.
	Sku *ArmStreamingEndpointCurrentSku `pulumi:"sku"`
	// The name of the streaming endpoint, maximum length is 24.
	StreamingEndpointName *string `pulumi:"streamingEndpointName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StreamingEndpoint resource.
type StreamingEndpointArgs struct {
	// The access control definition of the streaming endpoint.
	AccessControl StreamingEndpointAccessControlPtrInput
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The flag indicates if the resource should be automatically started on creation.
	AutoStart pulumi.BoolPtrInput
	// This feature is deprecated, do not set a value for this property.
	AvailabilitySetName pulumi.StringPtrInput
	// The CDN enabled flag.
	CdnEnabled pulumi.BoolPtrInput
	// The CDN profile name.
	CdnProfile pulumi.StringPtrInput
	// The CDN provider name.
	CdnProvider pulumi.StringPtrInput
	// The streaming endpoint access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesPtrInput
	// The custom host names of the streaming endpoint
	CustomHostNames pulumi.StringArrayInput
	// The streaming endpoint description.
	Description pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Max cache age
	MaxCacheAge pulumi.Float64PtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The number of scale units. Use the Scale operation to adjust this value.
	ScaleUnits pulumi.IntInput
	// The streaming endpoint sku.
	Sku ArmStreamingEndpointCurrentSkuPtrInput
	// The name of the streaming endpoint, maximum length is 24.
	StreamingEndpointName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (StreamingEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingEndpointArgs)(nil)).Elem()
}

type StreamingEndpointInput interface {
	pulumi.Input

	ToStreamingEndpointOutput() StreamingEndpointOutput
	ToStreamingEndpointOutputWithContext(ctx context.Context) StreamingEndpointOutput
}

func (*StreamingEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpoint)(nil)).Elem()
}

func (i *StreamingEndpoint) ToStreamingEndpointOutput() StreamingEndpointOutput {
	return i.ToStreamingEndpointOutputWithContext(context.Background())
}

func (i *StreamingEndpoint) ToStreamingEndpointOutputWithContext(ctx context.Context) StreamingEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointOutput)
}

type StreamingEndpointOutput struct{ *pulumi.OutputState }

func (StreamingEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpoint)(nil)).Elem()
}

func (o StreamingEndpointOutput) ToStreamingEndpointOutput() StreamingEndpointOutput {
	return o
}

func (o StreamingEndpointOutput) ToStreamingEndpointOutputWithContext(ctx context.Context) StreamingEndpointOutput {
	return o
}

// The access control definition of the streaming endpoint.
func (o StreamingEndpointOutput) AccessControl() StreamingEndpointAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) StreamingEndpointAccessControlResponsePtrOutput { return v.AccessControl }).(StreamingEndpointAccessControlResponsePtrOutput)
}

// This feature is deprecated, do not set a value for this property.
func (o StreamingEndpointOutput) AvailabilitySetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringPtrOutput { return v.AvailabilitySetName }).(pulumi.StringPtrOutput)
}

// The CDN enabled flag.
func (o StreamingEndpointOutput) CdnEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.BoolPtrOutput { return v.CdnEnabled }).(pulumi.BoolPtrOutput)
}

// The CDN profile name.
func (o StreamingEndpointOutput) CdnProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringPtrOutput { return v.CdnProfile }).(pulumi.StringPtrOutput)
}

// The CDN provider name.
func (o StreamingEndpointOutput) CdnProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringPtrOutput { return v.CdnProvider }).(pulumi.StringPtrOutput)
}

// The exact time the streaming endpoint was created.
func (o StreamingEndpointOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// The streaming endpoint access policies.
func (o StreamingEndpointOutput) CrossSiteAccessPolicies() CrossSiteAccessPoliciesResponsePtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) CrossSiteAccessPoliciesResponsePtrOutput { return v.CrossSiteAccessPolicies }).(CrossSiteAccessPoliciesResponsePtrOutput)
}

// The custom host names of the streaming endpoint
func (o StreamingEndpointOutput) CustomHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringArrayOutput { return v.CustomHostNames }).(pulumi.StringArrayOutput)
}

// The streaming endpoint description.
func (o StreamingEndpointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The free trial expiration time.
func (o StreamingEndpointOutput) FreeTrialEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.FreeTrialEndTime }).(pulumi.StringOutput)
}

// The streaming endpoint host name.
func (o StreamingEndpointOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// The exact time the streaming endpoint was last modified.
func (o StreamingEndpointOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o StreamingEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Max cache age
func (o StreamingEndpointOutput) MaxCacheAge() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.Float64PtrOutput { return v.MaxCacheAge }).(pulumi.Float64PtrOutput)
}

// The name of the resource
func (o StreamingEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the streaming endpoint.
func (o StreamingEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource state of the streaming endpoint.
func (o StreamingEndpointOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// The number of scale units. Use the Scale operation to adjust this value.
func (o StreamingEndpointOutput) ScaleUnits() pulumi.IntOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.IntOutput { return v.ScaleUnits }).(pulumi.IntOutput)
}

// The streaming endpoint sku.
func (o StreamingEndpointOutput) Sku() ArmStreamingEndpointCurrentSkuResponsePtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) ArmStreamingEndpointCurrentSkuResponsePtrOutput { return v.Sku }).(ArmStreamingEndpointCurrentSkuResponsePtrOutput)
}

// The system metadata relating to this resource.
func (o StreamingEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StreamingEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o StreamingEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o StreamingEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StreamingEndpointOutput{})
}
