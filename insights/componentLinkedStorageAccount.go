// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Application Insights component linked storage accounts
// Azure REST API version: 2020-03-01-preview. Prior API version in Azure Native 1.x: 2020-03-01-preview.
type ComponentLinkedStorageAccount struct {
	pulumi.CustomResourceState

	// Linked storage account resource ID
	LinkedStorageAccount pulumi.StringPtrOutput `pulumi:"linkedStorageAccount"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewComponentLinkedStorageAccount registers a new resource with the given unique name, arguments, and options.
func NewComponentLinkedStorageAccount(ctx *pulumi.Context,
	name string, args *ComponentLinkedStorageAccountArgs, opts ...pulumi.ResourceOption) (*ComponentLinkedStorageAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20200301preview:ComponentLinkedStorageAccount"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ComponentLinkedStorageAccount
	err := ctx.RegisterResource("azure-native:insights:ComponentLinkedStorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComponentLinkedStorageAccount gets an existing ComponentLinkedStorageAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComponentLinkedStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentLinkedStorageAccountState, opts ...pulumi.ResourceOption) (*ComponentLinkedStorageAccount, error) {
	var resource ComponentLinkedStorageAccount
	err := ctx.ReadResource("azure-native:insights:ComponentLinkedStorageAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComponentLinkedStorageAccount resources.
type componentLinkedStorageAccountState struct {
}

type ComponentLinkedStorageAccountState struct {
}

func (ComponentLinkedStorageAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentLinkedStorageAccountState)(nil)).Elem()
}

type componentLinkedStorageAccountArgs struct {
	// Linked storage account resource ID
	LinkedStorageAccount *string `pulumi:"linkedStorageAccount"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName string `pulumi:"resourceName"`
	// The type of the Application Insights component data source for the linked storage account.
	StorageType *string `pulumi:"storageType"`
}

// The set of arguments for constructing a ComponentLinkedStorageAccount resource.
type ComponentLinkedStorageAccountArgs struct {
	// Linked storage account resource ID
	LinkedStorageAccount pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringInput
	// The type of the Application Insights component data source for the linked storage account.
	StorageType pulumi.StringPtrInput
}

func (ComponentLinkedStorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentLinkedStorageAccountArgs)(nil)).Elem()
}

type ComponentLinkedStorageAccountInput interface {
	pulumi.Input

	ToComponentLinkedStorageAccountOutput() ComponentLinkedStorageAccountOutput
	ToComponentLinkedStorageAccountOutputWithContext(ctx context.Context) ComponentLinkedStorageAccountOutput
}

func (*ComponentLinkedStorageAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**ComponentLinkedStorageAccount)(nil)).Elem()
}

func (i *ComponentLinkedStorageAccount) ToComponentLinkedStorageAccountOutput() ComponentLinkedStorageAccountOutput {
	return i.ToComponentLinkedStorageAccountOutputWithContext(context.Background())
}

func (i *ComponentLinkedStorageAccount) ToComponentLinkedStorageAccountOutputWithContext(ctx context.Context) ComponentLinkedStorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentLinkedStorageAccountOutput)
}

type ComponentLinkedStorageAccountOutput struct{ *pulumi.OutputState }

func (ComponentLinkedStorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComponentLinkedStorageAccount)(nil)).Elem()
}

func (o ComponentLinkedStorageAccountOutput) ToComponentLinkedStorageAccountOutput() ComponentLinkedStorageAccountOutput {
	return o
}

func (o ComponentLinkedStorageAccountOutput) ToComponentLinkedStorageAccountOutputWithContext(ctx context.Context) ComponentLinkedStorageAccountOutput {
	return o
}

// Linked storage account resource ID
func (o ComponentLinkedStorageAccountOutput) LinkedStorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComponentLinkedStorageAccount) pulumi.StringPtrOutput { return v.LinkedStorageAccount }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ComponentLinkedStorageAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComponentLinkedStorageAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ComponentLinkedStorageAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ComponentLinkedStorageAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ComponentLinkedStorageAccountOutput{})
}
