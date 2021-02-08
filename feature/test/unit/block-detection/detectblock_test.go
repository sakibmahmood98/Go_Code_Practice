package detectblock

import "testing"

type IsLineTest struct {
	arg1     string
	expected bool
}

var IsLineTests = []IsLineTest{
	IsLineTest{"adsf", true},
	IsLineTest{"fdsfaaa", true},
	IsLineTest{"fdsf", true},
	IsLineTest{"fdsfsda", false},
}

var NGBlock NginxBlock

func TestIsLine(t *testing.T) {

	for _, test := range IsLineTests {
		if output := NGBlock.IsLine(test.arg1); output != test.expected {
			t.Errorf("Output for Line %v  %t", output, test.expected)
		}
	}
}

type HasCommentTest struct {
	arg1     string
	expected bool
}

var HasCommentTests = []HasCommentTest{
	HasCommentTest{"#adsf", true},
	HasCommentTest{"#fdsfaaa", true},
	HasCommentTest{"#dsf", true},
	HasCommentTest{"fdsfsda", false},
}

func TestHasComment(t *testing.T) {

	for _, test := range HasCommentTests {
		if output := NGBlock.HasComment(test.arg1); output != test.expected {
			t.Errorf("Output for Comment %v  %t", output, test.expected)
		}
	}
}

type IsBlockTest struct {
	arg1     string
	expected bool
}

var IsBlockTests = []IsBlockTest{
	IsBlockTest{"#adsf{fd}", true},
}

func TestIsBlock(t *testing.T) {

	for _, test := range IsBlockTests {
		if output := NGBlock.IsBlock(test.arg1); output != test.expected {
			t.Errorf("Output for Block %v  %t", output, test.expected)
		}
	}
}

type NginxBlockTest struct {
	//	arg1     string
	lines        []*string
	startIndex   int
	endIndex     int
	recursionMax int
	expected     int
}

var NginxBlockTests = []NginxBlockTest{
	//	NginxBlockTest{ "dasf" ,1, 5, 0, 3},
}

func TestGetNginxBlock(t *testing.T) {
	var output *NginxBlock
	for _, test := range NginxBlockTests {

		output = GetNginxBlock(test.lines, test.startIndex, test.endIndex, test.recursionMax)
		if output.TotalBlocksInside != test.expected {
			t.Errorf("Output for GetNginxBlock %v  %v", output.TotalBlocksInside, test.expected)
		}
	}
}

type NginxBlocksTest struct {
	arg1     string
	expected bool
}

var NginxBlocksTests = []NginxBlocksTest{
	NginxBlocksTest{"#adsf{fd}ds", true},
}

func TestGetNginxBlocks(t *testing.T) {
	var output *NginxBlocks
	for _, test := range NginxBlocksTests {

		output = GetNginxBlocks(test.arg1)
		if output.flag != test.expected {
			t.Errorf("Output for GetNginxBlocks %v  %t", output.flag, test.expected)
		}
	}
}
