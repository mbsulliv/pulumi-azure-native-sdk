// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datalakeanalytics

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A Data Lake Analytics account object, containing all information associated with the named Data Lake Analytics account.
// Azure REST API version: 2019-11-01-preview. Prior API version in Azure Native 1.x: 2016-11-01
type Account struct {
	pulumi.CustomResourceState

	// The unique identifier associated with this Data Lake Analytics account.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The list of compute policies associated with this account.
	ComputePolicies ComputePolicyResponseArrayOutput `pulumi:"computePolicies"`
	// The account creation time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The commitment tier in use for the current month.
	CurrentTier pulumi.StringOutput `pulumi:"currentTier"`
	// The list of Data Lake Store accounts associated with this account.
	DataLakeStoreAccounts DataLakeStoreAccountInformationResponseArrayOutput `pulumi:"dataLakeStoreAccounts"`
	// The current state of the DebugDataAccessLevel for this account.
	DebugDataAccessLevel pulumi.StringOutput `pulumi:"debugDataAccessLevel"`
	// The default Data Lake Store account associated with this account.
	DefaultDataLakeStoreAccount pulumi.StringOutput `pulumi:"defaultDataLakeStoreAccount"`
	// The type of the default Data Lake Store account associated with this account.
	DefaultDataLakeStoreAccountType pulumi.StringOutput `pulumi:"defaultDataLakeStoreAccountType"`
	// The full CName endpoint for this account.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
	FirewallAllowAzureIps pulumi.StringPtrOutput `pulumi:"firewallAllowAzureIps"`
	// The list of firewall rules associated with this account.
	FirewallRules FirewallRuleResponseArrayOutput `pulumi:"firewallRules"`
	// The current state of the IP address firewall for this account.
	FirewallState pulumi.StringPtrOutput `pulumi:"firewallState"`
	// The list of hiveMetastores associated with this account.
	HiveMetastores HiveMetastoreResponseArrayOutput `pulumi:"hiveMetastores"`
	// The account last modified time.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The maximum supported active jobs under the account at the same time.
	MaxActiveJobCountPerUser pulumi.IntOutput `pulumi:"maxActiveJobCountPerUser"`
	// The maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism pulumi.IntPtrOutput `pulumi:"maxDegreeOfParallelism"`
	// The maximum supported degree of parallelism per job for this account.
	MaxDegreeOfParallelismPerJob pulumi.IntPtrOutput `pulumi:"maxDegreeOfParallelismPerJob"`
	// The maximum supported jobs running under the account at the same time.
	MaxJobCount pulumi.IntPtrOutput `pulumi:"maxJobCount"`
	// The maximum supported active jobs under the account at the same time.
	MaxJobRunningTimeInMin pulumi.IntOutput `pulumi:"maxJobRunningTimeInMin"`
	// The maximum supported jobs queued under the account at the same time.
	MaxQueuedJobCountPerUser pulumi.IntOutput `pulumi:"maxQueuedJobCountPerUser"`
	// The minimum supported priority per job for this account.
	MinPriorityPerJob pulumi.IntOutput `pulumi:"minPriorityPerJob"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The commitment tier for the next month.
	NewTier pulumi.StringPtrOutput `pulumi:"newTier"`
	// The provisioning status of the Data Lake Analytics account.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The list of Data Lake Store accounts associated with this account.
	PublicDataLakeStoreAccounts DataLakeStoreAccountInformationResponseArrayOutput `pulumi:"publicDataLakeStoreAccounts"`
	// The number of days that job metadata is retained.
	QueryStoreRetention pulumi.IntPtrOutput `pulumi:"queryStoreRetention"`
	// The state of the Data Lake Analytics account.
	State pulumi.StringOutput `pulumi:"state"`
	// The list of Azure Blob Storage accounts associated with this account.
	StorageAccounts StorageAccountInformationResponseArrayOutput `pulumi:"storageAccounts"`
	// The system defined maximum supported degree of parallelism for this account, which restricts the maximum value of parallelism the user can set for the account.
	SystemMaxDegreeOfParallelism pulumi.IntOutput `pulumi:"systemMaxDegreeOfParallelism"`
	// The system defined maximum supported jobs running under the account at the same time, which restricts the maximum number of running jobs the user can set for the account.
	SystemMaxJobCount pulumi.IntOutput `pulumi:"systemMaxJobCount"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of virtualNetwork rules associated with this account.
	VirtualNetworkRules VirtualNetworkRuleResponseArrayOutput `pulumi:"virtualNetworkRules"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataLakeStoreAccounts == nil {
		return nil, errors.New("invalid value for required argument 'DataLakeStoreAccounts'")
	}
	if args.DefaultDataLakeStoreAccount == nil {
		return nil, errors.New("invalid value for required argument 'DefaultDataLakeStoreAccount'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.FirewallAllowAzureIps == nil {
		args.FirewallAllowAzureIps = FirewallAllowAzureIpsState("Disabled")
	}
	if args.FirewallState == nil {
		args.FirewallState = FirewallState("Disabled")
	}
	if args.MaxDegreeOfParallelism == nil {
		args.MaxDegreeOfParallelism = pulumi.IntPtr(30)
	}
	if args.MaxDegreeOfParallelismPerJob == nil {
		args.MaxDegreeOfParallelismPerJob = pulumi.IntPtr(32)
	}
	if args.MaxJobCount == nil {
		args.MaxJobCount = pulumi.IntPtr(3)
	}
	if args.NewTier == nil {
		args.NewTier = TierType("Consumption")
	}
	if args.QueryStoreRetention == nil {
		args.QueryStoreRetention = pulumi.IntPtr(30)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20151001preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20161101:Account"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20191101preview:Account"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Account
	err := ctx.RegisterResource("azure-native:datalakeanalytics:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:datalakeanalytics:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
}

type AccountState struct {
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName *string `pulumi:"accountName"`
	// The list of compute policies associated with this account.
	ComputePolicies []CreateComputePolicyWithAccountParameters `pulumi:"computePolicies"`
	// The list of Data Lake Store accounts associated with this account.
	DataLakeStoreAccounts []AddDataLakeStoreWithAccountParameters `pulumi:"dataLakeStoreAccounts"`
	// The default Data Lake Store account associated with this account.
	DefaultDataLakeStoreAccount string `pulumi:"defaultDataLakeStoreAccount"`
	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
	FirewallAllowAzureIps *FirewallAllowAzureIpsState `pulumi:"firewallAllowAzureIps"`
	// The list of firewall rules associated with this account.
	FirewallRules []CreateFirewallRuleWithAccountParameters `pulumi:"firewallRules"`
	// The current state of the IP address firewall for this account.
	FirewallState *FirewallState `pulumi:"firewallState"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism *int `pulumi:"maxDegreeOfParallelism"`
	// The maximum supported degree of parallelism per job for this account.
	MaxDegreeOfParallelismPerJob *int `pulumi:"maxDegreeOfParallelismPerJob"`
	// The maximum supported jobs running under the account at the same time.
	MaxJobCount *int `pulumi:"maxJobCount"`
	// The minimum supported priority per job for this account.
	MinPriorityPerJob *int `pulumi:"minPriorityPerJob"`
	// The commitment tier for the next month.
	NewTier *TierType `pulumi:"newTier"`
	// The number of days that job metadata is retained.
	QueryStoreRetention *int `pulumi:"queryStoreRetention"`
	// The name of the Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of Azure Blob Storage accounts associated with this account.
	StorageAccounts []AddStorageAccountWithAccountParameters `pulumi:"storageAccounts"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName pulumi.StringPtrInput
	// The list of compute policies associated with this account.
	ComputePolicies CreateComputePolicyWithAccountParametersArrayInput
	// The list of Data Lake Store accounts associated with this account.
	DataLakeStoreAccounts AddDataLakeStoreWithAccountParametersArrayInput
	// The default Data Lake Store account associated with this account.
	DefaultDataLakeStoreAccount pulumi.StringInput
	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
	FirewallAllowAzureIps FirewallAllowAzureIpsStatePtrInput
	// The list of firewall rules associated with this account.
	FirewallRules CreateFirewallRuleWithAccountParametersArrayInput
	// The current state of the IP address firewall for this account.
	FirewallState FirewallStatePtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism pulumi.IntPtrInput
	// The maximum supported degree of parallelism per job for this account.
	MaxDegreeOfParallelismPerJob pulumi.IntPtrInput
	// The maximum supported jobs running under the account at the same time.
	MaxJobCount pulumi.IntPtrInput
	// The minimum supported priority per job for this account.
	MinPriorityPerJob pulumi.IntPtrInput
	// The commitment tier for the next month.
	NewTier TierTypePtrInput
	// The number of days that job metadata is retained.
	QueryStoreRetention pulumi.IntPtrInput
	// The name of the Azure resource group.
	ResourceGroupName pulumi.StringInput
	// The list of Azure Blob Storage accounts associated with this account.
	StorageAccounts AddStorageAccountWithAccountParametersArrayInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

func (i *Account) ToOutput(ctx context.Context) pulumix.Output[*Account] {
	return pulumix.Output[*Account]{
		OutputState: i.ToAccountOutputWithContext(ctx).OutputState,
	}
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func (o AccountOutput) ToOutput(ctx context.Context) pulumix.Output[*Account] {
	return pulumix.Output[*Account]{
		OutputState: o.OutputState,
	}
}

// The unique identifier associated with this Data Lake Analytics account.
func (o AccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The list of compute policies associated with this account.
func (o AccountOutput) ComputePolicies() ComputePolicyResponseArrayOutput {
	return o.ApplyT(func(v *Account) ComputePolicyResponseArrayOutput { return v.ComputePolicies }).(ComputePolicyResponseArrayOutput)
}

// The account creation time.
func (o AccountOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The commitment tier in use for the current month.
func (o AccountOutput) CurrentTier() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CurrentTier }).(pulumi.StringOutput)
}

// The list of Data Lake Store accounts associated with this account.
func (o AccountOutput) DataLakeStoreAccounts() DataLakeStoreAccountInformationResponseArrayOutput {
	return o.ApplyT(func(v *Account) DataLakeStoreAccountInformationResponseArrayOutput { return v.DataLakeStoreAccounts }).(DataLakeStoreAccountInformationResponseArrayOutput)
}

// The current state of the DebugDataAccessLevel for this account.
func (o AccountOutput) DebugDataAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.DebugDataAccessLevel }).(pulumi.StringOutput)
}

// The default Data Lake Store account associated with this account.
func (o AccountOutput) DefaultDataLakeStoreAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.DefaultDataLakeStoreAccount }).(pulumi.StringOutput)
}

// The type of the default Data Lake Store account associated with this account.
func (o AccountOutput) DefaultDataLakeStoreAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.DefaultDataLakeStoreAccountType }).(pulumi.StringOutput)
}

// The full CName endpoint for this account.
func (o AccountOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
func (o AccountOutput) FirewallAllowAzureIps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.FirewallAllowAzureIps }).(pulumi.StringPtrOutput)
}

// The list of firewall rules associated with this account.
func (o AccountOutput) FirewallRules() FirewallRuleResponseArrayOutput {
	return o.ApplyT(func(v *Account) FirewallRuleResponseArrayOutput { return v.FirewallRules }).(FirewallRuleResponseArrayOutput)
}

// The current state of the IP address firewall for this account.
func (o AccountOutput) FirewallState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.FirewallState }).(pulumi.StringPtrOutput)
}

// The list of hiveMetastores associated with this account.
func (o AccountOutput) HiveMetastores() HiveMetastoreResponseArrayOutput {
	return o.ApplyT(func(v *Account) HiveMetastoreResponseArrayOutput { return v.HiveMetastores }).(HiveMetastoreResponseArrayOutput)
}

// The account last modified time.
func (o AccountOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The resource location.
func (o AccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The maximum supported active jobs under the account at the same time.
func (o AccountOutput) MaxActiveJobCountPerUser() pulumi.IntOutput {
	return o.ApplyT(func(v *Account) pulumi.IntOutput { return v.MaxActiveJobCountPerUser }).(pulumi.IntOutput)
}

// The maximum supported degree of parallelism for this account.
func (o AccountOutput) MaxDegreeOfParallelism() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.IntPtrOutput { return v.MaxDegreeOfParallelism }).(pulumi.IntPtrOutput)
}

// The maximum supported degree of parallelism per job for this account.
func (o AccountOutput) MaxDegreeOfParallelismPerJob() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.IntPtrOutput { return v.MaxDegreeOfParallelismPerJob }).(pulumi.IntPtrOutput)
}

// The maximum supported jobs running under the account at the same time.
func (o AccountOutput) MaxJobCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.IntPtrOutput { return v.MaxJobCount }).(pulumi.IntPtrOutput)
}

// The maximum supported active jobs under the account at the same time.
func (o AccountOutput) MaxJobRunningTimeInMin() pulumi.IntOutput {
	return o.ApplyT(func(v *Account) pulumi.IntOutput { return v.MaxJobRunningTimeInMin }).(pulumi.IntOutput)
}

// The maximum supported jobs queued under the account at the same time.
func (o AccountOutput) MaxQueuedJobCountPerUser() pulumi.IntOutput {
	return o.ApplyT(func(v *Account) pulumi.IntOutput { return v.MaxQueuedJobCountPerUser }).(pulumi.IntOutput)
}

// The minimum supported priority per job for this account.
func (o AccountOutput) MinPriorityPerJob() pulumi.IntOutput {
	return o.ApplyT(func(v *Account) pulumi.IntOutput { return v.MinPriorityPerJob }).(pulumi.IntOutput)
}

// The resource name.
func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The commitment tier for the next month.
func (o AccountOutput) NewTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.NewTier }).(pulumi.StringPtrOutput)
}

// The provisioning status of the Data Lake Analytics account.
func (o AccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The list of Data Lake Store accounts associated with this account.
func (o AccountOutput) PublicDataLakeStoreAccounts() DataLakeStoreAccountInformationResponseArrayOutput {
	return o.ApplyT(func(v *Account) DataLakeStoreAccountInformationResponseArrayOutput {
		return v.PublicDataLakeStoreAccounts
	}).(DataLakeStoreAccountInformationResponseArrayOutput)
}

// The number of days that job metadata is retained.
func (o AccountOutput) QueryStoreRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.IntPtrOutput { return v.QueryStoreRetention }).(pulumi.IntPtrOutput)
}

// The state of the Data Lake Analytics account.
func (o AccountOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The list of Azure Blob Storage accounts associated with this account.
func (o AccountOutput) StorageAccounts() StorageAccountInformationResponseArrayOutput {
	return o.ApplyT(func(v *Account) StorageAccountInformationResponseArrayOutput { return v.StorageAccounts }).(StorageAccountInformationResponseArrayOutput)
}

// The system defined maximum supported degree of parallelism for this account, which restricts the maximum value of parallelism the user can set for the account.
func (o AccountOutput) SystemMaxDegreeOfParallelism() pulumi.IntOutput {
	return o.ApplyT(func(v *Account) pulumi.IntOutput { return v.SystemMaxDegreeOfParallelism }).(pulumi.IntOutput)
}

// The system defined maximum supported jobs running under the account at the same time, which restricts the maximum number of running jobs the user can set for the account.
func (o AccountOutput) SystemMaxJobCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Account) pulumi.IntOutput { return v.SystemMaxJobCount }).(pulumi.IntOutput)
}

// The resource tags.
func (o AccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Account) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The list of virtualNetwork rules associated with this account.
func (o AccountOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *Account) VirtualNetworkRuleResponseArrayOutput { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
