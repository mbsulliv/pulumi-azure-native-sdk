// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Cluster details.
type Cluster struct {
	pulumi.CustomResourceState

	// Object id of cluster AAD identity.
	AadApplicationObjectId pulumi.StringPtrOutput `pulumi:"aadApplicationObjectId"`
	// App id of cluster AAD identity.
	AadClientId pulumi.StringPtrOutput `pulumi:"aadClientId"`
	// Id of cluster identity service principal.
	AadServicePrincipalObjectId pulumi.StringPtrOutput `pulumi:"aadServicePrincipalObjectId"`
	// Tenant id of cluster AAD identity.
	AadTenantId pulumi.StringPtrOutput `pulumi:"aadTenantId"`
	// Type of billing applied to the resource.
	BillingModel pulumi.StringOutput `pulumi:"billingModel"`
	// Unique, immutable resource id.
	CloudId pulumi.StringOutput `pulumi:"cloudId"`
	// Endpoint configured for management from the Azure portal.
	CloudManagementEndpoint pulumi.StringPtrOutput `pulumi:"cloudManagementEndpoint"`
	// Desired properties of the cluster.
	DesiredProperties ClusterDesiredPropertiesResponsePtrOutput `pulumi:"desiredProperties"`
	// Most recent billing meter timestamp.
	LastBillingTimestamp pulumi.StringOutput `pulumi:"lastBillingTimestamp"`
	// Most recent cluster sync timestamp.
	LastSyncTimestamp pulumi.StringOutput `pulumi:"lastSyncTimestamp"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// Provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// First cluster sync timestamp.
	RegistrationTimestamp pulumi.StringOutput `pulumi:"registrationTimestamp"`
	// Properties reported by cluster agent.
	ReportedProperties ClusterReportedPropertiesResponseOutput `pulumi:"reportedProperties"`
	// Object id of RP Service Principal
	ResourceProviderObjectId pulumi.StringOutput `pulumi:"resourceProviderObjectId"`
	// Region specific DataPath Endpoint of the cluster.
	ServiceEndpoint pulumi.StringOutput `pulumi:"serviceEndpoint"`
	// Software Assurance properties of the cluster.
	SoftwareAssuranceProperties SoftwareAssurancePropertiesResponsePtrOutput `pulumi:"softwareAssuranceProperties"`
	// Status of the cluster agent.
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Number of days remaining in the trial period.
	TrialDaysRemaining pulumi.Float64Output `pulumi:"trialDaysRemaining"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
	UserAssignedIdentities UserAssignedIdentityResponseMapOutput `pulumi:"userAssignedIdentities"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20200301preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20201001:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210101preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220101:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220301:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220501:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220901:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221001:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221201:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221215preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230201:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230601:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230801:Cluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:azurestackhci/v20230301:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:azurestackhci/v20230301:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Object id of cluster AAD identity.
	AadApplicationObjectId *string `pulumi:"aadApplicationObjectId"`
	// App id of cluster AAD identity.
	AadClientId *string `pulumi:"aadClientId"`
	// Id of cluster identity service principal.
	AadServicePrincipalObjectId *string `pulumi:"aadServicePrincipalObjectId"`
	// Tenant id of cluster AAD identity.
	AadTenantId *string `pulumi:"aadTenantId"`
	// Endpoint configured for management from the Azure portal.
	CloudManagementEndpoint *string `pulumi:"cloudManagementEndpoint"`
	// The name of the cluster.
	ClusterName *string `pulumi:"clusterName"`
	// Desired properties of the cluster.
	DesiredProperties *ClusterDesiredProperties `pulumi:"desiredProperties"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Software Assurance properties of the cluster.
	SoftwareAssuranceProperties *SoftwareAssuranceProperties `pulumi:"softwareAssuranceProperties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type string `pulumi:"type"`
	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
	UserAssignedIdentities []string `pulumi:"userAssignedIdentities"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Object id of cluster AAD identity.
	AadApplicationObjectId pulumi.StringPtrInput
	// App id of cluster AAD identity.
	AadClientId pulumi.StringPtrInput
	// Id of cluster identity service principal.
	AadServicePrincipalObjectId pulumi.StringPtrInput
	// Tenant id of cluster AAD identity.
	AadTenantId pulumi.StringPtrInput
	// Endpoint configured for management from the Azure portal.
	CloudManagementEndpoint pulumi.StringPtrInput
	// The name of the cluster.
	ClusterName pulumi.StringPtrInput
	// Desired properties of the cluster.
	DesiredProperties ClusterDesiredPropertiesPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Software Assurance properties of the cluster.
	SoftwareAssuranceProperties SoftwareAssurancePropertiesPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type pulumi.StringInput
	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
	UserAssignedIdentities pulumi.StringArrayInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

func (i *Cluster) ToOutput(ctx context.Context) pulumix.Output[*Cluster] {
	return pulumix.Output[*Cluster]{
		OutputState: i.ToClusterOutputWithContext(ctx).OutputState,
	}
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) ToOutput(ctx context.Context) pulumix.Output[*Cluster] {
	return pulumix.Output[*Cluster]{
		OutputState: o.OutputState,
	}
}

// Object id of cluster AAD identity.
func (o ClusterOutput) AadApplicationObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.AadApplicationObjectId }).(pulumi.StringPtrOutput)
}

// App id of cluster AAD identity.
func (o ClusterOutput) AadClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.AadClientId }).(pulumi.StringPtrOutput)
}

// Id of cluster identity service principal.
func (o ClusterOutput) AadServicePrincipalObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.AadServicePrincipalObjectId }).(pulumi.StringPtrOutput)
}

// Tenant id of cluster AAD identity.
func (o ClusterOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.AadTenantId }).(pulumi.StringPtrOutput)
}

// Type of billing applied to the resource.
func (o ClusterOutput) BillingModel() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.BillingModel }).(pulumi.StringOutput)
}

// Unique, immutable resource id.
func (o ClusterOutput) CloudId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CloudId }).(pulumi.StringOutput)
}

// Endpoint configured for management from the Azure portal.
func (o ClusterOutput) CloudManagementEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.CloudManagementEndpoint }).(pulumi.StringPtrOutput)
}

// Desired properties of the cluster.
func (o ClusterOutput) DesiredProperties() ClusterDesiredPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterDesiredPropertiesResponsePtrOutput { return v.DesiredProperties }).(ClusterDesiredPropertiesResponsePtrOutput)
}

// Most recent billing meter timestamp.
func (o ClusterOutput) LastBillingTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.LastBillingTimestamp }).(pulumi.StringOutput)
}

// Most recent cluster sync timestamp.
func (o ClusterOutput) LastSyncTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.LastSyncTimestamp }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o ClusterOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

// Provisioning state.
func (o ClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// First cluster sync timestamp.
func (o ClusterOutput) RegistrationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.RegistrationTimestamp }).(pulumi.StringOutput)
}

// Properties reported by cluster agent.
func (o ClusterOutput) ReportedProperties() ClusterReportedPropertiesResponseOutput {
	return o.ApplyT(func(v *Cluster) ClusterReportedPropertiesResponseOutput { return v.ReportedProperties }).(ClusterReportedPropertiesResponseOutput)
}

// Object id of RP Service Principal
func (o ClusterOutput) ResourceProviderObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ResourceProviderObjectId }).(pulumi.StringOutput)
}

// Region specific DataPath Endpoint of the cluster.
func (o ClusterOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// Software Assurance properties of the cluster.
func (o ClusterOutput) SoftwareAssuranceProperties() SoftwareAssurancePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) SoftwareAssurancePropertiesResponsePtrOutput { return v.SoftwareAssuranceProperties }).(SoftwareAssurancePropertiesResponsePtrOutput)
}

// Status of the cluster agent.
func (o ClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Cluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o ClusterOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Number of days remaining in the trial period.
func (o ClusterOutput) TrialDaysRemaining() pulumi.Float64Output {
	return o.ApplyT(func(v *Cluster) pulumi.Float64Output { return v.TrialDaysRemaining }).(pulumi.Float64Output)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
func (o ClusterOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *Cluster) UserAssignedIdentityResponseMapOutput { return v.UserAssignedIdentities }).(UserAssignedIdentityResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
