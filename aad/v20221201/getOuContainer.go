// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get OuContainer in DomainService instance.
func LookupOuContainer(ctx *pulumi.Context, args *LookupOuContainerArgs, opts ...pulumi.InvokeOption) (*LookupOuContainerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOuContainerResult
	err := ctx.Invoke("azure-native:aad/v20221201:getOuContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOuContainerArgs struct {
	// The name of the domain service.
	DomainServiceName string `pulumi:"domainServiceName"`
	// The name of the OuContainer.
	OuContainerName string `pulumi:"ouContainerName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Resource for OuContainer.
type LookupOuContainerResult struct {
	// The list of container accounts
	Accounts []ContainerAccountResponse `pulumi:"accounts"`
	// The OuContainer name
	ContainerId string `pulumi:"containerId"`
	// The Deployment id
	DeploymentId string `pulumi:"deploymentId"`
	// Distinguished Name of OuContainer instance
	DistinguishedName string `pulumi:"distinguishedName"`
	// The domain name of Domain Services.
	DomainName string `pulumi:"domainName"`
	// Resource etag
	Etag *string `pulumi:"etag"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// Status of OuContainer instance
	ServiceStatus string `pulumi:"serviceStatus"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure Active Directory tenant id
	TenantId string `pulumi:"tenantId"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupOuContainerOutput(ctx *pulumi.Context, args LookupOuContainerOutputArgs, opts ...pulumi.InvokeOption) LookupOuContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOuContainerResult, error) {
			args := v.(LookupOuContainerArgs)
			r, err := LookupOuContainer(ctx, &args, opts...)
			var s LookupOuContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOuContainerResultOutput)
}

type LookupOuContainerOutputArgs struct {
	// The name of the domain service.
	DomainServiceName pulumi.StringInput `pulumi:"domainServiceName"`
	// The name of the OuContainer.
	OuContainerName pulumi.StringInput `pulumi:"ouContainerName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOuContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOuContainerArgs)(nil)).Elem()
}

// Resource for OuContainer.
type LookupOuContainerResultOutput struct{ *pulumi.OutputState }

func (LookupOuContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOuContainerResult)(nil)).Elem()
}

func (o LookupOuContainerResultOutput) ToLookupOuContainerResultOutput() LookupOuContainerResultOutput {
	return o
}

func (o LookupOuContainerResultOutput) ToLookupOuContainerResultOutputWithContext(ctx context.Context) LookupOuContainerResultOutput {
	return o
}

func (o LookupOuContainerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOuContainerResult] {
	return pulumix.Output[LookupOuContainerResult]{
		OutputState: o.OutputState,
	}
}

// The list of container accounts
func (o LookupOuContainerResultOutput) Accounts() ContainerAccountResponseArrayOutput {
	return o.ApplyT(func(v LookupOuContainerResult) []ContainerAccountResponse { return v.Accounts }).(ContainerAccountResponseArrayOutput)
}

// The OuContainer name
func (o LookupOuContainerResultOutput) ContainerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.ContainerId }).(pulumi.StringOutput)
}

// The Deployment id
func (o LookupOuContainerResultOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.DeploymentId }).(pulumi.StringOutput)
}

// Distinguished Name of OuContainer instance
func (o LookupOuContainerResultOutput) DistinguishedName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.DistinguishedName }).(pulumi.StringOutput)
}

// The domain name of Domain Services.
func (o LookupOuContainerResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.DomainName }).(pulumi.StringOutput)
}

// Resource etag
func (o LookupOuContainerResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOuContainerResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupOuContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupOuContainerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOuContainerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o LookupOuContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current deployment or provisioning state, which only appears in the response.
func (o LookupOuContainerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Status of OuContainer instance
func (o LookupOuContainerResultOutput) ServiceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.ServiceStatus }).(pulumi.StringOutput)
}

// The system meta data relating to this resource.
func (o LookupOuContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOuContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupOuContainerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOuContainerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure Active Directory tenant id
func (o LookupOuContainerResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type
func (o LookupOuContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOuContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOuContainerResultOutput{})
}
