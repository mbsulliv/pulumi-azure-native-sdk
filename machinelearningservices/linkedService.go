// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Linked service.
// Azure REST API version: 2020-09-01-preview. Prior API version in Azure Native 1.x: 2020-09-01-preview.
type LinkedService struct {
	pulumi.CustomResourceState

	// Identity for the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// location of the linked service.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Friendly name of the linked service.
	Name pulumi.StringOutput `pulumi:"name"`
	// LinkedService specific properties.
	Properties LinkedServicePropsResponseOutput `pulumi:"properties"`
	// Resource type of linked service.
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
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:LinkedService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LinkedService
	err := ctx.RegisterResource("azure-native:machinelearningservices:LinkedService", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:machinelearningservices:LinkedService", name, id, state, &resource, opts...)
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
	// Identity for the resource.
	Identity *Identity `pulumi:"identity"`
	// Friendly name of the linked workspace
	LinkName *string `pulumi:"linkName"`
	// location of the linked service.
	Location *string `pulumi:"location"`
	// Friendly name of the linked service
	Name *string `pulumi:"name"`
	// LinkedService specific properties.
	Properties *LinkedServiceProps `pulumi:"properties"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a LinkedService resource.
type LinkedServiceArgs struct {
	// Identity for the resource.
	Identity IdentityPtrInput
	// Friendly name of the linked workspace
	LinkName pulumi.StringPtrInput
	// location of the linked service.
	Location pulumi.StringPtrInput
	// Friendly name of the linked service
	Name pulumi.StringPtrInput
	// LinkedService specific properties.
	Properties LinkedServicePropsPtrInput
	// Name of the resource group in which workspace is located.
	ResourceGroupName pulumi.StringInput
	// Name of Azure Machine Learning workspace.
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

// Identity for the resource.
func (o LinkedServiceOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *LinkedService) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// location of the linked service.
func (o LinkedServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Friendly name of the linked service.
func (o LinkedServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// LinkedService specific properties.
func (o LinkedServiceOutput) Properties() LinkedServicePropsResponseOutput {
	return o.ApplyT(func(v *LinkedService) LinkedServicePropsResponseOutput { return v.Properties }).(LinkedServicePropsResponseOutput)
}

// Resource type of linked service.
func (o LinkedServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceOutput{})
}
