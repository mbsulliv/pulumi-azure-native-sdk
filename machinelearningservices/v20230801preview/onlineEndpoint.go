// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type OnlineEndpoint struct {
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
	OnlineEndpointProperties OnlineEndpointResponseOutput `pulumi:"onlineEndpointProperties"`
	// Sku details required for ARM contract for Autoscaling.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOnlineEndpoint registers a new resource with the given unique name, arguments, and options.
func NewOnlineEndpoint(ctx *pulumi.Context,
	name string, args *OnlineEndpointArgs, opts ...pulumi.ResourceOption) (*OnlineEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OnlineEndpointProperties == nil {
		return nil, errors.New("invalid value for required argument 'OnlineEndpointProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.OnlineEndpointProperties = args.OnlineEndpointProperties.ToOnlineEndpointTypeOutput().ApplyT(func(v OnlineEndpointType) OnlineEndpointType { return *v.Defaults() }).(OnlineEndpointTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:OnlineEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource OnlineEndpoint
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20230801preview:OnlineEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOnlineEndpoint gets an existing OnlineEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOnlineEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OnlineEndpointState, opts ...pulumi.ResourceOption) (*OnlineEndpoint, error) {
	var resource OnlineEndpoint
	err := ctx.ReadResource("azure-native:machinelearningservices/v20230801preview:OnlineEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OnlineEndpoint resources.
type onlineEndpointState struct {
}

type OnlineEndpointState struct {
}

func (OnlineEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*onlineEndpointState)(nil)).Elem()
}

type onlineEndpointArgs struct {
	// Online Endpoint name.
	EndpointName *string `pulumi:"endpointName"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// [Required] Additional attributes of the entity.
	OnlineEndpointProperties OnlineEndpointType `pulumi:"onlineEndpointProperties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sku details required for ARM contract for Autoscaling.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a OnlineEndpoint resource.
type OnlineEndpointArgs struct {
	// Online Endpoint name.
	EndpointName pulumi.StringPtrInput
	// Managed service identity (system assigned and/or user assigned identities)
	Identity ManagedServiceIdentityPtrInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// [Required] Additional attributes of the entity.
	OnlineEndpointProperties OnlineEndpointTypeInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Sku details required for ARM contract for Autoscaling.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (OnlineEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*onlineEndpointArgs)(nil)).Elem()
}

type OnlineEndpointInput interface {
	pulumi.Input

	ToOnlineEndpointOutput() OnlineEndpointOutput
	ToOnlineEndpointOutputWithContext(ctx context.Context) OnlineEndpointOutput
}

func (*OnlineEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**OnlineEndpoint)(nil)).Elem()
}

func (i *OnlineEndpoint) ToOnlineEndpointOutput() OnlineEndpointOutput {
	return i.ToOnlineEndpointOutputWithContext(context.Background())
}

func (i *OnlineEndpoint) ToOnlineEndpointOutputWithContext(ctx context.Context) OnlineEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnlineEndpointOutput)
}

func (i *OnlineEndpoint) ToOutput(ctx context.Context) pulumix.Output[*OnlineEndpoint] {
	return pulumix.Output[*OnlineEndpoint]{
		OutputState: i.ToOnlineEndpointOutputWithContext(ctx).OutputState,
	}
}

type OnlineEndpointOutput struct{ *pulumi.OutputState }

func (OnlineEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnlineEndpoint)(nil)).Elem()
}

func (o OnlineEndpointOutput) ToOnlineEndpointOutput() OnlineEndpointOutput {
	return o
}

func (o OnlineEndpointOutput) ToOnlineEndpointOutputWithContext(ctx context.Context) OnlineEndpointOutput {
	return o
}

func (o OnlineEndpointOutput) ToOutput(ctx context.Context) pulumix.Output[*OnlineEndpoint] {
	return pulumix.Output[*OnlineEndpoint]{
		OutputState: o.OutputState,
	}
}

// Managed service identity (system assigned and/or user assigned identities)
func (o OnlineEndpointOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *OnlineEndpoint) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
func (o OnlineEndpointOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnlineEndpoint) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o OnlineEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o OnlineEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o OnlineEndpointOutput) OnlineEndpointProperties() OnlineEndpointResponseOutput {
	return o.ApplyT(func(v *OnlineEndpoint) OnlineEndpointResponseOutput { return v.OnlineEndpointProperties }).(OnlineEndpointResponseOutput)
}

// Sku details required for ARM contract for Autoscaling.
func (o OnlineEndpointOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *OnlineEndpoint) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o OnlineEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OnlineEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o OnlineEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OnlineEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o OnlineEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OnlineEndpointOutput{})
}
