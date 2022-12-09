// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Custom domain analysis.
func ListContainerAppCustomHostNameAnalysis(ctx *pulumi.Context, args *ListContainerAppCustomHostNameAnalysisArgs, opts ...pulumi.InvokeOption) (*ListContainerAppCustomHostNameAnalysisResult, error) {
	var rv ListContainerAppCustomHostNameAnalysisResult
	err := ctx.Invoke("azure-native:app/v20221001:listContainerAppCustomHostNameAnalysis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListContainerAppCustomHostNameAnalysisArgs struct {
	// Name of the Container App.
	ContainerAppName string `pulumi:"containerAppName"`
	// Custom hostname.
	CustomHostname *string `pulumi:"customHostname"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Custom domain analysis.
type ListContainerAppCustomHostNameAnalysisResult struct {
	// A records visible for this hostname.
	ARecords []string `pulumi:"aRecords"`
	// Alternate CName records visible for this hostname.
	AlternateCNameRecords []string `pulumi:"alternateCNameRecords"`
	// Alternate TXT records visible for this hostname.
	AlternateTxtRecords []string `pulumi:"alternateTxtRecords"`
	// CName records visible for this hostname.
	CNameRecords []string `pulumi:"cNameRecords"`
	// <code>true</code> if there is a conflict on the Container App's managed environment level custom domain; otherwise, <code>false</code>.
	ConflictWithEnvironmentCustomDomain bool `pulumi:"conflictWithEnvironmentCustomDomain"`
	// Name of the conflicting Container App on the Managed Environment if it's within the same subscription.
	ConflictingContainerAppResourceId string `pulumi:"conflictingContainerAppResourceId"`
	// Raw failure information if DNS verification fails.
	CustomDomainVerificationFailureInfo CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfo `pulumi:"customDomainVerificationFailureInfo"`
	// DNS verification test result.
	CustomDomainVerificationTest string `pulumi:"customDomainVerificationTest"`
	// <code>true</code> if there is a conflict on the Container App's managed environment; otherwise, <code>false</code>.
	HasConflictOnManagedEnvironment bool `pulumi:"hasConflictOnManagedEnvironment"`
	// Host name that was analyzed
	HostName string `pulumi:"hostName"`
	// <code>true</code> if hostname is already verified; otherwise, <code>false</code>.
	IsHostnameAlreadyVerified bool `pulumi:"isHostnameAlreadyVerified"`
	// TXT records visible for this hostname.
	TxtRecords []string `pulumi:"txtRecords"`
}

func ListContainerAppCustomHostNameAnalysisOutput(ctx *pulumi.Context, args ListContainerAppCustomHostNameAnalysisOutputArgs, opts ...pulumi.InvokeOption) ListContainerAppCustomHostNameAnalysisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListContainerAppCustomHostNameAnalysisResult, error) {
			args := v.(ListContainerAppCustomHostNameAnalysisArgs)
			r, err := ListContainerAppCustomHostNameAnalysis(ctx, &args, opts...)
			var s ListContainerAppCustomHostNameAnalysisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListContainerAppCustomHostNameAnalysisResultOutput)
}

type ListContainerAppCustomHostNameAnalysisOutputArgs struct {
	// Name of the Container App.
	ContainerAppName pulumi.StringInput `pulumi:"containerAppName"`
	// Custom hostname.
	CustomHostname pulumi.StringPtrInput `pulumi:"customHostname"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListContainerAppCustomHostNameAnalysisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListContainerAppCustomHostNameAnalysisArgs)(nil)).Elem()
}

// Custom domain analysis.
type ListContainerAppCustomHostNameAnalysisResultOutput struct{ *pulumi.OutputState }

func (ListContainerAppCustomHostNameAnalysisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListContainerAppCustomHostNameAnalysisResult)(nil)).Elem()
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) ToListContainerAppCustomHostNameAnalysisResultOutput() ListContainerAppCustomHostNameAnalysisResultOutput {
	return o
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) ToListContainerAppCustomHostNameAnalysisResultOutputWithContext(ctx context.Context) ListContainerAppCustomHostNameAnalysisResultOutput {
	return o
}

// A records visible for this hostname.
func (o ListContainerAppCustomHostNameAnalysisResultOutput) ARecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) []string { return v.ARecords }).(pulumi.StringArrayOutput)
}

// Alternate CName records visible for this hostname.
func (o ListContainerAppCustomHostNameAnalysisResultOutput) AlternateCNameRecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) []string { return v.AlternateCNameRecords }).(pulumi.StringArrayOutput)
}

// Alternate TXT records visible for this hostname.
func (o ListContainerAppCustomHostNameAnalysisResultOutput) AlternateTxtRecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) []string { return v.AlternateTxtRecords }).(pulumi.StringArrayOutput)
}

// CName records visible for this hostname.
func (o ListContainerAppCustomHostNameAnalysisResultOutput) CNameRecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) []string { return v.CNameRecords }).(pulumi.StringArrayOutput)
}

// <code>true</code> if there is a conflict on the Container App's managed environment level custom domain; otherwise, <code>false</code>.
func (o ListContainerAppCustomHostNameAnalysisResultOutput) ConflictWithEnvironmentCustomDomain() pulumi.BoolOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) bool {
		return v.ConflictWithEnvironmentCustomDomain
	}).(pulumi.BoolOutput)
}

// Name of the conflicting Container App on the Managed Environment if it's within the same subscription.
func (o ListContainerAppCustomHostNameAnalysisResultOutput) ConflictingContainerAppResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) string {
		return v.ConflictingContainerAppResourceId
	}).(pulumi.StringOutput)
}

// Raw failure information if DNS verification fails.
func (o ListContainerAppCustomHostNameAnalysisResultOutput) CustomDomainVerificationFailureInfo() CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfo {
		return v.CustomDomainVerificationFailureInfo
	}).(CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput)
}

// DNS verification test result.
func (o ListContainerAppCustomHostNameAnalysisResultOutput) CustomDomainVerificationTest() pulumi.StringOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) string { return v.CustomDomainVerificationTest }).(pulumi.StringOutput)
}

// <code>true</code> if there is a conflict on the Container App's managed environment; otherwise, <code>false</code>.
func (o ListContainerAppCustomHostNameAnalysisResultOutput) HasConflictOnManagedEnvironment() pulumi.BoolOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) bool { return v.HasConflictOnManagedEnvironment }).(pulumi.BoolOutput)
}

// Host name that was analyzed
func (o ListContainerAppCustomHostNameAnalysisResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) string { return v.HostName }).(pulumi.StringOutput)
}

// <code>true</code> if hostname is already verified; otherwise, <code>false</code>.
func (o ListContainerAppCustomHostNameAnalysisResultOutput) IsHostnameAlreadyVerified() pulumi.BoolOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) bool { return v.IsHostnameAlreadyVerified }).(pulumi.BoolOutput)
}

// TXT records visible for this hostname.
func (o ListContainerAppCustomHostNameAnalysisResultOutput) TxtRecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) []string { return v.TxtRecords }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListContainerAppCustomHostNameAnalysisResultOutput{})
}
