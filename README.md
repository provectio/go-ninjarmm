**Table of contents:**

- [Description](#description)
  - [Version](#version)
  - [Installation](#installation)
- [Examples](#examples)
  - [Login](#login)
  - [Find devices](#find-devices)
  - [Create organization](#create-organization)
- [Authors](#authors)

# Description

This is a Go client for the NinjaRMM API for EU region.

Login is required the first time you launch the client and information is stored internally in the package. The package automatically refresh the token when it's expired.

To get a valid `Client ID` and `Client Secret` you can read the [Ninja Documentation API](https://eu.ninjarmm.com/apidocs-beta/authorization/create-applications/machine-to-machine-apps).

## Version

NinjaAPI `v2.0.9-draft`

## Installation

```bash
go get github.com/provectio/go-ninjarmm@latest
```

Then you can import the package in your code:

```go
import "github.com/provectio/go-ninjarmm"
```

# Examples

## Login

You need call this function only one time. After, for each request, the package automatically refresh the token when it's expired whit same `client-id`, `client-secret` and `scope` from the first call.

```go
err := ninjarmm.Login("<client-id>", "<client-secret>", "monitoring management control")
if err != nil {
  panic(err)
}
```

## Find devices

```go
search := "my-hostname"
limit := 10

devices, err := ninjarmm.FindDevices(search, limit)
if err != nil {
  panic(err)
}

for _, device := range devices {
  fmt.Println(device)
}
```

## Create organization

```go
newOrg := ninjarmm.OrganizationDetailed{
  Name: "My new organization",
  Description: "My description",
  NodeApprovalMode: ninjarmm.ApprovalModeAutomatic,
  Tags: []string{"tag1", "tag2"},
  Fields: map[string]interface{}{
    "my_custom_field": "my custom value",
  },
  Locations: []ninjarmm.Location{
    {
      Name: "My location",
      Description: "My location description",
      Address: "My address",
      Tags: []string{"tag1", "tag2"},
      Fields: map[string]interface{}{
        "my_custom_field": "my custom value",
      },
    },
  },
  Policies: ninjarmm.OrganizationPolicyItem{
    NodeRoleID: 0,
    PolicyID: 0,
  }
}

org, err := ninjarmm.CreateOrganization(newOrg, 0)
if err != nil {
  panic(err)
}

fmt.Println(org)
```

# Authors

- [f41k4l](https://github.com/f41k4l)
- [Tryton](https://github.com/guillaumecollignon)

> Feel free to contribute to this project by creating a pull request or issue.
