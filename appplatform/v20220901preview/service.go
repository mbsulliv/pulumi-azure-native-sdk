// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Service resource
type Service struct {
	pulumi.CustomResourceState

	// The GEO location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the Service resource
	Properties ClusterResourcePropertiesResponseOutput `pulumi:"properties"`
	// Sku of the Service resource
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToClusterResourcePropertiesPtrOutput().ApplyT(func(v *ClusterResourceProperties) *ClusterResourceProperties { return v.Defaults() }).(ClusterResourcePropertiesPtrOutput)
	}
	if args.Sku != nil {
		args.Sku = args.Sku.ToSkuPtrOutput().ApplyT(func(v *Sku) *Sku { return v.Defaults() }).(SkuPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20200701:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210601preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:Service"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azure-native:appplatform/v20220901preview:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:appplatform/v20220901preview:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// The GEO location of the resource.
	Location *string `pulumi:"location"`
	// Properties of the Service resource
	Properties *ClusterResourceProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName *string `pulumi:"serviceName"`
	// Sku of the Service resource
	Sku *Sku `pulumi:"sku"`
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The GEO location of the resource.
	Location pulumi.StringPtrInput
	// Properties of the Service resource
	Properties ClusterResourcePropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringPtrInput
	// Sku of the Service resource
	Sku SkuPtrInput
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// The GEO location of the resource.
func (o ServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the Service resource
func (o ServiceOutput) Properties() ClusterResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *Service) ClusterResourcePropertiesResponseOutput { return v.Properties }).(ClusterResourcePropertiesResponseOutput)
}

// Sku of the Service resource
func (o ServiceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Service) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Service) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags of the service which is a list of key value pairs that describe the resource.
func (o ServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o ServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
