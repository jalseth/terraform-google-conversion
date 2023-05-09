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

package google

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ComputeRegionHealthCheckAssetType string = "compute.googleapis.com/RegionHealthCheck"

func resourceConverterComputeRegionHealthCheck() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeRegionHealthCheckAssetType,
		Convert:   GetComputeRegionHealthCheckCaiObject,
	}
}

func GetComputeRegionHealthCheckCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/healthChecks/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeRegionHealthCheckApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeRegionHealthCheckAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "RegionHealthCheck",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeRegionHealthCheckApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	checkIntervalSecProp, err := expandComputeRegionHealthCheckCheckIntervalSec(d.Get("check_interval_sec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("check_interval_sec"); !tpgresource.IsEmptyValue(reflect.ValueOf(checkIntervalSecProp)) && (ok || !reflect.DeepEqual(v, checkIntervalSecProp)) {
		obj["checkIntervalSec"] = checkIntervalSecProp
	}
	descriptionProp, err := expandComputeRegionHealthCheckDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); ok || !reflect.DeepEqual(v, descriptionProp) {
		obj["description"] = descriptionProp
	}
	healthyThresholdProp, err := expandComputeRegionHealthCheckHealthyThreshold(d.Get("healthy_threshold"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("healthy_threshold"); !tpgresource.IsEmptyValue(reflect.ValueOf(healthyThresholdProp)) && (ok || !reflect.DeepEqual(v, healthyThresholdProp)) {
		obj["healthyThreshold"] = healthyThresholdProp
	}
	nameProp, err := expandComputeRegionHealthCheckName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	unhealthyThresholdProp, err := expandComputeRegionHealthCheckUnhealthyThreshold(d.Get("unhealthy_threshold"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("unhealthy_threshold"); !tpgresource.IsEmptyValue(reflect.ValueOf(unhealthyThresholdProp)) && (ok || !reflect.DeepEqual(v, unhealthyThresholdProp)) {
		obj["unhealthyThreshold"] = unhealthyThresholdProp
	}
	timeoutSecProp, err := expandComputeRegionHealthCheckTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("timeout_sec"); !tpgresource.IsEmptyValue(reflect.ValueOf(timeoutSecProp)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	httpHealthCheckProp, err := expandComputeRegionHealthCheckHttpHealthCheck(d.Get("http_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("http_health_check"); !tpgresource.IsEmptyValue(reflect.ValueOf(httpHealthCheckProp)) && (ok || !reflect.DeepEqual(v, httpHealthCheckProp)) {
		obj["httpHealthCheck"] = httpHealthCheckProp
	}
	httpsHealthCheckProp, err := expandComputeRegionHealthCheckHttpsHealthCheck(d.Get("https_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("https_health_check"); !tpgresource.IsEmptyValue(reflect.ValueOf(httpsHealthCheckProp)) && (ok || !reflect.DeepEqual(v, httpsHealthCheckProp)) {
		obj["httpsHealthCheck"] = httpsHealthCheckProp
	}
	tcpHealthCheckProp, err := expandComputeRegionHealthCheckTcpHealthCheck(d.Get("tcp_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tcp_health_check"); !tpgresource.IsEmptyValue(reflect.ValueOf(tcpHealthCheckProp)) && (ok || !reflect.DeepEqual(v, tcpHealthCheckProp)) {
		obj["tcpHealthCheck"] = tcpHealthCheckProp
	}
	sslHealthCheckProp, err := expandComputeRegionHealthCheckSslHealthCheck(d.Get("ssl_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ssl_health_check"); !tpgresource.IsEmptyValue(reflect.ValueOf(sslHealthCheckProp)) && (ok || !reflect.DeepEqual(v, sslHealthCheckProp)) {
		obj["sslHealthCheck"] = sslHealthCheckProp
	}
	http2HealthCheckProp, err := expandComputeRegionHealthCheckHttp2HealthCheck(d.Get("http2_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("http2_health_check"); !tpgresource.IsEmptyValue(reflect.ValueOf(http2HealthCheckProp)) && (ok || !reflect.DeepEqual(v, http2HealthCheckProp)) {
		obj["http2HealthCheck"] = http2HealthCheckProp
	}
	grpcHealthCheckProp, err := expandComputeRegionHealthCheckGrpcHealthCheck(d.Get("grpc_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("grpc_health_check"); !tpgresource.IsEmptyValue(reflect.ValueOf(grpcHealthCheckProp)) && (ok || !reflect.DeepEqual(v, grpcHealthCheckProp)) {
		obj["grpcHealthCheck"] = grpcHealthCheckProp
	}
	logConfigProp, err := expandComputeRegionHealthCheckLogConfig(d.Get("log_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("log_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(logConfigProp)) && (ok || !reflect.DeepEqual(v, logConfigProp)) {
		obj["logConfig"] = logConfigProp
	}
	regionProp, err := expandComputeRegionHealthCheckRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return resourceComputeRegionHealthCheckEncoder(d, config, obj)
}

func resourceComputeRegionHealthCheckEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {

	if _, ok := d.GetOk("http_health_check"); ok {
		hc := d.Get("http_health_check").([]interface{})[0]
		ps := hc.(map[string]interface{})["port_specification"]
		pn := hc.(map[string]interface{})["port_name"]

		if ps == "USE_FIXED_PORT" || (ps == "" && pn == "") {
			m := obj["httpHealthCheck"].(map[string]interface{})
			if m["port"] == nil {
				m["port"] = 80
			}
		}
		obj["type"] = "HTTP"
		return obj, nil
	}
	if _, ok := d.GetOk("https_health_check"); ok {
		hc := d.Get("https_health_check").([]interface{})[0]
		ps := hc.(map[string]interface{})["port_specification"]
		pn := hc.(map[string]interface{})["port_name"]

		if ps == "USE_FIXED_PORT" || (ps == "" && pn == "") {
			m := obj["httpsHealthCheck"].(map[string]interface{})
			if m["port"] == nil {
				m["port"] = 443
			}
		}
		obj["type"] = "HTTPS"
		return obj, nil
	}
	if _, ok := d.GetOk("http2_health_check"); ok {
		hc := d.Get("http2_health_check").([]interface{})[0]
		ps := hc.(map[string]interface{})["port_specification"]
		pn := hc.(map[string]interface{})["port_name"]

		if ps == "USE_FIXED_PORT" || (ps == "" && pn == "") {
			m := obj["http2HealthCheck"].(map[string]interface{})
			if m["port"] == nil {
				m["port"] = 443
			}
		}
		obj["type"] = "HTTP2"
		return obj, nil
	}
	if _, ok := d.GetOk("tcp_health_check"); ok {
		hc := d.Get("tcp_health_check").([]interface{})[0]
		ps := hc.(map[string]interface{})["port_specification"]
		pn := hc.(map[string]interface{})["port_name"]

		if ps == "USE_FIXED_PORT" || (ps == "" && pn == "") {
			m := obj["tcpHealthCheck"].(map[string]interface{})
			if m["port"] == nil {
				m["port"] = 80
			}
		}
		obj["type"] = "TCP"
		return obj, nil
	}
	if _, ok := d.GetOk("ssl_health_check"); ok {
		hc := d.Get("ssl_health_check").([]interface{})[0]
		ps := hc.(map[string]interface{})["port_specification"]
		pn := hc.(map[string]interface{})["port_name"]

		if ps == "USE_FIXED_PORT" || (ps == "" && pn == "") {
			m := obj["sslHealthCheck"].(map[string]interface{})
			if m["port"] == nil {
				m["port"] = 443
			}
		}
		obj["type"] = "SSL"
		return obj, nil
	}

	if _, ok := d.GetOk("grpc_health_check"); ok {
		hc := d.Get("grpc_health_check").([]interface{})[0]
		ps := hc.(map[string]interface{})["port_specification"]
		pn := hc.(map[string]interface{})["port_name"]

		if ps == "USE_FIXED_PORT" || (ps == "" && pn == "") {
			m := obj["grpcHealthCheck"].(map[string]interface{})
			if m["port"] == nil {
				return nil, fmt.Errorf("error in HealthCheck %s: `port` must be set for GRPC health checks`.", d.Get("name").(string))
			}
		}
		obj["type"] = "GRPC"
		return obj, nil
	}

	return nil, fmt.Errorf("error in HealthCheck %s: No health check block specified.", d.Get("name").(string))
}

func expandComputeRegionHealthCheckCheckIntervalSec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHealthyThreshold(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckUnhealthyThreshold(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckTimeoutSec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpHealthCheck(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHost, err := expandComputeRegionHealthCheckHttpHealthCheckHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedRequestPath, err := expandComputeRegionHealthCheckHttpHealthCheckRequestPath(original["request_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["requestPath"] = transformedRequestPath
	}

	transformedResponse, err := expandComputeRegionHealthCheckHttpHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeRegionHealthCheckHttpHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPortName, err := expandComputeRegionHealthCheckHttpHealthCheckPortName(original["port_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["portName"] = transformedPortName
	}

	transformedProxyHeader, err := expandComputeRegionHealthCheckHttpHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	transformedPortSpecification, err := expandComputeRegionHealthCheckHttpHealthCheckPortSpecification(original["port_specification"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortSpecification); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["portSpecification"] = transformedPortSpecification
	}

	return transformed, nil
}

func expandComputeRegionHealthCheckHttpHealthCheckHost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpHealthCheckRequestPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpHealthCheckResponse(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpHealthCheckPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpHealthCheckPortName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpHealthCheckProxyHeader(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpHealthCheckPortSpecification(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpsHealthCheck(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHost, err := expandComputeRegionHealthCheckHttpsHealthCheckHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedRequestPath, err := expandComputeRegionHealthCheckHttpsHealthCheckRequestPath(original["request_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["requestPath"] = transformedRequestPath
	}

	transformedResponse, err := expandComputeRegionHealthCheckHttpsHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeRegionHealthCheckHttpsHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPortName, err := expandComputeRegionHealthCheckHttpsHealthCheckPortName(original["port_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["portName"] = transformedPortName
	}

	transformedProxyHeader, err := expandComputeRegionHealthCheckHttpsHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	transformedPortSpecification, err := expandComputeRegionHealthCheckHttpsHealthCheckPortSpecification(original["port_specification"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortSpecification); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["portSpecification"] = transformedPortSpecification
	}

	return transformed, nil
}

func expandComputeRegionHealthCheckHttpsHealthCheckHost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpsHealthCheckRequestPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpsHealthCheckResponse(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpsHealthCheckPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpsHealthCheckPortName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpsHealthCheckProxyHeader(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttpsHealthCheckPortSpecification(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckTcpHealthCheck(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequest, err := expandComputeRegionHealthCheckTcpHealthCheckRequest(original["request"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequest); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["request"] = transformedRequest
	}

	transformedResponse, err := expandComputeRegionHealthCheckTcpHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeRegionHealthCheckTcpHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPortName, err := expandComputeRegionHealthCheckTcpHealthCheckPortName(original["port_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["portName"] = transformedPortName
	}

	transformedProxyHeader, err := expandComputeRegionHealthCheckTcpHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	transformedPortSpecification, err := expandComputeRegionHealthCheckTcpHealthCheckPortSpecification(original["port_specification"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortSpecification); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["portSpecification"] = transformedPortSpecification
	}

	return transformed, nil
}

func expandComputeRegionHealthCheckTcpHealthCheckRequest(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckTcpHealthCheckResponse(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckTcpHealthCheckPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckTcpHealthCheckPortName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckTcpHealthCheckProxyHeader(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckTcpHealthCheckPortSpecification(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckSslHealthCheck(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequest, err := expandComputeRegionHealthCheckSslHealthCheckRequest(original["request"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequest); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["request"] = transformedRequest
	}

	transformedResponse, err := expandComputeRegionHealthCheckSslHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeRegionHealthCheckSslHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPortName, err := expandComputeRegionHealthCheckSslHealthCheckPortName(original["port_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["portName"] = transformedPortName
	}

	transformedProxyHeader, err := expandComputeRegionHealthCheckSslHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	transformedPortSpecification, err := expandComputeRegionHealthCheckSslHealthCheckPortSpecification(original["port_specification"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortSpecification); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["portSpecification"] = transformedPortSpecification
	}

	return transformed, nil
}

func expandComputeRegionHealthCheckSslHealthCheckRequest(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckSslHealthCheckResponse(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckSslHealthCheckPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckSslHealthCheckPortName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckSslHealthCheckProxyHeader(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckSslHealthCheckPortSpecification(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttp2HealthCheck(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHost, err := expandComputeRegionHealthCheckHttp2HealthCheckHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedRequestPath, err := expandComputeRegionHealthCheckHttp2HealthCheckRequestPath(original["request_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["requestPath"] = transformedRequestPath
	}

	transformedResponse, err := expandComputeRegionHealthCheckHttp2HealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeRegionHealthCheckHttp2HealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPortName, err := expandComputeRegionHealthCheckHttp2HealthCheckPortName(original["port_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["portName"] = transformedPortName
	}

	transformedProxyHeader, err := expandComputeRegionHealthCheckHttp2HealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	transformedPortSpecification, err := expandComputeRegionHealthCheckHttp2HealthCheckPortSpecification(original["port_specification"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortSpecification); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["portSpecification"] = transformedPortSpecification
	}

	return transformed, nil
}

func expandComputeRegionHealthCheckHttp2HealthCheckHost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttp2HealthCheckRequestPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttp2HealthCheckResponse(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttp2HealthCheckPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttp2HealthCheckPortName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttp2HealthCheckProxyHeader(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckHttp2HealthCheckPortSpecification(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckGrpcHealthCheck(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPort, err := expandComputeRegionHealthCheckGrpcHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPortName, err := expandComputeRegionHealthCheckGrpcHealthCheckPortName(original["port_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["portName"] = transformedPortName
	}

	transformedPortSpecification, err := expandComputeRegionHealthCheckGrpcHealthCheckPortSpecification(original["port_specification"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortSpecification); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["portSpecification"] = transformedPortSpecification
	}

	transformedGrpcServiceName, err := expandComputeRegionHealthCheckGrpcHealthCheckGrpcServiceName(original["grpc_service_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGrpcServiceName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["grpcServiceName"] = transformedGrpcServiceName
	}

	return transformed, nil
}

func expandComputeRegionHealthCheckGrpcHealthCheckPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckGrpcHealthCheckPortName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckGrpcHealthCheckPortSpecification(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckGrpcHealthCheckGrpcServiceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckLogConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnable, err := expandComputeRegionHealthCheckLogConfigEnable(original["enable"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnable); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enable"] = transformedEnable
	}

	return transformed, nil
}

func expandComputeRegionHealthCheckLogConfigEnable(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionHealthCheckRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
