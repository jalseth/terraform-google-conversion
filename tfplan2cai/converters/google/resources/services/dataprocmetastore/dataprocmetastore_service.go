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

package dataprocmetastore

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const DataprocMetastoreServiceAssetType string = "metastore.googleapis.com/Service"

func ResourceConverterDataprocMetastoreService() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: DataprocMetastoreServiceAssetType,
		Convert:   GetDataprocMetastoreServiceCaiObject,
	}
}

func GetDataprocMetastoreServiceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//metastore.googleapis.com/projects/{{project}}/locations/{{location}}/services/{{service_id}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetDataprocMetastoreServiceApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: DataprocMetastoreServiceAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/metastore/v1/rest",
				DiscoveryName:        "Service",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetDataprocMetastoreServiceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandDataprocMetastoreServiceLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	networkProp, err := expandDataprocMetastoreServiceNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	portProp, err := expandDataprocMetastoreServicePort(d.Get("port"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("port"); !tpgresource.IsEmptyValue(reflect.ValueOf(portProp)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	tierProp, err := expandDataprocMetastoreServiceTier(d.Get("tier"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tier"); !tpgresource.IsEmptyValue(reflect.ValueOf(tierProp)) && (ok || !reflect.DeepEqual(v, tierProp)) {
		obj["tier"] = tierProp
	}
	maintenanceWindowProp, err := expandDataprocMetastoreServiceMaintenanceWindow(d.Get("maintenance_window"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_window"); !tpgresource.IsEmptyValue(reflect.ValueOf(maintenanceWindowProp)) && (ok || !reflect.DeepEqual(v, maintenanceWindowProp)) {
		obj["maintenanceWindow"] = maintenanceWindowProp
	}
	encryptionConfigProp, err := expandDataprocMetastoreServiceEncryptionConfig(d.Get("encryption_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("encryption_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptionConfigProp)) && (ok || !reflect.DeepEqual(v, encryptionConfigProp)) {
		obj["encryptionConfig"] = encryptionConfigProp
	}
	hiveMetastoreConfigProp, err := expandDataprocMetastoreServiceHiveMetastoreConfig(d.Get("hive_metastore_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("hive_metastore_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(hiveMetastoreConfigProp)) && (ok || !reflect.DeepEqual(v, hiveMetastoreConfigProp)) {
		obj["hiveMetastoreConfig"] = hiveMetastoreConfigProp
	}
	networkConfigProp, err := expandDataprocMetastoreServiceNetworkConfig(d.Get("network_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkConfigProp)) && (ok || !reflect.DeepEqual(v, networkConfigProp)) {
		obj["networkConfig"] = networkConfigProp
	}
	databaseTypeProp, err := expandDataprocMetastoreServiceDatabaseType(d.Get("database_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("database_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(databaseTypeProp)) && (ok || !reflect.DeepEqual(v, databaseTypeProp)) {
		obj["databaseType"] = databaseTypeProp
	}
	releaseChannelProp, err := expandDataprocMetastoreServiceReleaseChannel(d.Get("release_channel"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("release_channel"); !tpgresource.IsEmptyValue(reflect.ValueOf(releaseChannelProp)) && (ok || !reflect.DeepEqual(v, releaseChannelProp)) {
		obj["releaseChannel"] = releaseChannelProp
	}
	telemetryConfigProp, err := expandDataprocMetastoreServiceTelemetryConfig(d.Get("telemetry_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("telemetry_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(telemetryConfigProp)) && (ok || !reflect.DeepEqual(v, telemetryConfigProp)) {
		obj["telemetryConfig"] = telemetryConfigProp
	}

	return obj, nil
}

func expandDataprocMetastoreServiceLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataprocMetastoreServiceNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServicePort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServiceTier(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServiceMaintenanceWindow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHourOfDay, err := expandDataprocMetastoreServiceMaintenanceWindowHourOfDay(original["hour_of_day"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHourOfDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hourOfDay"] = transformedHourOfDay
	}

	transformedDayOfWeek, err := expandDataprocMetastoreServiceMaintenanceWindowDayOfWeek(original["day_of_week"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDayOfWeek); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dayOfWeek"] = transformedDayOfWeek
	}

	return transformed, nil
}

func expandDataprocMetastoreServiceMaintenanceWindowHourOfDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServiceMaintenanceWindowDayOfWeek(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServiceEncryptionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKey, err := expandDataprocMetastoreServiceEncryptionConfigKmsKey(original["kms_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKey"] = transformedKmsKey
	}

	return transformed, nil
}

func expandDataprocMetastoreServiceEncryptionConfigKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServiceHiveMetastoreConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVersion, err := expandDataprocMetastoreServiceHiveMetastoreConfigVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	transformedConfigOverrides, err := expandDataprocMetastoreServiceHiveMetastoreConfigConfigOverrides(original["config_overrides"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfigOverrides); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["configOverrides"] = transformedConfigOverrides
	}

	transformedKerberosConfig, err := expandDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig(original["kerberos_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKerberosConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kerberosConfig"] = transformedKerberosConfig
	}

	return transformed, nil
}

func expandDataprocMetastoreServiceHiveMetastoreConfigVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServiceHiveMetastoreConfigConfigOverrides(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKeytab, err := expandDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab(original["keytab"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeytab); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["keytab"] = transformedKeytab
	}

	transformedPrincipal, err := expandDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigPrincipal(original["principal"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrincipal); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["principal"] = transformedPrincipal
	}

	transformedKrb5ConfigGcsUri, err := expandDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKrb5ConfigGcsUri(original["krb5_config_gcs_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKrb5ConfigGcsUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["krb5ConfigGcsUri"] = transformedKrb5ConfigGcsUri
	}

	return transformed, nil
}

func expandDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCloudSecret, err := expandDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytabCloudSecret(original["cloud_secret"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCloudSecret); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cloudSecret"] = transformedCloudSecret
	}

	return transformed, nil
}

func expandDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytabCloudSecret(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigPrincipal(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKrb5ConfigGcsUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServiceNetworkConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedConsumers, err := expandDataprocMetastoreServiceNetworkConfigConsumers(original["consumers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConsumers); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["consumers"] = transformedConsumers
	}

	return transformed, nil
}

func expandDataprocMetastoreServiceNetworkConfigConsumers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedEndpointUri, err := expandDataprocMetastoreServiceNetworkConfigConsumersEndpointUri(original["endpoint_uri"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEndpointUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["endpointUri"] = transformedEndpointUri
		}

		transformedSubnetwork, err := expandDataprocMetastoreServiceNetworkConfigConsumersSubnetwork(original["subnetwork"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSubnetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["subnetwork"] = transformedSubnetwork
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataprocMetastoreServiceNetworkConfigConsumersEndpointUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServiceNetworkConfigConsumersSubnetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServiceDatabaseType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServiceReleaseChannel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreServiceTelemetryConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLogFormat, err := expandDataprocMetastoreServiceTelemetryConfigLogFormat(original["log_format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLogFormat); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["logFormat"] = transformedLogFormat
	}

	return transformed, nil
}

func expandDataprocMetastoreServiceTelemetryConfigLogFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}