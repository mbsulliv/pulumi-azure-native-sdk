// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230815

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Class representing a cluster principal assignment.
type ClusterPrincipalAssignment struct {
	pulumi.CustomResourceState

	// The service principal object id in AAD (Azure active directory)
	AadObjectId pulumi.StringOutput `pulumi:"aadObjectId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// The principal name
	PrincipalName pulumi.StringOutput `pulumi:"principalName"`
	// Principal type.
	PrincipalType pulumi.StringOutput `pulumi:"principalType"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Cluster principal role.
	Role pulumi.StringOutput `pulumi:"role"`
	// The tenant id of the principal
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The tenant name of the principal
	TenantName pulumi.StringOutput `pulumi:"tenantName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewClusterPrincipalAssignment registers a new resource with the given unique name, arguments, and options.
func NewClusterPrincipalAssignment(ctx *pulumi.Context,
	name string, args *ClusterPrincipalAssignmentArgs, opts ...pulumi.ResourceOption) (*ClusterPrincipalAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	if args.PrincipalType == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221111:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221229:ClusterPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230502:ClusterPrincipalAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ClusterPrincipalAssignment
	err := ctx.RegisterResource("azure-native:kusto/v20230815:ClusterPrincipalAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterPrincipalAssignment gets an existing ClusterPrincipalAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterPrincipalAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterPrincipalAssignmentState, opts ...pulumi.ResourceOption) (*ClusterPrincipalAssignment, error) {
	var resource ClusterPrincipalAssignment
	err := ctx.ReadResource("azure-native:kusto/v20230815:ClusterPrincipalAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterPrincipalAssignment resources.
type clusterPrincipalAssignmentState struct {
}

type ClusterPrincipalAssignmentState struct {
}

func (ClusterPrincipalAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPrincipalAssignmentState)(nil)).Elem()
}

type clusterPrincipalAssignmentArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName *string `pulumi:"principalAssignmentName"`
	// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
	PrincipalId string `pulumi:"principalId"`
	// Principal type.
	PrincipalType string `pulumi:"principalType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cluster principal role.
	Role string `pulumi:"role"`
	// The tenant id of the principal
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a ClusterPrincipalAssignment resource.
type ClusterPrincipalAssignmentArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName pulumi.StringPtrInput
	// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
	PrincipalId pulumi.StringInput
	// Principal type.
	PrincipalType pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Cluster principal role.
	Role pulumi.StringInput
	// The tenant id of the principal
	TenantId pulumi.StringPtrInput
}

func (ClusterPrincipalAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPrincipalAssignmentArgs)(nil)).Elem()
}

type ClusterPrincipalAssignmentInput interface {
	pulumi.Input

	ToClusterPrincipalAssignmentOutput() ClusterPrincipalAssignmentOutput
	ToClusterPrincipalAssignmentOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentOutput
}

func (*ClusterPrincipalAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPrincipalAssignment)(nil)).Elem()
}

func (i *ClusterPrincipalAssignment) ToClusterPrincipalAssignmentOutput() ClusterPrincipalAssignmentOutput {
	return i.ToClusterPrincipalAssignmentOutputWithContext(context.Background())
}

func (i *ClusterPrincipalAssignment) ToClusterPrincipalAssignmentOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPrincipalAssignmentOutput)
}

func (i *ClusterPrincipalAssignment) ToOutput(ctx context.Context) pulumix.Output[*ClusterPrincipalAssignment] {
	return pulumix.Output[*ClusterPrincipalAssignment]{
		OutputState: i.ToClusterPrincipalAssignmentOutputWithContext(ctx).OutputState,
	}
}

type ClusterPrincipalAssignmentOutput struct{ *pulumi.OutputState }

func (ClusterPrincipalAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPrincipalAssignment)(nil)).Elem()
}

func (o ClusterPrincipalAssignmentOutput) ToClusterPrincipalAssignmentOutput() ClusterPrincipalAssignmentOutput {
	return o
}

func (o ClusterPrincipalAssignmentOutput) ToClusterPrincipalAssignmentOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentOutput {
	return o
}

func (o ClusterPrincipalAssignmentOutput) ToOutput(ctx context.Context) pulumix.Output[*ClusterPrincipalAssignment] {
	return pulumix.Output[*ClusterPrincipalAssignment]{
		OutputState: o.OutputState,
	}
}

// The service principal object id in AAD (Azure active directory)
func (o ClusterPrincipalAssignmentOutput) AadObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.AadObjectId }).(pulumi.StringOutput)
}

// The name of the resource
func (o ClusterPrincipalAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
func (o ClusterPrincipalAssignmentOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

// The principal name
func (o ClusterPrincipalAssignmentOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

// Principal type.
func (o ClusterPrincipalAssignmentOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o ClusterPrincipalAssignmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Cluster principal role.
func (o ClusterPrincipalAssignmentOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The tenant id of the principal
func (o ClusterPrincipalAssignmentOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The tenant name of the principal
func (o ClusterPrincipalAssignmentOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.TenantName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ClusterPrincipalAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPrincipalAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterPrincipalAssignmentOutput{})
}
