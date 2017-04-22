package main

import (
	"net/http"
	"strings"

	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/romainmenke/pusher/linkheader"
)

func main() {
	testing.Run(
		testing.Middleware{
			Name: "Handler",
			Func: func(h http.Handler) http.Handler {

				reader := strings.NewReader(`/
	</css/stylesheet.css>; rel=preload; as=style;
	</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;

	/alpha.html
	</css/stylesheet.css>; rel=preload; as=style;
	</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;
	</js/text_change.js>; rel=preload; as=script;

	/beta.html
	</css/stylesheet.css>; rel=preload; as=style;
	</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;
	</img/gopher.png>; rel=preload; as=image;
	</img/gopher1.png>; rel=preload; as=image;
	</img/gopher2.png>; rel=preload; as=image;
	</img/gopher3.png>; rel=preload; as=image;
	</img/gopher4.png>; rel=preload; as=image;
	</img/gopher5.png>; rel=preload; as=image;

	/gamma.html
	-

	/gamma-b.html
	</css/stylesheet.css>; rel=preload; as=style;
	</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;
	</call.json>; rel=preload;`)

				return linkheader.Handler(h, linkheader.RulesReaderOption(reader))
			},
		},
	)
}
