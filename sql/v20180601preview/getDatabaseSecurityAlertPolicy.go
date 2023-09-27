// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a  database's security alert policy.
func LookupDatabaseSecurityAlertPolicy(ctx *pulumi.Context, args *LookupDatabaseSecurityAlertPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseSecurityAlertPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabaseSecurityAlertPolicyResult
	err := ctx.Invoke("azure-native:sql/v20180601preview:getDatabaseSecurityAlertPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseSecurityAlertPolicyArgs struct {
	// The name of the  database for which the security alert policy is defined.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the security alert policy.
	SecurityAlertPolicyName string `pulumi:"securityAlertPolicyName"`
	// The name of the  server.
	ServerName string `pulumi:"serverName"`
}

// A database security alert policy.
type LookupDatabaseSecurityAlertPolicyResult struct {
	// Specifies the UTC creation time of the policy.
	CreationTime string `pulumi:"creationTime"`
	// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action
	DisabledAlerts []string `pulumi:"disabledAlerts"`
	// Specifies that the alert is sent to the account administrators.
	EmailAccountAdmins *bool `pulumi:"emailAccountAdmins"`
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses []string `pulumi:"emailAddresses"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays *int `pulumi:"retentionDays"`
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
	State string `pulumi:"state"`
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupDatabaseSecurityAlertPolicyOutput(ctx *pulumi.Context, args LookupDatabaseSecurityAlertPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseSecurityAlertPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseSecurityAlertPolicyResult, error) {
			args := v.(LookupDatabaseSecurityAlertPolicyArgs)
			r, err := LookupDatabaseSecurityAlertPolicy(ctx, &args, opts...)
			var s LookupDatabaseSecurityAlertPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseSecurityAlertPolicyResultOutput)
}

type LookupDatabaseSecurityAlertPolicyOutputArgs struct {
	// The name of the  database for which the security alert policy is defined.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the security alert policy.
	SecurityAlertPolicyName pulumi.StringInput `pulumi:"securityAlertPolicyName"`
	// The name of the  server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupDatabaseSecurityAlertPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseSecurityAlertPolicyArgs)(nil)).Elem()
}

// A database security alert policy.
type LookupDatabaseSecurityAlertPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseSecurityAlertPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseSecurityAlertPolicyResult)(nil)).Elem()
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) ToLookupDatabaseSecurityAlertPolicyResultOutput() LookupDatabaseSecurityAlertPolicyResultOutput {
	return o
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) ToLookupDatabaseSecurityAlertPolicyResultOutputWithContext(ctx context.Context) LookupDatabaseSecurityAlertPolicyResultOutput {
	return o
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDatabaseSecurityAlertPolicyResult] {
	return pulumix.Output[LookupDatabaseSecurityAlertPolicyResult]{
		OutputState: o.OutputState,
	}
}

// Specifies the UTC creation time of the policy.
func (o LookupDatabaseSecurityAlertPolicyResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action
func (o LookupDatabaseSecurityAlertPolicyResultOutput) DisabledAlerts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) []string { return v.DisabledAlerts }).(pulumi.StringArrayOutput)
}

// Specifies that the alert is sent to the account administrators.
func (o LookupDatabaseSecurityAlertPolicyResultOutput) EmailAccountAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) *bool { return v.EmailAccountAdmins }).(pulumi.BoolPtrOutput)
}

// Specifies an array of e-mail addresses to which the alert is sent.
func (o LookupDatabaseSecurityAlertPolicyResultOutput) EmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) []string { return v.EmailAddresses }).(pulumi.StringArrayOutput)
}

// Resource ID.
func (o LookupDatabaseSecurityAlertPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupDatabaseSecurityAlertPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the number of days to keep in the Threat Detection audit logs.
func (o LookupDatabaseSecurityAlertPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
func (o LookupDatabaseSecurityAlertPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

// Specifies the identifier key of the Threat Detection audit storage account.
func (o LookupDatabaseSecurityAlertPolicyResultOutput) StorageAccountAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) *string { return v.StorageAccountAccessKey }).(pulumi.StringPtrOutput)
}

// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
func (o LookupDatabaseSecurityAlertPolicyResultOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) *string { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupDatabaseSecurityAlertPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseSecurityAlertPolicyResultOutput{})
}
