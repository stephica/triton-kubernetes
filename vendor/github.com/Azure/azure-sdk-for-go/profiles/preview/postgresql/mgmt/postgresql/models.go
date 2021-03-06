// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package postgresql

import original "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"

type CheckNameAvailabilityClient = original.CheckNameAvailabilityClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ConfigurationsClient = original.ConfigurationsClient
type DatabasesClient = original.DatabasesClient
type FirewallRulesClient = original.FirewallRulesClient
type LocationBasedPerformanceTierClient = original.LocationBasedPerformanceTierClient
type LogFilesClient = original.LogFilesClient
type CreateMode = original.CreateMode

const (
	CreateModeDefault                   CreateMode = original.CreateModeDefault
	CreateModeGeoRestore                CreateMode = original.CreateModeGeoRestore
	CreateModePointInTimeRestore        CreateMode = original.CreateModePointInTimeRestore
	CreateModeServerPropertiesForCreate CreateMode = original.CreateModeServerPropertiesForCreate
)

type GeoRedundantBackup = original.GeoRedundantBackup

const (
	Disabled GeoRedundantBackup = original.Disabled
	Enabled  GeoRedundantBackup = original.Enabled
)

type OperationOrigin = original.OperationOrigin

const (
	NotSpecified OperationOrigin = original.NotSpecified
	System       OperationOrigin = original.System
	User         OperationOrigin = original.User
)

type ServerState = original.ServerState

const (
	ServerStateDisabled ServerState = original.ServerStateDisabled
	ServerStateDropping ServerState = original.ServerStateDropping
	ServerStateReady    ServerState = original.ServerStateReady
)

type ServerVersion = original.ServerVersion

const (
	NineFullStopFive ServerVersion = original.NineFullStopFive
	NineFullStopSix  ServerVersion = original.NineFullStopSix
)

type SkuTier = original.SkuTier

const (
	Basic           SkuTier = original.Basic
	GeneralPurpose  SkuTier = original.GeneralPurpose
	MemoryOptimized SkuTier = original.MemoryOptimized
)

type SslEnforcementEnum = original.SslEnforcementEnum

const (
	SslEnforcementEnumDisabled SslEnforcementEnum = original.SslEnforcementEnumDisabled
	SslEnforcementEnumEnabled  SslEnforcementEnum = original.SslEnforcementEnumEnabled
)

type Configuration = original.Configuration
type ConfigurationListResult = original.ConfigurationListResult
type ConfigurationProperties = original.ConfigurationProperties
type ConfigurationsCreateOrUpdateFuture = original.ConfigurationsCreateOrUpdateFuture
type Database = original.Database
type DatabaseListResult = original.DatabaseListResult
type DatabaseProperties = original.DatabaseProperties
type DatabasesCreateOrUpdateFuture = original.DatabasesCreateOrUpdateFuture
type DatabasesDeleteFuture = original.DatabasesDeleteFuture
type FirewallRule = original.FirewallRule
type FirewallRuleListResult = original.FirewallRuleListResult
type FirewallRuleProperties = original.FirewallRuleProperties
type FirewallRulesCreateOrUpdateFuture = original.FirewallRulesCreateOrUpdateFuture
type FirewallRulesDeleteFuture = original.FirewallRulesDeleteFuture
type LogFile = original.LogFile
type LogFileListResult = original.LogFileListResult
type LogFileProperties = original.LogFileProperties
type NameAvailability = original.NameAvailability
type NameAvailabilityRequest = original.NameAvailabilityRequest
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type PerformanceTierListResult = original.PerformanceTierListResult
type PerformanceTierProperties = original.PerformanceTierProperties
type PerformanceTierServiceLevelObjectives = original.PerformanceTierServiceLevelObjectives
type ProxyResource = original.ProxyResource
type Server = original.Server
type ServerForCreate = original.ServerForCreate
type ServerListResult = original.ServerListResult
type ServerProperties = original.ServerProperties
type BasicServerPropertiesForCreate = original.BasicServerPropertiesForCreate
type ServerPropertiesForCreate = original.ServerPropertiesForCreate
type ServerPropertiesForDefaultCreate = original.ServerPropertiesForDefaultCreate
type ServerPropertiesForGeoRestore = original.ServerPropertiesForGeoRestore
type ServerPropertiesForRestore = original.ServerPropertiesForRestore
type ServersCreateFuture = original.ServersCreateFuture
type ServersDeleteFuture = original.ServersDeleteFuture
type ServersUpdateFuture = original.ServersUpdateFuture
type ServerUpdateParameters = original.ServerUpdateParameters
type ServerUpdateParametersProperties = original.ServerUpdateParametersProperties
type Sku = original.Sku
type StorageProfile = original.StorageProfile
type TrackedResource = original.TrackedResource
type OperationsClient = original.OperationsClient
type ServersClient = original.ServersClient

func NewCheckNameAvailabilityClient(subscriptionID string) CheckNameAvailabilityClient {
	return original.NewCheckNameAvailabilityClient(subscriptionID)
}
func NewCheckNameAvailabilityClientWithBaseURI(baseURI string, subscriptionID string) CheckNameAvailabilityClient {
	return original.NewCheckNameAvailabilityClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClient(subscriptionID)
}
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationBasedPerformanceTierClient(subscriptionID string) LocationBasedPerformanceTierClient {
	return original.NewLocationBasedPerformanceTierClient(subscriptionID)
}
func NewLocationBasedPerformanceTierClientWithBaseURI(baseURI string, subscriptionID string) LocationBasedPerformanceTierClient {
	return original.NewLocationBasedPerformanceTierClientWithBaseURI(baseURI, subscriptionID)
}
func NewLogFilesClient(subscriptionID string) LogFilesClient {
	return original.NewLogFilesClient(subscriptionID)
}
func NewLogFilesClientWithBaseURI(baseURI string, subscriptionID string) LogFilesClient {
	return original.NewLogFilesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreateModeValues() []CreateMode {
	return original.PossibleCreateModeValues()
}
func PossibleGeoRedundantBackupValues() []GeoRedundantBackup {
	return original.PossibleGeoRedundantBackupValues()
}
func PossibleOperationOriginValues() []OperationOrigin {
	return original.PossibleOperationOriginValues()
}
func PossibleServerStateValues() []ServerState {
	return original.PossibleServerStateValues()
}
func PossibleServerVersionValues() []ServerVersion {
	return original.PossibleServerVersionValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleSslEnforcementEnumValues() []SslEnforcementEnum {
	return original.PossibleSslEnforcementEnumValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServersClient(subscriptionID string) ServersClient {
	return original.NewServersClient(subscriptionID)
}
func NewServersClientWithBaseURI(baseURI string, subscriptionID string) ServersClient {
	return original.NewServersClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
