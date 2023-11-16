// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a TestLine
func LookupTestLine(ctx *pulumi.Context, args *LookupTestLineArgs, opts ...pulumi.InvokeOption) (*LookupTestLineResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTestLineResult
	err := ctx.Invoke("azure-native:voiceservices/v20230901:getTestLine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTestLineArgs struct {
	// Unique identifier for this deployment
	CommunicationsGatewayName string `pulumi:"communicationsGatewayName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Unique identifier for this test line
	TestLineName string `pulumi:"testLineName"`
}

// A TestLine resource
type LookupTestLineResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The phone number
	PhoneNumber string `pulumi:"phoneNumber"`
	// Resource provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Purpose of this test line, e.g. automated or manual testing
	Purpose string `pulumi:"purpose"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupTestLineOutput(ctx *pulumi.Context, args LookupTestLineOutputArgs, opts ...pulumi.InvokeOption) LookupTestLineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTestLineResult, error) {
			args := v.(LookupTestLineArgs)
			r, err := LookupTestLine(ctx, &args, opts...)
			var s LookupTestLineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTestLineResultOutput)
}

type LookupTestLineOutputArgs struct {
	// Unique identifier for this deployment
	CommunicationsGatewayName pulumi.StringInput `pulumi:"communicationsGatewayName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Unique identifier for this test line
	TestLineName pulumi.StringInput `pulumi:"testLineName"`
}

func (LookupTestLineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestLineArgs)(nil)).Elem()
}

// A TestLine resource
type LookupTestLineResultOutput struct{ *pulumi.OutputState }

func (LookupTestLineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestLineResult)(nil)).Elem()
}

func (o LookupTestLineResultOutput) ToLookupTestLineResultOutput() LookupTestLineResultOutput {
	return o
}

func (o LookupTestLineResultOutput) ToLookupTestLineResultOutputWithContext(ctx context.Context) LookupTestLineResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupTestLineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupTestLineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupTestLineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.Name }).(pulumi.StringOutput)
}

// The phone number
func (o LookupTestLineResultOutput) PhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.PhoneNumber }).(pulumi.StringOutput)
}

// Resource provisioning state.
func (o LookupTestLineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Purpose of this test line, e.g. automated or manual testing
func (o LookupTestLineResultOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.Purpose }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupTestLineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTestLineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupTestLineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTestLineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupTestLineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTestLineResultOutput{})
}
