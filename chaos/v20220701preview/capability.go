// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Model that represents a Capability resource.
type Capability struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of a capability resource.
	Properties CapabilityPropertiesResponseOutput `pulumi:"properties"`
	// The standard system metadata of a resource type.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCapability registers a new resource with the given unique name, arguments, and options.
func NewCapability(ctx *pulumi.Context,
	name string, args *CapabilityArgs, opts ...pulumi.ResourceOption) (*Capability, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ParentProviderNamespace'")
	}
	if args.ParentResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourceName'")
	}
	if args.ParentResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourceType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TargetName == nil {
		return nil, errors.New("invalid value for required argument 'TargetName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:chaos:Capability"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20210915preview:Capability"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20221001preview:Capability"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20230401preview:Capability"),
		},
	})
	opts = append(opts, aliases)
	var resource Capability
	err := ctx.RegisterResource("azure-native:chaos/v20220701preview:Capability", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapability gets an existing Capability resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapability(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapabilityState, opts ...pulumi.ResourceOption) (*Capability, error) {
	var resource Capability
	err := ctx.ReadResource("azure-native:chaos/v20220701preview:Capability", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Capability resources.
type capabilityState struct {
}

type CapabilityState struct {
}

func (CapabilityState) ElementType() reflect.Type {
	return reflect.TypeOf((*capabilityState)(nil)).Elem()
}

type capabilityArgs struct {
	// String that represents a Capability resource name.
	CapabilityName *string `pulumi:"capabilityName"`
	// String that represents a resource provider namespace.
	ParentProviderNamespace string `pulumi:"parentProviderNamespace"`
	// String that represents a resource name.
	ParentResourceName string `pulumi:"parentResourceName"`
	// String that represents a resource type.
	ParentResourceType string `pulumi:"parentResourceType"`
	// String that represents an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// String that represents a Target resource name.
	TargetName string `pulumi:"targetName"`
}

// The set of arguments for constructing a Capability resource.
type CapabilityArgs struct {
	// String that represents a Capability resource name.
	CapabilityName pulumi.StringPtrInput
	// String that represents a resource provider namespace.
	ParentProviderNamespace pulumi.StringInput
	// String that represents a resource name.
	ParentResourceName pulumi.StringInput
	// String that represents a resource type.
	ParentResourceType pulumi.StringInput
	// String that represents an Azure resource group.
	ResourceGroupName pulumi.StringInput
	// String that represents a Target resource name.
	TargetName pulumi.StringInput
}

func (CapabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capabilityArgs)(nil)).Elem()
}

type CapabilityInput interface {
	pulumi.Input

	ToCapabilityOutput() CapabilityOutput
	ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput
}

func (*Capability) ElementType() reflect.Type {
	return reflect.TypeOf((**Capability)(nil)).Elem()
}

func (i *Capability) ToCapabilityOutput() CapabilityOutput {
	return i.ToCapabilityOutputWithContext(context.Background())
}

func (i *Capability) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityOutput)
}

type CapabilityOutput struct{ *pulumi.OutputState }

func (CapabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Capability)(nil)).Elem()
}

func (o CapabilityOutput) ToCapabilityOutput() CapabilityOutput {
	return o
}

func (o CapabilityOutput) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return o
}

// The name of the resource
func (o CapabilityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Capability) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The properties of a capability resource.
func (o CapabilityOutput) Properties() CapabilityPropertiesResponseOutput {
	return o.ApplyT(func(v *Capability) CapabilityPropertiesResponseOutput { return v.Properties }).(CapabilityPropertiesResponseOutput)
}

// The standard system metadata of a resource type.
func (o CapabilityOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Capability) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CapabilityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Capability) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CapabilityOutput{})
}
