// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains information about a database Threat Detection policy.
// Azure REST API version: 2014-04-01.
type DatabaseThreatDetectionPolicy struct {
	pulumi.CustomResourceState

	// Specifies the semicolon-separated list of alerts that are disabled, or empty string to disable no alerts. Possible values: Sql_Injection; Sql_Injection_Vulnerability; Access_Anomaly; Data_Exfiltration; Unsafe_Action.
	DisabledAlerts pulumi.StringPtrOutput `pulumi:"disabledAlerts"`
	// Specifies that the alert is sent to the account administrators.
	EmailAccountAdmins pulumi.StringPtrOutput `pulumi:"emailAccountAdmins"`
	// Specifies the semicolon-separated list of e-mail addresses to which the alert is sent.
	EmailAddresses pulumi.StringPtrOutput `pulumi:"emailAddresses"`
	// Resource kind.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
	// Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
	State pulumi.StringOutput `pulumi:"state"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. If state is Enabled, storageEndpoint is required.
	StorageEndpoint pulumi.StringPtrOutput `pulumi:"storageEndpoint"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies whether to use the default server policy.
	UseServerDefault pulumi.StringPtrOutput `pulumi:"useServerDefault"`
}

// NewDatabaseThreatDetectionPolicy registers a new resource with the given unique name, arguments, and options.
func NewDatabaseThreatDetectionPolicy(ctx *pulumi.Context,
	name string, args *DatabaseThreatDetectionPolicyArgs, opts ...pulumi.ResourceOption) (*DatabaseThreatDetectionPolicy, error) {
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
			Type: pulumi.String("azure-native:sql/v20140401:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230801preview:DatabaseThreatDetectionPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DatabaseThreatDetectionPolicy
	err := ctx.RegisterResource("azure-native:sql:DatabaseThreatDetectionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseThreatDetectionPolicy gets an existing DatabaseThreatDetectionPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseThreatDetectionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseThreatDetectionPolicyState, opts ...pulumi.ResourceOption) (*DatabaseThreatDetectionPolicy, error) {
	var resource DatabaseThreatDetectionPolicy
	err := ctx.ReadResource("azure-native:sql:DatabaseThreatDetectionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseThreatDetectionPolicy resources.
type databaseThreatDetectionPolicyState struct {
}

type DatabaseThreatDetectionPolicyState struct {
}

func (DatabaseThreatDetectionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseThreatDetectionPolicyState)(nil)).Elem()
}

type databaseThreatDetectionPolicyArgs struct {
	// The name of the database for which database Threat Detection policy is defined.
	DatabaseName string `pulumi:"databaseName"`
	// Specifies the semicolon-separated list of alerts that are disabled, or empty string to disable no alerts. Possible values: Sql_Injection; Sql_Injection_Vulnerability; Access_Anomaly; Data_Exfiltration; Unsafe_Action.
	DisabledAlerts *string `pulumi:"disabledAlerts"`
	// Specifies that the alert is sent to the account administrators.
	EmailAccountAdmins *string `pulumi:"emailAccountAdmins"`
	// Specifies the semicolon-separated list of e-mail addresses to which the alert is sent.
	EmailAddresses *string `pulumi:"emailAddresses"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays *int `pulumi:"retentionDays"`
	// The name of the security alert policy.
	SecurityAlertPolicyName *string `pulumi:"securityAlertPolicyName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
	State string `pulumi:"state"`
	// Specifies the identifier key of the Threat Detection audit storage account. If state is Enabled, storageAccountAccessKey is required.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. If state is Enabled, storageEndpoint is required.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
	// Specifies whether to use the default server policy.
	UseServerDefault *string `pulumi:"useServerDefault"`
}

// The set of arguments for constructing a DatabaseThreatDetectionPolicy resource.
type DatabaseThreatDetectionPolicyArgs struct {
	// The name of the database for which database Threat Detection policy is defined.
	DatabaseName pulumi.StringInput
	// Specifies the semicolon-separated list of alerts that are disabled, or empty string to disable no alerts. Possible values: Sql_Injection; Sql_Injection_Vulnerability; Access_Anomaly; Data_Exfiltration; Unsafe_Action.
	DisabledAlerts pulumi.StringPtrInput
	// Specifies that the alert is sent to the account administrators.
	EmailAccountAdmins pulumi.StringPtrInput
	// Specifies the semicolon-separated list of e-mail addresses to which the alert is sent.
	EmailAddresses pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays pulumi.IntPtrInput
	// The name of the security alert policy.
	SecurityAlertPolicyName pulumi.StringPtrInput
	// The name of the server.
	ServerName pulumi.StringInput
	// Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
	State pulumi.StringInput
	// Specifies the identifier key of the Threat Detection audit storage account. If state is Enabled, storageAccountAccessKey is required.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. If state is Enabled, storageEndpoint is required.
	StorageEndpoint pulumi.StringPtrInput
	// Specifies whether to use the default server policy.
	UseServerDefault pulumi.StringPtrInput
}

func (DatabaseThreatDetectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseThreatDetectionPolicyArgs)(nil)).Elem()
}

type DatabaseThreatDetectionPolicyInput interface {
	pulumi.Input

	ToDatabaseThreatDetectionPolicyOutput() DatabaseThreatDetectionPolicyOutput
	ToDatabaseThreatDetectionPolicyOutputWithContext(ctx context.Context) DatabaseThreatDetectionPolicyOutput
}

func (*DatabaseThreatDetectionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseThreatDetectionPolicy)(nil)).Elem()
}

func (i *DatabaseThreatDetectionPolicy) ToDatabaseThreatDetectionPolicyOutput() DatabaseThreatDetectionPolicyOutput {
	return i.ToDatabaseThreatDetectionPolicyOutputWithContext(context.Background())
}

func (i *DatabaseThreatDetectionPolicy) ToDatabaseThreatDetectionPolicyOutputWithContext(ctx context.Context) DatabaseThreatDetectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseThreatDetectionPolicyOutput)
}

type DatabaseThreatDetectionPolicyOutput struct{ *pulumi.OutputState }

func (DatabaseThreatDetectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseThreatDetectionPolicy)(nil)).Elem()
}

func (o DatabaseThreatDetectionPolicyOutput) ToDatabaseThreatDetectionPolicyOutput() DatabaseThreatDetectionPolicyOutput {
	return o
}

func (o DatabaseThreatDetectionPolicyOutput) ToDatabaseThreatDetectionPolicyOutputWithContext(ctx context.Context) DatabaseThreatDetectionPolicyOutput {
	return o
}

// Specifies the semicolon-separated list of alerts that are disabled, or empty string to disable no alerts. Possible values: Sql_Injection; Sql_Injection_Vulnerability; Access_Anomaly; Data_Exfiltration; Unsafe_Action.
func (o DatabaseThreatDetectionPolicyOutput) DisabledAlerts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseThreatDetectionPolicy) pulumi.StringPtrOutput { return v.DisabledAlerts }).(pulumi.StringPtrOutput)
}

// Specifies that the alert is sent to the account administrators.
func (o DatabaseThreatDetectionPolicyOutput) EmailAccountAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseThreatDetectionPolicy) pulumi.StringPtrOutput { return v.EmailAccountAdmins }).(pulumi.StringPtrOutput)
}

// Specifies the semicolon-separated list of e-mail addresses to which the alert is sent.
func (o DatabaseThreatDetectionPolicyOutput) EmailAddresses() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseThreatDetectionPolicy) pulumi.StringPtrOutput { return v.EmailAddresses }).(pulumi.StringPtrOutput)
}

// Resource kind.
func (o DatabaseThreatDetectionPolicyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseThreatDetectionPolicy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o DatabaseThreatDetectionPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseThreatDetectionPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o DatabaseThreatDetectionPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseThreatDetectionPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the number of days to keep in the Threat Detection audit logs.
func (o DatabaseThreatDetectionPolicyOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DatabaseThreatDetectionPolicy) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
func (o DatabaseThreatDetectionPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseThreatDetectionPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. If state is Enabled, storageEndpoint is required.
func (o DatabaseThreatDetectionPolicyOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseThreatDetectionPolicy) pulumi.StringPtrOutput { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o DatabaseThreatDetectionPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseThreatDetectionPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies whether to use the default server policy.
func (o DatabaseThreatDetectionPolicyOutput) UseServerDefault() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseThreatDetectionPolicy) pulumi.StringPtrOutput { return v.UseServerDefault }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseThreatDetectionPolicyOutput{})
}
