// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231027preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PrivateAccesses tracked resource.
type PrivateAccess struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// A readonly collection of private endpoint connection. Currently only one endpoint connection is supported.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateAccess registers a new resource with the given unique name, arguments, and options.
func NewPrivateAccess(ctx *pulumi.Context,
	name string, args *PrivateAccessArgs, opts ...pulumi.ResourceOption) (*PrivateAccess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:chaos:PrivateAccess"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateAccess
	err := ctx.RegisterResource("azure-native:chaos/v20231027preview:PrivateAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateAccess gets an existing PrivateAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateAccessState, opts ...pulumi.ResourceOption) (*PrivateAccess, error) {
	var resource PrivateAccess
	err := ctx.ReadResource("azure-native:chaos/v20231027preview:PrivateAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateAccess resources.
type privateAccessState struct {
}

type PrivateAccessState struct {
}

func (PrivateAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateAccessState)(nil)).Elem()
}

type privateAccessArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the private access resource that is being created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	PrivateAccessName *string `pulumi:"privateAccessName"`
	// String that represents an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrivateAccess resource.
type PrivateAccessArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the private access resource that is being created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	PrivateAccessName pulumi.StringPtrInput
	// String that represents an Azure resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (PrivateAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateAccessArgs)(nil)).Elem()
}

type PrivateAccessInput interface {
	pulumi.Input

	ToPrivateAccessOutput() PrivateAccessOutput
	ToPrivateAccessOutputWithContext(ctx context.Context) PrivateAccessOutput
}

func (*PrivateAccess) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateAccess)(nil)).Elem()
}

func (i *PrivateAccess) ToPrivateAccessOutput() PrivateAccessOutput {
	return i.ToPrivateAccessOutputWithContext(context.Background())
}

func (i *PrivateAccess) ToPrivateAccessOutputWithContext(ctx context.Context) PrivateAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateAccessOutput)
}

type PrivateAccessOutput struct{ *pulumi.OutputState }

func (PrivateAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateAccess)(nil)).Elem()
}

func (o PrivateAccessOutput) ToPrivateAccessOutput() PrivateAccessOutput {
	return o
}

func (o PrivateAccessOutput) ToPrivateAccessOutputWithContext(ctx context.Context) PrivateAccessOutput {
	return o
}

// The geo-location where the resource lives
func (o PrivateAccessOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateAccess) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o PrivateAccessOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateAccess) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A readonly collection of private endpoint connection. Currently only one endpoint connection is supported.
func (o PrivateAccessOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *PrivateAccess) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PrivateAccessOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateAccess) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o PrivateAccessOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateAccess) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateAccessOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateAccess) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateAccessOutput{})
}
