package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func RebootInstances(t *testing.T) {
	resp, err := veen.NewInstance(ak, sk).RebootInstances(&veen.RebootInstancesReq{
		InstanceIdentities: []string{"veen20020201213020320353", "veen12922733323321300332"},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}
