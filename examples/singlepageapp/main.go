package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/skirmish/html.git"
)

func main() {
	ctx := context.Background()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/healthz"))
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", HelloServer)
	r.Get("/favicon.png", FavIconHandler)

	srv := http.Server{Addr: ":8080", Handler: chi.ServerBaseContext(ctx, r)}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
		select {
		case <-time.After(4 * time.Second):
			fmt.Println("not all connections done")
		case <-ctx.Done():
		}
	}()
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

var page = []html.Element{
	html.DocType(),
	html.Html(html.Lang("en")).AddElements(
		html.Head().AddElements(
			html.Meta(html.Charset("utf-8")),
			html.Title().AddElements(html.Content("Page Title")),
			html.Link(html.Rel("icon"), html.Type("image/png"), html.Href("/favicon.png")),
			html.Meta(html.Name("viewport"), html.ContentAttr("width=device-width, initial-scale=1.0")),
			html.Style().AddElements(html.Content(
				"body {margin: 10px 15px 20px 30px; "+
					"font-family: Arial, Helvetica, sans-serif; padding-top: 60px;}\n"+
					"header {background-color: #292961;padding: 0 16px; z-index: 1000; margin-bottom: 0; "+
					"min-height: 40px; border: 0; border-bottom: 1px solid #e5e5e5; "+
					"position: fixed; top: 0; left: 0; right: 0; border-radius: 0;}\n"+
					"header h1 {color:#ffeeff; margin-block-start: 0.5em;margin-block-end: 0.5em;}\n"+
					"ul.samples {list-style: none; margin: 0; padding: 0;"+
					" margin-block-start: 1px; margin-block-end: 1px;}\n"+
					"ul.samples li {display: inline-block; width: 300px; text-shadow: 1px 1px 3px #cccccc;}\n"+
					"ul.samples li img {display: inline-block; width: 32px; height: 32px;"+
					"box-shadow: -2px -2px 2px 2px #c1c1c1;}\n"+
					"ul.samples li.good {background-color: #c0e0c0;}\n")),
		),
		html.Body().AddElements(
			html.Header().AddElements(
				html.Heading(html.Level(1)).AddElements(html.Content("The Title of the Page")),
			),
			html.P().AddElements(html.Content("This is the first sentence.")),
			html.Ul(html.Class("samples")).AddElements(
				html.Li().AddElements(
					html.Img(html.Src("data:image/png;base64,"+favicon)),
					html.Content("some content"),
				),
				html.Li(html.Class("good")).AddElements(
					html.Img(html.Src("data:image/png;base64,"+favicon)),
					html.Content("some more content"),
				),
			),
		),
	),
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	for _, el := range page {
		el.Render(w)
	}
}

var favicon = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAD4klEQVRYR8WXW2xUVRSGv3WmtReo00lba6sWMMi91bZT8RZi8EWwQ6" +
	"JEoxgehIQHxajBRGPQECVEEyUaqw8E5IGoJERJoIgPhhi8YqfFhosBjbQCLZep01qZaW3nbLPPdKYXzjlzqtOwXtft32vt/e+1BI+iwo/6YeAhYC" +
	"moO0BmAsUj7r2gOkB+Bg5B/gEJ7unzEloyGamWFXMxzJdQ8jhQkMl+RB9H1G5M4y1p2HfKzccRgAqHClFsRlgP5HpMPNFsCEUTwkYJ7o/ZxbAFoI" +
	"423kZC9gIL/2PiiW4n8KmHpbb514mKqwCocGMdIl+iKMtS8mQY4TJKPSjB5raxcccBsE5uyndZT57KqEEY6t6xlUgDsHoOP2Uqe+TPGG3Hu+k4F6" +
	"Wvf9AK7S/KY+bNAeqqKygN6DAuojiBcGfqTowCaAltRXjBybUnGuOzgydpO34BpZStmYhQt6iClcsXUFLs9mDUVgk2b0h2BrCemqhjTre9/eQFdu" +
	"45SnxgOJ04x2dQWJgLCmLxIYYTZlpXkJ/LmsdqqZlf7nSeIRKJRbL4i9NJAK2hj1A8ZWfdeqyL7bvbME3FrVUBgjWVzJ9dRnnpNHw+w3JJJEwuRq" +
	"7wy2+XCbd38fvZKIYhrFtVT+3CCnsQSu2UhuY1MsJw3XYk032pny1N3zBvdimhB+ZSdZPf08PoPNfL/q9OcfpMD6+sX8KNZdPt/OKQXyEqHFoFfG" +
	"xn0Xm+17poNfMcS+kKSLeu2F/ADGfgT2oA24G1no6WfaMdosKNYZD67Mf2ElG16gpEgBIv5lNg06MB6Lflm4LgXkImMgKI9sX5cFcLlyJXWLlsAU" +
	"sWz3ANfPhIp0VYN5RO4+nVDQT8rj+4BcC1Bds+aUVzgRafIbz+4lJHuo1EY7z29iESZpIp66srLS5wEd0C90v43KaDDAyOMuDqR27nvoYq25jftv" +
	"zBrs/b07r8vBze27TMJX/yEro+w2c2HhhHs0+sqOb+u/U0drV8/UMHn+7TjJ4UTdcfbNZTnKPoZ+hMRNrtnW3fW4yWEs1sTsTSeb6PLU2H07ZzZp" +
	"WwYd09bgA0EVnDpi0Va88zZ6O8u+NHqw269LoFbqJboFuhy//82ruYdUvAyTxJxVrr9hlp/d+xf/irf5DK8iLX5Cll18V+ri/KY3rhdc72qc/IAp" +
	"DhO/aUdXJG47/jJAj3gWRy8TNZTxhILAAeR7JMoTPqnUYyC8S1HEpTyK/pWJ4GoSsxLHuRLC0muuw5HheT0UroMV29AfLs/1rNUO+DvDqp1WzsRV" +
	"JHls/BMF5GJrmcKr2cmm/qydftYmbcjkcrMjXr+b/U/pXQYWvvwgAAAABJRU5ErkJggg=="

func FavIconHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/png")
	decoded, err := base64.StdEncoding.DecodeString(favicon)
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	w.Write(decoded)
}
