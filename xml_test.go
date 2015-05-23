package main

import (
	"testing"

	"github.com/jamesmura/go-cc-radiator/ccrad"
	"github.com/stretchr/testify/assert"
)

func TestParseXML(t *testing.T) {
	xml := `<Projects>
  <Project name="theSampleProject/angularProject (master) :: TESTS" activity="Sleeping" lastBuildLabel="1659" lastBuildStatus="Success" lastBuildTime="2015-05-22T14:30:34Z" webUrl="https://theci-server.com/theSampleProject/angularProject/branch/master/logs/defaultPipeline/1659/TESTS"/>
  <Project name="theSampleProject/angularProject (master) :: Deploy-QA" activity="Sleeping" lastBuildLabel="1659" lastBuildStatus="Success" lastBuildTime="2015-05-22T14:37:27Z" webUrl="https://theci-server.com/theSampleProject/angularProject/branch/master/logs/defaultPipeline/1659/Deploy-QA"/>
  <Project name="theSampleProject/angularProject (master) :: Deploy-PROD" activity="Sleeping" lastBuildLabel="1607" lastBuildStatus="Success" lastBuildTime="2015-05-14T09:07:08Z" webUrl="https://theci-server.com/theSampleProject/angularProject/branch/master/logs/defaultPipeline/1659/Deploy-PROD"/>
</Projects>
`
	query := ccrad.ParseXML([]byte(xml))
	assert.Equal(t, len(query.List), 3)
	assert.Equal(t, "theSampleProject/angularProject (master) :: TESTS", query.List[0].Name)
	assert.Equal(t, "theSampleProject/angularProject (master) :: Deploy-QA", query.List[1].Name)
	assert.Equal(t, "theSampleProject/angularProject (master) :: Deploy-PROD", query.List[2].Name)
	assert.Equal(t, "Sleeping", query.List[0].Activity)
	assert.Equal(t, "1659", query.List[0].LastBuildLabel)
	assert.Equal(t, "Success", query.List[0].LastBuildStatus)
	assert.Equal(t, "2015-05-22T14:30:34Z", query.List[0].LastBuildTime)
	assert.Equal(t, "https://theci-server.com/theSampleProject/angularProject/branch/master/logs/defaultPipeline/1659/TESTS", query.List[0].WebUrl)
}
