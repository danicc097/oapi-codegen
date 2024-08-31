module github.com/danicc097/oapi-codegen/v2

go 1.21.13

replace (
	// since it depends on itself we can just point to the local directory and go.mod luckily doesn't complain.
	// it if ends up not working at some point, a dummy tag push for each actual change is required
	// so that replace references the lib with changes
	github.com/oapi-codegen/oapi-codegen/v2 => ./
	// github.com/getkin/kin-openapi => github.com/danicc097/kin-openapi v0.123.1-0.20240320222651-5a7e849603fb
	github.com/oapi-codegen/runtime => github.com/danicc097/runtime v1.0.10004
)

require (
	github.com/getkin/kin-openapi v0.127.0
	github.com/oapi-codegen/oapi-codegen/v2 v2.3.0
	github.com/speakeasy-api/openapi-overlay v0.9.0
	github.com/stretchr/testify v1.9.0
	golang.org/x/text v0.17.0
	golang.org/x/tools v0.24.0
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dprotaso/go-yit v0.0.0-20240618133044-5a0af90af097 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/invopop/yaml v0.3.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/vmware-labs/yaml-jsonpath v0.3.2 // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
)
