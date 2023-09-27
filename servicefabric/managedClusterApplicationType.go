// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabric

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The application type name resource
// Azure REST API version: 2023-03-01-preview.
type ManagedClusterApplicationType struct {
	pulumi.CustomResourceState

	// Resource location depends on the parent resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedClusterApplicationType registers a new resource with the given unique name, arguments, and options.
func NewManagedClusterApplicationType(ctx *pulumi.Context,
	name string, args *ManagedClusterApplicationTypeArgs, opts ...pulumi.ResourceOption) (*ManagedClusterApplicationType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric/v20210101preview:ManagedClusterApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210501:ManagedClusterApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210701preview:ManagedClusterApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210901privatepreview:ManagedClusterApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20211101preview:ManagedClusterApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220101:ManagedClusterApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220201preview:ManagedClusterApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220601preview:ManagedClusterApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220801preview:ManagedClusterApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20221001preview:ManagedClusterApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230201preview:ManagedClusterApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230301preview:ManagedClusterApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230701preview:ManagedClusterApplicationType"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedClusterApplicationType
	err := ctx.RegisterResource("azure-native:servicefabric:ManagedClusterApplicationType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedClusterApplicationType gets an existing ManagedClusterApplicationType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedClusterApplicationType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterApplicationTypeState, opts ...pulumi.ResourceOption) (*ManagedClusterApplicationType, error) {
	var resource ManagedClusterApplicationType
	err := ctx.ReadResource("azure-native:servicefabric:ManagedClusterApplicationType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedClusterApplicationType resources.
type managedClusterApplicationTypeState struct {
}

type ManagedClusterApplicationTypeState struct {
}

func (ManagedClusterApplicationTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterApplicationTypeState)(nil)).Elem()
}

type managedClusterApplicationTypeArgs struct {
	// The name of the application type name resource.
	ApplicationTypeName *string `pulumi:"applicationTypeName"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// Resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ManagedClusterApplicationType resource.
type ManagedClusterApplicationTypeArgs struct {
	// The name of the application type name resource.
	ApplicationTypeName pulumi.StringPtrInput
	// The name of the cluster resource.
	ClusterName pulumi.StringInput
	// Resource location depends on the parent resource.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
}

func (ManagedClusterApplicationTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterApplicationTypeArgs)(nil)).Elem()
}

type ManagedClusterApplicationTypeInput interface {
	pulumi.Input

	ToManagedClusterApplicationTypeOutput() ManagedClusterApplicationTypeOutput
	ToManagedClusterApplicationTypeOutputWithContext(ctx context.Context) ManagedClusterApplicationTypeOutput
}

func (*ManagedClusterApplicationType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterApplicationType)(nil)).Elem()
}

func (i *ManagedClusterApplicationType) ToManagedClusterApplicationTypeOutput() ManagedClusterApplicationTypeOutput {
	return i.ToManagedClusterApplicationTypeOutputWithContext(context.Background())
}

func (i *ManagedClusterApplicationType) ToManagedClusterApplicationTypeOutputWithContext(ctx context.Context) ManagedClusterApplicationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterApplicationTypeOutput)
}

func (i *ManagedClusterApplicationType) ToOutput(ctx context.Context) pulumix.Output[*ManagedClusterApplicationType] {
	return pulumix.Output[*ManagedClusterApplicationType]{
		OutputState: i.ToManagedClusterApplicationTypeOutputWithContext(ctx).OutputState,
	}
}

type ManagedClusterApplicationTypeOutput struct{ *pulumi.OutputState }

func (ManagedClusterApplicationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterApplicationType)(nil)).Elem()
}

func (o ManagedClusterApplicationTypeOutput) ToManagedClusterApplicationTypeOutput() ManagedClusterApplicationTypeOutput {
	return o
}

func (o ManagedClusterApplicationTypeOutput) ToManagedClusterApplicationTypeOutputWithContext(ctx context.Context) ManagedClusterApplicationTypeOutput {
	return o
}

func (o ManagedClusterApplicationTypeOutput) ToOutput(ctx context.Context) pulumix.Output[*ManagedClusterApplicationType] {
	return pulumix.Output[*ManagedClusterApplicationType]{
		OutputState: o.OutputState,
	}
}

// Resource location depends on the parent resource.
func (o ManagedClusterApplicationTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationType) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Azure resource name.
func (o ManagedClusterApplicationTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current deployment or provisioning state, which only appears in the response.
func (o ManagedClusterApplicationTypeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationType) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ManagedClusterApplicationTypeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationType) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o ManagedClusterApplicationTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationType) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o ManagedClusterApplicationTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationType) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedClusterApplicationTypeOutput{})
}
