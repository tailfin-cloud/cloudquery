package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildSessions(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)

	var s types.Session
	if err := faker.FakeObject(&s); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeSessions(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&ssm.DescribeSessionsOutput{Sessions: []types.Session{s}},
		nil,
	)

	return client.Services{Ssm: mock}
}

func TestSessions(t *testing.T) {
	client.AwsMockTestHelper(t, Sessions(), buildSessions, client.TestOptions{})
}
