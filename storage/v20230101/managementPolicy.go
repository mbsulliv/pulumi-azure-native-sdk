// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Get Storage Account ManagementPolicies operation response.
type ManagementPolicy struct {
	pulumi.CustomResourceState

	// Returns the date and time the ManagementPolicies was last modified.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Storage Account ManagementPolicy, in JSON format. See more details in: https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
	Policy ManagementPolicySchemaResponseOutput `pulumi:"policy"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagementPolicy registers a new resource with the given unique name, arguments, and options.
func NewManagementPolicy(ctx *pulumi.Context,
	name string, args *ManagementPolicyArgs, opts ...pulumi.ResourceOption) (*ManagementPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180301preview:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20181101:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190401:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:ManagementPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220901:ManagementPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagementPolicy
	err := ctx.RegisterResource("azure-native:storage/v20230101:ManagementPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagementPolicy gets an existing ManagementPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagementPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementPolicyState, opts ...pulumi.ResourceOption) (*ManagementPolicy, error) {
	var resource ManagementPolicy
	err := ctx.ReadResource("azure-native:storage/v20230101:ManagementPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagementPolicy resources.
type managementPolicyState struct {
}

type ManagementPolicyState struct {
}

func (ManagementPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementPolicyState)(nil)).Elem()
}

type managementPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the Storage Account Management Policy. It should always be 'default'
	ManagementPolicyName *string `pulumi:"managementPolicyName"`
	// The Storage Account ManagementPolicy, in JSON format. See more details in: https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
	Policy ManagementPolicySchema `pulumi:"policy"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagementPolicy resource.
type ManagementPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// The name of the Storage Account Management Policy. It should always be 'default'
	ManagementPolicyName pulumi.StringPtrInput
	// The Storage Account ManagementPolicy, in JSON format. See more details in: https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
	Policy ManagementPolicySchemaInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ManagementPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementPolicyArgs)(nil)).Elem()
}

type ManagementPolicyInput interface {
	pulumi.Input

	ToManagementPolicyOutput() ManagementPolicyOutput
	ToManagementPolicyOutputWithContext(ctx context.Context) ManagementPolicyOutput
}

func (*ManagementPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicy)(nil)).Elem()
}

func (i *ManagementPolicy) ToManagementPolicyOutput() ManagementPolicyOutput {
	return i.ToManagementPolicyOutputWithContext(context.Background())
}

func (i *ManagementPolicy) ToManagementPolicyOutputWithContext(ctx context.Context) ManagementPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyOutput)
}

func (i *ManagementPolicy) ToOutput(ctx context.Context) pulumix.Output[*ManagementPolicy] {
	return pulumix.Output[*ManagementPolicy]{
		OutputState: i.ToManagementPolicyOutputWithContext(ctx).OutputState,
	}
}

type ManagementPolicyOutput struct{ *pulumi.OutputState }

func (ManagementPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicy)(nil)).Elem()
}

func (o ManagementPolicyOutput) ToManagementPolicyOutput() ManagementPolicyOutput {
	return o
}

func (o ManagementPolicyOutput) ToManagementPolicyOutputWithContext(ctx context.Context) ManagementPolicyOutput {
	return o
}

func (o ManagementPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*ManagementPolicy] {
	return pulumix.Output[*ManagementPolicy]{
		OutputState: o.OutputState,
	}
}

// Returns the date and time the ManagementPolicies was last modified.
func (o ManagementPolicyOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementPolicy) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o ManagementPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Storage Account ManagementPolicy, in JSON format. See more details in: https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
func (o ManagementPolicyOutput) Policy() ManagementPolicySchemaResponseOutput {
	return o.ApplyT(func(v *ManagementPolicy) ManagementPolicySchemaResponseOutput { return v.Policy }).(ManagementPolicySchemaResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ManagementPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementPolicyOutput{})
}
