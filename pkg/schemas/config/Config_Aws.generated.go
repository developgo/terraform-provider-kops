package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ConfigAws() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"profile":     OptionalString(),
			"region":      OptionalString(),
			"assume_role": OptionalStruct(ConfigAwsAssumeRole()),
		},
	}
}

func ExpandConfigAws(in map[string]interface{}) config.Aws {
	if in == nil {
		panic("expand Aws failure, in is nil")
	}
	return config.Aws{
		Profile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["profile"]),
		Region: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["region"]),
		AssumeRole: func(in interface{}) *config.AwsAssumeRole {
			return func(in interface{}) *config.AwsAssumeRole {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in config.AwsAssumeRole) *config.AwsAssumeRole {
					return &in
				}(func(in interface{}) config.AwsAssumeRole {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return config.AwsAssumeRole{}
					}
					return (ExpandConfigAwsAssumeRole(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["assume_role"]),
	}
}

func FlattenConfigAwsInto(in config.Aws, out map[string]interface{}) {
	out["profile"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Profile)
	out["region"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Region)
	out["assume_role"] = func(in *config.AwsAssumeRole) interface{} {
		return func(in *config.AwsAssumeRole) interface{} {
			if in == nil {
				return nil
			}
			return func(in config.AwsAssumeRole) interface{} {
				return func(in config.AwsAssumeRole) []map[string]interface{} {
					return []map[string]interface{}{FlattenConfigAwsAssumeRole(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AssumeRole)
}

func FlattenConfigAws(in config.Aws) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenConfigAwsInto(in, out)
	return out
}
