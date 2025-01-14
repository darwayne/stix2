// Copyright 2020 Joakim Kennedy. All rights reserved. Use of
// this source code is governed by the included BSD license.

package stix2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromJSON(t *testing.T) {
	assert := assert.New(t)
	f, err := getResource("apt1-report.json")
	require.NoError(t, err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	t.Run("parse", func(t *testing.T) {
		c, err := FromJSON(data)
		assert.NoError(err)
		assert.NotNil(c)

		assert.Len(c.AttackPatterns(), 7)
		assert.Len(c.Identities(), 5)
		assert.Len(c.Indicators(), 12)
		assert.Len(c.IntrusionSets(), 1)
		assert.Len(c.AllMalware(), 6)
		assert.Len(c.MarkingDefinitions(), 1)
		assert.Len(c.Relationships(), 30)
		assert.Len(c.Reports(), 1)
		assert.Len(c.ThreatActors(), 5)
		assert.Len(c.Tools(), 10)
	})

	t.Run("keep-order", func(t *testing.T) {
		c, err := FromJSON(data)
		assert.NoError(err)
		assert.NotNil(c)
		assert.Equal(c.AllObjects(), c.AllObjects())
	})

	t.Run("do-not-keep-order", func(t *testing.T) {
		c, err := FromJSON(data, NoSortOption())
		assert.NoError(err)
		assert.NotNil(c)

		objs := c.AllObjects()

		// Because it deterministic, we loop to ensure with don't have a random
		// match.
		pass := false
		for i := 0; i < 1000; i++ {
			if !reflect.DeepEqual(objs, c.AllObjects()) {
				// We have found a case that doesn't match so we can exit the loop.
				pass = true
				break
			}
		}
		assert.True(pass, "Order is not deterministic")
	})
}

func TestFromJSONAll(t *testing.T) {
	assert := assert.New(t)
	f, err := getResource("all.json")
	require.NoError(t, err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	c, err := FromJSON(data)
	assert.NoError(err)
	assert.NotNil(c)

	assert.NotNil(c.ASs())
	assert.NotNil(c.AS(Identifier("autonomous-system--f720c34b-98ae-597f-ade5-27dc241e8c74")))
	assert.Nil(c.AS(Identifier("")))
	assert.NotNil(c.Artifacts())
	assert.NotNil(c.Artifact(Identifier("artifact--4cce66f8-6eaa-53cb-85d5-3a85fca3a6c5")))
	assert.Nil(c.Artifact(Identifier("")))
	assert.NotNil(c.AttackPatterns())
	assert.NotNil(c.AttackPattern(Identifier("attack-pattern--7e33a43e-e34b-40ec-89da-36c9bb2cacd5")))
	assert.Nil(c.AttackPattern(Identifier("")))
	assert.NotNil(c.Campaigns())
	assert.NotNil(c.Campaign(Identifier("campaign--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f")))
	assert.Nil(c.Campaign(Identifier("")))
	assert.NotNil(c.CourseOfActions())
	assert.NotNil(c.CourseOfAction(Identifier("course-of-action--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f")))
	assert.Nil(c.CourseOfAction(Identifier("")))
	assert.NotNil(c.Directories())
	assert.NotNil(c.Directory(Identifier("directory--93c0a9b0-520d-545d-9094-1a08ddf46b05")))
	assert.Nil(c.Directory(Identifier("")))
	assert.NotNil(c.DomainNames())
	assert.NotNil(c.DomainName(Identifier("domain-name--3c10e93f-798e-5a26-a0c1-08156efab7f5")))
	assert.Nil(c.DomainName(Identifier("")))
	assert.NotNil(c.EmailAddresses())
	assert.NotNil(c.EmailAddress(Identifier("email-addr--2d77a846-6264-5d51-b586-e43822ea1ea3")))
	assert.Nil(c.EmailAddress(Identifier("")))
	assert.NotNil(c.EmailMessages())
	assert.NotNil(c.EmailMessage(Identifier("email-message--cf9b4b7f-14c8-5955-8065-020e0316b559")))
	assert.Nil(c.EmailMessage(Identifier("")))
	assert.NotNil(c.ExtensionDefinition(Identifier("extension-definition--9c59fd79-4215-4ba2-920d-3e4f320e1e62")))
	assert.Nil(c.ExtensionDefinition(Identifier("")))
	assert.NotNil(c.Files())
	assert.NotNil(c.File(Identifier("file--6ce09d9c-0ad3-5ebf-900c-e3cb288955b5")))
	assert.Nil(c.File(Identifier("")))
	assert.NotNil(c.Groups())
	assert.NotNil(c.Group(Identifier("grouping--84e4d88f-44ea-4bcd-bbf3-b2c1c320bcb3")))
	assert.Nil(c.Group(Identifier("")))
	assert.NotNil(c.IPv4Addresses())
	assert.NotNil(c.IPv4Address(Identifier("ipv4-addr--b4e29b62-2053-47c4-bab4-bbce39e5ed67")))
	assert.Nil(c.IPv4Address(Identifier("")))
	assert.NotNil(c.IPv6Addresses())
	assert.NotNil(c.IPv6Address(Identifier("ipv6-addr--1e61d36c-a16c-53b7-a80f-2a00161c96b1")))
	assert.Nil(c.IPv6Address(Identifier("")))
	assert.NotNil(c.Identities())
	assert.NotNil(c.Identity(Identifier("identity--e5f1b90a-d9b6-40ab-81a9-8a29df4b6b65")))
	assert.Nil(c.Identity(Identifier("")))
	assert.NotNil(c.Indicators())
	assert.NotNil(c.Indicator(Identifier("indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f")))
	assert.Nil(c.Indicator(Identifier("")))
	assert.NotNil(c.Infrastructures())
	assert.NotNil(c.Infrastructure(Identifier("infrastructure--38c47d93-d984-4fd9-b87b-d69d0841628d")))
	assert.Nil(c.Infrastructure(Identifier("")))
	assert.NotNil(c.IntrusionSets())
	assert.NotNil(c.IntrusionSet(Identifier("intrusion-set--4e78f46f-a023-4e5f-bc24-71b3ca22ec29")))
	assert.Nil(c.IntrusionSet(Identifier("")))
	assert.NotNil(c.LanguageContents())
	assert.NotNil(c.LanguageContent(Identifier("language-content--b86bd89f-98bb-4fa9-8cb2-9ad421da981d")))
	assert.Nil(c.LanguageContent(Identifier("")))
	assert.NotNil(c.Locations())
	assert.NotNil(c.Location(Identifier("location--a6e9345f-5a15-4c29-8bb3-7dcc5d168d64")))
	assert.Nil(c.Location(Identifier("")))
	assert.NotNil(c.MACs())
	assert.NotNil(c.MAC(Identifier("mac-addr--65cfcf98-8a6e-5a1b-8f61-379ac4f92d00")))
	assert.Nil(c.MAC(Identifier("")))
	assert.NotNil(c.AllMalware())
	assert.NotNil(c.Malware(Identifier("malware--31b940d4-6f7f-459a-80ea-9c1f17b5891b")))
	assert.Nil(c.Malware(Identifier("")))
	assert.NotNil(c.MalwareAnalyses())
	assert.NotNil(c.MalwareAnalysis(Identifier("malware-analysis--d0a5219b-4960-4b0c-a9ce-ed7b0552cc1b")))
	assert.Nil(c.MalwareAnalysis(Identifier("")))
	assert.NotNil(c.MarkingDefinitions())
	assert.NotNil(c.MarkingDefinition(Identifier("marking-definition--34098fce-860f-48ae-8e50-ebd3cc5e41da")))
	assert.Nil(c.MarkingDefinition(Identifier("")))
	assert.NotNil(c.Mutexes())
	assert.NotNil(c.Mutex(Identifier("mutex--eba44954-d4e4-5d3b-814c-2b17dd8de300")))
	assert.Nil(c.Mutex(Identifier("")))
	assert.NotNil(c.AllNetworkTraffic())
	assert.NotNil(c.NetworkTraffic(Identifier("network-traffic--2568d22a-8998-58eb-99ec-3c8ca74f527d")))
	assert.Nil(c.NetworkTraffic(Identifier("")))
	assert.NotNil(c.Notes())
	assert.NotNil(c.Note(Identifier("note--0c7b5b88-8ff7-4a4d-aa9d-feb398cd0061")))
	assert.Nil(c.Note(Identifier("")))
	assert.NotNil(c.AllObservedData())
	assert.NotNil(c.ObservedData(Identifier("observed-data--b67d30ff-02ac-498a-92f9-32f845f448cf")))
	assert.Nil(c.ObservedData(Identifier("")))
	assert.NotNil(c.Opinions())
	assert.NotNil(c.Opinion(Identifier("opinion--b01efc25-77b4-4003-b18b-f6e24b5cd9f7")))
	assert.Nil(c.Opinion(Identifier("")))
	assert.NotNil(c.Processes())
	assert.NotNil(c.Process(Identifier("process--f52a906a-0dfc-40bd-92f1-e7778ead38a9")))
	assert.Nil(c.Process(Identifier("")))
	assert.NotNil(c.RegistryKeys())
	assert.NotNil(c.RegistryKey(Identifier("windows-registry-key--2ba37ae7-2745-5082-9dfd-9486dad41016")))
	assert.Nil(c.RegistryKey(Identifier("")))
	assert.NotNil(c.Relationships())
	assert.NotNil(c.Relationship(Identifier("relationship--57b56a43-b8b0-4cba-9deb-34e3e1faed9e")))
	assert.Nil(c.Relationship(Identifier("")))
	assert.NotNil(c.Reports())
	assert.NotNil(c.Report(Identifier("report--84e4d88f-44ea-4bcd-bbf3-b2c1c320bcb3")))
	assert.Nil(c.Report(Identifier("")))
	assert.NotNil(c.Sightings())
	assert.NotNil(c.Sighting(Identifier("sighting--ee20065d-2555-424f-ad9e-0f8428623c75")))
	assert.Nil(c.Sighting(Identifier("")))
	assert.NotNil(c.AllSoftware())
	assert.NotNil(c.Software(Identifier("software--a1827f6d-ca53-5605-9e93-4316cd22a00a")))
	assert.Nil(c.Software(Identifier("")))
	assert.NotNil(c.ThreatActors())
	assert.NotNil(c.ThreatActor(Identifier("threat-actor--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f")))
	assert.Nil(c.ThreatActor(Identifier("")))
	assert.NotNil(c.Tools())
	assert.NotNil(c.Tool(Identifier("tool--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f")))
	assert.Nil(c.Tool(Identifier("")))
	assert.NotNil(c.URLs())
	assert.NotNil(c.URL(Identifier("url--c1477287-23ac-5971-a010-5c287877fa60")))
	assert.Nil(c.URL(Identifier("")))
	assert.NotNil(c.UserAccounts())
	assert.NotNil(c.UserAccount(Identifier("user-account--0d5b424b-93b8-5cd8-ac36-306e1789d63c")))
	assert.Nil(c.UserAccount(Identifier("")))
	assert.NotNil(c.Vulnerabilities())
	assert.NotNil(c.Vulnerability(Identifier("vulnerability--0c7b5b88-8ff7-4a4d-aa9d-feb398cd0061")))
	assert.Nil(c.Vulnerability(Identifier("")))
	assert.NotNil(c.X509Certificates())
	assert.NotNil(c.X509Certificate(Identifier("x509-certificate--b595eaf0-0b28-5dad-9e8e-0fab9c1facc9")))
	assert.Nil(c.X509Certificate(Identifier("")))
}

func TestAllObjectsCollection(t *testing.T) {
	assert := assert.New(t)
	c := &Collection{}
	ip, err := NewIPv4Address("10.0.0.1")
	assert.NoError(err)
	c.Add(ip)
	ip, err = NewIPv4Address("10.0.0.2")
	assert.NoError(err)
	c.Add(ip)

	objs := c.AllObjects()
	assert.Len(objs, 2)
}

func TestGetCreatedAndModified(t *testing.T) {
	assert := assert.New(t)
	ts := time.Now()

	tests := []struct {
		object     STIXObject
		tscreated  *time.Time
		tsmodified *time.Time
	}{
		{&DomainName{Value: "example.com"}, nil, nil},
		{&AttackPattern{}, nil, nil},
		{&AttackPattern{
			Name: "example",
			STIXDomainObject: STIXDomainObject{
				Created:  &Timestamp{ts},
				Modified: &Timestamp{ts},
			},
		}, &ts, &ts},
		{&Relationship{}, nil, nil},
		{&Relationship{
			STIXRelationshipObject: STIXRelationshipObject{
				Created:  &Timestamp{ts},
				Modified: &Timestamp{ts},
			},
		}, &ts, &ts},
		{&LanguageContent{}, nil, nil},
		{&LanguageContent{
			Created:  &Timestamp{ts},
			Modified: &Timestamp{ts},
		}, &ts, &ts},
		{&MarkingDefinition{}, nil, nil},
		{&MarkingDefinition{
			Created: &Timestamp{ts},
		}, &ts, nil},
	}

	for _, test := range tests {
		assert.Equal(test.tscreated, test.object.GetCreated())
		assert.Equal(test.tsmodified, test.object.GetModified())
	}
}

func TestGetFromCollection(t *testing.T) {
	c := &Collection{}
	m, _ := NewMalware(true, OptionName("Test object"))
	c.Add(m)

	tests := []struct {
		id      Identifier
		noMatch bool
	}{
		{m.ID, false},
		{Identifier("incorrect-id-format"), true},
		{Identifier(fmt.Sprintf("type-not-in-bucket--%s", uuid.New())), true},
		{Identifier(fmt.Sprintf("malware--%s", uuid.New())), true},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			obj := c.Get(test.id)
			if test.noMatch {
				assert.Nil(t, obj)
			} else {
				assert.NotNil(t, obj)
			}
		})
	}
}

func TestDuplicateInCollection(t *testing.T) {
	d, _ := NewDomainName("example.com")
	col := New()
	col.Add(d)
	col.Add(d)

	assert.Len(t, col.AllObjects(), 1)
}

func TestJSONMarshal(t *testing.T) {
	assert := assert.New(t)
	f, err := getResource("all.json")
	require.NoError(t, err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	c, err := FromJSON(data)
	assert.NoError(err)
	assert.NotNil(c)

	bundle, err := c.ToBundle()
	assert.NoError(err)

	buf, err := json.Marshal(bundle)
	assert.NoError(err)
	assert.NotNil(buf)
}

func TestMarshalHandleErrorConditions(t *testing.T) {
	assert := assert.New(t)

	t.Run("error-when-required-field-is-missing", func(t *testing.T) {
		d, err := NewDomainName("example.com")
		assert.NoError(err)
		d.Value = ""
		_, err = marshalToJSONHelper(d)
		assert.Error(err)
	})

	t.Run("error-when-required-field-is-missing-in-parent", func(t *testing.T) {
		d, err := NewDomainName("example.com")
		assert.NoError(err)
		d.ID = ""
		_, err = marshalToJSONHelper(d)
		assert.Error(err)
	})
}

func getResource(file string) (*os.File, error) {
	pth, err := filepath.Abs(filepath.Join("testresources", file))
	if err != nil {
		return nil, err
	}
	return os.OpenFile(pth, os.O_RDONLY, 0600)
}
