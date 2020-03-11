package xmlreader

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// TestSuite is for test run
type TestSuite struct {
	XMLName   xml.Name `xml:"test-suite"`
	Text      string   `xml:",chardata"`
	Ns2       string   `xml:"ns2,attr"`
	Start     string   `xml:"start,attr"`
	Stop      string   `xml:"stop,attr"`
	Name      string   `xml:"name"`
	Title     string   `xml:"title"`
	TestCases struct {
		Text     string `xml:",chardata"`
		TestCase []struct {
			Text   string `xml:",chardata"`
			Start  string `xml:"start,attr"`
			Status string `xml:"status,attr"`
			Stop   string `xml:"stop,attr"`
			Name   string `xml:"name"`
			Title  string `xml:"title"`
			Labels struct {
				Text  string `xml:",chardata"`
				Label []struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name,attr"`
					Value string `xml:"value,attr"`
				} `xml:"label"`
			} `xml:"labels"`
			Parameters struct {
				Text      string `xml:",chardata"`
				Parameter struct {
					Text  string `xml:",chardata"`
					Kind  string `xml:"kind,attr"`
					Name  string `xml:"name,attr"`
					Value string `xml:"value,attr"`
				} `xml:"parameter"`
			} `xml:"parameters"`
			Steps       string `xml:"steps"`
			Attachments string `xml:"attachments"`
			Failure     struct {
				Text       string `xml:",chardata"`
				Message    string `xml:"message"`
				StackTrace string `xml:"stack-trace"`
			} `xml:"failure"`
		} `xml:"test-case"`
	} `xml:"test-cases"`
}

// TestCaseDetail is to store temp data
type TestCaseDetail struct {
	TestSuiteName string
	TestCaseName  string
	FeatureName   string
	Outcome       string
	StartDateTime string
	EndDateTime   string
	Duration      float64
	ErroMessage   string
}

// ConvertString the string of numbers recieved from the xml file
func ConvertString(x string) (string, int64) {

	ParsedInt, err := strconv.ParseInt(x, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ParsedStr := (time.Unix(0, ParsedInt*int64(time.Millisecond)))
	startDateTime := ParsedStr.Format(layoutUS)
	return startDateTime, ParsedInt

}

const (
	// layoutISO =  "2006-01-02T15:04:05.999999999Z07:00" //"2006-01-02"
	layoutUS = "02-Jan-2006 15:04:05"
)

//ReadDataFromFile is reading data from file.
func ReadDataFromFile() []TestCaseDetail {
	var testcases []TestCaseDetail
	// Open our xmlFile
	xmlFile, err := os.Open("1ec8da69-5cad-405a-bc39-5c4529d59f90-testsuite.xml")
	// xmlFile, err := os.Open("*.xml")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened the xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our array
	var testsuite TestSuite

	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'testsuite' which we defined above
	xml.Unmarshal(byteValue, &testsuite)

	// fmt.Println("Test Suite Name: **" + testsuite.Name + "**")

	// Loops through the xml file and pulls out the values to be read/manipulated
	for i := 0; i < len(testsuite.TestCases.TestCase); i++ {
		var testsuitename = testsuite.Name
		var testcasename = testsuite.TestCases.TestCase[i].Name
		var featurename = testsuite.TestCases.TestCase[i].Labels.Label[3].Value
		var outcome = testsuite.TestCases.TestCase[i].Status
		var errorMessage = testsuite.TestCases.TestCase[i].Failure.Message

		var StartTime = testsuite.TestCases.TestCase[i].Start

		var EndTime = (testsuite.TestCases.TestCase[i].Stop)

		startDateTime, StartParsedInt := ConvertString(StartTime)
		endDateTime, EndParsendInt := ConvertString(EndTime)

		// duration := (EndParsendInt.Sub(StartParsedInt))
		intDuration := (EndParsendInt - StartParsedInt)
		duration := (float64(intDuration)) / 1000

		testcases = append(testcases, TestCaseDetail{
			testsuitename,
			testcasename,
			featurename,
			outcome,
			startDateTime,
			endDateTime,
			duration,
			errorMessage,
		})

		print := fmt.Sprintf("TestCaseName: %s, ProductName: %s, Outcome: %s, StartTime: %s, EndTime: %s, Duration: %v \n", testcasename, featurename, outcome, startDateTime, endDateTime, duration)
		fmt.Println(print)

		if errorMessage != "" {
			fmt.Println(errorMessage)
		}
	}
	return testcases
}
