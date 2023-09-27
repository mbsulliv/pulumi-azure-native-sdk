// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A Big Data pool
type BigDataPool struct {
	pulumi.CustomResourceState

	// Auto-pausing properties
	AutoPause AutoPausePropertiesResponsePtrOutput `pulumi:"autoPause"`
	// Auto-scaling properties
	AutoScale AutoScalePropertiesResponsePtrOutput `pulumi:"autoScale"`
	// The cache size
	CacheSize pulumi.IntOutput `pulumi:"cacheSize"`
	// The time when the Big Data pool was created.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// List of custom libraries/packages associated with the spark pool.
	CustomLibraries LibraryInfoResponseArrayOutput `pulumi:"customLibraries"`
	// The default folder where Spark logs will be written.
	DefaultSparkLogFolder pulumi.StringPtrOutput `pulumi:"defaultSparkLogFolder"`
	// Dynamic Executor Allocation
	DynamicExecutorAllocation DynamicExecutorAllocationResponsePtrOutput `pulumi:"dynamicExecutorAllocation"`
	// Whether autotune is required or not.
	IsAutotuneEnabled pulumi.BoolPtrOutput `pulumi:"isAutotuneEnabled"`
	// Whether compute isolation is required or not.
	IsComputeIsolationEnabled pulumi.BoolPtrOutput `pulumi:"isComputeIsolationEnabled"`
	// The time when the Big Data pool was updated successfully.
	LastSucceededTimestamp pulumi.StringOutput `pulumi:"lastSucceededTimestamp"`
	// Library version requirements
	LibraryRequirements LibraryRequirementsResponsePtrOutput `pulumi:"libraryRequirements"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of nodes in the Big Data pool.
	NodeCount pulumi.IntPtrOutput `pulumi:"nodeCount"`
	// The level of compute power that each node in the Big Data pool has.
	NodeSize pulumi.StringPtrOutput `pulumi:"nodeSize"`
	// The kind of nodes that the Big Data pool provides.
	NodeSizeFamily pulumi.StringPtrOutput `pulumi:"nodeSizeFamily"`
	// The state of the Big Data pool.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Whether session level packages enabled.
	SessionLevelPackagesEnabled pulumi.BoolPtrOutput `pulumi:"sessionLevelPackagesEnabled"`
	// Spark configuration file to specify additional properties
	SparkConfigProperties SparkConfigPropertiesResponsePtrOutput `pulumi:"sparkConfigProperties"`
	// The Spark events folder
	SparkEventsFolder pulumi.StringPtrOutput `pulumi:"sparkEventsFolder"`
	// The Apache Spark version.
	SparkVersion pulumi.StringPtrOutput `pulumi:"sparkVersion"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBigDataPool registers a new resource with the given unique name, arguments, and options.
func NewBigDataPool(ctx *pulumi.Context,
	name string, args *BigDataPoolArgs, opts ...pulumi.ResourceOption) (*BigDataPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:BigDataPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:BigDataPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:BigDataPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:BigDataPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:BigDataPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:BigDataPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:BigDataPool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BigDataPool
	err := ctx.RegisterResource("azure-native:synapse/v20210601preview:BigDataPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigDataPool gets an existing BigDataPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigDataPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigDataPoolState, opts ...pulumi.ResourceOption) (*BigDataPool, error) {
	var resource BigDataPool
	err := ctx.ReadResource("azure-native:synapse/v20210601preview:BigDataPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigDataPool resources.
type bigDataPoolState struct {
}

type BigDataPoolState struct {
}

func (BigDataPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigDataPoolState)(nil)).Elem()
}

type bigDataPoolArgs struct {
	// Auto-pausing properties
	AutoPause *AutoPauseProperties `pulumi:"autoPause"`
	// Auto-scaling properties
	AutoScale *AutoScaleProperties `pulumi:"autoScale"`
	// Big Data pool name
	BigDataPoolName *string `pulumi:"bigDataPoolName"`
	// List of custom libraries/packages associated with the spark pool.
	CustomLibraries []LibraryInfo `pulumi:"customLibraries"`
	// The default folder where Spark logs will be written.
	DefaultSparkLogFolder *string `pulumi:"defaultSparkLogFolder"`
	// Dynamic Executor Allocation
	DynamicExecutorAllocation *DynamicExecutorAllocation `pulumi:"dynamicExecutorAllocation"`
	// Whether to stop any running jobs in the Big Data pool
	Force *bool `pulumi:"force"`
	// Whether autotune is required or not.
	IsAutotuneEnabled *bool `pulumi:"isAutotuneEnabled"`
	// Whether compute isolation is required or not.
	IsComputeIsolationEnabled *bool `pulumi:"isComputeIsolationEnabled"`
	// Library version requirements
	LibraryRequirements *LibraryRequirements `pulumi:"libraryRequirements"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The number of nodes in the Big Data pool.
	NodeCount *int `pulumi:"nodeCount"`
	// The level of compute power that each node in the Big Data pool has.
	NodeSize *string `pulumi:"nodeSize"`
	// The kind of nodes that the Big Data pool provides.
	NodeSizeFamily *string `pulumi:"nodeSizeFamily"`
	// The state of the Big Data pool.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Whether session level packages enabled.
	SessionLevelPackagesEnabled *bool `pulumi:"sessionLevelPackagesEnabled"`
	// Spark configuration file to specify additional properties
	SparkConfigProperties *SparkConfigProperties `pulumi:"sparkConfigProperties"`
	// The Spark events folder
	SparkEventsFolder *string `pulumi:"sparkEventsFolder"`
	// The Apache Spark version.
	SparkVersion *string `pulumi:"sparkVersion"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a BigDataPool resource.
type BigDataPoolArgs struct {
	// Auto-pausing properties
	AutoPause AutoPausePropertiesPtrInput
	// Auto-scaling properties
	AutoScale AutoScalePropertiesPtrInput
	// Big Data pool name
	BigDataPoolName pulumi.StringPtrInput
	// List of custom libraries/packages associated with the spark pool.
	CustomLibraries LibraryInfoArrayInput
	// The default folder where Spark logs will be written.
	DefaultSparkLogFolder pulumi.StringPtrInput
	// Dynamic Executor Allocation
	DynamicExecutorAllocation DynamicExecutorAllocationPtrInput
	// Whether to stop any running jobs in the Big Data pool
	Force pulumi.BoolPtrInput
	// Whether autotune is required or not.
	IsAutotuneEnabled pulumi.BoolPtrInput
	// Whether compute isolation is required or not.
	IsComputeIsolationEnabled pulumi.BoolPtrInput
	// Library version requirements
	LibraryRequirements LibraryRequirementsPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The number of nodes in the Big Data pool.
	NodeCount pulumi.IntPtrInput
	// The level of compute power that each node in the Big Data pool has.
	NodeSize pulumi.StringPtrInput
	// The kind of nodes that the Big Data pool provides.
	NodeSizeFamily pulumi.StringPtrInput
	// The state of the Big Data pool.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Whether session level packages enabled.
	SessionLevelPackagesEnabled pulumi.BoolPtrInput
	// Spark configuration file to specify additional properties
	SparkConfigProperties SparkConfigPropertiesPtrInput
	// The Spark events folder
	SparkEventsFolder pulumi.StringPtrInput
	// The Apache Spark version.
	SparkVersion pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (BigDataPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigDataPoolArgs)(nil)).Elem()
}

type BigDataPoolInput interface {
	pulumi.Input

	ToBigDataPoolOutput() BigDataPoolOutput
	ToBigDataPoolOutputWithContext(ctx context.Context) BigDataPoolOutput
}

func (*BigDataPool) ElementType() reflect.Type {
	return reflect.TypeOf((**BigDataPool)(nil)).Elem()
}

func (i *BigDataPool) ToBigDataPoolOutput() BigDataPoolOutput {
	return i.ToBigDataPoolOutputWithContext(context.Background())
}

func (i *BigDataPool) ToBigDataPoolOutputWithContext(ctx context.Context) BigDataPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigDataPoolOutput)
}

func (i *BigDataPool) ToOutput(ctx context.Context) pulumix.Output[*BigDataPool] {
	return pulumix.Output[*BigDataPool]{
		OutputState: i.ToBigDataPoolOutputWithContext(ctx).OutputState,
	}
}

type BigDataPoolOutput struct{ *pulumi.OutputState }

func (BigDataPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigDataPool)(nil)).Elem()
}

func (o BigDataPoolOutput) ToBigDataPoolOutput() BigDataPoolOutput {
	return o
}

func (o BigDataPoolOutput) ToBigDataPoolOutputWithContext(ctx context.Context) BigDataPoolOutput {
	return o
}

func (o BigDataPoolOutput) ToOutput(ctx context.Context) pulumix.Output[*BigDataPool] {
	return pulumix.Output[*BigDataPool]{
		OutputState: o.OutputState,
	}
}

// Auto-pausing properties
func (o BigDataPoolOutput) AutoPause() AutoPausePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *BigDataPool) AutoPausePropertiesResponsePtrOutput { return v.AutoPause }).(AutoPausePropertiesResponsePtrOutput)
}

// Auto-scaling properties
func (o BigDataPoolOutput) AutoScale() AutoScalePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *BigDataPool) AutoScalePropertiesResponsePtrOutput { return v.AutoScale }).(AutoScalePropertiesResponsePtrOutput)
}

// The cache size
func (o BigDataPoolOutput) CacheSize() pulumi.IntOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.IntOutput { return v.CacheSize }).(pulumi.IntOutput)
}

// The time when the Big Data pool was created.
func (o BigDataPoolOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// List of custom libraries/packages associated with the spark pool.
func (o BigDataPoolOutput) CustomLibraries() LibraryInfoResponseArrayOutput {
	return o.ApplyT(func(v *BigDataPool) LibraryInfoResponseArrayOutput { return v.CustomLibraries }).(LibraryInfoResponseArrayOutput)
}

// The default folder where Spark logs will be written.
func (o BigDataPoolOutput) DefaultSparkLogFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringPtrOutput { return v.DefaultSparkLogFolder }).(pulumi.StringPtrOutput)
}

// Dynamic Executor Allocation
func (o BigDataPoolOutput) DynamicExecutorAllocation() DynamicExecutorAllocationResponsePtrOutput {
	return o.ApplyT(func(v *BigDataPool) DynamicExecutorAllocationResponsePtrOutput { return v.DynamicExecutorAllocation }).(DynamicExecutorAllocationResponsePtrOutput)
}

// Whether autotune is required or not.
func (o BigDataPoolOutput) IsAutotuneEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.BoolPtrOutput { return v.IsAutotuneEnabled }).(pulumi.BoolPtrOutput)
}

// Whether compute isolation is required or not.
func (o BigDataPoolOutput) IsComputeIsolationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.BoolPtrOutput { return v.IsComputeIsolationEnabled }).(pulumi.BoolPtrOutput)
}

// The time when the Big Data pool was updated successfully.
func (o BigDataPoolOutput) LastSucceededTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringOutput { return v.LastSucceededTimestamp }).(pulumi.StringOutput)
}

// Library version requirements
func (o BigDataPoolOutput) LibraryRequirements() LibraryRequirementsResponsePtrOutput {
	return o.ApplyT(func(v *BigDataPool) LibraryRequirementsResponsePtrOutput { return v.LibraryRequirements }).(LibraryRequirementsResponsePtrOutput)
}

// The geo-location where the resource lives
func (o BigDataPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o BigDataPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of nodes in the Big Data pool.
func (o BigDataPoolOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.IntPtrOutput { return v.NodeCount }).(pulumi.IntPtrOutput)
}

// The level of compute power that each node in the Big Data pool has.
func (o BigDataPoolOutput) NodeSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringPtrOutput { return v.NodeSize }).(pulumi.StringPtrOutput)
}

// The kind of nodes that the Big Data pool provides.
func (o BigDataPoolOutput) NodeSizeFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringPtrOutput { return v.NodeSizeFamily }).(pulumi.StringPtrOutput)
}

// The state of the Big Data pool.
func (o BigDataPoolOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Whether session level packages enabled.
func (o BigDataPoolOutput) SessionLevelPackagesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.BoolPtrOutput { return v.SessionLevelPackagesEnabled }).(pulumi.BoolPtrOutput)
}

// Spark configuration file to specify additional properties
func (o BigDataPoolOutput) SparkConfigProperties() SparkConfigPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *BigDataPool) SparkConfigPropertiesResponsePtrOutput { return v.SparkConfigProperties }).(SparkConfigPropertiesResponsePtrOutput)
}

// The Spark events folder
func (o BigDataPoolOutput) SparkEventsFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringPtrOutput { return v.SparkEventsFolder }).(pulumi.StringPtrOutput)
}

// The Apache Spark version.
func (o BigDataPoolOutput) SparkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringPtrOutput { return v.SparkVersion }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o BigDataPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BigDataPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BigDataPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BigDataPoolOutput{})
}
