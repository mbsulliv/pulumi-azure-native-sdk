// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A pool of Virtual Machines.
type Pool struct {
	pulumi.CustomResourceState

	// Name of a Dev Box definition in parent Project of this Pool
	DevBoxDefinitionName pulumi.StringOutput `pulumi:"devBoxDefinitionName"`
	// Specifies the license type indicating the caller has already acquired licenses for the Dev Boxes that will be created.
	LicenseType pulumi.StringOutput `pulumi:"licenseType"`
	// Indicates whether owners of Dev Boxes in this pool are added as local administrators on the Dev Box.
	LocalAdministrator pulumi.StringOutput `pulumi:"localAdministrator"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of a Network Connection in parent Project of this Pool
	NetworkConnectionName pulumi.StringOutput `pulumi:"networkConnectionName"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
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
			Type: pulumi.String("azure-native:devcenter/v20220901preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221111preview:Pool"),
		},
	})
	opts = append(opts, aliases)
	var resource Pool
	err := ctx.RegisterResource("azure-native:devcenter/v20220801preview:Pool", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:devcenter/v20220801preview:Pool", name, id, state, &resource, opts...)
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
	// Specifies the license type indicating the caller has already acquired licenses for the Dev Boxes that will be created.
	LicenseType string `pulumi:"licenseType"`
	// Indicates whether owners of Dev Boxes in this pool are added as local administrators on the Dev Box.
	LocalAdministrator string `pulumi:"localAdministrator"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of a Network Connection in parent Project of this Pool
	NetworkConnectionName string `pulumi:"networkConnectionName"`
	// Name of the pool.
	PoolName *string `pulumi:"poolName"`
	// The name of the project.
	ProjectName string `pulumi:"projectName"`
	// Name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// Name of a Dev Box definition in parent Project of this Pool
	DevBoxDefinitionName pulumi.StringInput
	// Specifies the license type indicating the caller has already acquired licenses for the Dev Boxes that will be created.
	LicenseType pulumi.StringInput
	// Indicates whether owners of Dev Boxes in this pool are added as local administrators on the Dev Box.
	LocalAdministrator pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of a Network Connection in parent Project of this Pool
	NetworkConnectionName pulumi.StringInput
	// Name of the pool.
	PoolName pulumi.StringPtrInput
	// The name of the project.
	ProjectName pulumi.StringInput
	// Name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
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

// Name of a Dev Box definition in parent Project of this Pool
func (o PoolOutput) DevBoxDefinitionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.DevBoxDefinitionName }).(pulumi.StringOutput)
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

func init() {
	pulumi.RegisterOutputType(PoolOutput{})
}
