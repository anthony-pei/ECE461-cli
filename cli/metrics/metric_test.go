package metrics

import (
	"math"
	"testing"
	"time"
)

type MockModule struct {
	URL             string
	License         string
	OpenIssues      int
	StargazersCount int
	Contributors    int
	Name            string
	FakeIssues      []IssueNode
	RampUp          float64
}

func (m MockModule) GetGitHubUrl() string {
	return m.URL
}

func (m MockModule) GetLicense() string {
	return m.License
}

func (m MockModule) GetOpenIssuesCount() int {
	return m.OpenIssues
}

func (m MockModule) GetStargazerCount() int {
	return m.StargazersCount
}

func (m MockModule) GetContributorCount() int {
	return m.Contributors
}

func (m MockModule) GetName() string {
	return m.Name
}

func (m MockModule) Clone(dir string) {
}

func (m MockModule) GetLast10ClosedIssues() []IssueNode {
	return m.FakeIssues
}

func TestResponsivness1Issue(t *testing.T) {
	fakeIssues := []IssueNode{{ClosedAt: time.Now(), CreatedAt: time.Now().AddDate(0, 0, -1)}}
	m := MockModule{FakeIssues: fakeIssues}
	responsivnessMetric := ResponsiveMaintainerMetric{}

	assertEquals(t, "", responsivnessMetric.CalculateScore(m), .99)
}

func TestResponsiveness0Issue(t *testing.T) {
	fakeIssues := []IssueNode{}
	m := MockModule{FakeIssues: fakeIssues}
	responsivnessMetric := ResponsiveMaintainerMetric{}

	assertEquals(t, "", responsivnessMetric.CalculateScore(m), 1.00)
}

func TestResponsives10Issue(t *testing.T) {
	fakeIssues := []IssueNode{}
	for i := 0; i < 10; i++ {
		fakeIssues = append(fakeIssues, IssueNode{ClosedAt: time.Now(), CreatedAt: time.Now().AddDate(0, 0, -1)})
	}
	m := MockModule{FakeIssues: fakeIssues}
	responsivnessMetric := ResponsiveMaintainerMetric{}

	assertEquals(t, "", math.Round(responsivnessMetric.CalculateScore(m)*100)/100, .90)

}
func TestRampUpNoComments(t *testing.T) {
	m := MockModule{}
	rampUpMetric := RampUpMetric{}

	analyzeDirFunction = func(dir string) (int64, int64, int64, int64) {
		return 0, 100, 0, 0
	}
	assertEquals(t, "", rampUpMetric.CalculateScore(m), 0.0)
}
func TestRampUpEqualCommentsCode(t *testing.T) {
	m := MockModule{}
	rampUpMetric := RampUpMetric{}

	analyzeDirFunction = func(dir string) (int64, int64, int64, int64) {
		return 0, 100, 100, 0
	}
	assertEquals(t, "", rampUpMetric.CalculateScore(m), 1.0)
}

func TestRampUpNoCode(t *testing.T) {
	m := MockModule{}
	rampUpMetric := RampUpMetric{}

	analyzeDirFunction = func(dir string) (int64, int64, int64, int64) {
		return 0, 0, 0, 0
	}
	assertEquals(t, "", rampUpMetric.CalculateScore(m), 0.0)
}
func TestCorrectnessZeroIssues(t *testing.T) {
	m := MockModule{OpenIssues: 0, StargazersCount: 10}
	correctnessMetric := CorrectnessMetric{}

	assertEquals(t, "Correctness (0, 10)", correctnessMetric.CalculateScore(m), 0.99999000001)
}

func TestCorrectnessZeroStargazers(t *testing.T) {
	m := MockModule{OpenIssues: 10, StargazersCount: 0}
	correctnessMetric := CorrectnessMetric{}

	assertEquals(t, "Correctness (0, 10)", correctnessMetric.CalculateScore(m), 0.0)
}
func TestCorrectnessEqualStargazersOpenIssues(t *testing.T) {
	m := MockModule{OpenIssues: 10, StargazersCount: 10}
	correctnessMetric := CorrectnessMetric{}

	assertEquals(t, "Correctness (0, 10)", correctnessMetric.CalculateScore(m), 0.0)
}

func TestBusFactorOneContributors(t *testing.T) {
	m := MockModule{Contributors: 1}
	busfactorMetric = BusFactorMetric{}
	assertEquals(t, "", busfactorMetric.CalculateScore(m), 0.0)
}
func TestBusFactorZeroContributors(t *testing.T) {
	m := MockModule{Contributors: 0}
	busfactorMetric = BusFactorMetric{}
	assertEquals(t, "", busfactorMetric.CalculateScore(m), 0.0)
}
func TestBusFactorTwoContributors(t *testing.T) {
	m := MockModule{Contributors: 2}
	busfactorMetric = BusFactorMetric{}
	assertEquals(t, "", busfactorMetric.CalculateScore(m), 0.5)

}
func TestBusFactorTenContributors(t *testing.T) {
	m := MockModule{Contributors: 10}
	busfactorMetric = BusFactorMetric{}
	assertEquals(t, "", busfactorMetric.CalculateScore(m), 0.9)
}
func TestBusFactor500Contributors(t *testing.T) {
	m := MockModule{Contributors: 500}
	busfactorMetric = BusFactorMetric{}
	assertEquals(t, "", busfactorMetric.CalculateScore(m), 1.0-1.0/500)
}
func TestBusFactor1750Contributors(t *testing.T) {
	m := MockModule{Contributors: 1750}
	busfactorMetric = BusFactorMetric{}
	assertEquals(t, "", busfactorMetric.CalculateScore(m), 1.0-1.0/1750)
}
func TestLicenseAccept(t *testing.T) {
	m1 := MockModule{License: "mit"}
	m2 := MockModule{License: "lgpl-2.1"}
	m3 := MockModule{License: "unlicense"}
	licenseMetric = LicenseMetric{}

	assertEquals(t, "", licenseMetric.CalculateScore((m1)), 1.0)
	assertEquals(t, "", licenseMetric.CalculateScore((m2)), 1.0)
	assertEquals(t, "", licenseMetric.CalculateScore((m3)), 1.0)
}
func TestLicenseDeny(t *testing.T) {
	m1 := MockModule{License: ""}
	m2 := MockModule{License: "gpl-3.0"}
	m3 := MockModule{License: "agpl-3.0"}
	licenseMetric = LicenseMetric{}

	assertEquals(t, "", licenseMetric.CalculateScore((m1)), 0.0)
	assertEquals(t, "", licenseMetric.CalculateScore((m2)), 0.0)
	assertEquals(t, "", licenseMetric.CalculateScore((m3)), 0.0)
}
func TestNetScoreMiddle(t *testing.T) {
	fakeIssues := []IssueNode{}
	m := MockModule{URL: "https://github.com/anthony-pei/ECE461", License: "mit", OpenIssues: 10, StargazersCount: 10, Contributors: 10, FakeIssues: fakeIssues}
	netScoreMetric := NetScoreMetric{}
	assertEquals(t, "", netScoreMetric.CalculateScore(m), 0.78)
}
func TestNetScoreLow(t *testing.T) {
	fakeIssues := []IssueNode{}
	for i := 0; i < 10; i++ {
		fakeIssues = append(fakeIssues, IssueNode{ClosedAt: time.Now(), CreatedAt: time.Now().AddDate(0, 0, -1)})
	}
	m := MockModule{URL: "https://github.com/cloudinary/cloudinary_npm", License: "agpl-3.0", OpenIssues: 10, StargazersCount: 0, Contributors: 0, FakeIssues: fakeIssues}
	netScoreMetric := NetScoreMetric{}
	assertEquals(t, "", netScoreMetric.CalculateScore(m), 0.36)
}
func TestNetScoreHigh(t *testing.T) {
	fakeIssues := []IssueNode{}
	// rampUpScore := 0.9
	m := MockModule{URL: "https://www.npmjs.com/package/express", License: "lgpl-2.1", OpenIssues: 0, StargazersCount: 1000, Contributors: 1000, FakeIssues: fakeIssues}
	netScoreMetric := NetScoreMetric{}
	assertEquals(t, "", math.Round(netScoreMetric.CalculateScore(m)*100)/100, 0.90)
}
func assertEquals(t *testing.T, desc string, got interface{}, want interface{}) {
	if got != want {
		t.Errorf("%v Got: %v (%T), Want:%v (%T)", desc, got, got, want, want)
	}
}
