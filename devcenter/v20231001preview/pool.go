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

// A pool of Virtual Machines.
type Pool struct {
	pulumi.CustomResourceState

	// Indicates the number of provisioned Dev Boxes in this pool.
	DevBoxCount pulumi.IntOutput `pulumi:"devBoxCount"`
	// Name of a Dev Box definition in parent Project of this Pool
	DevBoxDefinitionName pulumi.StringOutput `pulumi:"devBoxDefinitionName"`
	// The display name of the pool.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Overall health status of the Pool. Indicates whether or not the Pool is available to create Dev Boxes.
	HealthStatus pulumi.StringOutput `pulumi:"healthStatus"`
	// Details on the Pool health status to help diagnose issues. This is only populated when the pool status indicates the pool is in a non-healthy state
	HealthStatusDetails HealthStatusDetailResponseArrayOutput `pulumi:"healthStatusDetails"`
	// Specifies the license type indicating the caller has already acquired licenses for the Dev Boxes that will be created.
	LicenseType pulumi.StringOutput `pulumi:"licenseType"`
	// Indicates whether owners of Dev Boxes in this pool are added as local administrators on the Dev Box.
	LocalAdministrator pulumi.StringOutput `pulumi:"localAdministrator"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The regions of the managed virtual network (required when managedNetworkType is Managed).
	ManagedVirtualNetworkRegions pulumi.StringArrayOutput `pulumi:"managedVirtualNetworkRegions"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of a Network Connection in parent Project of this Pool
	NetworkConnectionName pulumi.StringOutput `pulumi:"networkConnectionName"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Indicates whether Dev Boxes in this pool are created with single sign on enabled. The also requires that single sign on be enabled on the tenant.
	SingleSignOnStatus pulumi.StringPtrOutput `pulumi:"singleSignOnStatus"`
	// Stop on disconnect configuration settings for Dev Boxes created in this pool.
	StopOnDisconnect StopOnDisconnectConfigurationResponsePtrOutput `pulumi:"stopOnDisconnect"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Indicates whether the pool uses a Virtual Network managed by Microsoft or a customer provided network.
	VirtualNetworkType pulumi.StringPtrOutput `pulumi:"virtualNetworkType"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevBoxDefinitionName == nil {
		return nil, errors.New("invalid value for required argument 'DevBoxDefinitionName'")
	}
	if args.LicenseType == nil {
		return nil, errors.New("invalid value for required argument 'LicenseType'")
	}
	if args.LocalAdministrator == nil {
		return nil, errors.New("invalid value for required argument 'LocalAdministrator'")
	}
	if args.NetworkConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkConnectionName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221111preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230101preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230401:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230801preview:Pool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Pool
	err := ctx.RegisterResource("azure-native:devcenter/v20231001preview:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPool gets an existing Pool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("azure-native:devcenter/v20231001preview:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pool resources.
type poolState struct {
}

type PoolState struct {
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	// Name of a Dev Box definition in parent Project of this Pool
	DevBoxDefinitionName string `pulumi:"devBoxDefinitionName"`
	// The display name of the pool.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the license type indicating the caller has already acquired licenses for the Dev Boxes that will be created.
	LicenseType string `pulumi:"licenseType"`
	// Indicates whether owners of Dev Boxes in this pool are added as local administrators on the Dev Box.
	LocalAdministrator string `pulumi:"localAdministrator"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The regions of the managed virtual network (required when managedNetworkType is Managed).
	ManagedVirtualNetworkRegions []string `pulumi:"managedVirtualNetworkRegions"`
	// Name of a Network Connection in parent Project of this Pool
	NetworkConnectionName string `pulumi:"networkConnectionName"`
	// Name of the pool.
	PoolName *string `pulumi:"poolName"`
	// The name of the project.
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Indicates whether Dev Boxes in this pool are created with single sign on enabled. The also requires that single sign on be enabled on the tenant.
	SingleSignOnStatus *string `pulumi:"singleSignOnStatus"`
	// Stop on disconnect configuration settings for Dev Boxes created in this pool.
	StopOnDisconnect *StopOnDisconnectConfiguration `pulumi:"stopOnDisconnect"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Indicates whether the pool uses a Virtual Network managed by Microsoft or a customer provided network.
	VirtualNetworkType *string `pulumi:"virtualNetworkType"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// Name of a Dev Box definition in parent Project of this Pool
	DevBoxDefinitionName pulumi.StringInput
	// The display name of the pool.
	DisplayName pulumi.StringPtrInput
	// Specifies the license type indicating the caller has already acquired licenses for the Dev Boxes that will be created.
	LicenseType pulumi.StringInput
	// Indicates whether owners of Dev Boxes in this pool are added as local administrators on the Dev Box.
	LocalAdministrator pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The regions of the managed virtual network (required when managedNetworkType is Managed).
	ManagedVirtualNetworkRegions pulumi.StringArrayInput
	// Name of a Network Connection in parent Project of this Pool
	NetworkConnectionName pulumi.StringInput
	// Name of the pool.
	PoolName pulumi.StringPtrInput
	// The name of the project.
	ProjectName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Indicates whether Dev Boxes in this pool are created with single sign on enabled. The also requires that single sign on be enabled on the tenant.
	SingleSignOnStatus pulumi.StringPtrInput
	// Stop on disconnect configuration settings for Dev Boxes created in this pool.
	StopOnDisconnect StopOnDisconnectConfigurationPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Indicates whether the pool uses a Virtual Network managed by Microsoft or a customer provided network.
	VirtualNetworkType pulumi.StringPtrInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

type PoolOutput struct{ *pulumi.OutputState }

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

// Indicates the number of provisioned Dev Boxes in this pool.
func (o PoolOutput) DevBoxCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Pool) pulumi.IntOutput { return v.DevBoxCount }).(pulumi.IntOutput)
}

// Name of a Dev Box definition in parent Project of this Pool
func (o PoolOutput) DevBoxDefinitionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.DevBoxDefinitionName }).(pulumi.StringOutput)
}

// The display name of the pool.
func (o PoolOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Overall health status of the Pool. Indicates whether or not the Pool is available to create Dev Boxes.
func (o PoolOutput) HealthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.HealthStatus }).(pulumi.StringOutput)
}

// Details on the Pool health status to help diagnose issues. This is only populated when the pool status indicates the pool is in a non-healthy state
func (o PoolOutput) HealthStatusDetails() HealthStatusDetailResponseArrayOutput {
	return o.ApplyT(func(v *Pool) HealthStatusDetailResponseArrayOutput { return v.HealthStatusDetails }).(HealthStatusDetailResponseArrayOutput)
}

// Specifies the license type indicating the caller has already acquired licenses for the Dev Boxes that will be created.
func (o PoolOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.LicenseType }).(pulumi.StringOutput)
}

// Indicates whether owners of Dev Boxes in this pool are added as local administrators on the Dev Box.
func (o PoolOutput) LocalAdministrator() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.LocalAdministrator }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o PoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The regions of the managed virtual network (required when managedNetworkType is Managed).
func (o PoolOutput) ManagedVirtualNetworkRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringArrayOutput { return v.ManagedVirtualNetworkRegions }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o PoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name of a Network Connection in parent Project of this Pool
func (o PoolOutput) NetworkConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.NetworkConnectionName }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o PoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Indicates whether Dev Boxes in this pool are created with single sign on enabled. The also requires that single sign on be enabled on the tenant.
func (o PoolOutput) SingleSignOnStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringPtrOutput { return v.SingleSignOnStatus }).(pulumi.StringPtrOutput)
}

// Stop on disconnect configuration settings for Dev Boxes created in this pool.
func (o PoolOutput) StopOnDisconnect() StopOnDisconnectConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Pool) StopOnDisconnectConfigurationResponsePtrOutput { return v.StopOnDisconnect }).(StopOnDisconnectConfigurationResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Pool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o PoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Indicates whether the pool uses a Virtual Network managed by Microsoft or a customer provided network.
func (o PoolOutput) VirtualNetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringPtrOutput { return v.VirtualNetworkType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PoolOutput{})
}
