// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managedidentity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the federated identity credential.
// Azure REST API version: 2023-01-31.
func LookupFederatedIdentityCredential(ctx *pulumi.Context, args *LookupFederatedIdentityCredentialArgs, opts ...pulumi.InvokeOption) (*LookupFederatedIdentityCredentialResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedIdentityCredentialResult
	err := ctx.Invoke("azure-native:managedidentity:getFederatedIdentityCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFederatedIdentityCredentialArgs struct {
	// The name of the federated identity credential resource.
	FederatedIdentityCredentialResourceName string `pulumi:"federatedIdentityCredentialResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the identity resource.
	ResourceName string `pulumi:"resourceName"`
}

// Describes a federated identity credential.
type LookupFederatedIdentityCredentialResult struct {
	// The list of audiences that can appear in the issued token.
	Audiences []string `pulumi:"audiences"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The URL of the issuer to be trusted.
	Issuer string `pulumi:"issuer"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The identifier of the external identity.
	Subject string `pulumi:"subject"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupFederatedIdentityCredentialOutput(ctx *pulumi.Context, args LookupFederatedIdentityCredentialOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedIdentityCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedIdentityCredentialResult, error) {
			args := v.(LookupFederatedIdentityCredentialArgs)
			r, err := LookupFederatedIdentityCredential(ctx, &args, opts...)
			var s LookupFederatedIdentityCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedIdentityCredentialResultOutput)
}

type LookupFederatedIdentityCredentialOutputArgs struct {
	// The name of the federated identity credential resource.
	FederatedIdentityCredentialResourceName pulumi.StringInput `pulumi:"federatedIdentityCredentialResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the identity resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupFederatedIdentityCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedIdentityCredentialArgs)(nil)).Elem()
}

// Describes a federated identity credential.
type LookupFederatedIdentityCredentialResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedIdentityCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedIdentityCredentialResult)(nil)).Elem()
}

func (o LookupFederatedIdentityCredentialResultOutput) ToLookupFederatedIdentityCredentialResultOutput() LookupFederatedIdentityCredentialResultOutput {
	return o
}

func (o LookupFederatedIdentityCredentialResultOutput) ToLookupFederatedIdentityCredentialResultOutputWithContext(ctx context.Context) LookupFederatedIdentityCredentialResultOutput {
	return o
}

func (o LookupFederatedIdentityCredentialResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFederatedIdentityCredentialResult] {
	return pulumix.Output[LookupFederatedIdentityCredentialResult]{
		OutputState: o.OutputState,
	}
}

// The list of audiences that can appear in the issued token.
func (o LookupFederatedIdentityCredentialResultOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedIdentityCredentialResult) []string { return v.Audiences }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupFederatedIdentityCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIdentityCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

// The URL of the issuer to be trusted.
func (o LookupFederatedIdentityCredentialResultOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIdentityCredentialResult) string { return v.Issuer }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupFederatedIdentityCredentialResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIdentityCredentialResult) string { return v.Name }).(pulumi.StringOutput)
}

// The identifier of the external identity.
func (o LookupFederatedIdentityCredentialResultOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIdentityCredentialResult) string { return v.Subject }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFederatedIdentityCredentialResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFederatedIdentityCredentialResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFederatedIdentityCredentialResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIdentityCredentialResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedIdentityCredentialResultOutput{})
}
