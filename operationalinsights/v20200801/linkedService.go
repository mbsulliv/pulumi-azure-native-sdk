// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The top level Linked service resource container.
type LinkedService struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the linked service.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
	WriteAccessResourceId pulumi.StringPtrOutput `pulumi:"writeAccessResourceId"`
}

// NewLinkedService registers a new resource with the given unique name, arguments, and options.
func NewLinkedService(ctx *pulumi.Context,
	name string, args *LinkedServiceArgs, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20151101preview:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190801preview:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:LinkedService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LinkedService
	err := ctx.RegisterResource("azure-native:operationalinsights/v20200801:LinkedService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedService gets an existing LinkedService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceState, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	var resource LinkedService
	err := ctx.ReadResource("azure-native:operationalinsights/v20200801:LinkedService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedService resources.
type linkedServiceState struct {
}

type LinkedServiceState struct {
}

func (LinkedServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceState)(nil)).Elem()
}

type linkedServiceArgs struct {
	// Name of the linkedServices resource
	LinkedServiceName *string `pulumi:"linkedServiceName"`
	// The provisioning state of the linked service.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
	ResourceId *string `pulumi:"resourceId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
	WriteAccessResourceId *string `pulumi:"writeAccessResourceId"`
}

// The set of arguments for constructing a LinkedService resource.
type LinkedServiceArgs struct {
	// Name of the linkedServices resource
	LinkedServiceName pulumi.StringPtrInput
	// The provisioning state of the linked service.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
	ResourceId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
	WriteAccessResourceId pulumi.StringPtrInput
}

func (LinkedServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceArgs)(nil)).Elem()
}

type LinkedServiceInput interface {
	pulumi.Input

	ToLinkedServiceOutput() LinkedServiceOutput
	ToLinkedServiceOutputWithContext(ctx context.Context) LinkedServiceOutput
}

func (*LinkedService) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedService)(nil)).Elem()
}

func (i *LinkedService) ToLinkedServiceOutput() LinkedServiceOutput {
	return i.ToLinkedServiceOutputWithContext(context.Background())
}

func (i *LinkedService) ToLinkedServiceOutputWithContext(ctx context.Context) LinkedServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceOutput)
}

func (i *LinkedService) ToOutput(ctx context.Context) pulumix.Output[*LinkedService] {
	return pulumix.Output[*LinkedService]{
		OutputState: i.ToLinkedServiceOutputWithContext(ctx).OutputState,
	}
}

type LinkedServiceOutput struct{ *pulumi.OutputState }

func (LinkedServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedService)(nil)).Elem()
}

func (o LinkedServiceOutput) ToLinkedServiceOutput() LinkedServiceOutput {
	return o
}

func (o LinkedServiceOutput) ToLinkedServiceOutputWithContext(ctx context.Context) LinkedServiceOutput {
	return o
}

func (o LinkedServiceOutput) ToOutput(ctx context.Context) pulumix.Output[*LinkedService] {
	return pulumix.Output[*LinkedService]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o LinkedServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the linked service.
func (o LinkedServiceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
func (o LinkedServiceOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LinkedServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LinkedServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
func (o LinkedServiceOutput) WriteAccessResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringPtrOutput { return v.WriteAccessResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceOutput{})
}
