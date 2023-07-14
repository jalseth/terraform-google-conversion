// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package looker

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const LookerInstanceAssetType string = "looker.googleapis.com/Instance"

func ResourceConverterLookerInstance() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: LookerInstanceAssetType,
		Convert:   GetLookerInstanceCaiObject,
	}
}

func GetLookerInstanceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//looker.googleapis.com/projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetLookerInstanceApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: LookerInstanceAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/looker/v1/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetLookerInstanceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	adminSettingsProp, err := expandLookerInstanceAdminSettings(d.Get("admin_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("admin_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(adminSettingsProp)) && (ok || !reflect.DeepEqual(v, adminSettingsProp)) {
		obj["adminSettings"] = adminSettingsProp
	}
	consumerNetworkProp, err := expandLookerInstanceConsumerNetwork(d.Get("consumer_network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("consumer_network"); !tpgresource.IsEmptyValue(reflect.ValueOf(consumerNetworkProp)) && (ok || !reflect.DeepEqual(v, consumerNetworkProp)) {
		obj["consumerNetwork"] = consumerNetworkProp
	}
	denyMaintenancePeriodProp, err := expandLookerInstanceDenyMaintenancePeriod(d.Get("deny_maintenance_period"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deny_maintenance_period"); !tpgresource.IsEmptyValue(reflect.ValueOf(denyMaintenancePeriodProp)) && (ok || !reflect.DeepEqual(v, denyMaintenancePeriodProp)) {
		obj["denyMaintenancePeriod"] = denyMaintenancePeriodProp
	}
	encryptionConfigProp, err := expandLookerInstanceEncryptionConfig(d.Get("encryption_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("encryption_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptionConfigProp)) && (ok || !reflect.DeepEqual(v, encryptionConfigProp)) {
		obj["encryptionConfig"] = encryptionConfigProp
	}
	maintenanceWindowProp, err := expandLookerInstanceMaintenanceWindow(d.Get("maintenance_window"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_window"); !tpgresource.IsEmptyValue(reflect.ValueOf(maintenanceWindowProp)) && (ok || !reflect.DeepEqual(v, maintenanceWindowProp)) {
		obj["maintenanceWindow"] = maintenanceWindowProp
	}
	oauthConfigProp, err := expandLookerInstanceOauthConfig(d.Get("oauth_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("oauth_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(oauthConfigProp)) && (ok || !reflect.DeepEqual(v, oauthConfigProp)) {
		obj["oauthConfig"] = oauthConfigProp
	}
	platformEditionProp, err := expandLookerInstancePlatformEdition(d.Get("platform_edition"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("platform_edition"); !tpgresource.IsEmptyValue(reflect.ValueOf(platformEditionProp)) && (ok || !reflect.DeepEqual(v, platformEditionProp)) {
		obj["platformEdition"] = platformEditionProp
	}
	privateIpEnabledProp, err := expandLookerInstancePrivateIpEnabled(d.Get("private_ip_enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("private_ip_enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(privateIpEnabledProp)) && (ok || !reflect.DeepEqual(v, privateIpEnabledProp)) {
		obj["privateIpEnabled"] = privateIpEnabledProp
	}
	publicIpEnabledProp, err := expandLookerInstancePublicIpEnabled(d.Get("public_ip_enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("public_ip_enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(publicIpEnabledProp)) && (ok || !reflect.DeepEqual(v, publicIpEnabledProp)) {
		obj["publicIpEnabled"] = publicIpEnabledProp
	}
	reservedRangeProp, err := expandLookerInstanceReservedRange(d.Get("reserved_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reserved_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(reservedRangeProp)) && (ok || !reflect.DeepEqual(v, reservedRangeProp)) {
		obj["reservedRange"] = reservedRangeProp
	}
	userMetadataProp, err := expandLookerInstanceUserMetadata(d.Get("user_metadata"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("user_metadata"); !tpgresource.IsEmptyValue(reflect.ValueOf(userMetadataProp)) && (ok || !reflect.DeepEqual(v, userMetadataProp)) {
		obj["userMetadata"] = userMetadataProp
	}

	return obj, nil
}

func expandLookerInstanceAdminSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowedEmailDomains, err := expandLookerInstanceAdminSettingsAllowedEmailDomains(original["allowed_email_domains"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedEmailDomains); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedEmailDomains"] = transformedAllowedEmailDomains
	}

	return transformed, nil
}

func expandLookerInstanceAdminSettingsAllowedEmailDomains(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceConsumerNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceDenyMaintenancePeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStartDate, err := expandLookerInstanceDenyMaintenancePeriodStartDate(original["start_date"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartDate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startDate"] = transformedStartDate
	}

	transformedEndDate, err := expandLookerInstanceDenyMaintenancePeriodEndDate(original["end_date"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndDate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["endDate"] = transformedEndDate
	}

	transformedTime, err := expandLookerInstanceDenyMaintenancePeriodTime(original["time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["time"] = transformedTime
	}

	return transformed, nil
}

func expandLookerInstanceDenyMaintenancePeriodStartDate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedYear, err := expandLookerInstanceDenyMaintenancePeriodStartDateYear(original["year"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedYear); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["year"] = transformedYear
	}

	transformedMonth, err := expandLookerInstanceDenyMaintenancePeriodStartDateMonth(original["month"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonth); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["month"] = transformedMonth
	}

	transformedDay, err := expandLookerInstanceDenyMaintenancePeriodStartDateDay(original["day"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["day"] = transformedDay
	}

	return transformed, nil
}

func expandLookerInstanceDenyMaintenancePeriodStartDateYear(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceDenyMaintenancePeriodStartDateMonth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceDenyMaintenancePeriodStartDateDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceDenyMaintenancePeriodEndDate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedYear, err := expandLookerInstanceDenyMaintenancePeriodEndDateYear(original["year"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedYear); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["year"] = transformedYear
	}

	transformedMonth, err := expandLookerInstanceDenyMaintenancePeriodEndDateMonth(original["month"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonth); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["month"] = transformedMonth
	}

	transformedDay, err := expandLookerInstanceDenyMaintenancePeriodEndDateDay(original["day"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["day"] = transformedDay
	}

	return transformed, nil
}

func expandLookerInstanceDenyMaintenancePeriodEndDateYear(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceDenyMaintenancePeriodEndDateMonth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceDenyMaintenancePeriodEndDateDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceDenyMaintenancePeriodTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHours, err := expandLookerInstanceDenyMaintenancePeriodTimeHours(original["hours"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHours); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hours"] = transformedHours
	}

	transformedMinutes, err := expandLookerInstanceDenyMaintenancePeriodTimeMinutes(original["minutes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinutes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minutes"] = transformedMinutes
	}

	transformedSeconds, err := expandLookerInstanceDenyMaintenancePeriodTimeSeconds(original["seconds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeconds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["seconds"] = transformedSeconds
	}

	transformedNanos, err := expandLookerInstanceDenyMaintenancePeriodTimeNanos(original["nanos"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNanos); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nanos"] = transformedNanos
	}

	return transformed, nil
}

func expandLookerInstanceDenyMaintenancePeriodTimeHours(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceDenyMaintenancePeriodTimeMinutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceDenyMaintenancePeriodTimeSeconds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceDenyMaintenancePeriodTimeNanos(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceEncryptionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandLookerInstanceEncryptionConfigKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	transformedKmsKeyState, err := expandLookerInstanceEncryptionConfigKmsKeyState(original["kms_key_state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyState"] = transformedKmsKeyState
	}

	transformedKmsKeyNameVersion, err := expandLookerInstanceEncryptionConfigKmsKeyNameVersion(original["kms_key_name_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyNameVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyNameVersion"] = transformedKmsKeyNameVersion
	}

	return transformed, nil
}

func expandLookerInstanceEncryptionConfigKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceEncryptionConfigKmsKeyState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceEncryptionConfigKmsKeyNameVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceMaintenanceWindow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDayOfWeek, err := expandLookerInstanceMaintenanceWindowDayOfWeek(original["day_of_week"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDayOfWeek); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dayOfWeek"] = transformedDayOfWeek
	}

	transformedStartTime, err := expandLookerInstanceMaintenanceWindowStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	return transformed, nil
}

func expandLookerInstanceMaintenanceWindowDayOfWeek(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceMaintenanceWindowStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHours, err := expandLookerInstanceMaintenanceWindowStartTimeHours(original["hours"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHours); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hours"] = transformedHours
	}

	transformedMinutes, err := expandLookerInstanceMaintenanceWindowStartTimeMinutes(original["minutes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinutes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minutes"] = transformedMinutes
	}

	transformedSeconds, err := expandLookerInstanceMaintenanceWindowStartTimeSeconds(original["seconds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeconds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["seconds"] = transformedSeconds
	}

	transformedNanos, err := expandLookerInstanceMaintenanceWindowStartTimeNanos(original["nanos"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNanos); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nanos"] = transformedNanos
	}

	return transformed, nil
}

func expandLookerInstanceMaintenanceWindowStartTimeHours(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceMaintenanceWindowStartTimeMinutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceMaintenanceWindowStartTimeSeconds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceMaintenanceWindowStartTimeNanos(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceOauthConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedClientId, err := expandLookerInstanceOauthConfigClientId(original["client_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientId"] = transformedClientId
	}

	transformedClientSecret, err := expandLookerInstanceOauthConfigClientSecret(original["client_secret"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientSecret); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientSecret"] = transformedClientSecret
	}

	return transformed, nil
}

func expandLookerInstanceOauthConfigClientId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceOauthConfigClientSecret(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstancePlatformEdition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstancePrivateIpEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstancePublicIpEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceReservedRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceUserMetadata(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAdditionalViewerUserCount, err := expandLookerInstanceUserMetadataAdditionalViewerUserCount(original["additional_viewer_user_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdditionalViewerUserCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["additionalViewerUserCount"] = transformedAdditionalViewerUserCount
	}

	transformedAdditionalStandardUserCount, err := expandLookerInstanceUserMetadataAdditionalStandardUserCount(original["additional_standard_user_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdditionalStandardUserCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["additionalStandardUserCount"] = transformedAdditionalStandardUserCount
	}

	transformedAdditionalDeveloperUserCount, err := expandLookerInstanceUserMetadataAdditionalDeveloperUserCount(original["additional_developer_user_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdditionalDeveloperUserCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["additionalDeveloperUserCount"] = transformedAdditionalDeveloperUserCount
	}

	return transformed, nil
}

func expandLookerInstanceUserMetadataAdditionalViewerUserCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceUserMetadataAdditionalStandardUserCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLookerInstanceUserMetadataAdditionalDeveloperUserCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
