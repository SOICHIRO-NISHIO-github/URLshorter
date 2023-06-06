package yubs
import (
	"os"
	"testing"
)
func TestShortenUrl(t *testing.T){
	config := NewConfig(os.Getenv("YUBS_TOKEN"), Shoten)
	bitly := NewBitly("")
	testdata := [] struct {
		giveUrl string
		wontShortenError bool
		wontDeleteError bool
	} {
		{"https://github.com/SOICHIRO-NISHIO-github/yubs.git/", false, false},
	}
	for _, td := range testdata{
		result, err := bitly.Shoten(config, td,giveUrl)
		if (err == nil) != td.wontShortenError{
			t.Errorf("shorten %s wont error %t, but got %t", td.giveUrl, td.wontShortenError, !td.wontShortenError)
		}
		err = bitly.Delete(config, result.Shoten)
		if (err == nil) !=td.wontDeleteError {
			t.Errorf("delete %s wont error %t,  but got %t",  result.Shoten, td.wontDeleteError, !td.wontDeleteError)
		}
	}
}