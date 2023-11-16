// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211030preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Project Database Details
type DatabaseInfo struct {
	// Name of the database
	SourceDatabaseName string `pulumi:"sourceDatabaseName"`
}

// DatabaseInfoInput is an input type that accepts DatabaseInfoArgs and DatabaseInfoOutput values.
// You can construct a concrete instance of `DatabaseInfoInput` via:
//
//	DatabaseInfoArgs{...}
type DatabaseInfoInput interface {
	pulumi.Input

	ToDatabaseInfoOutput() DatabaseInfoOutput
	ToDatabaseInfoOutputWithContext(context.Context) DatabaseInfoOutput
}

// Project Database Details
type DatabaseInfoArgs struct {
	// Name of the database
	SourceDatabaseName pulumi.StringInput `pulumi:"sourceDatabaseName"`
}

func (DatabaseInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseInfo)(nil)).Elem()
}

func (i DatabaseInfoArgs) ToDatabaseInfoOutput() DatabaseInfoOutput {
	return i.ToDatabaseInfoOutputWithContext(context.Background())
}

func (i DatabaseInfoArgs) ToDatabaseInfoOutputWithContext(ctx context.Context) DatabaseInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInfoOutput)
}

// DatabaseInfoArrayInput is an input type that accepts DatabaseInfoArray and DatabaseInfoArrayOutput values.
// You can construct a concrete instance of `DatabaseInfoArrayInput` via:
//
//	DatabaseInfoArray{ DatabaseInfoArgs{...} }
type DatabaseInfoArrayInput interface {
	pulumi.Input

	ToDatabaseInfoArrayOutput() DatabaseInfoArrayOutput
	ToDatabaseInfoArrayOutputWithContext(context.Context) DatabaseInfoArrayOutput
}

type DatabaseInfoArray []DatabaseInfoInput

func (DatabaseInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseInfo)(nil)).Elem()
}

func (i DatabaseInfoArray) ToDatabaseInfoArrayOutput() DatabaseInfoArrayOutput {
	return i.ToDatabaseInfoArrayOutputWithContext(context.Background())
}

func (i DatabaseInfoArray) ToDatabaseInfoArrayOutputWithContext(ctx context.Context) DatabaseInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInfoArrayOutput)
}

// Project Database Details
type DatabaseInfoOutput struct{ *pulumi.OutputState }

func (DatabaseInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseInfo)(nil)).Elem()
}

func (o DatabaseInfoOutput) ToDatabaseInfoOutput() DatabaseInfoOutput {
	return o
}

func (o DatabaseInfoOutput) ToDatabaseInfoOutputWithContext(ctx context.Context) DatabaseInfoOutput {
	return o
}

// Name of the database
func (o DatabaseInfoOutput) SourceDatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseInfo) string { return v.SourceDatabaseName }).(pulumi.StringOutput)
}

type DatabaseInfoArrayOutput struct{ *pulumi.OutputState }

func (DatabaseInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseInfo)(nil)).Elem()
}

func (o DatabaseInfoArrayOutput) ToDatabaseInfoArrayOutput() DatabaseInfoArrayOutput {
	return o
}

func (o DatabaseInfoArrayOutput) ToDatabaseInfoArrayOutputWithContext(ctx context.Context) DatabaseInfoArrayOutput {
	return o
}

func (o DatabaseInfoArrayOutput) Index(i pulumi.IntInput) DatabaseInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseInfo {
		return vs[0].([]DatabaseInfo)[vs[1].(int)]
	}).(DatabaseInfoOutput)
}

// Project Database Details
type DatabaseInfoResponse struct {
	// Name of the database
	SourceDatabaseName string `pulumi:"sourceDatabaseName"`
}

// Project Database Details
type DatabaseInfoResponseOutput struct{ *pulumi.OutputState }

func (DatabaseInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseInfoResponse)(nil)).Elem()
}

func (o DatabaseInfoResponseOutput) ToDatabaseInfoResponseOutput() DatabaseInfoResponseOutput {
	return o
}

func (o DatabaseInfoResponseOutput) ToDatabaseInfoResponseOutputWithContext(ctx context.Context) DatabaseInfoResponseOutput {
	return o
}

// Name of the database
func (o DatabaseInfoResponseOutput) SourceDatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseInfoResponse) string { return v.SourceDatabaseName }).(pulumi.StringOutput)
}

type DatabaseInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseInfoResponse)(nil)).Elem()
}

func (o DatabaseInfoResponseArrayOutput) ToDatabaseInfoResponseArrayOutput() DatabaseInfoResponseArrayOutput {
	return o
}

func (o DatabaseInfoResponseArrayOutput) ToDatabaseInfoResponseArrayOutputWithContext(ctx context.Context) DatabaseInfoResponseArrayOutput {
	return o
}

func (o DatabaseInfoResponseArrayOutput) Index(i pulumi.IntInput) DatabaseInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseInfoResponse {
		return vs[0].([]DatabaseInfoResponse)[vs[1].(int)]
	}).(DatabaseInfoResponseOutput)
}

// Properties required to create a connection to Azure SQL database Managed instance
type MiSqlConnectionInfo struct {
	// Resource id for Azure SQL database Managed instance
	ManagedInstanceResourceId string `pulumi:"managedInstanceResourceId"`
	// Password credential.
	Password *string `pulumi:"password"`
	// Type of connection info
	// Expected value is 'MiSqlConnectionInfo'.
	Type string `pulumi:"type"`
	// User name
	UserName *string `pulumi:"userName"`
}

// Properties required to create a connection to Azure SQL database Managed instance
type MiSqlConnectionInfoResponse struct {
	// Resource id for Azure SQL database Managed instance
	ManagedInstanceResourceId string `pulumi:"managedInstanceResourceId"`
	// Password credential.
	Password *string `pulumi:"password"`
	// Type of connection info
	// Expected value is 'MiSqlConnectionInfo'.
	Type string `pulumi:"type"`
	// User name
	UserName *string `pulumi:"userName"`
}

// Describes a connection to a MongoDB data source
type MongoDbConnectionInfo struct {
	// Additional connection settings
	AdditionalSettings *string `pulumi:"additionalSettings"`
	// A MongoDB connection string or blob container URL. The user name and password can be specified here or in the userName and password properties
	ConnectionString string `pulumi:"connectionString"`
	// Data source
	DataSource *string `pulumi:"dataSource"`
	// Whether to encrypt the connection
	EncryptConnection *bool `pulumi:"encryptConnection"`
	EnforceSSL        *bool `pulumi:"enforceSSL"`
	// Password credential.
	Password *string `pulumi:"password"`
	// port for server
	Port *int `pulumi:"port"`
	// server brand version
	ServerBrandVersion *string `pulumi:"serverBrandVersion"`
	// Type of connection info
	// Expected value is 'MongoDbConnectionInfo'.
	Type string `pulumi:"type"`
	// User name
	UserName *string `pulumi:"userName"`
}

// Describes a connection to a MongoDB data source
type MongoDbConnectionInfoResponse struct {
	// Additional connection settings
	AdditionalSettings *string `pulumi:"additionalSettings"`
	// A MongoDB connection string or blob container URL. The user name and password can be specified here or in the userName and password properties
	ConnectionString string `pulumi:"connectionString"`
	// Data source
	DataSource *string `pulumi:"dataSource"`
	// Whether to encrypt the connection
	EncryptConnection *bool `pulumi:"encryptConnection"`
	EnforceSSL        *bool `pulumi:"enforceSSL"`
	// Password credential.
	Password *string `pulumi:"password"`
	// port for server
	Port *int `pulumi:"port"`
	// server brand version
	ServerBrandVersion *string `pulumi:"serverBrandVersion"`
	// Type of connection info
	// Expected value is 'MongoDbConnectionInfo'.
	Type string `pulumi:"type"`
	// User name
	UserName *string `pulumi:"userName"`
}

// Information for connecting to MySQL server
type MySqlConnectionInfo struct {
	// Data source
	DataSource *string `pulumi:"dataSource"`
	// Whether to encrypt the connection
	EncryptConnection *bool `pulumi:"encryptConnection"`
	// Password credential.
	Password *string `pulumi:"password"`
	// Port for Server
	Port int `pulumi:"port"`
	// Name of the server
	ServerName string `pulumi:"serverName"`
	// Type of connection info
	// Expected value is 'MySqlConnectionInfo'.
	Type string `pulumi:"type"`
	// User name
	UserName *string `pulumi:"userName"`
}

// Defaults sets the appropriate defaults for MySqlConnectionInfo
func (val *MySqlConnectionInfo) Defaults() *MySqlConnectionInfo {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EncryptConnection == nil {
		encryptConnection_ := true
		tmp.EncryptConnection = &encryptConnection_
	}
	return &tmp
}

// Information for connecting to MySQL server
type MySqlConnectionInfoResponse struct {
	// Data source
	DataSource *string `pulumi:"dataSource"`
	// Whether to encrypt the connection
	EncryptConnection *bool `pulumi:"encryptConnection"`
	// Password credential.
	Password *string `pulumi:"password"`
	// Port for Server
	Port int `pulumi:"port"`
	// Name of the server
	ServerName string `pulumi:"serverName"`
	// Type of connection info
	// Expected value is 'MySqlConnectionInfo'.
	Type string `pulumi:"type"`
	// User name
	UserName *string `pulumi:"userName"`
}

// Defaults sets the appropriate defaults for MySqlConnectionInfoResponse
func (val *MySqlConnectionInfoResponse) Defaults() *MySqlConnectionInfoResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EncryptConnection == nil {
		encryptConnection_ := true
		tmp.EncryptConnection = &encryptConnection_
	}
	return &tmp
}

// Information for connecting to Oracle server
type OracleConnectionInfo struct {
	// EZConnect or TNSName connection string.
	DataSource string `pulumi:"dataSource"`
	// Password credential.
	Password *string `pulumi:"password"`
	// Type of connection info
	// Expected value is 'OracleConnectionInfo'.
	Type string `pulumi:"type"`
	// User name
	UserName *string `pulumi:"userName"`
}

// Information for connecting to Oracle server
type OracleConnectionInfoResponse struct {
	// EZConnect or TNSName connection string.
	DataSource string `pulumi:"dataSource"`
	// Password credential.
	Password *string `pulumi:"password"`
	// Type of connection info
	// Expected value is 'OracleConnectionInfo'.
	Type string `pulumi:"type"`
	// User name
	UserName *string `pulumi:"userName"`
}

// Information for connecting to PostgreSQL server
type PostgreSqlConnectionInfo struct {
	// Data source
	DataSource *string `pulumi:"dataSource"`
	// Name of the database
	DatabaseName *string `pulumi:"databaseName"`
	// Whether to encrypt the connection
	EncryptConnection *bool `pulumi:"encryptConnection"`
	// Password credential.
	Password *string `pulumi:"password"`
	// Port for Server
	Port int `pulumi:"port"`
	// Name of the server
	ServerName string `pulumi:"serverName"`
	// server version
	ServerVersion *string `pulumi:"serverVersion"`
	// Whether to trust the server certificate
	TrustServerCertificate *bool `pulumi:"trustServerCertificate"`
	// Type of connection info
	// Expected value is 'PostgreSqlConnectionInfo'.
	Type string `pulumi:"type"`
	// User name
	UserName *string `pulumi:"userName"`
}

// Defaults sets the appropriate defaults for PostgreSqlConnectionInfo
func (val *PostgreSqlConnectionInfo) Defaults() *PostgreSqlConnectionInfo {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EncryptConnection == nil {
		encryptConnection_ := true
		tmp.EncryptConnection = &encryptConnection_
	}
	if tmp.TrustServerCertificate == nil {
		trustServerCertificate_ := false
		tmp.TrustServerCertificate = &trustServerCertificate_
	}
	return &tmp
}

// Information for connecting to PostgreSQL server
type PostgreSqlConnectionInfoResponse struct {
	// Data source
	DataSource *string `pulumi:"dataSource"`
	// Name of the database
	DatabaseName *string `pulumi:"databaseName"`
	// Whether to encrypt the connection
	EncryptConnection *bool `pulumi:"encryptConnection"`
	// Password credential.
	Password *string `pulumi:"password"`
	// Port for Server
	Port int `pulumi:"port"`
	// Name of the server
	ServerName string `pulumi:"serverName"`
	// server version
	ServerVersion *string `pulumi:"serverVersion"`
	// Whether to trust the server certificate
	TrustServerCertificate *bool `pulumi:"trustServerCertificate"`
	// Type of connection info
	// Expected value is 'PostgreSqlConnectionInfo'.
	Type string `pulumi:"type"`
	// User name
	UserName *string `pulumi:"userName"`
}

// Defaults sets the appropriate defaults for PostgreSqlConnectionInfoResponse
func (val *PostgreSqlConnectionInfoResponse) Defaults() *PostgreSqlConnectionInfoResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EncryptConnection == nil {
		encryptConnection_ := true
		tmp.EncryptConnection = &encryptConnection_
	}
	if tmp.TrustServerCertificate == nil {
		trustServerCertificate_ := false
		tmp.TrustServerCertificate = &trustServerCertificate_
	}
	return &tmp
}

// Information for connecting to SQL database server
type SqlConnectionInfo struct {
	// Additional connection settings
	AdditionalSettings *string `pulumi:"additionalSettings"`
	// Authentication type to use for connection
	Authentication *string `pulumi:"authentication"`
	// Data source in the format Protocol:MachineName\SQLServerInstanceName,PortNumber
	DataSource string `pulumi:"dataSource"`
	// Whether to encrypt the connection
	EncryptConnection *bool `pulumi:"encryptConnection"`
	// Password credential.
	Password *string `pulumi:"password"`
	// Server platform type for connection
	Platform *string `pulumi:"platform"`
	// port for server
	Port *string `pulumi:"port"`
	// Represents the ID of an HTTP resource represented by an Azure resource provider.
	ResourceId *string `pulumi:"resourceId"`
	// name of the server
	ServerName *string `pulumi:"serverName"`
	// Whether to trust the server certificate
	TrustServerCertificate *bool `pulumi:"trustServerCertificate"`
	// Type of connection info
	// Expected value is 'SqlConnectionInfo'.
	Type string `pulumi:"type"`
	// User name
	UserName *string `pulumi:"userName"`
}

// Defaults sets the appropriate defaults for SqlConnectionInfo
func (val *SqlConnectionInfo) Defaults() *SqlConnectionInfo {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EncryptConnection == nil {
		encryptConnection_ := true
		tmp.EncryptConnection = &encryptConnection_
	}
	if tmp.TrustServerCertificate == nil {
		trustServerCertificate_ := false
		tmp.TrustServerCertificate = &trustServerCertificate_
	}
	return &tmp
}

// Information for connecting to SQL database server
type SqlConnectionInfoResponse struct {
	// Additional connection settings
	AdditionalSettings *string `pulumi:"additionalSettings"`
	// Authentication type to use for connection
	Authentication *string `pulumi:"authentication"`
	// Data source in the format Protocol:MachineName\SQLServerInstanceName,PortNumber
	DataSource string `pulumi:"dataSource"`
	// Whether to encrypt the connection
	EncryptConnection *bool `pulumi:"encryptConnection"`
	// Password credential.
	Password *string `pulumi:"password"`
	// Server platform type for connection
	Platform *string `pulumi:"platform"`
	// port for server
	Port *string `pulumi:"port"`
	// Represents the ID of an HTTP resource represented by an Azure resource provider.
	ResourceId *string `pulumi:"resourceId"`
	// name of the server
	ServerName *string `pulumi:"serverName"`
	// Whether to trust the server certificate
	TrustServerCertificate *bool `pulumi:"trustServerCertificate"`
	// Type of connection info
	// Expected value is 'SqlConnectionInfo'.
	Type string `pulumi:"type"`
	// User name
	UserName *string `pulumi:"userName"`
}

// Defaults sets the appropriate defaults for SqlConnectionInfoResponse
func (val *SqlConnectionInfoResponse) Defaults() *SqlConnectionInfoResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EncryptConnection == nil {
		encryptConnection_ := true
		tmp.EncryptConnection = &encryptConnection_
	}
	if tmp.TrustServerCertificate == nil {
		trustServerCertificate_ := false
		tmp.TrustServerCertificate = &trustServerCertificate_
	}
	return &tmp
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseInfoOutput{})
	pulumi.RegisterOutputType(DatabaseInfoArrayOutput{})
	pulumi.RegisterOutputType(DatabaseInfoResponseOutput{})
	pulumi.RegisterOutputType(DatabaseInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
