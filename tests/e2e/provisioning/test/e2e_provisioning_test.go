package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_E2E_Provisioning(t *testing.T) {
	ts := newTestSuite(t)
	if ts.IsDummyTest {
		return
	}
	operationID, err := ts.brokerClient.ProvisionRuntime()
	require.NoError(t, err)
	defer ts.TearDown()

	err = ts.brokerClient.AwaitOperationSucceeded(operationID, ts.ProvisionTimeout)
	require.NoError(t, err)

	dashboardURL, err := ts.brokerClient.FetchDashboardURL()
	require.NoError(t, err)

	err = ts.dashboardChecker.AssertRedirectedToUAA(dashboardURL)
	assert.NoError(t, err)
}
