// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cache

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Response to an operation on access policy assignment
// Azure REST API version: 2023-05-01-preview.
type AccessPolicyAssignment struct {
	pulumi.CustomResourceState

	// The name of the access policy that is being assigned
	AccessPolicyName pulumi.StringOutput `pulumi:"accessPolicyName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Object Id to assign access policy to
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// User friendly name for object id. Also represents username for token based authentication
	ObjectIdAlias pulumi.StringOutput `pulumi:"objectIdAlias"`
	// Provisioning state of an access policy assignment set
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccessPolicyAssignment registers a new resource with the given unique name, arguments, and options.
func NewAccessPolicyAssignment(ctx *pulumi.Context,
	name string, args *AccessPolicyAssignmentArgs, opts ...pulumi.ResourceOption) (*AccessPolicyAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'AccessPolicyName'")
	}
	if args.CacheName == nil {
		return nil, errors.New("invalid value for required argument 'CacheName'")
	}
	if args.ObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ObjectId'")
	}
	if args.ObjectIdAlias == nil {
		return nil, errors.New("invalid value for required argument 'ObjectIdAlias'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cache/v20230501preview:AccessPolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20230801:AccessPolicyAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AccessPolicyAssignment
	err := ctx.RegisterResource("azure-native:cache:AccessPolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPolicyAssignment gets an existing AccessPolicyAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPolicyAssignmentState, opts ...pulumi.ResourceOption) (*AccessPolicyAssignment, error) {
	var resource AccessPolicyAssignment
	err := ctx.ReadResource("azure-native:cache:AccessPolicyAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPolicyAssignment resources.
type accessPolicyAssignmentState struct {
}

type AccessPolicyAssignmentState struct {
}

func (AccessPolicyAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyAssignmentState)(nil)).Elem()
}

type accessPolicyAssignmentArgs struct {
	// The name of the access policy assignment.
	AccessPolicyAssignmentName *string `pulumi:"accessPolicyAssignmentName"`
	// The name of the access policy that is being assigned
	AccessPolicyName string `pulumi:"accessPolicyName"`
	// The name of the Redis cache.
	CacheName string `pulumi:"cacheName"`
	// Object Id to assign access policy to
	ObjectId string `pulumi:"objectId"`
	// User friendly name for object id. Also represents username for token based authentication
	ObjectIdAlias string `pulumi:"objectIdAlias"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a AccessPolicyAssignment resource.
type AccessPolicyAssignmentArgs struct {
	// The name of the access policy assignment.
	AccessPolicyAssignmentName pulumi.StringPtrInput
	// The name of the access policy that is being assigned
	AccessPolicyName pulumi.StringInput
	// The name of the Redis cache.
	CacheName pulumi.StringInput
	// Object Id to assign access policy to
	ObjectId pulumi.StringInput
	// User friendly name for object id. Also represents username for token based authentication
	ObjectIdAlias pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (AccessPolicyAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyAssignmentArgs)(nil)).Elem()
}

type AccessPolicyAssignmentInput interface {
	pulumi.Input

	ToAccessPolicyAssignmentOutput() AccessPolicyAssignmentOutput
	ToAccessPolicyAssignmentOutputWithContext(ctx context.Context) AccessPolicyAssignmentOutput
}

func (*AccessPolicyAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicyAssignment)(nil)).Elem()
}

func (i *AccessPolicyAssignment) ToAccessPolicyAssignmentOutput() AccessPolicyAssignmentOutput {
	return i.ToAccessPolicyAssignmentOutputWithContext(context.Background())
}

func (i *AccessPolicyAssignment) ToAccessPolicyAssignmentOutputWithContext(ctx context.Context) AccessPolicyAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyAssignmentOutput)
}

func (i *AccessPolicyAssignment) ToOutput(ctx context.Context) pulumix.Output[*AccessPolicyAssignment] {
	return pulumix.Output[*AccessPolicyAssignment]{
		OutputState: i.ToAccessPolicyAssignmentOutputWithContext(ctx).OutputState,
	}
}

type AccessPolicyAssignmentOutput struct{ *pulumi.OutputState }

func (AccessPolicyAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicyAssignment)(nil)).Elem()
}

func (o AccessPolicyAssignmentOutput) ToAccessPolicyAssignmentOutput() AccessPolicyAssignmentOutput {
	return o
}

func (o AccessPolicyAssignmentOutput) ToAccessPolicyAssignmentOutputWithContext(ctx context.Context) AccessPolicyAssignmentOutput {
	return o
}

func (o AccessPolicyAssignmentOutput) ToOutput(ctx context.Context) pulumix.Output[*AccessPolicyAssignment] {
	return pulumix.Output[*AccessPolicyAssignment]{
		OutputState: o.OutputState,
	}
}

// The name of the access policy that is being assigned
func (o AccessPolicyAssignmentOutput) AccessPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicyAssignment) pulumi.StringOutput { return v.AccessPolicyName }).(pulumi.StringOutput)
}

// The name of the resource
func (o AccessPolicyAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicyAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Object Id to assign access policy to
func (o AccessPolicyAssignmentOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicyAssignment) pulumi.StringOutput { return v.ObjectId }).(pulumi.StringOutput)
}

// User friendly name for object id. Also represents username for token based authentication
func (o AccessPolicyAssignmentOutput) ObjectIdAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicyAssignment) pulumi.StringOutput { return v.ObjectIdAlias }).(pulumi.StringOutput)
}

// Provisioning state of an access policy assignment set
func (o AccessPolicyAssignmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicyAssignment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AccessPolicyAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicyAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessPolicyAssignmentOutput{})
}
