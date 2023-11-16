// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VMware collector resource.
// Azure REST API version: 2023-03-15.
type VmwareCollectorsOperation struct {
	pulumi.CustomResourceState

	// Gets or sets the collector agent properties.
	AgentProperties CollectorAgentPropertiesBaseResponsePtrOutput `pulumi:"agentProperties"`
	// Gets the Timestamp when collector was created.
	CreatedTimestamp pulumi.StringOutput `pulumi:"createdTimestamp"`
	// Gets the discovery site id.
	DiscoverySiteId pulumi.StringPtrOutput `pulumi:"discoverySiteId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Timestamp when collector was last updated.
	UpdatedTimestamp pulumi.StringOutput `pulumi:"updatedTimestamp"`
}

// NewVmwareCollectorsOperation registers a new resource with the given unique name, arguments, and options.
func NewVmwareCollectorsOperation(ctx *pulumi.Context,
	name string, args *VmwareCollectorsOperationArgs, opts ...pulumi.ResourceOption) (*VmwareCollectorsOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20191001:VmwareCollectorsOperation"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230315:VmwareCollectorsOperation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VmwareCollectorsOperation
	err := ctx.RegisterResource("azure-native:migrate:VmwareCollectorsOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVmwareCollectorsOperation gets an existing VmwareCollectorsOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVmwareCollectorsOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmwareCollectorsOperationState, opts ...pulumi.ResourceOption) (*VmwareCollectorsOperation, error) {
	var resource VmwareCollectorsOperation
	err := ctx.ReadResource("azure-native:migrate:VmwareCollectorsOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VmwareCollectorsOperation resources.
type vmwareCollectorsOperationState struct {
}

type VmwareCollectorsOperationState struct {
}

func (VmwareCollectorsOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmwareCollectorsOperationState)(nil)).Elem()
}

type vmwareCollectorsOperationArgs struct {
	// Gets or sets the collector agent properties.
	AgentProperties *CollectorAgentPropertiesBase `pulumi:"agentProperties"`
	// Gets the discovery site id.
	DiscoverySiteId *string `pulumi:"discoverySiteId"`
	// Assessment Project Name
	ProjectName string `pulumi:"projectName"`
	// The status of the last operation.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// VMware collector ARM name
	VmWareCollectorName *string `pulumi:"vmWareCollectorName"`
}

// The set of arguments for constructing a VmwareCollectorsOperation resource.
type VmwareCollectorsOperationArgs struct {
	// Gets or sets the collector agent properties.
	AgentProperties CollectorAgentPropertiesBasePtrInput
	// Gets the discovery site id.
	DiscoverySiteId pulumi.StringPtrInput
	// Assessment Project Name
	ProjectName pulumi.StringInput
	// The status of the last operation.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// VMware collector ARM name
	VmWareCollectorName pulumi.StringPtrInput
}

func (VmwareCollectorsOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmwareCollectorsOperationArgs)(nil)).Elem()
}

type VmwareCollectorsOperationInput interface {
	pulumi.Input

	ToVmwareCollectorsOperationOutput() VmwareCollectorsOperationOutput
	ToVmwareCollectorsOperationOutputWithContext(ctx context.Context) VmwareCollectorsOperationOutput
}

func (*VmwareCollectorsOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**VmwareCollectorsOperation)(nil)).Elem()
}

func (i *VmwareCollectorsOperation) ToVmwareCollectorsOperationOutput() VmwareCollectorsOperationOutput {
	return i.ToVmwareCollectorsOperationOutputWithContext(context.Background())
}

func (i *VmwareCollectorsOperation) ToVmwareCollectorsOperationOutputWithContext(ctx context.Context) VmwareCollectorsOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmwareCollectorsOperationOutput)
}

type VmwareCollectorsOperationOutput struct{ *pulumi.OutputState }

func (VmwareCollectorsOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmwareCollectorsOperation)(nil)).Elem()
}

func (o VmwareCollectorsOperationOutput) ToVmwareCollectorsOperationOutput() VmwareCollectorsOperationOutput {
	return o
}

func (o VmwareCollectorsOperationOutput) ToVmwareCollectorsOperationOutputWithContext(ctx context.Context) VmwareCollectorsOperationOutput {
	return o
}

// Gets or sets the collector agent properties.
func (o VmwareCollectorsOperationOutput) AgentProperties() CollectorAgentPropertiesBaseResponsePtrOutput {
	return o.ApplyT(func(v *VmwareCollectorsOperation) CollectorAgentPropertiesBaseResponsePtrOutput {
		return v.AgentProperties
	}).(CollectorAgentPropertiesBaseResponsePtrOutput)
}

// Gets the Timestamp when collector was created.
func (o VmwareCollectorsOperationOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareCollectorsOperation) pulumi.StringOutput { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// Gets the discovery site id.
func (o VmwareCollectorsOperationOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VmwareCollectorsOperation) pulumi.StringPtrOutput { return v.DiscoverySiteId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o VmwareCollectorsOperationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareCollectorsOperation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o VmwareCollectorsOperationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VmwareCollectorsOperation) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o VmwareCollectorsOperationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VmwareCollectorsOperation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o VmwareCollectorsOperationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareCollectorsOperation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Timestamp when collector was last updated.
func (o VmwareCollectorsOperationOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareCollectorsOperation) pulumi.StringOutput { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VmwareCollectorsOperationOutput{})
}
