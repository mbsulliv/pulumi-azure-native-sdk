// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package offazure

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A MasterSite
// Azure REST API version: 2023-06-06.
type MasterSitesController struct {
	pulumi.CustomResourceState

	// Gets or sets a value indicating whether multiple sites per site type are
	// allowed.
	AllowMultipleSites pulumi.BoolPtrOutput `pulumi:"allowMultipleSites"`
	// Gets or sets a value for customer storage account ARM id.
	CustomerStorageAccountArmId pulumi.StringPtrOutput `pulumi:"customerStorageAccountArmId"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the nested sites under Master Site.
	NestedSites pulumi.StringArrayOutput `pulumi:"nestedSites"`
	// Gets the private endpoint connections.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// provisioning state enum
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets or sets the state of public network access.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Gets or sets the sites that are a part of Master Site.
	//             The key
	// should contain the Site ARM name.
	Sites pulumi.StringArrayOutput `pulumi:"sites"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMasterSitesController registers a new resource with the given unique name, arguments, and options.
func NewMasterSitesController(ctx *pulumi.Context,
	name string, args *MasterSitesControllerArgs, opts ...pulumi.ResourceOption) (*MasterSitesController, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:offazure/v20200707:MasterSitesController"),
		},
		{
			Type: pulumi.String("azure-native:offazure/v20230606:MasterSitesController"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MasterSitesController
	err := ctx.RegisterResource("azure-native:offazure:MasterSitesController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMasterSitesController gets an existing MasterSitesController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMasterSitesController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MasterSitesControllerState, opts ...pulumi.ResourceOption) (*MasterSitesController, error) {
	var resource MasterSitesController
	err := ctx.ReadResource("azure-native:offazure:MasterSitesController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MasterSitesController resources.
type masterSitesControllerState struct {
}

type MasterSitesControllerState struct {
}

func (MasterSitesControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*masterSitesControllerState)(nil)).Elem()
}

type masterSitesControllerArgs struct {
	// Gets or sets a value indicating whether multiple sites per site type are
	// allowed.
	AllowMultipleSites *bool `pulumi:"allowMultipleSites"`
	// Gets or sets a value for customer storage account ARM id.
	CustomerStorageAccountArmId *string `pulumi:"customerStorageAccountArmId"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Gets or sets the state of public network access.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name
	SiteName *string `pulumi:"siteName"`
	// Gets or sets the sites that are a part of Master Site.
	//             The key
	// should contain the Site ARM name.
	Sites []string `pulumi:"sites"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MasterSitesController resource.
type MasterSitesControllerArgs struct {
	// Gets or sets a value indicating whether multiple sites per site type are
	// allowed.
	AllowMultipleSites pulumi.BoolPtrInput
	// Gets or sets a value for customer storage account ARM id.
	CustomerStorageAccountArmId pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Gets or sets the state of public network access.
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Site name
	SiteName pulumi.StringPtrInput
	// Gets or sets the sites that are a part of Master Site.
	//             The key
	// should contain the Site ARM name.
	Sites pulumi.StringArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (MasterSitesControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*masterSitesControllerArgs)(nil)).Elem()
}

type MasterSitesControllerInput interface {
	pulumi.Input

	ToMasterSitesControllerOutput() MasterSitesControllerOutput
	ToMasterSitesControllerOutputWithContext(ctx context.Context) MasterSitesControllerOutput
}

func (*MasterSitesController) ElementType() reflect.Type {
	return reflect.TypeOf((**MasterSitesController)(nil)).Elem()
}

func (i *MasterSitesController) ToMasterSitesControllerOutput() MasterSitesControllerOutput {
	return i.ToMasterSitesControllerOutputWithContext(context.Background())
}

func (i *MasterSitesController) ToMasterSitesControllerOutputWithContext(ctx context.Context) MasterSitesControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MasterSitesControllerOutput)
}

func (i *MasterSitesController) ToOutput(ctx context.Context) pulumix.Output[*MasterSitesController] {
	return pulumix.Output[*MasterSitesController]{
		OutputState: i.ToMasterSitesControllerOutputWithContext(ctx).OutputState,
	}
}

type MasterSitesControllerOutput struct{ *pulumi.OutputState }

func (MasterSitesControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MasterSitesController)(nil)).Elem()
}

func (o MasterSitesControllerOutput) ToMasterSitesControllerOutput() MasterSitesControllerOutput {
	return o
}

func (o MasterSitesControllerOutput) ToMasterSitesControllerOutputWithContext(ctx context.Context) MasterSitesControllerOutput {
	return o
}

func (o MasterSitesControllerOutput) ToOutput(ctx context.Context) pulumix.Output[*MasterSitesController] {
	return pulumix.Output[*MasterSitesController]{
		OutputState: o.OutputState,
	}
}

// Gets or sets a value indicating whether multiple sites per site type are
// allowed.
func (o MasterSitesControllerOutput) AllowMultipleSites() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MasterSitesController) pulumi.BoolPtrOutput { return v.AllowMultipleSites }).(pulumi.BoolPtrOutput)
}

// Gets or sets a value for customer storage account ARM id.
func (o MasterSitesControllerOutput) CustomerStorageAccountArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MasterSitesController) pulumi.StringPtrOutput { return v.CustomerStorageAccountArmId }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o MasterSitesControllerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MasterSitesController) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o MasterSitesControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MasterSitesController) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the nested sites under Master Site.
func (o MasterSitesControllerOutput) NestedSites() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MasterSitesController) pulumi.StringArrayOutput { return v.NestedSites }).(pulumi.StringArrayOutput)
}

// Gets the private endpoint connections.
func (o MasterSitesControllerOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *MasterSitesController) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// provisioning state enum
func (o MasterSitesControllerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MasterSitesController) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets the state of public network access.
func (o MasterSitesControllerOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MasterSitesController) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Gets or sets the sites that are a part of Master Site.
//
//	The key
//
// should contain the Site ARM name.
func (o MasterSitesControllerOutput) Sites() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MasterSitesController) pulumi.StringArrayOutput { return v.Sites }).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MasterSitesControllerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MasterSitesController) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o MasterSitesControllerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MasterSitesController) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MasterSitesControllerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MasterSitesController) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MasterSitesControllerOutput{})
}
