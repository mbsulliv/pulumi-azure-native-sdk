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

type OnlineDeployment struct {
	pulumi.CustomResourceState

	// Managed service identity (system assigned and/or user assigned identities)
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	OnlineDeploymentDetails pulumi.AnyOutput `pulumi:"onlineDeploymentDetails"`
	// Sku details required for ARM contract for Autoscaling.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOnlineDeployment registers a new resource with the given unique name, arguments, and options.
func NewOnlineDeployment(ctx *pulumi.Context,
	name string, args *OnlineDeploymentArgs, opts ...pulumi.ResourceOption) (*OnlineDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.OnlineDeploymentDetails == nil {
		return nil, errors.New("invalid value for required argument 'OnlineDeploymentDetails'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:OnlineDeployment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource OnlineDeployment
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220201preview:OnlineDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOnlineDeployment gets an existing OnlineDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOnlineDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OnlineDeploymentState, opts ...pulumi.ResourceOption) (*OnlineDeployment, error) {
	var resource OnlineDeployment
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220201preview:OnlineDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OnlineDeployment resources.
type onlineDeploymentState struct {
}

type OnlineDeploymentState struct {
}

func (OnlineDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*onlineDeploymentState)(nil)).Elem()
}

type onlineDeploymentArgs struct {
	// Inference Endpoint Deployment name.
	DeploymentName *string `pulumi:"deploymentName"`
	// Inference endpoint name.
	EndpointName string `pulumi:"endpointName"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// [Required] Additional attributes of the entity.
	OnlineDeploymentDetails interface{} `pulumi:"onlineDeploymentDetails"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sku details required for ARM contract for Autoscaling.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a OnlineDeployment resource.
type OnlineDeploymentArgs struct {
	// Inference Endpoint Deployment name.
	DeploymentName pulumi.StringPtrInput
	// Inference endpoint name.
	EndpointName pulumi.StringInput
	// Managed service identity (system assigned and/or user assigned identities)
	Identity ManagedServiceIdentityPtrInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// [Required] Additional attributes of the entity.
	OnlineDeploymentDetails pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Sku details required for ARM contract for Autoscaling.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (OnlineDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*onlineDeploymentArgs)(nil)).Elem()
}

type OnlineDeploymentInput interface {
	pulumi.Input

	ToOnlineDeploymentOutput() OnlineDeploymentOutput
	ToOnlineDeploymentOutputWithContext(ctx context.Context) OnlineDeploymentOutput
}

func (*OnlineDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**OnlineDeployment)(nil)).Elem()
}

func (i *OnlineDeployment) ToOnlineDeploymentOutput() OnlineDeploymentOutput {
	return i.ToOnlineDeploymentOutputWithContext(context.Background())
}

func (i *OnlineDeployment) ToOnlineDeploymentOutputWithContext(ctx context.Context) OnlineDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnlineDeploymentOutput)
}

func (i *OnlineDeployment) ToOutput(ctx context.Context) pulumix.Output[*OnlineDeployment] {
	return pulumix.Output[*OnlineDeployment]{
		OutputState: i.ToOnlineDeploymentOutputWithContext(ctx).OutputState,
	}
}

type OnlineDeploymentOutput struct{ *pulumi.OutputState }

func (OnlineDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnlineDeployment)(nil)).Elem()
}

func (o OnlineDeploymentOutput) ToOnlineDeploymentOutput() OnlineDeploymentOutput {
	return o
}

func (o OnlineDeploymentOutput) ToOnlineDeploymentOutputWithContext(ctx context.Context) OnlineDeploymentOutput {
	return o
}

func (o OnlineDeploymentOutput) ToOutput(ctx context.Context) pulumix.Output[*OnlineDeployment] {
	return pulumix.Output[*OnlineDeployment]{
		OutputState: o.OutputState,
	}
}

// Managed service identity (system assigned and/or user assigned identities)
func (o OnlineDeploymentOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *OnlineDeployment) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
func (o OnlineDeploymentOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnlineDeployment) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o OnlineDeploymentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineDeployment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o OnlineDeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineDeployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o OnlineDeploymentOutput) OnlineDeploymentDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *OnlineDeployment) pulumi.AnyOutput { return v.OnlineDeploymentDetails }).(pulumi.AnyOutput)
}

// Sku details required for ARM contract for Autoscaling.
func (o OnlineDeploymentOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *OnlineDeployment) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o OnlineDeploymentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OnlineDeployment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o OnlineDeploymentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OnlineDeployment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o OnlineDeploymentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineDeployment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OnlineDeploymentOutput{})
}
