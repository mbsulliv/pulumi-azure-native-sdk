// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type BatchEndpoint struct {
	pulumi.CustomResourceState

	// [Required] Additional attributes of the entity.
	BatchEndpointProperties BatchEndpointResponseOutput `pulumi:"batchEndpointProperties"`
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

// NewBatchEndpoint registers a new resource with the given unique name, arguments, and options.
func NewBatchEndpoint(ctx *pulumi.Context,
	name string, args *BatchEndpointArgs, opts ...pulumi.ResourceOption) (*BatchEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BatchEndpointProperties == nil {
		return nil, errors.New("invalid value for required argument 'BatchEndpointProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:BatchEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BatchEndpoint
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230401:BatchEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBatchEndpoint gets an existing BatchEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBatchEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BatchEndpointState, opts ...pulumi.ResourceOption) (*BatchEndpoint, error) {
	var resource BatchEndpoint
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230401:BatchEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BatchEndpoint resources.
type batchEndpointState struct {
}

type BatchEndpointState struct {
}

func (BatchEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*batchEndpointState)(nil)).Elem()
}

type batchEndpointArgs struct {
	// [Required] Additional attributes of the entity.
	BatchEndpointProperties BatchEndpointType `pulumi:"batchEndpointProperties"`
	// Name for the Batch inference endpoint.
	EndpointName *string `pulumi:"endpointName"`
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

// The set of arguments for constructing a BatchEndpoint resource.
type BatchEndpointArgs struct {
	// [Required] Additional attributes of the entity.
	BatchEndpointProperties BatchEndpointTypeInput
	// Name for the Batch inference endpoint.
	EndpointName pulumi.StringPtrInput
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

func (BatchEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*batchEndpointArgs)(nil)).Elem()
}

type BatchEndpointInput interface {
	pulumi.Input

	ToBatchEndpointOutput() BatchEndpointOutput
	ToBatchEndpointOutputWithContext(ctx context.Context) BatchEndpointOutput
}

func (*BatchEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchEndpoint)(nil)).Elem()
}

func (i *BatchEndpoint) ToBatchEndpointOutput() BatchEndpointOutput {
	return i.ToBatchEndpointOutputWithContext(context.Background())
}

func (i *BatchEndpoint) ToBatchEndpointOutputWithContext(ctx context.Context) BatchEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchEndpointOutput)
}

func (i *BatchEndpoint) ToOutput(ctx context.Context) pulumix.Output[*BatchEndpoint] {
	return pulumix.Output[*BatchEndpoint]{
		OutputState: i.ToBatchEndpointOutputWithContext(ctx).OutputState,
	}
}

type BatchEndpointOutput struct{ *pulumi.OutputState }

func (BatchEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchEndpoint)(nil)).Elem()
}

func (o BatchEndpointOutput) ToBatchEndpointOutput() BatchEndpointOutput {
	return o
}

func (o BatchEndpointOutput) ToBatchEndpointOutputWithContext(ctx context.Context) BatchEndpointOutput {
	return o
}

func (o BatchEndpointOutput) ToOutput(ctx context.Context) pulumix.Output[*BatchEndpoint] {
	return pulumix.Output[*BatchEndpoint]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o BatchEndpointOutput) BatchEndpointProperties() BatchEndpointResponseOutput {
	return o.ApplyT(func(v *BatchEndpoint) BatchEndpointResponseOutput { return v.BatchEndpointProperties }).(BatchEndpointResponseOutput)
}

// Managed service identity (system assigned and/or user assigned identities)
func (o BatchEndpointOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *BatchEndpoint) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
func (o BatchEndpointOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchEndpoint) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o BatchEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o BatchEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Sku details required for ARM contract for Autoscaling.
func (o BatchEndpointOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *BatchEndpoint) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o BatchEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BatchEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o BatchEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BatchEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BatchEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BatchEndpointOutput{})
}
