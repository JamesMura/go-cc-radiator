package ccrad

import "encoding/xml"

type Project struct {
	Name            string `xml:"name,attr"`
	Activity        string `xml:"activity,attr"`
	LastBuildLabel  string `xml:"lastBuildLabel,attr"`
	LastBuildStatus string `xml:"lastBuildStatus,attr"`
	LastBuildTime   string `xml:"lastBuildTime,attr"`
	WebUrl          string `xml:"webUrl,attr"`
}

type Projects struct {
	List []Project `xml:"Project"`
}

func ParseXML(contents []byte) Projects {
	var q Projects
	xml.Unmarshal(contents, &q)
	return q
}
