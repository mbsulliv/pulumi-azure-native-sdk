// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The properties of a storage account’s Queue service.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2021-02-01.
//
// Other available API versions: 2023-01-01, 2023-04-01.
type QueueServiceProperties struct {
	pulumi.CustomResourceState

	// Specifies CORS rules for the Queue service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Queue service.
	Cors CorsRulesResponsePtrOutput `pulumi:"cors"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewQueueServiceProperties registers a new resource with the given unique name, arguments, and options.
func NewQueueServiceProperties(ctx *pulumi.Context,
	name string, args *QueueServicePropertiesArgs, opts ...pulumi.ResourceOption) (*QueueServiceProperties, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage/v20190601:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220901:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20230101:QueueServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20230401:QueueServiceProperties"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource QueueServiceProperties
	err := ctx.RegisterResource("azure-native:storage:QueueServiceProperties", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueueServiceProperties gets an existing QueueServiceProperties resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueueServiceProperties(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueServicePropertiesState, opts ...pulumi.ResourceOption) (*QueueServiceProperties, error) {
	var resource QueueServiceProperties
	err := ctx.ReadResource("azure-native:storage:QueueServiceProperties", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueueServiceProperties resources.
type queueServicePropertiesState struct {
}

type QueueServicePropertiesState struct {
}

func (QueueServicePropertiesState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueServicePropertiesState)(nil)).Elem()
}

type queueServicePropertiesArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// Specifies CORS rules for the Queue service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Queue service.
	Cors *CorsRules `pulumi:"cors"`
	// The name of the Queue Service within the specified storage account. Queue Service Name must be 'default'
	QueueServiceName *string `pulumi:"queueServiceName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a QueueServiceProperties resource.
type QueueServicePropertiesArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// Specifies CORS rules for the Queue service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Queue service.
	Cors CorsRulesPtrInput
	// The name of the Queue Service within the specified storage account. Queue Service Name must be 'default'
	QueueServiceName pulumi.StringPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (QueueServicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueServicePropertiesArgs)(nil)).Elem()
}

type QueueServicePropertiesInput interface {
	pulumi.Input

	ToQueueServicePropertiesOutput() QueueServicePropertiesOutput
	ToQueueServicePropertiesOutputWithContext(ctx context.Context) QueueServicePropertiesOutput
}

func (*QueueServiceProperties) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueServiceProperties)(nil)).Elem()
}

func (i *QueueServiceProperties) ToQueueServicePropertiesOutput() QueueServicePropertiesOutput {
	return i.ToQueueServicePropertiesOutputWithContext(context.Background())
}

func (i *QueueServiceProperties) ToQueueServicePropertiesOutputWithContext(ctx context.Context) QueueServicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueServicePropertiesOutput)
}

type QueueServicePropertiesOutput struct{ *pulumi.OutputState }

func (QueueServicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueServiceProperties)(nil)).Elem()
}

func (o QueueServicePropertiesOutput) ToQueueServicePropertiesOutput() QueueServicePropertiesOutput {
	return o
}

func (o QueueServicePropertiesOutput) ToQueueServicePropertiesOutputWithContext(ctx context.Context) QueueServicePropertiesOutput {
	return o
}

// Specifies CORS rules for the Queue service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Queue service.
func (o QueueServicePropertiesOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v *QueueServiceProperties) CorsRulesResponsePtrOutput { return v.Cors }).(CorsRulesResponsePtrOutput)
}

// The name of the resource
func (o QueueServicePropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueServiceProperties) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o QueueServicePropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueServiceProperties) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueServicePropertiesOutput{})
}
