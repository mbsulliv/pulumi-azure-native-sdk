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

// A DDoS protection plan in a resource group.
type DdosProtectionPlan struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the DDoS protection plan resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The list of public IPs associated with the DDoS protection plan resource. This list is read-only.
	PublicIPAddresses SubResourceResponseArrayOutput `pulumi:"publicIPAddresses"`
	// The resource GUID property of the DDoS protection plan resource. It uniquely identifies the resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of virtual networks associated with the DDoS protection plan resource. This list is read-only.
	VirtualNetworks SubResourceResponseArrayOutput `pulumi:"virtualNetworks"`
}

// NewDdosProtectionPlan registers a new resource with the given unique name, arguments, and options.
func NewDdosProtectionPlan(ctx *pulumi.Context,
	name string, args *DdosProtectionPlanArgs, opts ...pulumi.ResourceOption) (*DdosProtectionPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:DdosProtectionPlan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:DdosProtectionPlan"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DdosProtectionPlan
	err := ctx.RegisterResource("azure-native:network/v20230401:DdosProtectionPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDdosProtectionPlan gets an existing DdosProtectionPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDdosProtectionPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdosProtectionPlanState, opts ...pulumi.ResourceOption) (*DdosProtectionPlan, error) {
	var resource DdosProtectionPlan
	err := ctx.ReadResource("azure-native:network/v20230401:DdosProtectionPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DdosProtectionPlan resources.
type ddosProtectionPlanState struct {
}

type DdosProtectionPlanState struct {
}

func (DdosProtectionPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosProtectionPlanState)(nil)).Elem()
}

type ddosProtectionPlanArgs struct {
	// The name of the DDoS protection plan.
	DdosProtectionPlanName *string `pulumi:"ddosProtectionPlanName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DdosProtectionPlan resource.
type DdosProtectionPlanArgs struct {
	// The name of the DDoS protection plan.
	DdosProtectionPlanName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DdosProtectionPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosProtectionPlanArgs)(nil)).Elem()
}

type DdosProtectionPlanInput interface {
	pulumi.Input

	ToDdosProtectionPlanOutput() DdosProtectionPlanOutput
	ToDdosProtectionPlanOutputWithContext(ctx context.Context) DdosProtectionPlanOutput
}

func (*DdosProtectionPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosProtectionPlan)(nil)).Elem()
}

func (i *DdosProtectionPlan) ToDdosProtectionPlanOutput() DdosProtectionPlanOutput {
	return i.ToDdosProtectionPlanOutputWithContext(context.Background())
}

func (i *DdosProtectionPlan) ToDdosProtectionPlanOutputWithContext(ctx context.Context) DdosProtectionPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosProtectionPlanOutput)
}

func (i *DdosProtectionPlan) ToOutput(ctx context.Context) pulumix.Output[*DdosProtectionPlan] {
	return pulumix.Output[*DdosProtectionPlan]{
		OutputState: i.ToDdosProtectionPlanOutputWithContext(ctx).OutputState,
	}
}

type DdosProtectionPlanOutput struct{ *pulumi.OutputState }

func (DdosProtectionPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosProtectionPlan)(nil)).Elem()
}

func (o DdosProtectionPlanOutput) ToDdosProtectionPlanOutput() DdosProtectionPlanOutput {
	return o
}

func (o DdosProtectionPlanOutput) ToDdosProtectionPlanOutputWithContext(ctx context.Context) DdosProtectionPlanOutput {
	return o
}

func (o DdosProtectionPlanOutput) ToOutput(ctx context.Context) pulumix.Output[*DdosProtectionPlan] {
	return pulumix.Output[*DdosProtectionPlan]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o DdosProtectionPlanOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosProtectionPlan) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource location.
func (o DdosProtectionPlanOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdosProtectionPlan) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o DdosProtectionPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosProtectionPlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the DDoS protection plan resource.
func (o DdosProtectionPlanOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosProtectionPlan) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The list of public IPs associated with the DDoS protection plan resource. This list is read-only.
func (o DdosProtectionPlanOutput) PublicIPAddresses() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *DdosProtectionPlan) SubResourceResponseArrayOutput { return v.PublicIPAddresses }).(SubResourceResponseArrayOutput)
}

// The resource GUID property of the DDoS protection plan resource. It uniquely identifies the resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
func (o DdosProtectionPlanOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosProtectionPlan) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Resource tags.
func (o DdosProtectionPlanOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DdosProtectionPlan) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o DdosProtectionPlanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosProtectionPlan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The list of virtual networks associated with the DDoS protection plan resource. This list is read-only.
func (o DdosProtectionPlanOutput) VirtualNetworks() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *DdosProtectionPlan) SubResourceResponseArrayOutput { return v.VirtualNetworks }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DdosProtectionPlanOutput{})
}
