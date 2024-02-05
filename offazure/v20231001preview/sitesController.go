// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A VmwareSite
type SitesController struct {
	pulumi.CustomResourceState

	// Gets or sets the on-premises agent details.
	AgentDetails SiteAgentPropertiesResponsePtrOutput `pulumi:"agentDetails"`
	// Gets or sets the Appliance Name.
	ApplianceName pulumi.StringPtrOutput `pulumi:"applianceName"`
	// Gets or sets the ARM ID of migration hub solution for SDS.
	DiscoverySolutionId pulumi.StringPtrOutput `pulumi:"discoverySolutionId"`
	// If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	ETag pulumi.StringOutput `pulumi:"eTag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Gets the Master Site this site is linked to.
	MasterSiteId pulumi.StringOutput `pulumi:"masterSiteId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets the service endpoint.
	ServiceEndpoint pulumi.StringOutput `pulumi:"serviceEndpoint"`
	// Gets or sets the service principal identity details used by agent for
	// communication
	//             to the service.
	ServicePrincipalIdentityDetails SiteSpnPropertiesResponsePtrOutput `pulumi:"servicePrincipalIdentityDetails"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSitesController registers a new resource with the given unique name, arguments, and options.
func NewSitesController(ctx *pulumi.Context,
	name string, args *SitesControllerArgs, opts ...pulumi.ResourceOption) (*SitesController, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:offazure:SitesController"),
		},
		{
			Type: pulumi.String("azure-native:offazure/v20200101:SitesController"),
		},
		{
			Type: pulumi.String("azure-native:offazure/v20200707:SitesController"),
		},
		{
			Type: pulumi.String("azure-native:offazure/v20230606:SitesController"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SitesController
	err := ctx.RegisterResource("azure-native:offazure/v20231001preview:SitesController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSitesController gets an existing SitesController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSitesController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SitesControllerState, opts ...pulumi.ResourceOption) (*SitesController, error) {
	var resource SitesController
	err := ctx.ReadResource("azure-native:offazure/v20231001preview:SitesController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SitesController resources.
type sitesControllerState struct {
}

type SitesControllerState struct {
}

func (SitesControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sitesControllerState)(nil)).Elem()
}

type sitesControllerArgs struct {
	// Gets or sets the on-premises agent details.
	AgentDetails *SiteAgentProperties `pulumi:"agentDetails"`
	// Gets or sets the Appliance Name.
	ApplianceName *string `pulumi:"applianceName"`
	// Gets or sets the ARM ID of migration hub solution for SDS.
	DiscoverySolutionId *string `pulumi:"discoverySolutionId"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the service principal identity details used by agent for
	// communication
	//             to the service.
	ServicePrincipalIdentityDetails *SiteSpnProperties `pulumi:"servicePrincipalIdentityDetails"`
	// Site name
	SiteName *string `pulumi:"siteName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SitesController resource.
type SitesControllerArgs struct {
	// Gets or sets the on-premises agent details.
	AgentDetails SiteAgentPropertiesPtrInput
	// Gets or sets the Appliance Name.
	ApplianceName pulumi.StringPtrInput
	// Gets or sets the ARM ID of migration hub solution for SDS.
	DiscoverySolutionId pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the service principal identity details used by agent for
	// communication
	//             to the service.
	ServicePrincipalIdentityDetails SiteSpnPropertiesPtrInput
	// Site name
	SiteName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SitesControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sitesControllerArgs)(nil)).Elem()
}

type SitesControllerInput interface {
	pulumi.Input

	ToSitesControllerOutput() SitesControllerOutput
	ToSitesControllerOutputWithContext(ctx context.Context) SitesControllerOutput
}

func (*SitesController) ElementType() reflect.Type {
	return reflect.TypeOf((**SitesController)(nil)).Elem()
}

func (i *SitesController) ToSitesControllerOutput() SitesControllerOutput {
	return i.ToSitesControllerOutputWithContext(context.Background())
}

func (i *SitesController) ToSitesControllerOutputWithContext(ctx context.Context) SitesControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SitesControllerOutput)
}

type SitesControllerOutput struct{ *pulumi.OutputState }

func (SitesControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SitesController)(nil)).Elem()
}

func (o SitesControllerOutput) ToSitesControllerOutput() SitesControllerOutput {
	return o
}

func (o SitesControllerOutput) ToSitesControllerOutputWithContext(ctx context.Context) SitesControllerOutput {
	return o
}

// Gets or sets the on-premises agent details.
func (o SitesControllerOutput) AgentDetails() SiteAgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SitesController) SiteAgentPropertiesResponsePtrOutput { return v.AgentDetails }).(SiteAgentPropertiesResponsePtrOutput)
}

// Gets or sets the Appliance Name.
func (o SitesControllerOutput) ApplianceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SitesController) pulumi.StringPtrOutput { return v.ApplianceName }).(pulumi.StringPtrOutput)
}

// Gets or sets the ARM ID of migration hub solution for SDS.
func (o SitesControllerOutput) DiscoverySolutionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SitesController) pulumi.StringPtrOutput { return v.DiscoverySolutionId }).(pulumi.StringPtrOutput)
}

// If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o SitesControllerOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *SitesController) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o SitesControllerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SitesController) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Gets the Master Site this site is linked to.
func (o SitesControllerOutput) MasterSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v *SitesController) pulumi.StringOutput { return v.MasterSiteId }).(pulumi.StringOutput)
}

// The name of the resource
func (o SitesControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SitesController) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o SitesControllerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SitesController) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets the service endpoint.
func (o SitesControllerOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *SitesController) pulumi.StringOutput { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// Gets or sets the service principal identity details used by agent for
// communication
//
//	to the service.
func (o SitesControllerOutput) ServicePrincipalIdentityDetails() SiteSpnPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SitesController) SiteSpnPropertiesResponsePtrOutput { return v.ServicePrincipalIdentityDetails }).(SiteSpnPropertiesResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SitesControllerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SitesController) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SitesControllerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SitesController) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SitesControllerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SitesController) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SitesControllerOutput{})
}
