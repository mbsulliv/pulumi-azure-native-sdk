// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Subscription keys.
// Azure REST API version: 2022-08-01.
//
// Other available API versions: 2022-09-01-preview, 2023-03-01-preview.
func ListSubscriptionSecrets(ctx *pulumi.Context, args *ListSubscriptionSecretsArgs, opts ...pulumi.InvokeOption) (*ListSubscriptionSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListSubscriptionSecretsResult
	err := ctx.Invoke("azure-native:apimanagement:listSubscriptionSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSubscriptionSecretsArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
	Sid string `pulumi:"sid"`
}

// Subscription keys.
type ListSubscriptionSecretsResult struct {
	// Subscription primary key.
	PrimaryKey *string `pulumi:"primaryKey"`
	// Subscription secondary key.
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListSubscriptionSecretsOutput(ctx *pulumi.Context, args ListSubscriptionSecretsOutputArgs, opts ...pulumi.InvokeOption) ListSubscriptionSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSubscriptionSecretsResult, error) {
			args := v.(ListSubscriptionSecretsArgs)
			r, err := ListSubscriptionSecrets(ctx, &args, opts...)
			var s ListSubscriptionSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSubscriptionSecretsResultOutput)
}

type ListSubscriptionSecretsOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
	Sid pulumi.StringInput `pulumi:"sid"`
}

func (ListSubscriptionSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSubscriptionSecretsArgs)(nil)).Elem()
}

// Subscription keys.
type ListSubscriptionSecretsResultOutput struct{ *pulumi.OutputState }

func (ListSubscriptionSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSubscriptionSecretsResult)(nil)).Elem()
}

func (o ListSubscriptionSecretsResultOutput) ToListSubscriptionSecretsResultOutput() ListSubscriptionSecretsResultOutput {
	return o
}

func (o ListSubscriptionSecretsResultOutput) ToListSubscriptionSecretsResultOutputWithContext(ctx context.Context) ListSubscriptionSecretsResultOutput {
	return o
}

// Subscription primary key.
func (o ListSubscriptionSecretsResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSubscriptionSecretsResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

// Subscription secondary key.
func (o ListSubscriptionSecretsResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSubscriptionSecretsResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSubscriptionSecretsResultOutput{})
}
