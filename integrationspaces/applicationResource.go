// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package integrationspaces

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A resource under application.
// Azure REST API version: 2023-11-14-preview.
type ApplicationResource struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The Arm id of the application resource.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The kind of the application resource.
	ResourceKind pulumi.StringPtrOutput `pulumi:"resourceKind"`
	// The type of the application resource.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationResource registers a new resource with the given unique name, arguments, and options.
func NewApplicationResource(ctx *pulumi.Context,
	name string, args *ApplicationResourceArgs, opts ...pulumi.ResourceOption) (*ApplicationResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.SpaceName == nil {
		return nil, errors.New("invalid value for required argument 'SpaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:integrationspaces/v20231114preview:ApplicationResource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApplicationResource
	err := ctx.RegisterResource("azure-native:integrationspaces:ApplicationResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationResource gets an existing ApplicationResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationResourceState, opts ...pulumi.ResourceOption) (*ApplicationResource, error) {
	var resource ApplicationResource
	err := ctx.ReadResource("azure-native:integrationspaces:ApplicationResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationResource resources.
type applicationResourceState struct {
}

type ApplicationResourceState struct {
}

func (ApplicationResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationResourceState)(nil)).Elem()
}

type applicationResourceArgs struct {
	// The name of the Application
	ApplicationName string `pulumi:"applicationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Arm id of the application resource.
	ResourceId string `pulumi:"resourceId"`
	// The kind of the application resource.
	ResourceKind *string `pulumi:"resourceKind"`
	// The name of the application resource.
	ResourceName *string `pulumi:"resourceName"`
	// The type of the application resource.
	ResourceType string `pulumi:"resourceType"`
	// The name of the space
	SpaceName string `pulumi:"spaceName"`
}

// The set of arguments for constructing a ApplicationResource resource.
type ApplicationResourceArgs struct {
	// The name of the Application
	ApplicationName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The Arm id of the application resource.
	ResourceId pulumi.StringInput
	// The kind of the application resource.
	ResourceKind pulumi.StringPtrInput
	// The name of the application resource.
	ResourceName pulumi.StringPtrInput
	// The type of the application resource.
	ResourceType pulumi.StringInput
	// The name of the space
	SpaceName pulumi.StringInput
}

func (ApplicationResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationResourceArgs)(nil)).Elem()
}

type ApplicationResourceInput interface {
	pulumi.Input

	ToApplicationResourceOutput() ApplicationResourceOutput
	ToApplicationResourceOutputWithContext(ctx context.Context) ApplicationResourceOutput
}

func (*ApplicationResource) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationResource)(nil)).Elem()
}

func (i *ApplicationResource) ToApplicationResourceOutput() ApplicationResourceOutput {
	return i.ToApplicationResourceOutputWithContext(context.Background())
}

func (i *ApplicationResource) ToApplicationResourceOutputWithContext(ctx context.Context) ApplicationResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationResourceOutput)
}

func (i *ApplicationResource) ToOutput(ctx context.Context) pulumix.Output[*ApplicationResource] {
	return pulumix.Output[*ApplicationResource]{
		OutputState: i.ToApplicationResourceOutputWithContext(ctx).OutputState,
	}
}

type ApplicationResourceOutput struct{ *pulumi.OutputState }

func (ApplicationResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationResource)(nil)).Elem()
}

func (o ApplicationResourceOutput) ToApplicationResourceOutput() ApplicationResourceOutput {
	return o
}

func (o ApplicationResourceOutput) ToApplicationResourceOutputWithContext(ctx context.Context) ApplicationResourceOutput {
	return o
}

func (o ApplicationResourceOutput) ToOutput(ctx context.Context) pulumix.Output[*ApplicationResource] {
	return pulumix.Output[*ApplicationResource]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o ApplicationResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o ApplicationResourceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationResource) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The Arm id of the application resource.
func (o ApplicationResourceOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationResource) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// The kind of the application resource.
func (o ApplicationResourceOutput) ResourceKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationResource) pulumi.StringPtrOutput { return v.ResourceKind }).(pulumi.StringPtrOutput)
}

// The type of the application resource.
func (o ApplicationResourceOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationResource) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ApplicationResourceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApplicationResource) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApplicationResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationResourceOutput{})
}
