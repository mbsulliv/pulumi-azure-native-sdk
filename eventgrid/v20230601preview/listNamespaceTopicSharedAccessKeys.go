// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List the two keys used to publish to a namespace topic.
func ListNamespaceTopicSharedAccessKeys(ctx *pulumi.Context, args *ListNamespaceTopicSharedAccessKeysArgs, opts ...pulumi.InvokeOption) (*ListNamespaceTopicSharedAccessKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListNamespaceTopicSharedAccessKeysResult
	err := ctx.Invoke("azure-native:eventgrid/v20230601preview:listNamespaceTopicSharedAccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNamespaceTopicSharedAccessKeysArgs struct {
	// Name of the namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the topic.
	TopicName string `pulumi:"topicName"`
}

// Shared access keys of the Topic
type ListNamespaceTopicSharedAccessKeysResult struct {
	// Shared access key1 for the topic.
	Key1 *string `pulumi:"key1"`
	// Shared access key2 for the topic.
	Key2 *string `pulumi:"key2"`
}

func ListNamespaceTopicSharedAccessKeysOutput(ctx *pulumi.Context, args ListNamespaceTopicSharedAccessKeysOutputArgs, opts ...pulumi.InvokeOption) ListNamespaceTopicSharedAccessKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNamespaceTopicSharedAccessKeysResult, error) {
			args := v.(ListNamespaceTopicSharedAccessKeysArgs)
			r, err := ListNamespaceTopicSharedAccessKeys(ctx, &args, opts...)
			var s ListNamespaceTopicSharedAccessKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNamespaceTopicSharedAccessKeysResultOutput)
}

type ListNamespaceTopicSharedAccessKeysOutputArgs struct {
	// Name of the namespace.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the topic.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (ListNamespaceTopicSharedAccessKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNamespaceTopicSharedAccessKeysArgs)(nil)).Elem()
}

// Shared access keys of the Topic
type ListNamespaceTopicSharedAccessKeysResultOutput struct{ *pulumi.OutputState }

func (ListNamespaceTopicSharedAccessKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNamespaceTopicSharedAccessKeysResult)(nil)).Elem()
}

func (o ListNamespaceTopicSharedAccessKeysResultOutput) ToListNamespaceTopicSharedAccessKeysResultOutput() ListNamespaceTopicSharedAccessKeysResultOutput {
	return o
}

func (o ListNamespaceTopicSharedAccessKeysResultOutput) ToListNamespaceTopicSharedAccessKeysResultOutputWithContext(ctx context.Context) ListNamespaceTopicSharedAccessKeysResultOutput {
	return o
}

// Shared access key1 for the topic.
func (o ListNamespaceTopicSharedAccessKeysResultOutput) Key1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNamespaceTopicSharedAccessKeysResult) *string { return v.Key1 }).(pulumi.StringPtrOutput)
}

// Shared access key2 for the topic.
func (o ListNamespaceTopicSharedAccessKeysResultOutput) Key2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNamespaceTopicSharedAccessKeysResult) *string { return v.Key2 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNamespaceTopicSharedAccessKeysResultOutput{})
}
