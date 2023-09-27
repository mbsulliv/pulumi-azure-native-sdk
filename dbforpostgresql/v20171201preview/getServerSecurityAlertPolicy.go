// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20171201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a server's security alert policy.
func LookupServerSecurityAlertPolicy(ctx *pulumi.Context, args *LookupServerSecurityAlertPolicyArgs, opts ...pulumi.InvokeOption) (*LookupServerSecurityAlertPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServerSecurityAlertPolicyResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20171201preview:getServerSecurityAlertPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerSecurityAlertPolicyArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the security alert policy.
	SecurityAlertPolicyName string `pulumi:"securityAlertPolicyName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A server security alert policy.
type LookupServerSecurityAlertPolicyResult struct {
	// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly
	DisabledAlerts []string `pulumi:"disabledAlerts"`
	// Specifies that the alert is sent to the account administrators.
	EmailAccountAdmins *bool `pulumi:"emailAccountAdmins"`
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses []string `pulumi:"emailAddresses"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays *int `pulumi:"retentionDays"`
	// Specifies the state of the policy, whether it is enabled or disabled.
	State string `pulumi:"state"`
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupServerSecurityAlertPolicyOutput(ctx *pulumi.Context, args LookupServerSecurityAlertPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupServerSecurityAlertPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerSecurityAlertPolicyResult, error) {
			args := v.(LookupServerSecurityAlertPolicyArgs)
			r, err := LookupServerSecurityAlertPolicy(ctx, &args, opts...)
			var s LookupServerSecurityAlertPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerSecurityAlertPolicyResultOutput)
}

type LookupServerSecurityAlertPolicyOutputArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the security alert policy.
	SecurityAlertPolicyName pulumi.StringInput `pulumi:"securityAlertPolicyName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerSecurityAlertPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerSecurityAlertPolicyArgs)(nil)).Elem()
}

// A server security alert policy.
type LookupServerSecurityAlertPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupServerSecurityAlertPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerSecurityAlertPolicyResult)(nil)).Elem()
}

func (o LookupServerSecurityAlertPolicyResultOutput) ToLookupServerSecurityAlertPolicyResultOutput() LookupServerSecurityAlertPolicyResultOutput {
	return o
}

func (o LookupServerSecurityAlertPolicyResultOutput) ToLookupServerSecurityAlertPolicyResultOutputWithContext(ctx context.Context) LookupServerSecurityAlertPolicyResultOutput {
	return o
}

func (o LookupServerSecurityAlertPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupServerSecurityAlertPolicyResult] {
	return pulumix.Output[LookupServerSecurityAlertPolicyResult]{
		OutputState: o.OutputState,
	}
}

// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly
func (o LookupServerSecurityAlertPolicyResultOutput) DisabledAlerts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) []string { return v.DisabledAlerts }).(pulumi.StringArrayOutput)
}

// Specifies that the alert is sent to the account administrators.
func (o LookupServerSecurityAlertPolicyResultOutput) EmailAccountAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) *bool { return v.EmailAccountAdmins }).(pulumi.BoolPtrOutput)
}

// Specifies an array of e-mail addresses to which the alert is sent.
func (o LookupServerSecurityAlertPolicyResultOutput) EmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) []string { return v.EmailAddresses }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupServerSecurityAlertPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupServerSecurityAlertPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the number of days to keep in the Threat Detection audit logs.
func (o LookupServerSecurityAlertPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Specifies the state of the policy, whether it is enabled or disabled.
func (o LookupServerSecurityAlertPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

// Specifies the identifier key of the Threat Detection audit storage account.
func (o LookupServerSecurityAlertPolicyResultOutput) StorageAccountAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) *string { return v.StorageAccountAccessKey }).(pulumi.StringPtrOutput)
}

// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
func (o LookupServerSecurityAlertPolicyResultOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) *string { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupServerSecurityAlertPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerSecurityAlertPolicyResultOutput{})
}
