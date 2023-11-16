// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a network manager security admin configuration rule collection.
func LookupAdminRuleCollection(ctx *pulumi.Context, args *LookupAdminRuleCollectionArgs, opts ...pulumi.InvokeOption) (*LookupAdminRuleCollectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAdminRuleCollectionResult
	err := ctx.Invoke("azure-native:network/v20230201:getAdminRuleCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdminRuleCollectionArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName string `pulumi:"ruleCollectionName"`
}

// Defines the admin rule collection.
type LookupAdminRuleCollectionResult struct {
	// Groups for configuration
	AppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"appliesToGroups"`
	// A description of the admin rule collection.
	Description *string `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Unique identifier for this resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupAdminRuleCollectionOutput(ctx *pulumi.Context, args LookupAdminRuleCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupAdminRuleCollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAdminRuleCollectionResult, error) {
			args := v.(LookupAdminRuleCollectionArgs)
			r, err := LookupAdminRuleCollection(ctx, &args, opts...)
			var s LookupAdminRuleCollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAdminRuleCollectionResultOutput)
}

type LookupAdminRuleCollectionOutputArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName pulumi.StringInput `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName pulumi.StringInput `pulumi:"ruleCollectionName"`
}

func (LookupAdminRuleCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdminRuleCollectionArgs)(nil)).Elem()
}

// Defines the admin rule collection.
type LookupAdminRuleCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupAdminRuleCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdminRuleCollectionResult)(nil)).Elem()
}

func (o LookupAdminRuleCollectionResultOutput) ToLookupAdminRuleCollectionResultOutput() LookupAdminRuleCollectionResultOutput {
	return o
}

func (o LookupAdminRuleCollectionResultOutput) ToLookupAdminRuleCollectionResultOutputWithContext(ctx context.Context) LookupAdminRuleCollectionResultOutput {
	return o
}

// Groups for configuration
func (o LookupAdminRuleCollectionResultOutput) AppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) []NetworkManagerSecurityGroupItemResponse {
		return v.AppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

// A description of the admin rule collection.
func (o LookupAdminRuleCollectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupAdminRuleCollectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupAdminRuleCollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupAdminRuleCollectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupAdminRuleCollectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Unique identifier for this resource.
func (o LookupAdminRuleCollectionResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o LookupAdminRuleCollectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupAdminRuleCollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAdminRuleCollectionResultOutput{})
}
