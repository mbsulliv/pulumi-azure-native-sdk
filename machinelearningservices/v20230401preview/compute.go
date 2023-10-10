// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Machine Learning compute object wrapped into ARM resource envelope.
type Compute struct {
	pulumi.CustomResourceState

	// The identity of the resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Specifies the location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Compute properties
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The sku of the workspace.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCompute registers a new resource with the given unique name, arguments, and options.
func NewCompute(ctx *pulumi.Context,
	name string, args *ComputeArgs, opts ...pulumi.ResourceOption) (*Compute, error) {
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
			Type: pulumi.String("azure-native:machinelearningservices:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20180301preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20181119:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20190501:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20190601:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20191101:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200101:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200218preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200301:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200401:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200515preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200601:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200801:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210101:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210701:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220101preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:Compute"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Compute
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230401preview:Compute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCompute gets an existing Compute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCompute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeState, opts ...pulumi.ResourceOption) (*Compute, error) {
	var resource Compute
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230401preview:Compute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Compute resources.
type computeState struct {
}

type ComputeState struct {
}

func (ComputeState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeState)(nil)).Elem()
}

type computeArgs struct {
	// Name of the Azure Machine Learning compute.
	ComputeName *string `pulumi:"computeName"`
	// The identity of the resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Compute properties
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku of the workspace.
	Sku *Sku `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Compute resource.
type ComputeArgs struct {
	// Name of the Azure Machine Learning compute.
	ComputeName pulumi.StringPtrInput
	// The identity of the resource.
	Identity ManagedServiceIdentityPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Compute properties
	Properties pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The sku of the workspace.
	Sku SkuPtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (ComputeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeArgs)(nil)).Elem()
}

type ComputeInput interface {
	pulumi.Input

	ToComputeOutput() ComputeOutput
	ToComputeOutputWithContext(ctx context.Context) ComputeOutput
}

func (*Compute) ElementType() reflect.Type {
	return reflect.TypeOf((**Compute)(nil)).Elem()
}

func (i *Compute) ToComputeOutput() ComputeOutput {
	return i.ToComputeOutputWithContext(context.Background())
}

func (i *Compute) ToComputeOutputWithContext(ctx context.Context) ComputeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeOutput)
}

func (i *Compute) ToOutput(ctx context.Context) pulumix.Output[*Compute] {
	return pulumix.Output[*Compute]{
		OutputState: i.ToComputeOutputWithContext(ctx).OutputState,
	}
}

type ComputeOutput struct{ *pulumi.OutputState }

func (ComputeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Compute)(nil)).Elem()
}

func (o ComputeOutput) ToComputeOutput() ComputeOutput {
	return o
}

func (o ComputeOutput) ToComputeOutputWithContext(ctx context.Context) ComputeOutput {
	return o
}

func (o ComputeOutput) ToOutput(ctx context.Context) pulumix.Output[*Compute] {
	return pulumix.Output[*Compute]{
		OutputState: o.OutputState,
	}
}

// The identity of the resource.
func (o ComputeOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Compute) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Specifies the location of the resource.
func (o ComputeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ComputeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Compute properties
func (o ComputeOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Compute) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// The sku of the workspace.
func (o ComputeOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Compute) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ComputeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Compute) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Contains resource tags defined as key/value pairs.
func (o ComputeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ComputeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ComputeOutput{})
}
