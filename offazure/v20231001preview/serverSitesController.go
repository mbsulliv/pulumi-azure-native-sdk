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

// A ServerSiteResource
type ServerSitesController struct {
	pulumi.CustomResourceState

	// Gets or sets the on-premises agent details.
	AgentDetails SiteAgentPropertiesResponsePtrOutput `pulumi:"agentDetails"`
	// Gets or sets the Appliance Name.
	ApplianceName pulumi.StringPtrOutput `pulumi:"applianceName"`
	// Gets or sets the ARM ID of migration hub solution for SDS.
	DiscoverySolutionId pulumi.StringPtrOutput `pulumi:"discoverySolutionId"`
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

// NewServerSitesController registers a new resource with the given unique name, arguments, and options.
func NewServerSitesController(ctx *pulumi.Context,
	name string, args *ServerSitesControllerArgs, opts ...pulumi.ResourceOption) (*ServerSitesController, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:offazure:ServerSitesController"),
		},
		{
			Type: pulumi.String("azure-native:offazure/v20230606:ServerSitesController"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServerSitesController
	err := ctx.RegisterResource("azure-native:offazure/v20231001preview:ServerSitesController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerSitesController gets an existing ServerSitesController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerSitesController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerSitesControllerState, opts ...pulumi.ResourceOption) (*ServerSitesController, error) {
	var resource ServerSitesController
	err := ctx.ReadResource("azure-native:offazure/v20231001preview:ServerSitesController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerSitesController resources.
type serverSitesControllerState struct {
}

type ServerSitesControllerState struct {
}

func (ServerSitesControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverSitesControllerState)(nil)).Elem()
}

type serverSitesControllerArgs struct {
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

// The set of arguments for constructing a ServerSitesController resource.
type ServerSitesControllerArgs struct {
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

func (ServerSitesControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverSitesControllerArgs)(nil)).Elem()
}

type ServerSitesControllerInput interface {
	pulumi.Input

	ToServerSitesControllerOutput() ServerSitesControllerOutput
	ToServerSitesControllerOutputWithContext(ctx context.Context) ServerSitesControllerOutput
}

func (*ServerSitesController) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerSitesController)(nil)).Elem()
}

func (i *ServerSitesController) ToServerSitesControllerOutput() ServerSitesControllerOutput {
	return i.ToServerSitesControllerOutputWithContext(context.Background())
}

func (i *ServerSitesController) ToServerSitesControllerOutputWithContext(ctx context.Context) ServerSitesControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerSitesControllerOutput)
}

type ServerSitesControllerOutput struct{ *pulumi.OutputState }

func (ServerSitesControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerSitesController)(nil)).Elem()
}

func (o ServerSitesControllerOutput) ToServerSitesControllerOutput() ServerSitesControllerOutput {
	return o
}

func (o ServerSitesControllerOutput) ToServerSitesControllerOutputWithContext(ctx context.Context) ServerSitesControllerOutput {
	return o
}

// Gets or sets the on-premises agent details.
func (o ServerSitesControllerOutput) AgentDetails() SiteAgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ServerSitesController) SiteAgentPropertiesResponsePtrOutput { return v.AgentDetails }).(SiteAgentPropertiesResponsePtrOutput)
}

// Gets or sets the Appliance Name.
func (o ServerSitesControllerOutput) ApplianceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerSitesController) pulumi.StringPtrOutput { return v.ApplianceName }).(pulumi.StringPtrOutput)
}

// Gets or sets the ARM ID of migration hub solution for SDS.
func (o ServerSitesControllerOutput) DiscoverySolutionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerSitesController) pulumi.StringPtrOutput { return v.DiscoverySolutionId }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o ServerSitesControllerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerSitesController) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Gets the Master Site this site is linked to.
func (o ServerSitesControllerOutput) MasterSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerSitesController) pulumi.StringOutput { return v.MasterSiteId }).(pulumi.StringOutput)
}

// The name of the resource
func (o ServerSitesControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerSitesController) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o ServerSitesControllerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerSitesController) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets the service endpoint.
func (o ServerSitesControllerOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerSitesController) pulumi.StringOutput { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// Gets or sets the service principal identity details used by agent for
// communication
//
//	to the service.
func (o ServerSitesControllerOutput) ServicePrincipalIdentityDetails() SiteSpnPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ServerSitesController) SiteSpnPropertiesResponsePtrOutput {
		return v.ServicePrincipalIdentityDetails
	}).(SiteSpnPropertiesResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ServerSitesControllerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ServerSitesController) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ServerSitesControllerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServerSitesController) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ServerSitesControllerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerSitesController) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerSitesControllerOutput{})
}
