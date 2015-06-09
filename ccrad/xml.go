package ccrad

import "encoding/xml"

type Project struct {
	Name            string `xml:"name,attr" json:"name"`
	Activity        string `xml:"activity,attr" json:"activity"`
	LastBuildLabel  string `xml:"lastBuildLabel,attr"  json:"lastBuildLabel"`
	LastBuildStatus string `xml:"lastBuildStatus,attr"  json:"lastBuildStatus"`
	LastBuildTime   string `xml:"lastBuildTime,attr"  json:"lastBuildTime"`
	WebUrl          string `xml:"webUrl,attr" json:"webUrl"`
}

type Projects struct {
	List []Project `xml:"Project" json:"items"`
}

func ParseXML(contents []byte) Projects {
	var q Projects
	xml.Unmarshal(contents, &q)
	return q
}
