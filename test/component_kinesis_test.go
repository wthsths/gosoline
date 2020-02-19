//+build integration

package test_test

import (
	pkgTest "github.com/applike/gosoline/pkg/test"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_kinesis(t *testing.T) {
	setup(t)

	pkgTest.Boot("test_configs/config.kinesis.test.yml")
	defer pkgTest.Shutdown()

	kinClient := pkgTest.ProvideKinesisClient("kinesis")
	o, err := kinClient.ListStreams(&kinesis.ListStreamsInput{})

	assert.NoError(t, err)
	assert.Len(t, o.StreamNames, 0)
}
