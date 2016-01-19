package main

import (
	log "github.com/sirupsen/logrus"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

type renderCtx struct {
	StateNum int
	States   []stateDesc
}
type stateDesc struct {
	Current bool
	Id      int
}

var tmpl = template.Must(template.New("main").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>USB Switch</title>
</head>
<body>
	<h2>Current State: {{.StateNum}}</h2>
	<hr>
	<ol>
	{{range .States}}
		<li>
			<form action="set-state" method="POST">
				<input type="hidden" name="id" value="{{.Id}}" />
				<button {{if .Current}}disabled{{end}} type="submit">Set</button>
			</form>
		</li>
	{{end}}
	</ol>
</body>
</html>
`))

func newRenderCtx() *renderCtx {
	states := make([]stateDesc, stateCount)
	for i := range states {
		states[i].Current = currentState == i
		states[i].Id = i
	}
	return &renderCtx{
		StateNum: currentState + 1,
		States:   states,
	}
}

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		tmpl.Execute(w, newRenderCtx())
	case "/set-state":
		err := req.ParseForm()
		if err != nil {
			log.Warnln(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newState, err := strconv.Atoi(req.FormValue("id"))
		if err != nil {
			log.Warnln(err)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, err.Error())
			return
		}
		setState(newState)
		http.Redirect(w, req, "/", 302)
	default:
		http.NotFound(w, req)
	}
}
