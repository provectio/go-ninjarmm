/*
Create 'env.json' file with the following content:
{
  "client_id": "x",
  "client_secret": "x"
}

Replace 'x' with your own values.

Else you can set the environment variables NINJARMM_CLIENT_ID and NINJARMM_CLIENT_SECRET.
*/

package ninjarmm

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"
)

// ⚠️ Use with Caution ⚠️
// Can't delete organizations with API
const testOrganizationCreation bool = false

func TestMain(t *testing.T) {
	// Getting environment variables
	clientID, clientSecret, err := envVars()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Login", func(t *testing.T) {
		err = Login(clientID, clientSecret, "monitoring management control")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("CreateOrganization", func(t *testing.T) {
		// ⚠️ Use with Caution ⚠️
		// Can't delete organizations with API
		if !testOrganizationCreation {
			t.Skip("Skipping organization creation test")
			return
		}

		copyFromOrganizationID := 6
		newOrganization := OrganizationDetailed{
			Name:             "Test Organization",
			NodeApprovalMode: ApprovalModeAutomatic,
		}

		createdOrganization, err := CreateOrganization(newOrganization, copyFromOrganizationID)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Created organization:\n%+v", createdOrganization)
		}
	})

	t.Run("ListOrganizations", func(t *testing.T) {
		organizations, err := ListOrganizations()
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Found %d organizations", len(organizations))
		}

		t.Run("UpdateOrganization", func(t *testing.T) {
			if len(organizations) < 1 {
				t.Skip("Skipping test (no organizations found)")
			}
			oldDescription := organizations[0].Description
			organizations[0].Description = "Testing organization update"
			err := UpdateOrganization(organizations[0])
			if err != nil {
				t.Error(err)
			} else {
				t.Logf("Updated organization:\n%+v", organizations[0])

				// Revert changes
				organizations[0].Description = oldDescription
				UpdateOrganization(organizations[0])
			}
		})

		t.Run("GetOrganizationLocations", func(t *testing.T) {
			if len(organizations) < 1 {
				t.Skip("Skipping test (no organizations found)")
			}
			locations, err := ListOrganizationLocations(organizations[0].ID)
			if err != nil {
				t.Error(err)
			} else {
				t.Logf("Found %d locations", len(locations))
				t.Logf("%+v", locations)
			}
		})

	})

	t.Run("ListDevices", func(t *testing.T) {
		devices, err := ListDevices("", false, 0, 0)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Found %d devices", len(devices))
		}
	})
}

func TestActivityLogQueryValues(t *testing.T) {
	options := ActivityLogOptions{
		AfterDate: "2021-01-01",
		PageSize:  300,
	}
	t.Log(options.queryString())
}

func TestTime(t *testing.T) {
	var testObject struct {
		Name string `json:"name"`
		T    Time   `json:"time"`
	}

	err := json.Unmarshal([]byte(`{"time": 1700666991.1700666, "name": "test object"}`), &testObject)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Time unmarshal: %s", testObject.T.String())

	data, err := json.Marshal(testObject)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Time marshal: %s", string(data))
}

func envVars() (clientID, clientSecret string, err error) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		if _, err := os.Stat("env.json"); os.IsNotExist(err) {
			clientID = os.Getenv("NINJARMM_CLIENT_ID")
			clientSecret = os.Getenv("NINJARMM_CLIENT_SECRET")
		} else {

			type envJSON struct {
				ClientID     string `json:"client_id"`
				ClientSecret string `json:"client_secret"`
			}

			var env envJSON
			data, err := os.ReadFile("env.json")
			if err != nil {
				return "", "", err
			}
			err = json.Unmarshal(data, &env)
			if err != nil {
				return "", "", err
			}
			clientID = env.ClientID
			clientSecret = env.ClientSecret
		}

	} else {
		data, err := os.ReadFile(".env")
		if err != nil {
			return "", "", err
		}
		for _, line := range strings.Split(string(data), "\n") {
			if strings.HasPrefix(line, "NINJARMM_CLIENT_ID=") {
				clientID = strings.TrimPrefix(line, "NINJARMM_CLIENT_ID=")
			}
			if strings.HasPrefix(line, "NINJARMM_CLIENT_SECRET=") {
				clientSecret = strings.TrimPrefix(line, "NINJARMM_CLIENT_SECRET=")
			}
		}
	}

	if clientID == "" || clientSecret == "" {
		err = errors.New("no test vars found")
	}

	return
}
