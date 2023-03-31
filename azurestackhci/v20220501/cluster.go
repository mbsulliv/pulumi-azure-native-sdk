// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
	// The timestamp of resource creation (UTC).
	CreatedAt pulumi.StringPtrOutput `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy pulumi.StringPtrOutput `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType pulumi.StringPtrOutput `pulumi:"createdByType"`
	// Desired properties of the cluster.
	DesiredProperties ClusterDesiredPropertiesResponsePtrOutput `pulumi:"desiredProperties"`
	// Most recent billing meter timestamp.
	LastBillingTimestamp pulumi.StringOutput `pulumi:"lastBillingTimestamp"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt pulumi.StringPtrOutput `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy pulumi.StringPtrOutput `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType pulumi.StringPtrOutput `pulumi:"lastModifiedByType"`
	// Most recent cluster sync timestamp.
	LastSyncTimestamp pulumi.StringOutput `pulumi:"lastSyncTimestamp"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// First cluster sync timestamp.
	RegistrationTimestamp pulumi.StringOutput `pulumi:"registrationTimestamp"`
	// Properties reported by cluster agent.
	ReportedProperties ClusterReportedPropertiesResponseOutput `pulumi:"reportedProperties"`
	// Region specific DataPath Endpoint of the cluster.
	ServiceEndpoint pulumi.StringOutput `pulumi:"serviceEndpoint"`
	// Status of the cluster agent.
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Number of days remaining in the trial period.
	TrialDaysRemaining pulumi.Float64Output `pulumi:"trialDaysRemaining"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
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
			Type: pulumi.String("azure-native:azurestackhci/v20220901:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221001:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221201:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230201:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:azurestackhci/v20220501:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:azurestackhci/v20220501:Cluster", name, id, state, &resource, opts...)
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
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// Desired properties of the cluster.
	DesiredProperties *ClusterDesiredProperties `pulumi:"desiredProperties"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
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
	// The timestamp of resource creation (UTC).
	CreatedAt pulumi.StringPtrInput
	// The identity that created the resource.
	CreatedBy pulumi.StringPtrInput
	// The type of identity that created the resource.
	CreatedByType pulumi.StringPtrInput
	// Desired properties of the cluster.
	DesiredProperties ClusterDesiredPropertiesPtrInput
	// The timestamp of resource last modification (UTC)
	LastModifiedAt pulumi.StringPtrInput
	// The identity that last modified the resource.
	LastModifiedBy pulumi.StringPtrInput
	// The type of identity that last modified the resource.
	LastModifiedByType pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
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

// The timestamp of resource creation (UTC).
func (o ClusterOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o ClusterOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o ClusterOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// Desired properties of the cluster.
func (o ClusterOutput) DesiredProperties() ClusterDesiredPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterDesiredPropertiesResponsePtrOutput { return v.DesiredProperties }).(ClusterDesiredPropertiesResponsePtrOutput)
}

// Most recent billing meter timestamp.
func (o ClusterOutput) LastBillingTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.LastBillingTimestamp }).(pulumi.StringOutput)
}

// The timestamp of resource last modification (UTC)
func (o ClusterOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o ClusterOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o ClusterOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
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

// Region specific DataPath Endpoint of the cluster.
func (o ClusterOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// Status of the cluster agent.
func (o ClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Resource tags.
func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Number of days remaining in the trial period.
func (o ClusterOutput) TrialDaysRemaining() pulumi.Float64Output {
	return o.ApplyT(func(v *Cluster) pulumi.Float64Output { return v.TrialDaysRemaining }).(pulumi.Float64Output)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
