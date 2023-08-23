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

package providerbreakdown

import (
	providerbreakdown "github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/rum/providerbreakdown/settings"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/settings"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/settings/services/settings20"
)

const SchemaVersion = "1.0.6"
const SchemaID = "builtin:rum.provider-breakdown"

func Service(credentials *settings.Credentials) settings.CRUDService[*providerbreakdown.Settings] {
	return settings20.Service[*providerbreakdown.Settings](credentials, SchemaID, SchemaVersion)
}