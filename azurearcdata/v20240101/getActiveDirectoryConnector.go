// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves an Active Directory connector resource
func LookupActiveDirectoryConnector(ctx *pulumi.Context, args *LookupActiveDirectoryConnectorArgs, opts ...pulumi.InvokeOption) (*LookupActiveDirectoryConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupActiveDirectoryConnectorResult
	err := ctx.Invoke("azure-native:azurearcdata/v20240101:getActiveDirectoryConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupActiveDirectoryConnectorArgs struct {
	// The name of the Active Directory connector instance
	ActiveDirectoryConnectorName string `pulumi:"activeDirectoryConnectorName"`
	// The name of the data controller
	DataControllerName string `pulumi:"dataControllerName"`
	// The name of the Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Active directory connector resource
type LookupActiveDirectoryConnectorResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// null
	Properties ActiveDirectoryConnectorPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupActiveDirectoryConnectorResult
func (val *LookupActiveDirectoryConnectorResult) Defaults() *LookupActiveDirectoryConnectorResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupActiveDirectoryConnectorOutput(ctx *pulumi.Context, args LookupActiveDirectoryConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupActiveDirectoryConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupActiveDirectoryConnectorResult, error) {
			args := v.(LookupActiveDirectoryConnectorArgs)
			r, err := LookupActiveDirectoryConnector(ctx, &args, opts...)
			var s LookupActiveDirectoryConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupActiveDirectoryConnectorResultOutput)
}

type LookupActiveDirectoryConnectorOutputArgs struct {
	// The name of the Active Directory connector instance
	ActiveDirectoryConnectorName pulumi.StringInput `pulumi:"activeDirectoryConnectorName"`
	// The name of the data controller
	DataControllerName pulumi.StringInput `pulumi:"dataControllerName"`
	// The name of the Azure resource group
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupActiveDirectoryConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActiveDirectoryConnectorArgs)(nil)).Elem()
}

// Active directory connector resource
type LookupActiveDirectoryConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupActiveDirectoryConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActiveDirectoryConnectorResult)(nil)).Elem()
}

func (o LookupActiveDirectoryConnectorResultOutput) ToLookupActiveDirectoryConnectorResultOutput() LookupActiveDirectoryConnectorResultOutput {
	return o
}

func (o LookupActiveDirectoryConnectorResultOutput) ToLookupActiveDirectoryConnectorResultOutputWithContext(ctx context.Context) LookupActiveDirectoryConnectorResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupActiveDirectoryConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActiveDirectoryConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupActiveDirectoryConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActiveDirectoryConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// null
func (o LookupActiveDirectoryConnectorResultOutput) Properties() ActiveDirectoryConnectorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupActiveDirectoryConnectorResult) ActiveDirectoryConnectorPropertiesResponse {
		return v.Properties
	}).(ActiveDirectoryConnectorPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupActiveDirectoryConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupActiveDirectoryConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupActiveDirectoryConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActiveDirectoryConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActiveDirectoryConnectorResultOutput{})
}
