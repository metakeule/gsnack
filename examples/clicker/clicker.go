package clicker

import (
	ŧ "fmt"
	. "github.com/metakeule/gsnack"
)

var js = B(`
	$(
		function(){
			$('@@cssSelector@@').click(
				function() { $(this).css('background-color', '@@color@@'); }
			);
		}
	);
`)
var html = B(`<div @@bodyAttr@@>@@content@@</div>`)
var css = B(`
	@@cssSelector@@ {
		width: @@width@@;
		height: @@height@@;
		background-color: gray;
	}`)
var New Snacker

func init() {
	proto := NewSnack(map[string]Version{"jquery": Version{1, 8, 0}})
	proto.Html = html
	proto.Css = css
	proto.Js = js
	errs := proto.ParseErrors()
	if len(errs) > 0 {
		panic(ŧ.Sprintln(errs))
	}
	proto.Defaults = map[string][]byte{
		"width":  B(`200px`),
		"height": B(`400px`)}
	proto.Libs = []*JsLib{Jquery(1, 8, 3)}
	New = proto
}
