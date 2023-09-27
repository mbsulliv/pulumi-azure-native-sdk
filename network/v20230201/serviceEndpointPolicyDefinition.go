// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Service Endpoint policy definitions.
type ServiceEndpointPolicyDefinition struct {
	pulumi.CustomResourceState

	// A description for this rule. Restricted to 140 chars.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning state of the service endpoint policy definition resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Service endpoint name.
	Service pulumi.StringPtrOutput `pulumi:"service"`
	// A list of service resources.
	ServiceResources pulumi.StringArrayOutput `pulumi:"serviceResources"`
	// The type of the resource.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewServiceEndpointPolicyDefinition registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointPolicyDefinition(ctx *pulumi.Context,
	name string, args *ServiceEndpointPolicyDefinitionArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointPolicyDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceEndpointPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceEndpointPolicyName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:ServiceEndpointPolicyDefinition"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:ServiceEndpointPolicyDefinition"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServiceEndpointPolicyDefinition
	err := ctx.RegisterResource("azure-native:network/v20230201:ServiceEndpointPolicyDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointPolicyDefinition gets an existing ServiceEndpointPolicyDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointPolicyDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointPolicyDefinitionState, opts ...pulumi.ResourceOption) (*ServiceEndpointPolicyDefinition, error) {
	var resource ServiceEndpointPolicyDefinition
	err := ctx.ReadResource("azure-native:network/v20230201:ServiceEndpointPolicyDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointPolicyDefinition resources.
type serviceEndpointPolicyDefinitionState struct {
}

type ServiceEndpointPolicyDefinitionState struct {
}

func (ServiceEndpointPolicyDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPolicyDefinitionState)(nil)).Elem()
}

type serviceEndpointPolicyDefinitionArgs struct {
	// A description for this rule. Restricted to 140 chars.
	Description *string `pulumi:"description"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Service endpoint name.
	Service *string `pulumi:"service"`
	// The name of the service endpoint policy definition name.
	ServiceEndpointPolicyDefinitionName *string `pulumi:"serviceEndpointPolicyDefinitionName"`
	// The name of the service endpoint policy.
	ServiceEndpointPolicyName string `pulumi:"serviceEndpointPolicyName"`
	// A list of service resources.
	ServiceResources []string `pulumi:"serviceResources"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a ServiceEndpointPolicyDefinition resource.
type ServiceEndpointPolicyDefinitionArgs struct {
	// A description for this rule. Restricted to 140 chars.
	Description pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Service endpoint name.
	Service pulumi.StringPtrInput
	// The name of the service endpoint policy definition name.
	ServiceEndpointPolicyDefinitionName pulumi.StringPtrInput
	// The name of the service endpoint policy.
	ServiceEndpointPolicyName pulumi.StringInput
	// A list of service resources.
	ServiceResources pulumi.StringArrayInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (ServiceEndpointPolicyDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPolicyDefinitionArgs)(nil)).Elem()
}

type ServiceEndpointPolicyDefinitionInput interface {
	pulumi.Input

	ToServiceEndpointPolicyDefinitionOutput() ServiceEndpointPolicyDefinitionOutput
	ToServiceEndpointPolicyDefinitionOutputWithContext(ctx context.Context) ServiceEndpointPolicyDefinitionOutput
}

func (*ServiceEndpointPolicyDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointPolicyDefinition)(nil)).Elem()
}

func (i *ServiceEndpointPolicyDefinition) ToServiceEndpointPolicyDefinitionOutput() ServiceEndpointPolicyDefinitionOutput {
	return i.ToServiceEndpointPolicyDefinitionOutputWithContext(context.Background())
}

func (i *ServiceEndpointPolicyDefinition) ToServiceEndpointPolicyDefinitionOutputWithContext(ctx context.Context) ServiceEndpointPolicyDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPolicyDefinitionOutput)
}

func (i *ServiceEndpointPolicyDefinition) ToOutput(ctx context.Context) pulumix.Output[*ServiceEndpointPolicyDefinition] {
	return pulumix.Output[*ServiceEndpointPolicyDefinition]{
		OutputState: i.ToServiceEndpointPolicyDefinitionOutputWithContext(ctx).OutputState,
	}
}

type ServiceEndpointPolicyDefinitionOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPolicyDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointPolicyDefinition)(nil)).Elem()
}

func (o ServiceEndpointPolicyDefinitionOutput) ToServiceEndpointPolicyDefinitionOutput() ServiceEndpointPolicyDefinitionOutput {
	return o
}

func (o ServiceEndpointPolicyDefinitionOutput) ToServiceEndpointPolicyDefinitionOutputWithContext(ctx context.Context) ServiceEndpointPolicyDefinitionOutput {
	return o
}

func (o ServiceEndpointPolicyDefinitionOutput) ToOutput(ctx context.Context) pulumix.Output[*ServiceEndpointPolicyDefinition] {
	return pulumix.Output[*ServiceEndpointPolicyDefinition]{
		OutputState: o.OutputState,
	}
}

// A description for this rule. Restricted to 140 chars.
func (o ServiceEndpointPolicyDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicyDefinition) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o ServiceEndpointPolicyDefinitionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicyDefinition) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o ServiceEndpointPolicyDefinitionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicyDefinition) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the service endpoint policy definition resource.
func (o ServiceEndpointPolicyDefinitionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicyDefinition) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Service endpoint name.
func (o ServiceEndpointPolicyDefinitionOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicyDefinition) pulumi.StringPtrOutput { return v.Service }).(pulumi.StringPtrOutput)
}

// A list of service resources.
func (o ServiceEndpointPolicyDefinitionOutput) ServiceResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicyDefinition) pulumi.StringArrayOutput { return v.ServiceResources }).(pulumi.StringArrayOutput)
}

// The type of the resource.
func (o ServiceEndpointPolicyDefinitionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicyDefinition) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointPolicyDefinitionOutput{})
}
