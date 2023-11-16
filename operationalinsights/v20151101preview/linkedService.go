// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20151101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The top level Linked service resource container.
type LinkedService struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource id of the resource that will be linked to the workspace.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
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
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190801preview:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:LinkedService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LinkedService
	err := ctx.RegisterResource("azure-native:operationalinsights/v20151101preview:LinkedService", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:operationalinsights/v20151101preview:LinkedService", name, id, state, &resource, opts...)
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
	// The name of the resource group to get. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource id of the resource that will be linked to the workspace.
	ResourceId string `pulumi:"resourceId"`
	// Name of the Log Analytics Workspace that will contain the linkedServices resource
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a LinkedService resource.
type LinkedServiceArgs struct {
	// Name of the linkedServices resource
	LinkedServiceName pulumi.StringPtrInput
	// The name of the resource group to get. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The resource id of the resource that will be linked to the workspace.
	ResourceId pulumi.StringInput
	// Name of the Log Analytics Workspace that will contain the linkedServices resource
	WorkspaceName pulumi.StringInput
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

// Resource name.
func (o LinkedServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource id of the resource that will be linked to the workspace.
func (o LinkedServiceOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Resource type.
func (o LinkedServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceOutput{})
}
