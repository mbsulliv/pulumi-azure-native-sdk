// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an ApplicationGroup for a Namespace.
// Azure REST API version: 2022-10-01-preview.
//
// Other available API versions: 2023-01-01-preview.
func LookupApplicationGroup(ctx *pulumi.Context, args *LookupApplicationGroupArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationGroupResult
	err := ctx.Invoke("azure-native:eventhub:getApplicationGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationGroupArgs struct {
	// The Application Group name
	ApplicationGroupName string `pulumi:"applicationGroupName"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Application Group object
type LookupApplicationGroupResult struct {
	// The Unique identifier for application group.Supports SAS(SASKeyName=KeyName) or AAD(AADAppID=Guid)
	ClientAppGroupIdentifier string `pulumi:"clientAppGroupIdentifier"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Determines if Application Group is allowed to create connection with namespace or not. Once the isEnabled is set to false, all the existing connections of application group gets dropped and no new connections will be allowed
	IsEnabled *bool `pulumi:"isEnabled"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of group policies that define the behavior of application group. The policies can support resource governance scenarios such as limiting ingress or egress traffic.
	Policies []ThrottlingPolicyResponse `pulumi:"policies"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type string `pulumi:"type"`
}

func LookupApplicationGroupOutput(ctx *pulumi.Context, args LookupApplicationGroupOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationGroupResult, error) {
			args := v.(LookupApplicationGroupArgs)
			r, err := LookupApplicationGroup(ctx, &args, opts...)
			var s LookupApplicationGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationGroupResultOutput)
}

type LookupApplicationGroupOutputArgs struct {
	// The Application Group name
	ApplicationGroupName pulumi.StringInput `pulumi:"applicationGroupName"`
	// The Namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGroupArgs)(nil)).Elem()
}

// The Application Group object
type LookupApplicationGroupResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGroupResult)(nil)).Elem()
}

func (o LookupApplicationGroupResultOutput) ToLookupApplicationGroupResultOutput() LookupApplicationGroupResultOutput {
	return o
}

func (o LookupApplicationGroupResultOutput) ToLookupApplicationGroupResultOutputWithContext(ctx context.Context) LookupApplicationGroupResultOutput {
	return o
}

// The Unique identifier for application group.Supports SAS(SASKeyName=KeyName) or AAD(AADAppID=Guid)
func (o LookupApplicationGroupResultOutput) ClientAppGroupIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.ClientAppGroupIdentifier }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApplicationGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Determines if Application Group is allowed to create connection with namespace or not. Once the isEnabled is set to false, all the existing connections of application group gets dropped and no new connections will be allowed
func (o LookupApplicationGroupResultOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// The geo-location where the resource lives
func (o LookupApplicationGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupApplicationGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of group policies that define the behavior of application group. The policies can support resource governance scenarios such as limiting ingress or egress traffic.
func (o LookupApplicationGroupResultOutput) Policies() ThrottlingPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) []ThrottlingPolicyResponse { return v.Policies }).(ThrottlingPolicyResponseArrayOutput)
}

// The system meta data relating to this resource.
func (o LookupApplicationGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o LookupApplicationGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationGroupResultOutput{})
}
