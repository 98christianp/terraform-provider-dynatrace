/**
* @license
* Copyright 2020 Dynatrace LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package export

import (
	"reflect"
	"strings"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/activegatetoken"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/alerting/connectivityalerts"
	database_anomalies_v2 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/databases"
	disk_anomalies_v2 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/infrastructure/disks"
	disk_specific_anomalies_v2 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/infrastructure/disks/perdiskoverride"
	host_anomalies_v2 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/infrastructure/hosts"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/kubernetes/cluster"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/kubernetes/namespace"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/kubernetes/node"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/kubernetes/pvc"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/kubernetes/workload"
	custom_app_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/rum/custom"
	custom_app_crash_rate "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/rum/custom/crashrate"
	mobile_app_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/rum/mobile"
	mobile_app_crash_rate "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/rum/mobile/crashrate"
	web_app_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/rum/web"
	apidetection "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/apis/detectionrules"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/auditlog"
	cloudfoundryv2 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/cloud/cloudfoundry"
	kubernetesv2 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/cloud/kubernetes"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/container/builtinmonitoringrule"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/container/monitoringrule"
	containertechnology "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/container/technology"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/custommetrics"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/customunit"
	dashboardsgeneral "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/dashboards/general"
	dashboardsallowlist "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/dashboards/image/allowlist"
	dashboardspresets "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/dashboards/presets"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/declarativegrouping"
	activegateupdates "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/deployment/activegate/updates"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/deployment/management/updatewindows"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/deployment/oneagent/defaultversion"
	oneagentupdates "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/deployment/oneagent/updates"
	diskanalytics "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/disk/analytics/extension"
	diskoptions "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/disk/options"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/eec/local"
	eecremote "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/eec/remote"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/eulasettings"
	networktraffic "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/exclude/network/traffic"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/geosettings"
	hostmonitoring "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/host/monitoring"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/host/monitoring/aixkernelextension"
	hostprocessgroupmonitoring "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/host/processgroups/monitoringstate"
	issuetracking "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/issuetracking/integration"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/logmonitoring/customlogsourcesettings"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/logmonitoring/logagentconfiguration"
	logcustomattributes "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/logmonitoring/logcustomattributes"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/logmonitoring/logdpprules"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/logmonitoring/logevents"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/logmonitoring/logsongrailactivate"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/logmonitoring/logstoragesettings"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/logmonitoring/schemalesslogmetric"
	sensitivedatamasking "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/logmonitoring/sensitivedatamaskingsettings"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/logmonitoring/timestampconfiguration"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/mainframe/txmonitoring"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/mainframe/txstartfilters"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/metric/metadata"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/metric/query"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredentities/generic/relation"
	generictypes "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredentities/generic/type"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/apache"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/dotnet"
	golang "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/go"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/iis"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/java"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/nginx"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/nodejs"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/opentracingnative"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/php"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/varnish"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoredtechnologies/wsmb"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/monitoring/slo/normalization"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/nettracer/traffic"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/oneagent/features"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/opentelemetrymetrics"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/osservicesmonitoring"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/ownership/teams"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/preferences/privacy"
	processmonitoring "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/process/monitoring"
	customprocessmonitoring "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/process/monitoring/custom"
	processavailability "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processavailability"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processgroup/advanceddetectionrule"
	workloaddetection "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processgroup/cloudapplication/workloaddetection"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processgroup/detectionflags"
	processgroupmonitoring "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processgroup/monitoring/state"
	processgroupsimpledetection "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processgroup/simpledetectionrule"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/processvisibility"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/remote/environment"
	rumcustomenablement "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/custom/enablement"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/hostheaders"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/ipdetermination"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/ipmappings"
	rummobileenablement "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/mobile/enablement"
	mobilerequesterrors "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/mobile/requesterrors"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/overloadprevention"
	rumprocessgroup "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/processgroup"
	rumproviderbreakdown "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/providerbreakdown"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/resourcetimingorigins"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/userexperiencescore"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/web/appdetection"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/web/beacondomainorigins"
	webappcustomerrors "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/web/customerrors"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/web/customrumjavascriptversion"
	rumwebenablement "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/web/enablement"
	webapprequesterrors "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/web/requesterrors"
	webappresourcecleanup "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/web/resourcecleanuprules"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/web/resourcetypes"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/web/rumjavascriptupdates"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/servicedetection/externalwebrequest"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/servicedetection/externalwebservice"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/servicedetection/fullwebrequest"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/servicedetection/fullwebservice"
	sessionreplaywebprivacy "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/sessionreplay/web/privacypreferences"
	sessionreplayresourcecapture "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/sessionreplay/web/resourcecapturing"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/settings/mutedrequests"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/synthetic/availability"
	browseroutagehandling "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/synthetic/browser/outagehandling"
	browserperformancethresholds "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/synthetic/browser/performancethresholds"
	httpcookies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/synthetic/http/cookies"
	httpoutagehandling "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/synthetic/http/outagehandling"
	httpperformancethresholds "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/synthetic/http/performancethresholds"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/tokens/tokensettings"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/usability/analytics"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/useractioncustommetrics"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/usersettings"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/iam/bindings"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/iam/groups"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/iam/permissions"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/iam/policies"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/iam/users"
	alertingv1 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/alerting"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/dashboards"
	maintenancev1 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/maintenance"
	managementzonesv1 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/managementzones"
	notificationsv1 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/notifications"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/requestnaming/order"
	locations "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/synthetic/locations/private"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/alerting"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/anomalies/frequentissues"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/anomalies/metricevents"
	service_anomalies_v2 "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/anomalies/services"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/availability/processgroupalerting"
	ddupool "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/ddupool"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/ibmmq/filters"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/ibmmq/imsbridges"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/ibmmq/queuemanagers"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/ibmmq/queuesharinggroups"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/keyrequests"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/networkzones"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/ansible"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/email"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/jira"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/opsgenie"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/pagerduty"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/servicenow"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/slack"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/trello"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/victorops"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/webhook"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/notifications/xmatters"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/slo"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/spans/attributes"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/spans/capture"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/spans/ctxprop"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/spans/entrypoints"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/spans/resattr"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/settings"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/settings/services/cache"

	v2managementzones "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/managementzones"

	application_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/applications"
	database_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/databaseservices"
	disk_event_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/diskevents"
	host_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/hosts"
	custom_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/metricevents"
	pg_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/processgroups"
	service_anomalies "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/anomalies/services"

	host_naming "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/naming/hosts"
	processgroup_naming "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/naming/processgroups"
	service_naming "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/naming/services"
	networkzone "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/networkzones"

	envparameters "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/failuredetection/environment/parameters"
	envrules "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/failuredetection/environment/rules"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/failuredetection/service/generalparameters"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/failuredetection/service/httpparameters"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/applications/mobile"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/applications/web"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/applications/web/dataprivacy"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/applications/web/detection"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/applications/web/errors"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/autotags"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/credentials/aws"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/credentials/azure"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/credentials/cloudfoundry"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/credentials/kubernetes"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/credentials/vault"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/customservices"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/dashboards/sharing"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/jsondashboards"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/requestattributes"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/requestnaming"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/synthetic/monitors/browser"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/synthetic/monitors/http"

	calculated_service_metrics "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v1/config/metrics/calculated/service"
	v2maintenance "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/v2/maintenance"
)

func NewResourceDescriptor[T settings.Settings](fn func(credentials *settings.Credentials) settings.CRUDService[T], dependencies ...Dependency) ResourceDescriptor {
	return ResourceDescriptor{
		Service: func(credentials *settings.Credentials) settings.CRUDService[settings.Settings] {
			return &settings.GenericCRUDService[T]{Service: cache.CRUD(fn(credentials))}
		},
		protoType:    newSettings(fn),
		Dependencies: dependencies,
	}
}

func newSettings[T settings.Settings](sfn func(credentials *settings.Credentials) settings.CRUDService[T]) T {
	var proto T
	return reflect.New(reflect.TypeOf(proto).Elem()).Interface().(T)
}

type ResourceDescriptor struct {
	Dependencies []Dependency
	Service      func(credentials *settings.Credentials) settings.CRUDService[settings.Settings]
	protoType    settings.Settings
	except       func(id string, name string) bool
}

func (me ResourceDescriptor) Specify(t notifications.Type) ResourceDescriptor {
	if notification, ok := me.protoType.(*notifications.Notification); ok {
		notification.Type = t
	}
	return me
}

func (me ResourceDescriptor) Except(except func(id string, name string) bool) ResourceDescriptor {
	me.except = except
	return me
}

func (me ResourceDescriptor) NewSettings() settings.Settings {
	res := reflect.New(reflect.TypeOf(me.protoType).Elem()).Interface().(settings.Settings)
	if notification, ok := res.(*notifications.Notification); ok {
		notification.Type = me.protoType.(*notifications.Notification).Type
	}
	return res
}

var AllResources = map[ResourceType]ResourceDescriptor{
	ResourceTypes.Alerting: NewResourceDescriptor(
		alerting.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
	),
	ResourceTypes.AnsibleTowerNotification: NewResourceDescriptor(
		ansible.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.AnsibleTower),
	ResourceTypes.ApplicationAnomalies: NewResourceDescriptor(application_anomalies.Service),
	ResourceTypes.ApplicationDataPrivacy: NewResourceDescriptor(
		dataprivacy.Service,
		Dependencies.ID(ResourceTypes.WebApplication),
	),
	ResourceTypes.ApplicationDetection: NewResourceDescriptor(
		detection.Service,
		Dependencies.ID(ResourceTypes.WebApplication),
	),
	ResourceTypes.ApplicationErrorRules: NewResourceDescriptor(
		errors.Service,
		Dependencies.ID(ResourceTypes.WebApplication),
	),
	ResourceTypes.AutoTag: NewResourceDescriptor(
		autotags.Service,
		Coalesce(Dependencies.Service),
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
		Coalesce(Dependencies.ProcessGroup),
		Coalesce(Dependencies.ProcessGroupInstance),
	),
	ResourceTypes.AWSCredentials: NewResourceDescriptor(
		aws.Service,
		Dependencies.ID(ResourceTypes.AWSCredentials),
	),
	ResourceTypes.AzureCredentials: NewResourceDescriptor(azure.Service),
	ResourceTypes.BrowserMonitor: NewResourceDescriptor(
		browser.Service,
		Dependencies.ID(ResourceTypes.SyntheticLocation),
		Dependencies.ID(ResourceTypes.WebApplication),
		Dependencies.ID(ResourceTypes.Credentials),
	).Except(func(id string, name string) bool {
		return strings.HasPrefix(name, "Monitor synchronizing credentials with")
	}),
	ResourceTypes.CalculatedServiceMetric: NewResourceDescriptor(
		calculated_service_metrics.Service,
		Dependencies.ManagementZone,
		Dependencies.RequestAttribute,
		Dependencies.ManagementZone,
		Dependencies.Service,
		Dependencies.Host,
		Dependencies.HostGroup,
		Dependencies.ProcessGroup,
		Dependencies.ProcessGroupInstance,
	),
	ResourceTypes.CloudFoundryCredentials: NewResourceDescriptor(cloudfoundry.Service),
	ResourceTypes.CustomAnomalies: NewResourceDescriptor(
		custom_anomalies.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
	).Except(func(id string, name string) bool {
		return strings.HasPrefix(id, "builtin:") || strings.HasPrefix(id, "ruxit.") || strings.HasPrefix(id, "dynatrace.") || strings.HasPrefix(id, "custom.remote.python.") || strings.HasPrefix(id, "custom.python.")
	}),
	ResourceTypes.CustomAppAnomalies: NewResourceDescriptor(
		custom_app_anomalies.Service,
		Coalesce(Dependencies.DeviceApplicationMethod),
		Coalesce(Dependencies.CustomApplication),
	),
	ResourceTypes.CustomAppCrashRate: NewResourceDescriptor(
		custom_app_crash_rate.Service,
		Coalesce(Dependencies.CustomApplication),
	),
	ResourceTypes.MobileAppAnomalies: NewResourceDescriptor(
		mobile_app_anomalies.Service,
		Coalesce(Dependencies.DeviceApplicationMethod),
		Coalesce(Dependencies.MobileApplication),
	),
	ResourceTypes.MobileAppCrashRate: NewResourceDescriptor(
		mobile_app_crash_rate.Service,
		Coalesce(Dependencies.MobileApplication),
	),
	ResourceTypes.WebAppAnomalies: NewResourceDescriptor(
		web_app_anomalies.Service,
		Coalesce(Dependencies.Application),
		Coalesce(Dependencies.ApplicationMethod),
	),
	ResourceTypes.CustomService: NewResourceDescriptor(customservices.Service),
	ResourceTypes.Credentials: NewResourceDescriptor(
		vault.Service,
		Dependencies.ID(ResourceTypes.Credentials),
	),
	ResourceTypes.JSONDashboard: NewResourceDescriptor(
		jsondashboards.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
		Dependencies.ManagementZone,
		// Dependencies.Service,
		Dependencies.ID(ResourceTypes.SLO),
		Dependencies.ID(ResourceTypes.WebApplication),
		Dependencies.ID(ResourceTypes.MobileApplication),
		Dependencies.ID(ResourceTypes.SyntheticLocation),
		Dependencies.ID(ResourceTypes.HTTPMonitor),
		Dependencies.ID(ResourceTypes.CalculatedServiceMetric),
		Dependencies.ID(ResourceTypes.BrowserMonitor),
	),
	ResourceTypes.DashboardSharing: NewResourceDescriptor(
		sharing.Service,
		Dependencies.ID(ResourceTypes.JSONDashboard),
	),
	ResourceTypes.DatabaseAnomalies:  NewResourceDescriptor(database_anomalies.Service),
	ResourceTypes.DiskEventAnomalies: NewResourceDescriptor(disk_event_anomalies.Service),
	ResourceTypes.DiskAnomaliesV2: NewResourceDescriptor(
		disk_anomalies_v2.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.DiskSpecificAnomaliesV2: NewResourceDescriptor(
		disk_specific_anomalies_v2.Service,
		Coalesce(Dependencies.Disk),
	),
	ResourceTypes.EmailNotification: NewResourceDescriptor(
		email.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.Email),
	ResourceTypes.FrequentIssues: NewResourceDescriptor(frequentissues.Service),
	ResourceTypes.HostAnomalies:  NewResourceDescriptor(host_anomalies.Service),
	ResourceTypes.HostAnomaliesV2: NewResourceDescriptor(
		host_anomalies_v2.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.HTTPMonitor: NewResourceDescriptor(
		http.Service,
		Dependencies.ID(ResourceTypes.SyntheticLocation),
		Dependencies.ID(ResourceTypes.WebApplication),
		Dependencies.ID(ResourceTypes.Credentials),
	),
	ResourceTypes.HostNaming:   NewResourceDescriptor(host_naming.Service),
	ResourceTypes.IBMMQFilters: NewResourceDescriptor(filters.Service),
	ResourceTypes.IMSBridge:    NewResourceDescriptor(imsbridges.Service),
	ResourceTypes.JiraNotification: NewResourceDescriptor(
		jira.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.Jira),
	ResourceTypes.KeyRequests: NewResourceDescriptor(
		keyrequests.Service,
		Coalesce(Dependencies.Service),
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
		Coalesce(Dependencies.ProcessGroup),
		Coalesce(Dependencies.ProcessGroupInstance),
	),
	ResourceTypes.KubernetesCredentials: NewResourceDescriptor(kubernetes.Service),
	ResourceTypes.Maintenance: NewResourceDescriptor(
		v2maintenance.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
	),
	ResourceTypes.ManagementZoneV2: NewResourceDescriptor(v2managementzones.Service),
	ResourceTypes.MetricEvents:     NewResourceDescriptor(metricevents.Service),
	ResourceTypes.MobileApplication: NewResourceDescriptor(
		mobile.Service,
		Dependencies.ID(ResourceTypes.RequestAttribute),
	),
	ResourceTypes.MutedRequests: NewResourceDescriptor(
		mutedrequests.Service,
		Coalesce(Dependencies.Service),
	),
	ResourceTypes.NetworkZone:  NewResourceDescriptor(networkzone.Service),
	ResourceTypes.NetworkZones: NewResourceDescriptor(networkzones.Service),
	ResourceTypes.OpsGenieNotification: NewResourceDescriptor(
		opsgenie.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.OpsGenie),
	ResourceTypes.PagerDutyNotification: NewResourceDescriptor(
		pagerduty.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.PagerDuty),
	ResourceTypes.ProcessGroupNaming: NewResourceDescriptor(processgroup_naming.Service),
	ResourceTypes.QueueManager:       NewResourceDescriptor(queuemanagers.Service),
	ResourceTypes.RequestAttribute: NewResourceDescriptor(
		requestattributes.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
		Coalesce(Dependencies.ProcessGroup),
		Coalesce(Dependencies.ProcessGroupInstance),
		Coalesce(Dependencies.Service),
	),
	ResourceTypes.RequestNaming: NewResourceDescriptor(
		requestnaming.Service,
		Dependencies.RequestAttribute,
	),
	ResourceTypes.ResourceAttributes: NewResourceDescriptor(resattr.Service),
	ResourceTypes.ServiceAnomalies:   NewResourceDescriptor(service_anomalies.Service),
	ResourceTypes.ServiceAnomaliesV2: NewResourceDescriptor(
		service_anomalies_v2.Service,
		Coalesce(Dependencies.ServiceMethod),
		Coalesce(Dependencies.Service),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.ServiceNaming: NewResourceDescriptor(service_naming.Service),
	ResourceTypes.ServiceNowNotification: NewResourceDescriptor(
		servicenow.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.ServiceNow),
	ResourceTypes.SlackNotification: NewResourceDescriptor(
		slack.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.Slack),
	ResourceTypes.SLO: NewResourceDescriptor(
		slo.Service,
		Dependencies.ManagementZone,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
		Dependencies.ID(ResourceTypes.CalculatedServiceMetric),
	),
	ResourceTypes.SpanAttribute:          NewResourceDescriptor(attributes.Service),
	ResourceTypes.SpanCaptureRule:        NewResourceDescriptor(capture.Service),
	ResourceTypes.SpanContextPropagation: NewResourceDescriptor(ctxprop.Service),
	ResourceTypes.SpanEntryPoint:         NewResourceDescriptor(entrypoints.Service),
	ResourceTypes.SyntheticLocation:      NewResourceDescriptor(locations.Service),
	ResourceTypes.TrelloNotification: NewResourceDescriptor(
		trello.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.Trello),
	ResourceTypes.VictorOpsNotification: NewResourceDescriptor(
		victorops.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.VictorOps),
	ResourceTypes.WebApplication: NewResourceDescriptor(
		web.Service,
		Dependencies.ID(ResourceTypes.RequestAttribute),
	),
	ResourceTypes.WebHookNotification: NewResourceDescriptor(
		webhook.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.WebHook),
	ResourceTypes.XMattersNotification: NewResourceDescriptor(
		xmatters.Service,
		Dependencies.ID(ResourceTypes.Alerting),
	).Specify(notifications.Types.XMatters),

	ResourceTypes.MaintenanceWindow: NewResourceDescriptor(
		maintenancev1.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
	),
	ResourceTypes.ManagementZone: NewResourceDescriptor(managementzonesv1.Service),
	ResourceTypes.Dashboard: NewResourceDescriptor(
		dashboards.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
		Dependencies.ManagementZone,
		Dependencies.ID(ResourceTypes.SLO),
		Dependencies.ID(ResourceTypes.WebApplication),
		Dependencies.ID(ResourceTypes.SyntheticLocation),
	),
	ResourceTypes.Notification: NewResourceDescriptor(
		notificationsv1.Service,
		Dependencies.LegacyID(ResourceTypes.Alerting),
	),
	ResourceTypes.QueueSharingGroups: NewResourceDescriptor(queuesharinggroups.Service),
	ResourceTypes.AlertingProfile: NewResourceDescriptor(
		alertingv1.Service,
		Dependencies.LegacyID(ResourceTypes.ManagementZoneV2),
	),
	ResourceTypes.RequestNamings: NewResourceDescriptor(
		order.Service,
		Dependencies.ID(ResourceTypes.RequestNaming),
	),
	ResourceTypes.IAMUser:           NewResourceDescriptor(users.Service),
	ResourceTypes.IAMGroup:          NewResourceDescriptor(groups.Service),
	ResourceTypes.IAMPermission:     NewResourceDescriptor(permissions.Service),
	ResourceTypes.IAMPolicy:         NewResourceDescriptor(policies.Service),
	ResourceTypes.IAMPolicyBindings: NewResourceDescriptor(bindings.Service),
	ResourceTypes.DDUPool:           NewResourceDescriptor(ddupool.Service),
	ResourceTypes.ProcessGroupAnomalies: NewResourceDescriptor(
		pg_anomalies.Service,
		Coalesce(Dependencies.ProcessGroup),
		Coalesce(Dependencies.ProcessGroupInstance),
	),
	ResourceTypes.ProcessGroupAlerting: NewResourceDescriptor(
		processgroupalerting.Service,
		Coalesce(Dependencies.ProcessGroup),
	),
	ResourceTypes.DatabaseAnomaliesV2: NewResourceDescriptor(
		database_anomalies_v2.Service,
		Coalesce(Dependencies.ServiceMethod),
		Coalesce(Dependencies.Service),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.ProcessMonitoringRule: NewResourceDescriptor(
		customprocessmonitoring.Service,
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.ProcessMonitoring: NewResourceDescriptor(
		processmonitoring.Service,
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.ProcessAvailability: NewResourceDescriptor(
		processavailability.Service,
		Coalesce(Dependencies.HostGroup),
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.AdvancedProcessGroupDetectionRule: NewResourceDescriptor(advanceddetectionrule.Service),
	ResourceTypes.ConnectivityAlerts: NewResourceDescriptor(
		connectivityalerts.Service,
		Coalesce(Dependencies.ProcessGroup),
	),
	ResourceTypes.DeclarativeGrouping: NewResourceDescriptor(
		declarativegrouping.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.HostMonitoring: NewResourceDescriptor(
		hostmonitoring.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.HostProcessGroupMonitoring: NewResourceDescriptor(
		hostprocessgroupmonitoring.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.ProcessGroup),
	),
	ResourceTypes.RUMIPLocations: NewResourceDescriptor(ipmappings.Service),
	ResourceTypes.CustomAppEnablement: NewResourceDescriptor(
		rumcustomenablement.Service,
		Coalesce(Dependencies.MobileApplication),
	),
	ResourceTypes.MobileAppEnablement: NewResourceDescriptor(
		rummobileenablement.Service,
		Coalesce(Dependencies.MobileApplication),
	),
	ResourceTypes.WebAppEnablement: NewResourceDescriptor(
		rumwebenablement.Service,
		Coalesce(Dependencies.Application),
	),
	ResourceTypes.RUMProcessGroup: NewResourceDescriptor(
		rumprocessgroup.Service,
		Coalesce(Dependencies.ProcessGroup),
	),
	ResourceTypes.RUMProviderBreakdown:  NewResourceDescriptor(rumproviderbreakdown.Service),
	ResourceTypes.UserExperienceScore:   NewResourceDescriptor(userexperiencescore.Service),
	ResourceTypes.WebAppResourceCleanup: NewResourceDescriptor(webappresourcecleanup.Service),
	ResourceTypes.UpdateWindows:         NewResourceDescriptor(updatewindows.Service),
	ResourceTypes.ProcessGroupDetectionFlags: NewResourceDescriptor(
		detectionflags.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.ProcessGroupMonitoring: NewResourceDescriptor(
		processgroupmonitoring.Service,
		Coalesce(Dependencies.ProcessGroup),
	),
	ResourceTypes.ProcessGroupSimpleDetection: NewResourceDescriptor(processgroupsimpledetection.Service),
	ResourceTypes.LogMetrics:                  NewResourceDescriptor(schemalesslogmetric.Service),
	ResourceTypes.BrowserMonitorPerformanceThresholds: NewResourceDescriptor(
		browserperformancethresholds.Service,
		Coalesce(Dependencies.SyntheticTest),
	),
	ResourceTypes.HttpMonitorPerformanceThresholds: NewResourceDescriptor(
		httpperformancethresholds.Service,
		Coalesce(Dependencies.HttpCheck),
	),
	ResourceTypes.HttpMonitorCookies: NewResourceDescriptor(
		httpcookies.Service,
		Coalesce(Dependencies.HttpCheck),
	),
	ResourceTypes.SessionReplayWebPrivacy: NewResourceDescriptor(
		sessionreplaywebprivacy.Service,
		Coalesce(Dependencies.Application),
	),
	ResourceTypes.SessionReplayResourceCapture: NewResourceDescriptor(
		sessionreplayresourcecapture.Service,
		Coalesce(Dependencies.Application),
	),
	ResourceTypes.UsabilityAnalytics: NewResourceDescriptor(
		analytics.Service,
		Coalesce(Dependencies.Application),
	),
	ResourceTypes.SyntheticAvailability: NewResourceDescriptor(availability.Service),
	ResourceTypes.BrowserMonitorOutageHandling: NewResourceDescriptor(
		browseroutagehandling.Service,
		Coalesce(Dependencies.SyntheticTest),
	),
	ResourceTypes.HttpMonitorOutageHandling: NewResourceDescriptor(
		httpoutagehandling.Service,
		Coalesce(Dependencies.HttpCheck),
	),
	ResourceTypes.CloudAppWorkloadDetection:      NewResourceDescriptor(workloaddetection.Service),
	ResourceTypes.MainframeTransactionMonitoring: NewResourceDescriptor(txmonitoring.Service),
	ResourceTypes.MonitoredTechnologiesApache: NewResourceDescriptor(
		apache.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesDotNet: NewResourceDescriptor(
		dotnet.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesGo: NewResourceDescriptor(
		golang.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesIIS: NewResourceDescriptor(
		iis.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesJava: NewResourceDescriptor(
		java.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesNGINX: NewResourceDescriptor(
		nginx.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesNodeJS: NewResourceDescriptor(
		nodejs.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesOpenTracing: NewResourceDescriptor(
		opentracingnative.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesPHP: NewResourceDescriptor(
		php.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesVarnish: NewResourceDescriptor(
		varnish.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MonitoredTechnologiesWSMB: NewResourceDescriptor(
		wsmb.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.ProcessVisibility: NewResourceDescriptor(
		processvisibility.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.RUMHostHeaders:     NewResourceDescriptor(hostheaders.Service),
	ResourceTypes.RUMIPDetermination: NewResourceDescriptor(ipdetermination.Service),
	ResourceTypes.MobileAppRequestErrors: NewResourceDescriptor(
		mobilerequesterrors.Service,
		Coalesce(Dependencies.MobileApplication),
		Coalesce(Dependencies.CustomApplication),
	),
	ResourceTypes.TransactionStartFilters: NewResourceDescriptor(txstartfilters.Service),
	ResourceTypes.OneAgentFeatures: NewResourceDescriptor(
		features.Service,
		Coalesce(Dependencies.ProcessGroup),
		Coalesce(Dependencies.ProcessGroupInstance),
	),
	ResourceTypes.RUMOverloadPrevention:  NewResourceDescriptor(overloadprevention.Service),
	ResourceTypes.RUMAdvancedCorrelation: NewResourceDescriptor(resourcetimingorigins.Service),
	ResourceTypes.WebAppBeaconOrigins:    NewResourceDescriptor(beacondomainorigins.Service),
	ResourceTypes.WebAppResourceTypes:    NewResourceDescriptor(resourcetypes.Service),
	ResourceTypes.GenericTypes:           NewResourceDescriptor(generictypes.Service),
	ResourceTypes.GenericRelationships:   NewResourceDescriptor(relation.Service),
	ResourceTypes.SLONormalization:       NewResourceDescriptor(normalization.Service),
	ResourceTypes.DataPrivacy: NewResourceDescriptor(
		privacy.Service,
		Coalesce(Dependencies.Application),
	),
	ResourceTypes.ServiceFailure: NewResourceDescriptor(
		generalparameters.Service,
		Coalesce(Dependencies.Service),
		Dependencies.ID(ResourceTypes.RequestAttribute),
	),
	ResourceTypes.ServiceHTTPFailure: NewResourceDescriptor(
		httpparameters.Service,
		Coalesce(Dependencies.Service),
	),
	ResourceTypes.DiskOptions: NewResourceDescriptor(
		diskoptions.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.OSServices: NewResourceDescriptor(
		osservicesmonitoring.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.ExtensionExecutionController: NewResourceDescriptor(
		local.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.NetTracerTraffic: NewResourceDescriptor(
		traffic.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.AIXExtension: NewResourceDescriptor(
		aixkernelextension.Service,
		Coalesce(Dependencies.Host),
	),
	ResourceTypes.MetricMetadata:  NewResourceDescriptor(metadata.Service),
	ResourceTypes.MetricQuery:     NewResourceDescriptor(query.Service),
	ResourceTypes.ActiveGateToken: NewResourceDescriptor(activegatetoken.Service),
	ResourceTypes.AuditLog:        NewResourceDescriptor(auditlog.Service),
	ResourceTypes.K8sClusterAnomalies: NewResourceDescriptor(
		cluster.Service,
		Coalesce(Dependencies.K8sCluster),
	),
	ResourceTypes.K8sNamespaceAnomalies: NewResourceDescriptor(
		namespace.Service,
		Coalesce(Dependencies.CloudApplicationNamespace),
		Coalesce(Dependencies.K8sCluster),
	),
	ResourceTypes.K8sNodeAnomalies: NewResourceDescriptor(
		node.Service,
		Coalesce(Dependencies.K8sCluster),
	),
	ResourceTypes.K8sWorkloadAnomalies: NewResourceDescriptor(
		workload.Service,
		Coalesce(Dependencies.CloudApplicationNamespace),
		Coalesce(Dependencies.K8sCluster),
	),
	ResourceTypes.ContainerBuiltinRule: NewResourceDescriptor(builtinmonitoringrule.Service),
	ResourceTypes.ContainerRule:        NewResourceDescriptor(monitoringrule.Service),
	ResourceTypes.ContainerTechnology: NewResourceDescriptor(
		containertechnology.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.RemoteEnvironments: NewResourceDescriptor(environment.Service),
	ResourceTypes.WebAppCustomErrors: NewResourceDescriptor(
		webappcustomerrors.Service,
		Coalesce(Dependencies.Application),
	),
	ResourceTypes.WebAppRequestErrors: NewResourceDescriptor(
		webapprequesterrors.Service,
		Coalesce(Dependencies.Application),
	),
	ResourceTypes.UserSettings:               NewResourceDescriptor(usersettings.Service),
	ResourceTypes.DashboardsGeneral:          NewResourceDescriptor(dashboardsgeneral.Service),
	ResourceTypes.DashboardsPresets:          NewResourceDescriptor(dashboardspresets.Service),
	ResourceTypes.LogProcessing:              NewResourceDescriptor(logdpprules.Service),
	ResourceTypes.LogEvents:                  NewResourceDescriptor(logevents.Service),
	ResourceTypes.LogTimestamp:               NewResourceDescriptor(timestampconfiguration.Service),
	ResourceTypes.LogGrail:                   NewResourceDescriptor(logsongrailactivate.Service),
	ResourceTypes.LogCustomAttribute:         NewResourceDescriptor(logcustomattributes.Service),
	ResourceTypes.LogSensitiveDataMasking:    NewResourceDescriptor(sensitivedatamasking.Service),
	ResourceTypes.EULASettings:               NewResourceDescriptor(eulasettings.Service),
	ResourceTypes.APIDetectionRules:          NewResourceDescriptor(apidetection.Service),
	ResourceTypes.ServiceExternalWebRequest:  NewResourceDescriptor(externalwebrequest.Service),
	ResourceTypes.ServiceExternalWebService:  NewResourceDescriptor(externalwebservice.Service),
	ResourceTypes.ServiceFullWebRequest:      NewResourceDescriptor(fullwebrequest.Service),
	ResourceTypes.ServiceFullWebService:      NewResourceDescriptor(fullwebservice.Service),
	ResourceTypes.DashboardsAllowlist:        NewResourceDescriptor(dashboardsallowlist.Service),
	ResourceTypes.FailureDetectionParameters: NewResourceDescriptor(envparameters.Service),
	ResourceTypes.FailureDetectionRules: NewResourceDescriptor(
		envrules.Service,
		Dependencies.ID(ResourceTypes.FailureDetectionParameters),
	),
	ResourceTypes.LogOneAgent:              NewResourceDescriptor(logagentconfiguration.Service),
	ResourceTypes.IssueTracking:            NewResourceDescriptor(issuetracking.Service),
	ResourceTypes.GeolocationSettings:      NewResourceDescriptor(geosettings.Service),
	ResourceTypes.UserSessionCustomMetrics: NewResourceDescriptor(custommetrics.Service),
	ResourceTypes.CustomUnits:              NewResourceDescriptor(customunit.Service),
	ResourceTypes.DiskAnalytics:            NewResourceDescriptor(diskanalytics.Service),
	ResourceTypes.NetworkTraffic:           NewResourceDescriptor(networktraffic.Service),
	ResourceTypes.TokenSettings:            NewResourceDescriptor(tokensettings.Service),
	ResourceTypes.ExtensionExecutionRemote: NewResourceDescriptor(
		eecremote.Service,
		Coalesce(Dependencies.EnvironmentActiveGate),
	),
	ResourceTypes.K8sPVCAnomalies: NewResourceDescriptor(
		pvc.Service,
		Coalesce(Dependencies.CloudApplicationNamespace),
		Coalesce(Dependencies.K8sCluster),
	),
	ResourceTypes.UserActionCustomMetrics: NewResourceDescriptor(useractioncustommetrics.Service),
	ResourceTypes.WebAppJavascriptVersion: NewResourceDescriptor(customrumjavascriptversion.Service),
	ResourceTypes.WebAppJavascriptUpdates: NewResourceDescriptor(
		rumjavascriptupdates.Service,
		Coalesce(Dependencies.Application),
	),
	ResourceTypes.OpenTelemetryMetrics: NewResourceDescriptor(opentelemetrymetrics.Service),
	ResourceTypes.ActiveGateUpdates: NewResourceDescriptor(
		activegateupdates.Service,
		Coalesce(Dependencies.EnvironmentActiveGate),
	),
	ResourceTypes.OneAgentDefaultVersion: NewResourceDescriptor(defaultversion.Service),
	ResourceTypes.OneAgentUpdates: NewResourceDescriptor(
		oneagentupdates.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
		Dependencies.ID(ResourceTypes.UpdateWindows),
	),
	ResourceTypes.LogStorage: NewResourceDescriptor(
		logstoragesettings.Service,
		Coalesce(Dependencies.Host),
		Coalesce(Dependencies.HostGroup),
	),
	ResourceTypes.OwnershipTeams:  NewResourceDescriptor(teams.Service),
	ResourceTypes.LogCustomSource: NewResourceDescriptor(customlogsourcesettings.Service),
	ResourceTypes.ApplicationDetectionV2: NewResourceDescriptor(
		appdetection.Service,
		Coalesce(Dependencies.Application),
	),
	ResourceTypes.Kubernetes: NewResourceDescriptor(
		kubernetesv2.Service,
		Coalesce(Dependencies.K8sCluster),
	),
	ResourceTypes.CloudFoundry: NewResourceDescriptor(cloudfoundryv2.Service),
}

var BlackListedResources = []ResourceType{
	ResourceTypes.MaintenanceWindow, // legacy
	ResourceTypes.ManagementZone,    // legacy
	ResourceTypes.Notification,      // legacy
	ResourceTypes.AlertingProfile,   // legacy
	ResourceTypes.Dashboard,         // taken care of dynatrace_json_dashboard
	ResourceTypes.IAMUser,           // not sure whether to migrate
	ResourceTypes.IAMGroup,          // not sure whether to migrate
	ResourceTypes.IAMPermission,     // not sure whether to migrate
	ResourceTypes.IAMPolicy,         // not sure whether to migrate
	ResourceTypes.IAMPolicyBindings, // not sure whether to migrate

	// excluding by default
	ResourceTypes.JSONDashboard, // may replace dynatrace_dashboard in the future
	ResourceTypes.DashboardSharing,

	ResourceTypes.ProcessGroupAnomalies, // there could be 100k process groups

	ResourceTypes.WebAppEnablement,             // overlap with ResourceTypes.MobileApplication
	ResourceTypes.MobileAppEnablement,          // overlap with ResourceTypes.MobileApplication
	ResourceTypes.CustomAppEnablement,          // overlap with ResourceTypes.MobileApplication
	ResourceTypes.SessionReplayWebPrivacy,      // overlap with ResourceTypes.ApplicationDataPrivacy
	ResourceTypes.SessionReplayResourceCapture, // overlap with ResourceTypes.WebApplication
	ResourceTypes.BrowserMonitorOutageHandling, // overlap with ResourceTypes.BrowserMonitor
	ResourceTypes.HttpMonitorOutageHandling,    // overlap with ResourceTypes.HTTPMonitor
	ResourceTypes.DataPrivacy,                  // overlap with ResourceTypes.ApplicationDataPrivacy
	ResourceTypes.WebAppCustomErrors,           // overlap with ResourceTypes.ApplicationErrorRules
	ResourceTypes.WebAppRequestErrors,          // overlap with ResourceTypes.ApplicationErrorRules
	ResourceTypes.UserSettings,                 // requires personal token
	ResourceTypes.LogGrail,                     // phased rollout
	ResourceTypes.ApplicationDetectionV2,       // overlap with ResourceTypes.ApplicationDetection

	ResourceTypes.KubernetesCredentials,   // overlap with Settings 2.0 ResourceTypes.Kubernetes
	ResourceTypes.CloudFoundryCredentials, // overlap with Settings 2.0 ResourceTypes.CloudFoundry
}

func Service(credentials *settings.Credentials, resourceType ResourceType) settings.CRUDService[settings.Settings] {
	return AllResources[resourceType].Service(credentials)
}

// func DSService(credentials *settings.Credentials, dataSourceType DataSourceType) settings.RService[settings.Settings] {
// 	return AllDataSources[dataSourceType].Service(credentials)
// }
