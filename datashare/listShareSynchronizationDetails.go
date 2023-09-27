// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// List synchronization details
// Azure REST API version: 2021-08-01.
func ListShareSynchronizationDetails(ctx *pulumi.Context, args *ListShareSynchronizationDetailsArgs, opts ...pulumi.InvokeOption) (*ListShareSynchronizationDetailsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListShareSynchronizationDetailsResult
	err := ctx.Invoke("azure-native:datashare:listShareSynchronizationDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListShareSynchronizationDetailsArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// Email of the user who created the synchronization
	ConsumerEmail *string `pulumi:"consumerEmail"`
	// Name of the user who created the synchronization
	ConsumerName *string `pulumi:"consumerName"`
	// Tenant name of the consumer who created the synchronization
	ConsumerTenantName *string `pulumi:"consumerTenantName"`
	// synchronization duration
	DurationMs *int `pulumi:"durationMs"`
	// End time of synchronization
	EndTime *string `pulumi:"endTime"`
	// Filters the results using OData syntax.
	Filter *string `pulumi:"filter"`
	// message of synchronization
	Message *string `pulumi:"message"`
	// Sorts the results using OData syntax.
	Orderby *string `pulumi:"orderby"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName string `pulumi:"shareName"`
	// Continuation token
	SkipToken *string `pulumi:"skipToken"`
	// start time of synchronization
	StartTime *string `pulumi:"startTime"`
	// Raw Status
	Status *string `pulumi:"status"`
	// Synchronization id
	SynchronizationId *string `pulumi:"synchronizationId"`
}

// details of synchronization
type ListShareSynchronizationDetailsResult struct {
	// The Url of next result page.
	NextLink *string `pulumi:"nextLink"`
	// Collection of items of type DataTransferObjects.
	Value []SynchronizationDetailsResponse `pulumi:"value"`
}

func ListShareSynchronizationDetailsOutput(ctx *pulumi.Context, args ListShareSynchronizationDetailsOutputArgs, opts ...pulumi.InvokeOption) ListShareSynchronizationDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListShareSynchronizationDetailsResult, error) {
			args := v.(ListShareSynchronizationDetailsArgs)
			r, err := ListShareSynchronizationDetails(ctx, &args, opts...)
			var s ListShareSynchronizationDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListShareSynchronizationDetailsResultOutput)
}

type ListShareSynchronizationDetailsOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Email of the user who created the synchronization
	ConsumerEmail pulumi.StringPtrInput `pulumi:"consumerEmail"`
	// Name of the user who created the synchronization
	ConsumerName pulumi.StringPtrInput `pulumi:"consumerName"`
	// Tenant name of the consumer who created the synchronization
	ConsumerTenantName pulumi.StringPtrInput `pulumi:"consumerTenantName"`
	// synchronization duration
	DurationMs pulumi.IntPtrInput `pulumi:"durationMs"`
	// End time of synchronization
	EndTime pulumi.StringPtrInput `pulumi:"endTime"`
	// Filters the results using OData syntax.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// message of synchronization
	Message pulumi.StringPtrInput `pulumi:"message"`
	// Sorts the results using OData syntax.
	Orderby pulumi.StringPtrInput `pulumi:"orderby"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName pulumi.StringInput `pulumi:"shareName"`
	// Continuation token
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
	// start time of synchronization
	StartTime pulumi.StringPtrInput `pulumi:"startTime"`
	// Raw Status
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Synchronization id
	SynchronizationId pulumi.StringPtrInput `pulumi:"synchronizationId"`
}

func (ListShareSynchronizationDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListShareSynchronizationDetailsArgs)(nil)).Elem()
}

// details of synchronization
type ListShareSynchronizationDetailsResultOutput struct{ *pulumi.OutputState }

func (ListShareSynchronizationDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListShareSynchronizationDetailsResult)(nil)).Elem()
}

func (o ListShareSynchronizationDetailsResultOutput) ToListShareSynchronizationDetailsResultOutput() ListShareSynchronizationDetailsResultOutput {
	return o
}

func (o ListShareSynchronizationDetailsResultOutput) ToListShareSynchronizationDetailsResultOutputWithContext(ctx context.Context) ListShareSynchronizationDetailsResultOutput {
	return o
}

func (o ListShareSynchronizationDetailsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListShareSynchronizationDetailsResult] {
	return pulumix.Output[ListShareSynchronizationDetailsResult]{
		OutputState: o.OutputState,
	}
}

// The Url of next result page.
func (o ListShareSynchronizationDetailsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListShareSynchronizationDetailsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// Collection of items of type DataTransferObjects.
func (o ListShareSynchronizationDetailsResultOutput) Value() SynchronizationDetailsResponseArrayOutput {
	return o.ApplyT(func(v ListShareSynchronizationDetailsResult) []SynchronizationDetailsResponse { return v.Value }).(SynchronizationDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListShareSynchronizationDetailsResultOutput{})
}
