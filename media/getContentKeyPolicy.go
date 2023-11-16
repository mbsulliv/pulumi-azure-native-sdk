// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the details of a Content Key Policy in the Media Services account
// Azure REST API version: 2023-01-01.
func LookupContentKeyPolicy(ctx *pulumi.Context, args *LookupContentKeyPolicyArgs, opts ...pulumi.InvokeOption) (*LookupContentKeyPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupContentKeyPolicyResult
	err := ctx.Invoke("azure-native:media:getContentKeyPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContentKeyPolicyArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Content Key Policy name.
	ContentKeyPolicyName string `pulumi:"contentKeyPolicyName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Content Key Policy resource.
type LookupContentKeyPolicyResult struct {
	// The creation date of the Policy
	Created string `pulumi:"created"`
	// A description for the Policy.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The last modified date of the Policy
	LastModified string `pulumi:"lastModified"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The Key Policy options.
	Options []ContentKeyPolicyOptionResponse `pulumi:"options"`
	// The legacy Policy ID.
	PolicyId string `pulumi:"policyId"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupContentKeyPolicyOutput(ctx *pulumi.Context, args LookupContentKeyPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupContentKeyPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContentKeyPolicyResult, error) {
			args := v.(LookupContentKeyPolicyArgs)
			r, err := LookupContentKeyPolicy(ctx, &args, opts...)
			var s LookupContentKeyPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContentKeyPolicyResultOutput)
}

type LookupContentKeyPolicyOutputArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The Content Key Policy name.
	ContentKeyPolicyName pulumi.StringInput `pulumi:"contentKeyPolicyName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContentKeyPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentKeyPolicyArgs)(nil)).Elem()
}

// A Content Key Policy resource.
type LookupContentKeyPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupContentKeyPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentKeyPolicyResult)(nil)).Elem()
}

func (o LookupContentKeyPolicyResultOutput) ToLookupContentKeyPolicyResultOutput() LookupContentKeyPolicyResultOutput {
	return o
}

func (o LookupContentKeyPolicyResultOutput) ToLookupContentKeyPolicyResultOutputWithContext(ctx context.Context) LookupContentKeyPolicyResultOutput {
	return o
}

// The creation date of the Policy
func (o LookupContentKeyPolicyResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) string { return v.Created }).(pulumi.StringOutput)
}

// A description for the Policy.
func (o LookupContentKeyPolicyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupContentKeyPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The last modified date of the Policy
func (o LookupContentKeyPolicyResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) string { return v.LastModified }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupContentKeyPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Key Policy options.
func (o LookupContentKeyPolicyResultOutput) Options() ContentKeyPolicyOptionResponseArrayOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) []ContentKeyPolicyOptionResponse { return v.Options }).(ContentKeyPolicyOptionResponseArrayOutput)
}

// The legacy Policy ID.
func (o LookupContentKeyPolicyResultOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) string { return v.PolicyId }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o LookupContentKeyPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupContentKeyPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentKeyPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContentKeyPolicyResultOutput{})
}
