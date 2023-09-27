// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210303preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an read-only access token for application insights diagnostic service data.
func GetDiagnosticServiceTokenReadOnly(ctx *pulumi.Context, args *GetDiagnosticServiceTokenReadOnlyArgs, opts ...pulumi.InvokeOption) (*GetDiagnosticServiceTokenReadOnlyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetDiagnosticServiceTokenReadOnlyResult
	err := ctx.Invoke("azure-native:insights/v20210303preview:getDiagnosticServiceTokenReadOnly", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDiagnosticServiceTokenReadOnlyArgs struct {
	// The identifier of the resource.
	ResourceUri string `pulumi:"resourceUri"`
}

// The response to a diagnostic services token query.
type GetDiagnosticServiceTokenReadOnlyResult struct {
	// JWT token for accessing application insights diagnostic service data.
	Token *string `pulumi:"token"`
}

func GetDiagnosticServiceTokenReadOnlyOutput(ctx *pulumi.Context, args GetDiagnosticServiceTokenReadOnlyOutputArgs, opts ...pulumi.InvokeOption) GetDiagnosticServiceTokenReadOnlyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDiagnosticServiceTokenReadOnlyResult, error) {
			args := v.(GetDiagnosticServiceTokenReadOnlyArgs)
			r, err := GetDiagnosticServiceTokenReadOnly(ctx, &args, opts...)
			var s GetDiagnosticServiceTokenReadOnlyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDiagnosticServiceTokenReadOnlyResultOutput)
}

type GetDiagnosticServiceTokenReadOnlyOutputArgs struct {
	// The identifier of the resource.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (GetDiagnosticServiceTokenReadOnlyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiagnosticServiceTokenReadOnlyArgs)(nil)).Elem()
}

// The response to a diagnostic services token query.
type GetDiagnosticServiceTokenReadOnlyResultOutput struct{ *pulumi.OutputState }

func (GetDiagnosticServiceTokenReadOnlyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiagnosticServiceTokenReadOnlyResult)(nil)).Elem()
}

func (o GetDiagnosticServiceTokenReadOnlyResultOutput) ToGetDiagnosticServiceTokenReadOnlyResultOutput() GetDiagnosticServiceTokenReadOnlyResultOutput {
	return o
}

func (o GetDiagnosticServiceTokenReadOnlyResultOutput) ToGetDiagnosticServiceTokenReadOnlyResultOutputWithContext(ctx context.Context) GetDiagnosticServiceTokenReadOnlyResultOutput {
	return o
}

func (o GetDiagnosticServiceTokenReadOnlyResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetDiagnosticServiceTokenReadOnlyResult] {
	return pulumix.Output[GetDiagnosticServiceTokenReadOnlyResult]{
		OutputState: o.OutputState,
	}
}

// JWT token for accessing application insights diagnostic service data.
func (o GetDiagnosticServiceTokenReadOnlyResultOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDiagnosticServiceTokenReadOnlyResult) *string { return v.Token }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDiagnosticServiceTokenReadOnlyResultOutput{})
}
