// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Service Registry resource
type ServiceRegistry struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Service Registry properties payload
	Properties ServiceRegistryPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceRegistry registers a new resource with the given unique name, arguments, and options.
func NewServiceRegistry(ctx *pulumi.Context,
	name string, args *ServiceRegistryArgs, opts ...pulumi.ResourceOption) (*ServiceRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:ServiceRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:ServiceRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:ServiceRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:ServiceRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:ServiceRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:ServiceRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:ServiceRegistry"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:ServiceRegistry"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceRegistry
	err := ctx.RegisterResource("azure-native:appplatform/v20220401:ServiceRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceRegistry gets an existing ServiceRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceRegistryState, opts ...pulumi.ResourceOption) (*ServiceRegistry, error) {
	var resource ServiceRegistry
	err := ctx.ReadResource("azure-native:appplatform/v20220401:ServiceRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceRegistry resources.
type serviceRegistryState struct {
}

type ServiceRegistryState struct {
}

func (ServiceRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceRegistryState)(nil)).Elem()
}

type serviceRegistryArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
	// The name of Service Registry.
	ServiceRegistryName *string `pulumi:"serviceRegistryName"`
}

// The set of arguments for constructing a ServiceRegistry resource.
type ServiceRegistryArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
	// The name of Service Registry.
	ServiceRegistryName pulumi.StringPtrInput
}

func (ServiceRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceRegistryArgs)(nil)).Elem()
}

type ServiceRegistryInput interface {
	pulumi.Input

	ToServiceRegistryOutput() ServiceRegistryOutput
	ToServiceRegistryOutputWithContext(ctx context.Context) ServiceRegistryOutput
}

func (*ServiceRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceRegistry)(nil)).Elem()
}

func (i *ServiceRegistry) ToServiceRegistryOutput() ServiceRegistryOutput {
	return i.ToServiceRegistryOutputWithContext(context.Background())
}

func (i *ServiceRegistry) ToServiceRegistryOutputWithContext(ctx context.Context) ServiceRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceRegistryOutput)
}

type ServiceRegistryOutput struct{ *pulumi.OutputState }

func (ServiceRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceRegistry)(nil)).Elem()
}

func (o ServiceRegistryOutput) ToServiceRegistryOutput() ServiceRegistryOutput {
	return o
}

func (o ServiceRegistryOutput) ToServiceRegistryOutputWithContext(ctx context.Context) ServiceRegistryOutput {
	return o
}

// The name of the resource.
func (o ServiceRegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceRegistry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Service Registry properties payload
func (o ServiceRegistryOutput) Properties() ServiceRegistryPropertiesResponseOutput {
	return o.ApplyT(func(v *ServiceRegistry) ServiceRegistryPropertiesResponseOutput { return v.Properties }).(ServiceRegistryPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ServiceRegistryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ServiceRegistry) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o ServiceRegistryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceRegistry) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceRegistryOutput{})
}
