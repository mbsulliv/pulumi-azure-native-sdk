// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A stored credential that can be used by a job to connect to target databases.
type JobCredential struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The credential user name.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewJobCredential registers a new resource with the given unique name, arguments, and options.
func NewJobCredential(ctx *pulumi.Context,
	name string, args *JobCredentialArgs, opts ...pulumi.ResourceOption) (*JobCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobAgentName == nil {
		return nil, errors.New("invalid value for required argument 'JobAgentName'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:JobCredential"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:JobCredential"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource JobCredential
	err := ctx.RegisterResource("azure-native:sql/v20211101:JobCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobCredential gets an existing JobCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobCredentialState, opts ...pulumi.ResourceOption) (*JobCredential, error) {
	var resource JobCredential
	err := ctx.ReadResource("azure-native:sql/v20211101:JobCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobCredential resources.
type jobCredentialState struct {
}

type JobCredentialState struct {
}

func (JobCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobCredentialState)(nil)).Elem()
}

type jobCredentialArgs struct {
	// The name of the credential.
	CredentialName *string `pulumi:"credentialName"`
	// The name of the job agent.
	JobAgentName string `pulumi:"jobAgentName"`
	// The credential password.
	Password string `pulumi:"password"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The credential user name.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a JobCredential resource.
type JobCredentialArgs struct {
	// The name of the credential.
	CredentialName pulumi.StringPtrInput
	// The name of the job agent.
	JobAgentName pulumi.StringInput
	// The credential password.
	Password pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The credential user name.
	Username pulumi.StringInput
}

func (JobCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobCredentialArgs)(nil)).Elem()
}

type JobCredentialInput interface {
	pulumi.Input

	ToJobCredentialOutput() JobCredentialOutput
	ToJobCredentialOutputWithContext(ctx context.Context) JobCredentialOutput
}

func (*JobCredential) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCredential)(nil)).Elem()
}

func (i *JobCredential) ToJobCredentialOutput() JobCredentialOutput {
	return i.ToJobCredentialOutputWithContext(context.Background())
}

func (i *JobCredential) ToJobCredentialOutputWithContext(ctx context.Context) JobCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCredentialOutput)
}

type JobCredentialOutput struct{ *pulumi.OutputState }

func (JobCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCredential)(nil)).Elem()
}

func (o JobCredentialOutput) ToJobCredentialOutput() JobCredentialOutput {
	return o
}

func (o JobCredentialOutput) ToJobCredentialOutputWithContext(ctx context.Context) JobCredentialOutput {
	return o
}

// Resource name.
func (o JobCredentialOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobCredential) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o JobCredentialOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobCredential) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The credential user name.
func (o JobCredentialOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *JobCredential) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobCredentialOutput{})
}
