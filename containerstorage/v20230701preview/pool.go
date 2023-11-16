// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Pool resource
type Pool struct {
	pulumi.CustomResourceState

	// List of resources that should have access to the pool. Typically ARM references to AKS clusters or ACI Container Groups. For local and standard this must be a single reference. For ElasticSAN there can be many.
	Assignments AssignmentResponseArrayOutput `pulumi:"assignments"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Type of the Pool: ephemeralDisk, azureDisk, or elasticsan.
	PoolType PoolTypeResponseOutput `pulumi:"poolType"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// ReclaimPolicy defines what happens to the backend storage when StoragePool is deleted
	ReclaimPolicy pulumi.StringPtrOutput `pulumi:"reclaimPolicy"`
	// Resources represent the resources the pool should have.
	Resources ResourcesResponsePtrOutput `pulumi:"resources"`
	// The operational status of the resource
	Status ResourceOperationalStatusResponseOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// List of availability zones that resources can be created in.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PoolType == nil {
		return nil, errors.New("invalid value for required argument 'PoolType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.PoolType = args.PoolType.ToPoolTypeOutput().ApplyT(func(v PoolType) PoolType { return *v.Defaults() }).(PoolTypeOutput)
	if args.Resources != nil {
		args.Resources = args.Resources.ToResourcesPtrOutput().ApplyT(func(v *Resources) *Resources { return v.Defaults() }).(ResourcesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerstorage:Pool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Pool
	err := ctx.RegisterResource("azure-native:containerstorage/v20230701preview:Pool", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:containerstorage/v20230701preview:Pool", name, id, state, &resource, opts...)
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
	// List of resources that should have access to the pool. Typically ARM references to AKS clusters or ACI Container Groups. For local and standard this must be a single reference. For ElasticSAN there can be many.
	Assignments []Assignment `pulumi:"assignments"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Pool Object
	PoolName *string `pulumi:"poolName"`
	// Type of the Pool: ephemeralDisk, azureDisk, or elasticsan.
	PoolType PoolType `pulumi:"poolType"`
	// ReclaimPolicy defines what happens to the backend storage when StoragePool is deleted
	ReclaimPolicy *string `pulumi:"reclaimPolicy"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resources represent the resources the pool should have.
	Resources *Resources `pulumi:"resources"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// List of availability zones that resources can be created in.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// List of resources that should have access to the pool. Typically ARM references to AKS clusters or ACI Container Groups. For local and standard this must be a single reference. For ElasticSAN there can be many.
	Assignments AssignmentArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Pool Object
	PoolName pulumi.StringPtrInput
	// Type of the Pool: ephemeralDisk, azureDisk, or elasticsan.
	PoolType PoolTypeInput
	// ReclaimPolicy defines what happens to the backend storage when StoragePool is deleted
	ReclaimPolicy pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resources represent the resources the pool should have.
	Resources ResourcesPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// List of availability zones that resources can be created in.
	Zones pulumi.StringArrayInput
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

// List of resources that should have access to the pool. Typically ARM references to AKS clusters or ACI Container Groups. For local and standard this must be a single reference. For ElasticSAN there can be many.
func (o PoolOutput) Assignments() AssignmentResponseArrayOutput {
	return o.ApplyT(func(v *Pool) AssignmentResponseArrayOutput { return v.Assignments }).(AssignmentResponseArrayOutput)
}

// The geo-location where the resource lives
func (o PoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o PoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Type of the Pool: ephemeralDisk, azureDisk, or elasticsan.
func (o PoolOutput) PoolType() PoolTypeResponseOutput {
	return o.ApplyT(func(v *Pool) PoolTypeResponseOutput { return v.PoolType }).(PoolTypeResponseOutput)
}

// The status of the last operation.
func (o PoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// ReclaimPolicy defines what happens to the backend storage when StoragePool is deleted
func (o PoolOutput) ReclaimPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringPtrOutput { return v.ReclaimPolicy }).(pulumi.StringPtrOutput)
}

// Resources represent the resources the pool should have.
func (o PoolOutput) Resources() ResourcesResponsePtrOutput {
	return o.ApplyT(func(v *Pool) ResourcesResponsePtrOutput { return v.Resources }).(ResourcesResponsePtrOutput)
}

// The operational status of the resource
func (o PoolOutput) Status() ResourceOperationalStatusResponseOutput {
	return o.ApplyT(func(v *Pool) ResourceOperationalStatusResponseOutput { return v.Status }).(ResourceOperationalStatusResponseOutput)
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

// List of availability zones that resources can be created in.
func (o PoolOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(PoolOutput{})
}
