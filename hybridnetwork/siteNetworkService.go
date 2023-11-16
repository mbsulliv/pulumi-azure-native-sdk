// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridnetwork

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Site network service resource.
// Azure REST API version: 2023-09-01.
type SiteNetworkService struct {
	pulumi.CustomResourceState

	// The managed identity of the Site network service, if configured.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Site network service properties.
	Properties SiteNetworkServicePropertiesFormatResponseOutput `pulumi:"properties"`
	// Sku of the site network service.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSiteNetworkService registers a new resource with the given unique name, arguments, and options.
func NewSiteNetworkService(ctx *pulumi.Context,
	name string, args *SiteNetworkServiceArgs, opts ...pulumi.ResourceOption) (*SiteNetworkService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20230901:SiteNetworkService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SiteNetworkService
	err := ctx.RegisterResource("azure-native:hybridnetwork:SiteNetworkService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteNetworkService gets an existing SiteNetworkService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteNetworkService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteNetworkServiceState, opts ...pulumi.ResourceOption) (*SiteNetworkService, error) {
	var resource SiteNetworkService
	err := ctx.ReadResource("azure-native:hybridnetwork:SiteNetworkService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteNetworkService resources.
type siteNetworkServiceState struct {
}

type SiteNetworkServiceState struct {
}

func (SiteNetworkServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteNetworkServiceState)(nil)).Elem()
}

type siteNetworkServiceArgs struct {
	// The managed identity of the Site network service, if configured.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Site network service properties.
	Properties *SiteNetworkServicePropertiesFormat `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the site network service.
	SiteNetworkServiceName *string `pulumi:"siteNetworkServiceName"`
	// Sku of the site network service.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SiteNetworkService resource.
type SiteNetworkServiceArgs struct {
	// The managed identity of the Site network service, if configured.
	Identity ManagedServiceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Site network service properties.
	Properties SiteNetworkServicePropertiesFormatPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the site network service.
	SiteNetworkServiceName pulumi.StringPtrInput
	// Sku of the site network service.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SiteNetworkServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteNetworkServiceArgs)(nil)).Elem()
}

type SiteNetworkServiceInput interface {
	pulumi.Input

	ToSiteNetworkServiceOutput() SiteNetworkServiceOutput
	ToSiteNetworkServiceOutputWithContext(ctx context.Context) SiteNetworkServiceOutput
}

func (*SiteNetworkService) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteNetworkService)(nil)).Elem()
}

func (i *SiteNetworkService) ToSiteNetworkServiceOutput() SiteNetworkServiceOutput {
	return i.ToSiteNetworkServiceOutputWithContext(context.Background())
}

func (i *SiteNetworkService) ToSiteNetworkServiceOutputWithContext(ctx context.Context) SiteNetworkServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteNetworkServiceOutput)
}

type SiteNetworkServiceOutput struct{ *pulumi.OutputState }

func (SiteNetworkServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteNetworkService)(nil)).Elem()
}

func (o SiteNetworkServiceOutput) ToSiteNetworkServiceOutput() SiteNetworkServiceOutput {
	return o
}

func (o SiteNetworkServiceOutput) ToSiteNetworkServiceOutputWithContext(ctx context.Context) SiteNetworkServiceOutput {
	return o
}

// The managed identity of the Site network service, if configured.
func (o SiteNetworkServiceOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SiteNetworkService) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o SiteNetworkServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteNetworkService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SiteNetworkServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteNetworkService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Site network service properties.
func (o SiteNetworkServiceOutput) Properties() SiteNetworkServicePropertiesFormatResponseOutput {
	return o.ApplyT(func(v *SiteNetworkService) SiteNetworkServicePropertiesFormatResponseOutput { return v.Properties }).(SiteNetworkServicePropertiesFormatResponseOutput)
}

// Sku of the site network service.
func (o SiteNetworkServiceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *SiteNetworkService) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SiteNetworkServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SiteNetworkService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SiteNetworkServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteNetworkService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SiteNetworkServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteNetworkService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteNetworkServiceOutput{})
}
