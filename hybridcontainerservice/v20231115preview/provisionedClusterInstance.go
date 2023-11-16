// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231115preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provisionedClusterInstances resource definition.
type ProvisionedClusterInstance struct {
	pulumi.CustomResourceState

	// Extended Location definition
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// All properties of the provisioned cluster
	Properties ProvisionedClusterPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProvisionedClusterInstance registers a new resource with the given unique name, arguments, and options.
func NewProvisionedClusterInstance(ctx *pulumi.Context,
	name string, args *ProvisionedClusterInstanceArgs, opts ...pulumi.ResourceOption) (*ProvisionedClusterInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectedClusterResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ConnectedClusterResourceUri'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToProvisionedClusterPropertiesPtrOutput().ApplyT(func(v *ProvisionedClusterProperties) *ProvisionedClusterProperties { return v.Defaults() }).(ProvisionedClusterPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20231115preview:provisionedClusterInstance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProvisionedClusterInstance
	err := ctx.RegisterResource("azure-native:hybridcontainerservice/v20231115preview:ProvisionedClusterInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProvisionedClusterInstance gets an existing ProvisionedClusterInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProvisionedClusterInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProvisionedClusterInstanceState, opts ...pulumi.ResourceOption) (*ProvisionedClusterInstance, error) {
	var resource ProvisionedClusterInstance
	err := ctx.ReadResource("azure-native:hybridcontainerservice/v20231115preview:ProvisionedClusterInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProvisionedClusterInstance resources.
type provisionedClusterInstanceState struct {
}

type ProvisionedClusterInstanceState struct {
}

func (ProvisionedClusterInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*provisionedClusterInstanceState)(nil)).Elem()
}

type provisionedClusterInstanceArgs struct {
	// The fully qualified Azure Resource manager identifier of the connected cluster resource.
	ConnectedClusterResourceUri string `pulumi:"connectedClusterResourceUri"`
	// Extended Location definition
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// All properties of the provisioned cluster
	Properties *ProvisionedClusterProperties `pulumi:"properties"`
}

// The set of arguments for constructing a ProvisionedClusterInstance resource.
type ProvisionedClusterInstanceArgs struct {
	// The fully qualified Azure Resource manager identifier of the connected cluster resource.
	ConnectedClusterResourceUri pulumi.StringInput
	// Extended Location definition
	ExtendedLocation ExtendedLocationPtrInput
	// All properties of the provisioned cluster
	Properties ProvisionedClusterPropertiesPtrInput
}

func (ProvisionedClusterInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*provisionedClusterInstanceArgs)(nil)).Elem()
}

type ProvisionedClusterInstanceInput interface {
	pulumi.Input

	ToProvisionedClusterInstanceOutput() ProvisionedClusterInstanceOutput
	ToProvisionedClusterInstanceOutputWithContext(ctx context.Context) ProvisionedClusterInstanceOutput
}

func (*ProvisionedClusterInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClusterInstance)(nil)).Elem()
}

func (i *ProvisionedClusterInstance) ToProvisionedClusterInstanceOutput() ProvisionedClusterInstanceOutput {
	return i.ToProvisionedClusterInstanceOutputWithContext(context.Background())
}

func (i *ProvisionedClusterInstance) ToProvisionedClusterInstanceOutputWithContext(ctx context.Context) ProvisionedClusterInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClusterInstanceOutput)
}

type ProvisionedClusterInstanceOutput struct{ *pulumi.OutputState }

func (ProvisionedClusterInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedClusterInstance)(nil)).Elem()
}

func (o ProvisionedClusterInstanceOutput) ToProvisionedClusterInstanceOutput() ProvisionedClusterInstanceOutput {
	return o
}

func (o ProvisionedClusterInstanceOutput) ToProvisionedClusterInstanceOutputWithContext(ctx context.Context) ProvisionedClusterInstanceOutput {
	return o
}

// Extended Location definition
func (o ProvisionedClusterInstanceOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *ProvisionedClusterInstance) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The name of the resource
func (o ProvisionedClusterInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedClusterInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// All properties of the provisioned cluster
func (o ProvisionedClusterInstanceOutput) Properties() ProvisionedClusterPropertiesResponseOutput {
	return o.ApplyT(func(v *ProvisionedClusterInstance) ProvisionedClusterPropertiesResponseOutput { return v.Properties }).(ProvisionedClusterPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ProvisionedClusterInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ProvisionedClusterInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ProvisionedClusterInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedClusterInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProvisionedClusterInstanceOutput{})
}
