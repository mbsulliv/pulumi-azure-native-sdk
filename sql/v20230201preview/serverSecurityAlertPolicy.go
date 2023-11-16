// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A server security alert policy.
type ServerSecurityAlertPolicy struct {
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

// NewServerSecurityAlertPolicy registers a new resource with the given unique name, arguments, and options.
func NewServerSecurityAlertPolicy(ctx *pulumi.Context,
	name string, args *ServerSecurityAlertPolicyArgs, opts ...pulumi.ResourceOption) (*ServerSecurityAlertPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
			Type: pulumi.String("azure-native:sql:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:ServerSecurityAlertPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServerSecurityAlertPolicy
	err := ctx.RegisterResource("azure-native:sql/v20230201preview:ServerSecurityAlertPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerSecurityAlertPolicy gets an existing ServerSecurityAlertPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerSecurityAlertPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerSecurityAlertPolicyState, opts ...pulumi.ResourceOption) (*ServerSecurityAlertPolicy, error) {
	var resource ServerSecurityAlertPolicy
	err := ctx.ReadResource("azure-native:sql/v20230201preview:ServerSecurityAlertPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerSecurityAlertPolicy resources.
type serverSecurityAlertPolicyState struct {
}

type ServerSecurityAlertPolicyState struct {
}

func (ServerSecurityAlertPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverSecurityAlertPolicyState)(nil)).Elem()
}

type serverSecurityAlertPolicyArgs struct {
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
	// The name of the threat detection policy.
	SecurityAlertPolicyName *string `pulumi:"securityAlertPolicyName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
	State SecurityAlertsPolicyState `pulumi:"state"`
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
}

// The set of arguments for constructing a ServerSecurityAlertPolicy resource.
type ServerSecurityAlertPolicyArgs struct {
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
	// The name of the threat detection policy.
	SecurityAlertPolicyName pulumi.StringPtrInput
	// The name of the server.
	ServerName pulumi.StringInput
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
	State SecurityAlertsPolicyStateInput
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint pulumi.StringPtrInput
}

func (ServerSecurityAlertPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverSecurityAlertPolicyArgs)(nil)).Elem()
}

type ServerSecurityAlertPolicyInput interface {
	pulumi.Input

	ToServerSecurityAlertPolicyOutput() ServerSecurityAlertPolicyOutput
	ToServerSecurityAlertPolicyOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyOutput
}

func (*ServerSecurityAlertPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerSecurityAlertPolicy)(nil)).Elem()
}

func (i *ServerSecurityAlertPolicy) ToServerSecurityAlertPolicyOutput() ServerSecurityAlertPolicyOutput {
	return i.ToServerSecurityAlertPolicyOutputWithContext(context.Background())
}

func (i *ServerSecurityAlertPolicy) ToServerSecurityAlertPolicyOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerSecurityAlertPolicyOutput)
}

type ServerSecurityAlertPolicyOutput struct{ *pulumi.OutputState }

func (ServerSecurityAlertPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerSecurityAlertPolicy)(nil)).Elem()
}

func (o ServerSecurityAlertPolicyOutput) ToServerSecurityAlertPolicyOutput() ServerSecurityAlertPolicyOutput {
	return o
}

func (o ServerSecurityAlertPolicyOutput) ToServerSecurityAlertPolicyOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyOutput {
	return o
}

// Specifies the UTC creation time of the policy.
func (o ServerSecurityAlertPolicyOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action, Brute_Force
func (o ServerSecurityAlertPolicyOutput) DisabledAlerts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringArrayOutput { return v.DisabledAlerts }).(pulumi.StringArrayOutput)
}

// Specifies that the alert is sent to the account administrators.
func (o ServerSecurityAlertPolicyOutput) EmailAccountAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.BoolPtrOutput { return v.EmailAccountAdmins }).(pulumi.BoolPtrOutput)
}

// Specifies an array of e-mail addresses to which the alert is sent.
func (o ServerSecurityAlertPolicyOutput) EmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringArrayOutput { return v.EmailAddresses }).(pulumi.StringArrayOutput)
}

// Resource name.
func (o ServerSecurityAlertPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the number of days to keep in the Threat Detection audit logs.
func (o ServerSecurityAlertPolicyOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
func (o ServerSecurityAlertPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Specifies the identifier key of the Threat Detection audit storage account.
func (o ServerSecurityAlertPolicyOutput) StorageAccountAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringPtrOutput { return v.StorageAccountAccessKey }).(pulumi.StringPtrOutput)
}

// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
func (o ServerSecurityAlertPolicyOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringPtrOutput { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

// SystemData of SecurityAlertPolicyResource.
func (o ServerSecurityAlertPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o ServerSecurityAlertPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerSecurityAlertPolicyOutput{})
}
