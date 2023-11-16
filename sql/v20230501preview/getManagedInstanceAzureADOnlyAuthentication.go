// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a specific Azure Active Directory only authentication property.
func LookupManagedInstanceAzureADOnlyAuthentication(ctx *pulumi.Context, args *LookupManagedInstanceAzureADOnlyAuthenticationArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstanceAzureADOnlyAuthenticationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedInstanceAzureADOnlyAuthenticationResult
	err := ctx.Invoke("azure-native:sql/v20230501preview:getManagedInstanceAzureADOnlyAuthentication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstanceAzureADOnlyAuthenticationArgs struct {
	// The name of server azure active directory only authentication.
	AuthenticationName string `pulumi:"authenticationName"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Azure Active Directory only authentication.
type LookupManagedInstanceAzureADOnlyAuthenticationResult struct {
	// Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication bool `pulumi:"azureADOnlyAuthentication"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupManagedInstanceAzureADOnlyAuthenticationOutput(ctx *pulumi.Context, args LookupManagedInstanceAzureADOnlyAuthenticationOutputArgs, opts ...pulumi.InvokeOption) LookupManagedInstanceAzureADOnlyAuthenticationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedInstanceAzureADOnlyAuthenticationResult, error) {
			args := v.(LookupManagedInstanceAzureADOnlyAuthenticationArgs)
			r, err := LookupManagedInstanceAzureADOnlyAuthentication(ctx, &args, opts...)
			var s LookupManagedInstanceAzureADOnlyAuthenticationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedInstanceAzureADOnlyAuthenticationResultOutput)
}

type LookupManagedInstanceAzureADOnlyAuthenticationOutputArgs struct {
	// The name of server azure active directory only authentication.
	AuthenticationName pulumi.StringInput `pulumi:"authenticationName"`
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedInstanceAzureADOnlyAuthenticationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceAzureADOnlyAuthenticationArgs)(nil)).Elem()
}

// Azure Active Directory only authentication.
type LookupManagedInstanceAzureADOnlyAuthenticationResultOutput struct{ *pulumi.OutputState }

func (LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceAzureADOnlyAuthenticationResult)(nil)).Elem()
}

func (o LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) ToLookupManagedInstanceAzureADOnlyAuthenticationResultOutput() LookupManagedInstanceAzureADOnlyAuthenticationResultOutput {
	return o
}

func (o LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) ToLookupManagedInstanceAzureADOnlyAuthenticationResultOutputWithContext(ctx context.Context) LookupManagedInstanceAzureADOnlyAuthenticationResultOutput {
	return o
}

// Azure Active Directory only Authentication enabled.
func (o LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) AzureADOnlyAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupManagedInstanceAzureADOnlyAuthenticationResult) bool { return v.AzureADOnlyAuthentication }).(pulumi.BoolOutput)
}

// Resource ID.
func (o LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAzureADOnlyAuthenticationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAzureADOnlyAuthenticationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAzureADOnlyAuthenticationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedInstanceAzureADOnlyAuthenticationResultOutput{})
}
