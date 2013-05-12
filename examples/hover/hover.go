package hover

import (
	ŧ "fmt"
	. "github.com/metakeule/gsnack"
)

var js = B(`
	$(
		function(){
			$('@@cssSelector@@').hover(
				function() { $(this).css('background-color', '@@color-active@@' ); },
				function() { $(this).css('background-color', '@@color-passive@@'); }
			);
		}
	);
`)
var html = B(`<h1 @@bodyAttr@@>@@content@@</h1>`)
var css = B(`
	@@cssSelector@@ {
		width: @@width@@;
		height: @@height@@;
		background-color: @@color-passive@@;
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
		"width":         B(`200px`),
		"height":        B(`400px`),
		"color-passive": B(`gray`)}
	proto.Libs = []*JsLib{Jquery(1, 8, 3)}
	New = proto
}
