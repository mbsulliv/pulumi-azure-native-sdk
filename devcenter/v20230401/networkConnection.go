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

// Network related settings
type NetworkConnection struct {
	pulumi.CustomResourceState

	// AAD Join type.
	DomainJoinType pulumi.StringOutput `pulumi:"domainJoinType"`
	// Active Directory domain name
	DomainName pulumi.StringPtrOutput `pulumi:"domainName"`
	// The password for the account used to join domain
	DomainPassword pulumi.StringPtrOutput `pulumi:"domainPassword"`
	// The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com.
	DomainUsername pulumi.StringPtrOutput `pulumi:"domainUsername"`
	// Overall health status of the network connection. Health checks are run on creation, update, and periodically to validate the network connection.
	HealthCheckStatus pulumi.StringOutput `pulumi:"healthCheckStatus"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The name for resource group where NICs will be placed.
	NetworkingResourceGroupName pulumi.StringPtrOutput `pulumi:"networkingResourceGroupName"`
	// Active Directory domain Organization Unit (OU)
	OrganizationUnit pulumi.StringPtrOutput `pulumi:"organizationUnit"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The subnet to attach Virtual Machines to
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkConnection registers a new resource with the given unique name, arguments, and options.
func NewNetworkConnection(ctx *pulumi.Context,
	name string, args *NetworkConnectionArgs, opts ...pulumi.ResourceOption) (*NetworkConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainJoinType == nil {
		return nil, errors.New("invalid value for required argument 'DomainJoinType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:NetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:NetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:NetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:NetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221111preview:NetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230101preview:NetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230801preview:NetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20231001preview:NetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20240201:NetworkConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NetworkConnection
	err := ctx.RegisterResource("azure-native:devcenter/v20230401:NetworkConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkConnection gets an existing NetworkConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkConnectionState, opts ...pulumi.ResourceOption) (*NetworkConnection, error) {
	var resource NetworkConnection
	err := ctx.ReadResource("azure-native:devcenter/v20230401:NetworkConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkConnection resources.
type networkConnectionState struct {
}

type NetworkConnectionState struct {
}

func (NetworkConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectionState)(nil)).Elem()
}

type networkConnectionArgs struct {
	// AAD Join type.
	DomainJoinType string `pulumi:"domainJoinType"`
	// Active Directory domain name
	DomainName *string `pulumi:"domainName"`
	// The password for the account used to join domain
	DomainPassword *string `pulumi:"domainPassword"`
	// The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com.
	DomainUsername *string `pulumi:"domainUsername"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of the Network Connection that can be applied to a Pool.
	NetworkConnectionName *string `pulumi:"networkConnectionName"`
	// The name for resource group where NICs will be placed.
	NetworkingResourceGroupName *string `pulumi:"networkingResourceGroupName"`
	// Active Directory domain Organization Unit (OU)
	OrganizationUnit *string `pulumi:"organizationUnit"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The subnet to attach Virtual Machines to
	SubnetId string `pulumi:"subnetId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkConnection resource.
type NetworkConnectionArgs struct {
	// AAD Join type.
	DomainJoinType pulumi.StringInput
	// Active Directory domain name
	DomainName pulumi.StringPtrInput
	// The password for the account used to join domain
	DomainPassword pulumi.StringPtrInput
	// The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com.
	DomainUsername pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of the Network Connection that can be applied to a Pool.
	NetworkConnectionName pulumi.StringPtrInput
	// The name for resource group where NICs will be placed.
	NetworkingResourceGroupName pulumi.StringPtrInput
	// Active Directory domain Organization Unit (OU)
	OrganizationUnit pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The subnet to attach Virtual Machines to
	SubnetId pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectionArgs)(nil)).Elem()
}

type NetworkConnectionInput interface {
	pulumi.Input

	ToNetworkConnectionOutput() NetworkConnectionOutput
	ToNetworkConnectionOutputWithContext(ctx context.Context) NetworkConnectionOutput
}

func (*NetworkConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConnection)(nil)).Elem()
}

func (i *NetworkConnection) ToNetworkConnectionOutput() NetworkConnectionOutput {
	return i.ToNetworkConnectionOutputWithContext(context.Background())
}

func (i *NetworkConnection) ToNetworkConnectionOutputWithContext(ctx context.Context) NetworkConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConnectionOutput)
}

type NetworkConnectionOutput struct{ *pulumi.OutputState }

func (NetworkConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConnection)(nil)).Elem()
}

func (o NetworkConnectionOutput) ToNetworkConnectionOutput() NetworkConnectionOutput {
	return o
}

func (o NetworkConnectionOutput) ToNetworkConnectionOutputWithContext(ctx context.Context) NetworkConnectionOutput {
	return o
}

// AAD Join type.
func (o NetworkConnectionOutput) DomainJoinType() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.DomainJoinType }).(pulumi.StringOutput)
}

// Active Directory domain name
func (o NetworkConnectionOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringPtrOutput { return v.DomainName }).(pulumi.StringPtrOutput)
}

// The password for the account used to join domain
func (o NetworkConnectionOutput) DomainPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringPtrOutput { return v.DomainPassword }).(pulumi.StringPtrOutput)
}

// The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com.
func (o NetworkConnectionOutput) DomainUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringPtrOutput { return v.DomainUsername }).(pulumi.StringPtrOutput)
}

// Overall health status of the network connection. Health checks are run on creation, update, and periodically to validate the network connection.
func (o NetworkConnectionOutput) HealthCheckStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.HealthCheckStatus }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o NetworkConnectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o NetworkConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name for resource group where NICs will be placed.
func (o NetworkConnectionOutput) NetworkingResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringPtrOutput { return v.NetworkingResourceGroupName }).(pulumi.StringPtrOutput)
}

// Active Directory domain Organization Unit (OU)
func (o NetworkConnectionOutput) OrganizationUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringPtrOutput { return v.OrganizationUnit }).(pulumi.StringPtrOutput)
}

// The provisioning state of the resource.
func (o NetworkConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The subnet to attach Virtual Machines to
func (o NetworkConnectionOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o NetworkConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o NetworkConnectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NetworkConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkConnectionOutput{})
}
