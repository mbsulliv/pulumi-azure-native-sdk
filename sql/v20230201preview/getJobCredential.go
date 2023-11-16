// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a jobs credential.
func LookupJobCredential(ctx *pulumi.Context, args *LookupJobCredentialArgs, opts ...pulumi.InvokeOption) (*LookupJobCredentialResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupJobCredentialResult
	err := ctx.Invoke("azure-native:sql/v20230201preview:getJobCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobCredentialArgs struct {
	// The name of the credential.
	CredentialName string `pulumi:"credentialName"`
	// The name of the job agent.
	JobAgentName string `pulumi:"jobAgentName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A stored credential that can be used by a job to connect to target databases.
type LookupJobCredentialResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
	// The credential user name.
	Username string `pulumi:"username"`
}

func LookupJobCredentialOutput(ctx *pulumi.Context, args LookupJobCredentialOutputArgs, opts ...pulumi.InvokeOption) LookupJobCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobCredentialResult, error) {
			args := v.(LookupJobCredentialArgs)
			r, err := LookupJobCredential(ctx, &args, opts...)
			var s LookupJobCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobCredentialResultOutput)
}

type LookupJobCredentialOutputArgs struct {
	// The name of the credential.
	CredentialName pulumi.StringInput `pulumi:"credentialName"`
	// The name of the job agent.
	JobAgentName pulumi.StringInput `pulumi:"jobAgentName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupJobCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobCredentialArgs)(nil)).Elem()
}

// A stored credential that can be used by a job to connect to target databases.
type LookupJobCredentialResultOutput struct{ *pulumi.OutputState }

func (LookupJobCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobCredentialResult)(nil)).Elem()
}

func (o LookupJobCredentialResultOutput) ToLookupJobCredentialResultOutput() LookupJobCredentialResultOutput {
	return o
}

func (o LookupJobCredentialResultOutput) ToLookupJobCredentialResultOutputWithContext(ctx context.Context) LookupJobCredentialResultOutput {
	return o
}

// Resource ID.
func (o LookupJobCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupJobCredentialResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobCredentialResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupJobCredentialResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobCredentialResult) string { return v.Type }).(pulumi.StringOutput)
}

// The credential user name.
func (o LookupJobCredentialResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobCredentialResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobCredentialResultOutput{})
}
