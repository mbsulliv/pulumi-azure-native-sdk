// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The StreamingEndpoint.
type StreamingEndpoint struct {
	pulumi.CustomResourceState

	// The access control definition of the StreamingEndpoint.
	AccessControl StreamingEndpointAccessControlResponsePtrOutput `pulumi:"accessControl"`
	// AvailabilitySet name
	AvailabilitySetName pulumi.StringPtrOutput `pulumi:"availabilitySetName"`
	// The CDN enabled flag.
	CdnEnabled pulumi.BoolPtrOutput `pulumi:"cdnEnabled"`
	// The CDN profile name.
	CdnProfile pulumi.StringPtrOutput `pulumi:"cdnProfile"`
	// The CDN provider name.
	CdnProvider pulumi.StringPtrOutput `pulumi:"cdnProvider"`
	// The exact time the StreamingEndpoint was created.
	Created pulumi.StringOutput `pulumi:"created"`
	// The StreamingEndpoint access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesResponsePtrOutput `pulumi:"crossSiteAccessPolicies"`
	// The custom host names of the StreamingEndpoint
	CustomHostNames pulumi.StringArrayOutput `pulumi:"customHostNames"`
	// The StreamingEndpoint description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The free trial expiration time.
	FreeTrialEndTime pulumi.StringOutput `pulumi:"freeTrialEndTime"`
	// The StreamingEndpoint host name.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The exact time the StreamingEndpoint was last modified.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The Azure Region of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Max cache age
	MaxCacheAge pulumi.Float64PtrOutput `pulumi:"maxCacheAge"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the StreamingEndpoint.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource state of the StreamingEndpoint.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The number of scale units.
	ScaleUnits pulumi.IntPtrOutput `pulumi:"scaleUnits"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:StreamingEndpoint"),
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
		{
			Type: pulumi.String("azure-native:media/v20221101:StreamingEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StreamingEndpoint
	err := ctx.RegisterResource("azure-native:media/v20180601preview:StreamingEndpoint", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:media/v20180601preview:StreamingEndpoint", name, id, state, &resource, opts...)
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
	// The access control definition of the StreamingEndpoint.
	AccessControl *StreamingEndpointAccessControl `pulumi:"accessControl"`
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The flag indicates if auto start the Live Event.
	AutoStart *bool `pulumi:"autoStart"`
	// AvailabilitySet name
	AvailabilitySetName *string `pulumi:"availabilitySetName"`
	// The CDN enabled flag.
	CdnEnabled *bool `pulumi:"cdnEnabled"`
	// The CDN profile name.
	CdnProfile *string `pulumi:"cdnProfile"`
	// The CDN provider name.
	CdnProvider *string `pulumi:"cdnProvider"`
	// The StreamingEndpoint access policies.
	CrossSiteAccessPolicies *CrossSiteAccessPolicies `pulumi:"crossSiteAccessPolicies"`
	// The custom host names of the StreamingEndpoint
	CustomHostNames []string `pulumi:"customHostNames"`
	// The StreamingEndpoint description.
	Description *string `pulumi:"description"`
	// The Azure Region of the resource.
	Location *string `pulumi:"location"`
	// Max cache age
	MaxCacheAge *float64 `pulumi:"maxCacheAge"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The number of scale units.
	ScaleUnits *int `pulumi:"scaleUnits"`
	// The name of the StreamingEndpoint.
	StreamingEndpointName *string `pulumi:"streamingEndpointName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StreamingEndpoint resource.
type StreamingEndpointArgs struct {
	// The access control definition of the StreamingEndpoint.
	AccessControl StreamingEndpointAccessControlPtrInput
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The flag indicates if auto start the Live Event.
	AutoStart pulumi.BoolPtrInput
	// AvailabilitySet name
	AvailabilitySetName pulumi.StringPtrInput
	// The CDN enabled flag.
	CdnEnabled pulumi.BoolPtrInput
	// The CDN profile name.
	CdnProfile pulumi.StringPtrInput
	// The CDN provider name.
	CdnProvider pulumi.StringPtrInput
	// The StreamingEndpoint access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesPtrInput
	// The custom host names of the StreamingEndpoint
	CustomHostNames pulumi.StringArrayInput
	// The StreamingEndpoint description.
	Description pulumi.StringPtrInput
	// The Azure Region of the resource.
	Location pulumi.StringPtrInput
	// Max cache age
	MaxCacheAge pulumi.Float64PtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The number of scale units.
	ScaleUnits pulumi.IntPtrInput
	// The name of the StreamingEndpoint.
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

func (i *StreamingEndpoint) ToOutput(ctx context.Context) pulumix.Output[*StreamingEndpoint] {
	return pulumix.Output[*StreamingEndpoint]{
		OutputState: i.ToStreamingEndpointOutputWithContext(ctx).OutputState,
	}
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

func (o StreamingEndpointOutput) ToOutput(ctx context.Context) pulumix.Output[*StreamingEndpoint] {
	return pulumix.Output[*StreamingEndpoint]{
		OutputState: o.OutputState,
	}
}

// The access control definition of the StreamingEndpoint.
func (o StreamingEndpointOutput) AccessControl() StreamingEndpointAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) StreamingEndpointAccessControlResponsePtrOutput { return v.AccessControl }).(StreamingEndpointAccessControlResponsePtrOutput)
}

// AvailabilitySet name
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

// The exact time the StreamingEndpoint was created.
func (o StreamingEndpointOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// The StreamingEndpoint access policies.
func (o StreamingEndpointOutput) CrossSiteAccessPolicies() CrossSiteAccessPoliciesResponsePtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) CrossSiteAccessPoliciesResponsePtrOutput { return v.CrossSiteAccessPolicies }).(CrossSiteAccessPoliciesResponsePtrOutput)
}

// The custom host names of the StreamingEndpoint
func (o StreamingEndpointOutput) CustomHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringArrayOutput { return v.CustomHostNames }).(pulumi.StringArrayOutput)
}

// The StreamingEndpoint description.
func (o StreamingEndpointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The free trial expiration time.
func (o StreamingEndpointOutput) FreeTrialEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.FreeTrialEndTime }).(pulumi.StringOutput)
}

// The StreamingEndpoint host name.
func (o StreamingEndpointOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// The exact time the StreamingEndpoint was last modified.
func (o StreamingEndpointOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// The Azure Region of the resource.
func (o StreamingEndpointOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Max cache age
func (o StreamingEndpointOutput) MaxCacheAge() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.Float64PtrOutput { return v.MaxCacheAge }).(pulumi.Float64PtrOutput)
}

// The name of the resource.
func (o StreamingEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the StreamingEndpoint.
func (o StreamingEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource state of the StreamingEndpoint.
func (o StreamingEndpointOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// The number of scale units.
func (o StreamingEndpointOutput) ScaleUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.IntPtrOutput { return v.ScaleUnits }).(pulumi.IntPtrOutput)
}

// Resource tags.
func (o StreamingEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o StreamingEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StreamingEndpointOutput{})
}
