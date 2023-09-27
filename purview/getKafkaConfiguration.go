// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package purview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the kafka configuration for the account
// Azure REST API version: 2021-12-01.
func LookupKafkaConfiguration(ctx *pulumi.Context, args *LookupKafkaConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupKafkaConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKafkaConfigurationResult
	err := ctx.Invoke("azure-native:purview:getKafkaConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupKafkaConfigurationArgs struct {
	// The name of the account.
	AccountName string `pulumi:"accountName"`
	// Name of kafka configuration.
	KafkaConfigurationName string `pulumi:"kafkaConfigurationName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The configuration of the event streaming service resource attached to the Purview account for kafka notifications.
type LookupKafkaConfigurationResult struct {
	// Consumer group for hook event hub.
	ConsumerGroup *string `pulumi:"consumerGroup"`
	// Credentials to access event hub.
	Credentials *CredentialsResponse `pulumi:"credentials"`
	// Optional partition Id for notification event hub. If not set, all partitions will be leveraged.
	EventHubPartitionId *string `pulumi:"eventHubPartitionId"`
	EventHubResourceId  *string `pulumi:"eventHubResourceId"`
	// The event hub type.
	EventHubType *string `pulumi:"eventHubType"`
	// The state of the event streaming service
	EventStreamingState *string `pulumi:"eventStreamingState"`
	// The event streaming service type
	EventStreamingType *string `pulumi:"eventStreamingType"`
	// Gets or sets the identifier.
	Id string `pulumi:"id"`
	// Gets or sets the name.
	Name string `pulumi:"name"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData ProxyResourceResponseSystemData `pulumi:"systemData"`
	// Gets or sets the type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupKafkaConfigurationResult
func (val *LookupKafkaConfigurationResult) Defaults() *LookupKafkaConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EventStreamingState == nil {
		eventStreamingState_ := "Enabled"
		tmp.EventStreamingState = &eventStreamingState_
	}
	if tmp.EventStreamingType == nil {
		eventStreamingType_ := "None"
		tmp.EventStreamingType = &eventStreamingType_
	}
	return &tmp
}

func LookupKafkaConfigurationOutput(ctx *pulumi.Context, args LookupKafkaConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupKafkaConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKafkaConfigurationResult, error) {
			args := v.(LookupKafkaConfigurationArgs)
			r, err := LookupKafkaConfiguration(ctx, &args, opts...)
			var s LookupKafkaConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKafkaConfigurationResultOutput)
}

type LookupKafkaConfigurationOutputArgs struct {
	// The name of the account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Name of kafka configuration.
	KafkaConfigurationName pulumi.StringInput `pulumi:"kafkaConfigurationName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupKafkaConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKafkaConfigurationArgs)(nil)).Elem()
}

// The configuration of the event streaming service resource attached to the Purview account for kafka notifications.
type LookupKafkaConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupKafkaConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKafkaConfigurationResult)(nil)).Elem()
}

func (o LookupKafkaConfigurationResultOutput) ToLookupKafkaConfigurationResultOutput() LookupKafkaConfigurationResultOutput {
	return o
}

func (o LookupKafkaConfigurationResultOutput) ToLookupKafkaConfigurationResultOutputWithContext(ctx context.Context) LookupKafkaConfigurationResultOutput {
	return o
}

func (o LookupKafkaConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupKafkaConfigurationResult] {
	return pulumix.Output[LookupKafkaConfigurationResult]{
		OutputState: o.OutputState,
	}
}

// Consumer group for hook event hub.
func (o LookupKafkaConfigurationResultOutput) ConsumerGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *string { return v.ConsumerGroup }).(pulumi.StringPtrOutput)
}

// Credentials to access event hub.
func (o LookupKafkaConfigurationResultOutput) Credentials() CredentialsResponsePtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *CredentialsResponse { return v.Credentials }).(CredentialsResponsePtrOutput)
}

// Optional partition Id for notification event hub. If not set, all partitions will be leveraged.
func (o LookupKafkaConfigurationResultOutput) EventHubPartitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *string { return v.EventHubPartitionId }).(pulumi.StringPtrOutput)
}

func (o LookupKafkaConfigurationResultOutput) EventHubResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *string { return v.EventHubResourceId }).(pulumi.StringPtrOutput)
}

// The event hub type.
func (o LookupKafkaConfigurationResultOutput) EventHubType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *string { return v.EventHubType }).(pulumi.StringPtrOutput)
}

// The state of the event streaming service
func (o LookupKafkaConfigurationResultOutput) EventStreamingState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *string { return v.EventStreamingState }).(pulumi.StringPtrOutput)
}

// The event streaming service type
func (o LookupKafkaConfigurationResultOutput) EventStreamingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *string { return v.EventStreamingType }).(pulumi.StringPtrOutput)
}

// Gets or sets the identifier.
func (o LookupKafkaConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the name.
func (o LookupKafkaConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupKafkaConfigurationResultOutput) SystemData() ProxyResourceResponseSystemDataOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) ProxyResourceResponseSystemData { return v.SystemData }).(ProxyResourceResponseSystemDataOutput)
}

// Gets or sets the type.
func (o LookupKafkaConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKafkaConfigurationResultOutput{})
}
