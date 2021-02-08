package detectblock

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type NginxBlock struct {
	StartLine   string
	EndLine     string
	AllContents string
	// split lines by \n on AllContents,
	// use make to create *[],
	// first create make([]*Type..)
	// then use &var to make it *
	AllLines          []*string
	NestedBlocks      []*NginxBlock
	TotalBlocksInside int
}

//var ngBlock *NginxBlock

func (ngBlock *NginxBlock) IsBlock(line string) bool {
	ngBlock.AllContents = line
	// TODO Solve it using regex
	re := regexp.MustCompile(`\{.*\}`)
	// Braces open
	// store startindex

	// Braces close
	//store endindex

	return re.MatchString(line)

}

func (ngBlock *NginxBlock) IsLine(line string) bool {
	// TODO Solve it using regex
	re := regexp.MustCompile(`.*`)
	return re.MatchString(line)
}

func (ngBlock *NginxBlock) HasComment(line string) bool {
	// TODO Solve it using regex
	re := regexp.MustCompile(`#.*`)
	return re.MatchString(line)
}

type NginxBlocks struct {
	blocks      []*NginxBlock
	AllContents string
	// split lines by \n on AllContents
	AllLines []*string
	flag     bool
}

var NgBlocks *NginxBlocks

func GetNginxBlock(
	lines []*string,
	startIndex,
	endIndex,
	recursionMax int,
) *NginxBlock { //updating block
	var NGBlock *NginxBlock
	NGBlock.AllLines = lines
	size := len(NGBlock.AllLines)
	NGBlock.EndLine = *lines[size-1]
	NGBlock.StartLine = *lines[0]

	return NGBlock
}

func GetNginxBlocks(configContent string) *NginxBlocks {
	var ngBlock *NginxBlock
	NgBlocks.AllContents = configContent
	NgBlocks.flag = true

	AllLines := strings.Split(NgBlocks.AllContents, "\n") //splitted strings by new line
	comment := 0

	size := len(AllLines) //size of string
	for i := 0; i < size; i++ {
		//	lastindex = len(AllLines[i]) - 1 //last index of a line
		if ngBlock.IsLine(AllLines[i]) == true {

			NgBlocks.AllLines[i] = &AllLines[i]
		}
		if ngBlock.HasComment(AllLines[i]) == true {
			comment++
		}
	}

	if ngBlock.IsBlock(NgBlocks.AllContents) == true { //for first block

		NgBlocks.blocks[0] = GetNginxBlock(NgBlocks.AllLines, 0, 0, 1) //start index is the next index of Braces as well as for last index

		NgBlocks.blocks[0].AllContents = ngBlock.AllContents

	}

	i := 1
	recursionMax := 6 //updated from IsBlock function
	for i < recursionMax {
		if ngBlock.IsBlock(ngBlock.AllContents) == true {

			NgBlocks.blocks[i] = GetNginxBlock(ngBlock.AllLines, 0, 0, recursionMax) //updating the block & collecting child-block

			NgBlocks.blocks[i].AllContents = ngBlock.AllContents
			i++
		} else {
			break
		}

	}

	return NgBlocks

}

func main() {

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile("nginx.conf")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	NgBlocks = GetNginxBlocks(string(content))

}
