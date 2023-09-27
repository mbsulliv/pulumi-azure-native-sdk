// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type BatchDeployment struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	BatchDeploymentDetails BatchDeploymentResponseOutput `pulumi:"batchDeploymentDetails"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Sku details required for ARM contract for Autoscaling.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBatchDeployment registers a new resource with the given unique name, arguments, and options.
func NewBatchDeployment(ctx *pulumi.Context,
	name string, args *BatchDeploymentArgs, opts ...pulumi.ResourceOption) (*BatchDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BatchDeploymentDetails == nil {
		return nil, errors.New("invalid value for required argument 'BatchDeploymentDetails'")
	}
	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.BatchDeploymentDetails = args.BatchDeploymentDetails.ToBatchDeploymentTypeOutput().ApplyT(func(v BatchDeploymentType) BatchDeploymentType { return *v.Defaults() }).(BatchDeploymentTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:BatchDeployment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BatchDeployment
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220201preview:BatchDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBatchDeployment gets an existing BatchDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBatchDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BatchDeploymentState, opts ...pulumi.ResourceOption) (*BatchDeployment, error) {
	var resource BatchDeployment
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220201preview:BatchDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BatchDeployment resources.
type batchDeploymentState struct {
}

type BatchDeploymentState struct {
}

func (BatchDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*batchDeploymentState)(nil)).Elem()
}

type batchDeploymentArgs struct {
	// [Required] Additional attributes of the entity.
	BatchDeploymentDetails BatchDeploymentType `pulumi:"batchDeploymentDetails"`
	// The identifier for the Batch inference deployment.
	DeploymentName *string `pulumi:"deploymentName"`
	// Inference endpoint name
	EndpointName string `pulumi:"endpointName"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sku details required for ARM contract for Autoscaling.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a BatchDeployment resource.
type BatchDeploymentArgs struct {
	// [Required] Additional attributes of the entity.
	BatchDeploymentDetails BatchDeploymentTypeInput
	// The identifier for the Batch inference deployment.
	DeploymentName pulumi.StringPtrInput
	// Inference endpoint name
	EndpointName pulumi.StringInput
	// Managed service identity (system assigned and/or user assigned identities)
	Identity ManagedServiceIdentityPtrInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Sku details required for ARM contract for Autoscaling.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (BatchDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*batchDeploymentArgs)(nil)).Elem()
}

type BatchDeploymentInput interface {
	pulumi.Input

	ToBatchDeploymentOutput() BatchDeploymentOutput
	ToBatchDeploymentOutputWithContext(ctx context.Context) BatchDeploymentOutput
}

func (*BatchDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchDeployment)(nil)).Elem()
}

func (i *BatchDeployment) ToBatchDeploymentOutput() BatchDeploymentOutput {
	return i.ToBatchDeploymentOutputWithContext(context.Background())
}

func (i *BatchDeployment) ToBatchDeploymentOutputWithContext(ctx context.Context) BatchDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchDeploymentOutput)
}

func (i *BatchDeployment) ToOutput(ctx context.Context) pulumix.Output[*BatchDeployment] {
	return pulumix.Output[*BatchDeployment]{
		OutputState: i.ToBatchDeploymentOutputWithContext(ctx).OutputState,
	}
}

type BatchDeploymentOutput struct{ *pulumi.OutputState }

func (BatchDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchDeployment)(nil)).Elem()
}

func (o BatchDeploymentOutput) ToBatchDeploymentOutput() BatchDeploymentOutput {
	return o
}

func (o BatchDeploymentOutput) ToBatchDeploymentOutputWithContext(ctx context.Context) BatchDeploymentOutput {
	return o
}

func (o BatchDeploymentOutput) ToOutput(ctx context.Context) pulumix.Output[*BatchDeployment] {
	return pulumix.Output[*BatchDeployment]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o BatchDeploymentOutput) BatchDeploymentDetails() BatchDeploymentResponseOutput {
	return o.ApplyT(func(v *BatchDeployment) BatchDeploymentResponseOutput { return v.BatchDeploymentDetails }).(BatchDeploymentResponseOutput)
}

// Managed service identity (system assigned and/or user assigned identities)
func (o BatchDeploymentOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *BatchDeployment) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
func (o BatchDeploymentOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchDeployment) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o BatchDeploymentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchDeployment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o BatchDeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchDeployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Sku details required for ARM contract for Autoscaling.
func (o BatchDeploymentOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *BatchDeployment) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o BatchDeploymentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BatchDeployment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o BatchDeploymentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BatchDeployment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BatchDeploymentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchDeployment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BatchDeploymentOutput{})
}
