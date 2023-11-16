// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ServiceEndpoint resource details.
type ServiceEndpoint struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// ServiceEndpoint resource properties.
	Properties ServiceEndpointResourceResponsePropertiesOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceEndpoint registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpoint(ctx *pulumi.Context,
	name string, args *ServiceEndpointArgs, opts ...pulumi.ResourceOption) (*ServiceEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recommendationsservice:ServiceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:recommendationsservice/v20220301preview:ServiceEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServiceEndpoint
	err := ctx.RegisterResource("azure-native:recommendationsservice/v20220201:ServiceEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpoint gets an existing ServiceEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointState, opts ...pulumi.ResourceOption) (*ServiceEndpoint, error) {
	var resource ServiceEndpoint
	err := ctx.ReadResource("azure-native:recommendationsservice/v20220201:ServiceEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpoint resources.
type serviceEndpointState struct {
}

type ServiceEndpointState struct {
}

func (ServiceEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointState)(nil)).Elem()
}

type serviceEndpointArgs struct {
	// The name of the RecommendationsService Account resource.
	AccountName string `pulumi:"accountName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// ServiceEndpoint resource properties.
	Properties *ServiceEndpointResourceProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the ServiceEndpoint resource.
	ServiceEndpointName *string `pulumi:"serviceEndpointName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ServiceEndpoint resource.
type ServiceEndpointArgs struct {
	// The name of the RecommendationsService Account resource.
	AccountName pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// ServiceEndpoint resource properties.
	Properties ServiceEndpointResourcePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the ServiceEndpoint resource.
	ServiceEndpointName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ServiceEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointArgs)(nil)).Elem()
}

type ServiceEndpointInput interface {
	pulumi.Input

	ToServiceEndpointOutput() ServiceEndpointOutput
	ToServiceEndpointOutputWithContext(ctx context.Context) ServiceEndpointOutput
}

func (*ServiceEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpoint)(nil)).Elem()
}

func (i *ServiceEndpoint) ToServiceEndpointOutput() ServiceEndpointOutput {
	return i.ToServiceEndpointOutputWithContext(context.Background())
}

func (i *ServiceEndpoint) ToServiceEndpointOutputWithContext(ctx context.Context) ServiceEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointOutput)
}

type ServiceEndpointOutput struct{ *pulumi.OutputState }

func (ServiceEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpoint)(nil)).Elem()
}

func (o ServiceEndpointOutput) ToServiceEndpointOutput() ServiceEndpointOutput {
	return o
}

func (o ServiceEndpointOutput) ToServiceEndpointOutputWithContext(ctx context.Context) ServiceEndpointOutput {
	return o
}

// The geo-location where the resource lives
func (o ServiceEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ServiceEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ServiceEndpoint resource properties.
func (o ServiceEndpointOutput) Properties() ServiceEndpointResourceResponsePropertiesOutput {
	return o.ApplyT(func(v *ServiceEndpoint) ServiceEndpointResourceResponsePropertiesOutput { return v.Properties }).(ServiceEndpointResourceResponsePropertiesOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ServiceEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ServiceEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ServiceEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ServiceEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointOutput{})
}
