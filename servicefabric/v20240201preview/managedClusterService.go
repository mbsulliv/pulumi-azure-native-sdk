// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The service resource.
type ManagedClusterService struct {
	pulumi.CustomResourceState

	// Resource location depends on the parent resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The service resource properties.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedClusterService registers a new resource with the given unique name, arguments, and options.
func NewManagedClusterService(ctx *pulumi.Context,
	name string, args *ManagedClusterServiceArgs, opts ...pulumi.ResourceOption) (*ManagedClusterService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210101preview:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210501:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210701preview:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210901privatepreview:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20211101preview:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220101:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220201preview:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220601preview:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220801preview:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20221001preview:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230201preview:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230301preview:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230701preview:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230901preview:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20231101preview:ManagedClusterService"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20231201preview:ManagedClusterService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedClusterService
	err := ctx.RegisterResource("azure-native:servicefabric/v20240201preview:ManagedClusterService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedClusterService gets an existing ManagedClusterService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedClusterService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterServiceState, opts ...pulumi.ResourceOption) (*ManagedClusterService, error) {
	var resource ManagedClusterService
	err := ctx.ReadResource("azure-native:servicefabric/v20240201preview:ManagedClusterService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedClusterService resources.
type managedClusterServiceState struct {
}

type ManagedClusterServiceState struct {
}

func (ManagedClusterServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterServiceState)(nil)).Elem()
}

type managedClusterServiceArgs struct {
	// The name of the application resource.
	ApplicationName string `pulumi:"applicationName"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// Resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// The service resource properties.
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service resource in the format of {applicationName}~{serviceName}.
	ServiceName *string `pulumi:"serviceName"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ManagedClusterService resource.
type ManagedClusterServiceArgs struct {
	// The name of the application resource.
	ApplicationName pulumi.StringInput
	// The name of the cluster resource.
	ClusterName pulumi.StringInput
	// Resource location depends on the parent resource.
	Location pulumi.StringPtrInput
	// The service resource properties.
	Properties pulumi.Input
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the service resource in the format of {applicationName}~{serviceName}.
	ServiceName pulumi.StringPtrInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
}

func (ManagedClusterServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterServiceArgs)(nil)).Elem()
}

type ManagedClusterServiceInput interface {
	pulumi.Input

	ToManagedClusterServiceOutput() ManagedClusterServiceOutput
	ToManagedClusterServiceOutputWithContext(ctx context.Context) ManagedClusterServiceOutput
}

func (*ManagedClusterService) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterService)(nil)).Elem()
}

func (i *ManagedClusterService) ToManagedClusterServiceOutput() ManagedClusterServiceOutput {
	return i.ToManagedClusterServiceOutputWithContext(context.Background())
}

func (i *ManagedClusterService) ToManagedClusterServiceOutputWithContext(ctx context.Context) ManagedClusterServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServiceOutput)
}

type ManagedClusterServiceOutput struct{ *pulumi.OutputState }

func (ManagedClusterServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterService)(nil)).Elem()
}

func (o ManagedClusterServiceOutput) ToManagedClusterServiceOutput() ManagedClusterServiceOutput {
	return o
}

func (o ManagedClusterServiceOutput) ToManagedClusterServiceOutputWithContext(ctx context.Context) ManagedClusterServiceOutput {
	return o
}

// Resource location depends on the parent resource.
func (o ManagedClusterServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Azure resource name.
func (o ManagedClusterServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The service resource properties.
func (o ManagedClusterServiceOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ManagedClusterService) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ManagedClusterServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedClusterService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o ManagedClusterServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedClusterService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o ManagedClusterServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedClusterServiceOutput{})
}
