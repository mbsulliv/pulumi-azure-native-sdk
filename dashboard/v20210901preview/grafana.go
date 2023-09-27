// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The grafana resource type.
type Grafana struct {
	pulumi.CustomResourceState

	// The managed identity of the grafana resource.
	Identity ManagedIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the grafana resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the grafana resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties specific to the grafana resource.
	Properties ManagedGrafanaPropertiesResponseOutput `pulumi:"properties"`
	// The Sku of the grafana resource.
	Sku ResourceSkuResponsePtrOutput `pulumi:"sku"`
	// The system meta data relating to this grafana resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tags for grafana resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the grafana resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGrafana registers a new resource with the given unique name, arguments, and options.
func NewGrafana(ctx *pulumi.Context,
	name string, args *GrafanaArgs, opts ...pulumi.ResourceOption) (*Grafana, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dashboard:Grafana"),
		},
		{
			Type: pulumi.String("azure-native:dashboard/v20220501preview:Grafana"),
		},
		{
			Type: pulumi.String("azure-native:dashboard/v20220801:Grafana"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Grafana
	err := ctx.RegisterResource("azure-native:dashboard/v20210901preview:Grafana", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGrafana gets an existing Grafana resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGrafana(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GrafanaState, opts ...pulumi.ResourceOption) (*Grafana, error) {
	var resource Grafana
	err := ctx.ReadResource("azure-native:dashboard/v20210901preview:Grafana", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Grafana resources.
type grafanaState struct {
}

type GrafanaState struct {
}

func (GrafanaState) ElementType() reflect.Type {
	return reflect.TypeOf((*grafanaState)(nil)).Elem()
}

type grafanaArgs struct {
	// The managed identity of the grafana resource.
	Identity *ManagedIdentity `pulumi:"identity"`
	// The geo-location where the grafana resource lives
	Location *string `pulumi:"location"`
	// Properties specific to the grafana resource.
	Properties *ManagedGrafanaProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Sku of the grafana resource.
	Sku *ResourceSku `pulumi:"sku"`
	// The tags for grafana resource.
	Tags map[string]string `pulumi:"tags"`
	// The workspace name of Azure Managed Grafana.
	WorkspaceName *string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Grafana resource.
type GrafanaArgs struct {
	// The managed identity of the grafana resource.
	Identity ManagedIdentityPtrInput
	// The geo-location where the grafana resource lives
	Location pulumi.StringPtrInput
	// Properties specific to the grafana resource.
	Properties ManagedGrafanaPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The Sku of the grafana resource.
	Sku ResourceSkuPtrInput
	// The tags for grafana resource.
	Tags pulumi.StringMapInput
	// The workspace name of Azure Managed Grafana.
	WorkspaceName pulumi.StringPtrInput
}

func (GrafanaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*grafanaArgs)(nil)).Elem()
}

type GrafanaInput interface {
	pulumi.Input

	ToGrafanaOutput() GrafanaOutput
	ToGrafanaOutputWithContext(ctx context.Context) GrafanaOutput
}

func (*Grafana) ElementType() reflect.Type {
	return reflect.TypeOf((**Grafana)(nil)).Elem()
}

func (i *Grafana) ToGrafanaOutput() GrafanaOutput {
	return i.ToGrafanaOutputWithContext(context.Background())
}

func (i *Grafana) ToGrafanaOutputWithContext(ctx context.Context) GrafanaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrafanaOutput)
}

func (i *Grafana) ToOutput(ctx context.Context) pulumix.Output[*Grafana] {
	return pulumix.Output[*Grafana]{
		OutputState: i.ToGrafanaOutputWithContext(ctx).OutputState,
	}
}

type GrafanaOutput struct{ *pulumi.OutputState }

func (GrafanaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Grafana)(nil)).Elem()
}

func (o GrafanaOutput) ToGrafanaOutput() GrafanaOutput {
	return o
}

func (o GrafanaOutput) ToGrafanaOutputWithContext(ctx context.Context) GrafanaOutput {
	return o
}

func (o GrafanaOutput) ToOutput(ctx context.Context) pulumix.Output[*Grafana] {
	return pulumix.Output[*Grafana]{
		OutputState: o.OutputState,
	}
}

// The managed identity of the grafana resource.
func (o GrafanaOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Grafana) ManagedIdentityResponsePtrOutput { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

// The geo-location where the grafana resource lives
func (o GrafanaOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Grafana) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the grafana resource.
func (o GrafanaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Grafana) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties specific to the grafana resource.
func (o GrafanaOutput) Properties() ManagedGrafanaPropertiesResponseOutput {
	return o.ApplyT(func(v *Grafana) ManagedGrafanaPropertiesResponseOutput { return v.Properties }).(ManagedGrafanaPropertiesResponseOutput)
}

// The Sku of the grafana resource.
func (o GrafanaOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v *Grafana) ResourceSkuResponsePtrOutput { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// The system meta data relating to this grafana resource.
func (o GrafanaOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Grafana) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tags for grafana resource.
func (o GrafanaOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Grafana) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the grafana resource.
func (o GrafanaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Grafana) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GrafanaOutput{})
}
