// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A database security alert policy.
type DatabaseSecurityAlertPolicy struct {
	pulumi.CustomResourceState

	// Specifies the UTC creation time of the policy.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action, Brute_Force
	DisabledAlerts pulumi.StringArrayOutput `pulumi:"disabledAlerts"`
	// Specifies that the alert is sent to the account administrators.
	EmailAccountAdmins pulumi.BoolPtrOutput `pulumi:"emailAccountAdmins"`
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses pulumi.StringArrayOutput `pulumi:"emailAddresses"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
	State pulumi.StringOutput `pulumi:"state"`
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey pulumi.StringPtrOutput `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint pulumi.StringPtrOutput `pulumi:"storageEndpoint"`
	// SystemData of SecurityAlertPolicyResource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabaseSecurityAlertPolicy registers a new resource with the given unique name, arguments, and options.
func NewDatabaseSecurityAlertPolicy(ctx *pulumi.Context,
	name string, args *DatabaseSecurityAlertPolicyArgs, opts ...pulumi.ResourceOption) (*DatabaseSecurityAlertPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:DatabaseSecurityAlertPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DatabaseSecurityAlertPolicy
	err := ctx.RegisterResource("azure-native:sql/v20221101preview:DatabaseSecurityAlertPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseSecurityAlertPolicy gets an existing DatabaseSecurityAlertPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseSecurityAlertPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseSecurityAlertPolicyState, opts ...pulumi.ResourceOption) (*DatabaseSecurityAlertPolicy, error) {
	var resource DatabaseSecurityAlertPolicy
	err := ctx.ReadResource("azure-native:sql/v20221101preview:DatabaseSecurityAlertPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseSecurityAlertPolicy resources.
type databaseSecurityAlertPolicyState struct {
}

type DatabaseSecurityAlertPolicyState struct {
}

func (DatabaseSecurityAlertPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseSecurityAlertPolicyState)(nil)).Elem()
}

type databaseSecurityAlertPolicyArgs struct {
	// The name of the  database for which the security alert policy is defined.
	DatabaseName string `pulumi:"databaseName"`
	// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action, Brute_Force
	DisabledAlerts []string `pulumi:"disabledAlerts"`
	// Specifies that the alert is sent to the account administrators.
	EmailAccountAdmins *bool `pulumi:"emailAccountAdmins"`
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses []string `pulumi:"emailAddresses"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays *int `pulumi:"retentionDays"`
	// The name of the security alert policy.
	SecurityAlertPolicyName *string `pulumi:"securityAlertPolicyName"`
	// The name of the  server.
	ServerName string `pulumi:"serverName"`
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
	State SecurityAlertsPolicyState `pulumi:"state"`
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
}

// The set of arguments for constructing a DatabaseSecurityAlertPolicy resource.
type DatabaseSecurityAlertPolicyArgs struct {
	// The name of the  database for which the security alert policy is defined.
	DatabaseName pulumi.StringInput
	// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action, Brute_Force
	DisabledAlerts pulumi.StringArrayInput
	// Specifies that the alert is sent to the account administrators.
	EmailAccountAdmins pulumi.BoolPtrInput
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses pulumi.StringArrayInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays pulumi.IntPtrInput
	// The name of the security alert policy.
	SecurityAlertPolicyName pulumi.StringPtrInput
	// The name of the  server.
	ServerName pulumi.StringInput
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
	State SecurityAlertsPolicyStateInput
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint pulumi.StringPtrInput
}

func (DatabaseSecurityAlertPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseSecurityAlertPolicyArgs)(nil)).Elem()
}

type DatabaseSecurityAlertPolicyInput interface {
	pulumi.Input

	ToDatabaseSecurityAlertPolicyOutput() DatabaseSecurityAlertPolicyOutput
	ToDatabaseSecurityAlertPolicyOutputWithContext(ctx context.Context) DatabaseSecurityAlertPolicyOutput
}

func (*DatabaseSecurityAlertPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseSecurityAlertPolicy)(nil)).Elem()
}

func (i *DatabaseSecurityAlertPolicy) ToDatabaseSecurityAlertPolicyOutput() DatabaseSecurityAlertPolicyOutput {
	return i.ToDatabaseSecurityAlertPolicyOutputWithContext(context.Background())
}

func (i *DatabaseSecurityAlertPolicy) ToDatabaseSecurityAlertPolicyOutputWithContext(ctx context.Context) DatabaseSecurityAlertPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseSecurityAlertPolicyOutput)
}

type DatabaseSecurityAlertPolicyOutput struct{ *pulumi.OutputState }

func (DatabaseSecurityAlertPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseSecurityAlertPolicy)(nil)).Elem()
}

func (o DatabaseSecurityAlertPolicyOutput) ToDatabaseSecurityAlertPolicyOutput() DatabaseSecurityAlertPolicyOutput {
	return o
}

func (o DatabaseSecurityAlertPolicyOutput) ToDatabaseSecurityAlertPolicyOutputWithContext(ctx context.Context) DatabaseSecurityAlertPolicyOutput {
	return o
}

// Specifies the UTC creation time of the policy.
func (o DatabaseSecurityAlertPolicyOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action, Brute_Force
func (o DatabaseSecurityAlertPolicyOutput) DisabledAlerts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringArrayOutput { return v.DisabledAlerts }).(pulumi.StringArrayOutput)
}

// Specifies that the alert is sent to the account administrators.
func (o DatabaseSecurityAlertPolicyOutput) EmailAccountAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.BoolPtrOutput { return v.EmailAccountAdmins }).(pulumi.BoolPtrOutput)
}

// Specifies an array of e-mail addresses to which the alert is sent.
func (o DatabaseSecurityAlertPolicyOutput) EmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringArrayOutput { return v.EmailAddresses }).(pulumi.StringArrayOutput)
}

// Resource name.
func (o DatabaseSecurityAlertPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the number of days to keep in the Threat Detection audit logs.
func (o DatabaseSecurityAlertPolicyOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
func (o DatabaseSecurityAlertPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Specifies the identifier key of the Threat Detection audit storage account.
func (o DatabaseSecurityAlertPolicyOutput) StorageAccountAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringPtrOutput { return v.StorageAccountAccessKey }).(pulumi.StringPtrOutput)
}

// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
func (o DatabaseSecurityAlertPolicyOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringPtrOutput { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

// SystemData of SecurityAlertPolicyResource.
func (o DatabaseSecurityAlertPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o DatabaseSecurityAlertPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseSecurityAlertPolicyOutput{})
}
