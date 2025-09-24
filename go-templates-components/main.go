package main

import (
	"bytes"
	"html/template"
	"io"
	"os"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

const tmpl = `
{{define "button"}}
{{ $id := uniqueid }}
<o-button>
	<button id="{{$id}}" class="{{.Class}}">{{.Text}}</button>

	<script type="module">
		document.getElementById("{{$id}}").addEventListener('click', function() {
			alert('Button {{.Text}} clicked!');
		});
	</script>

	<style>
		@scope {
			button {
				color: red;
			}
		}
	</style>
 </o-button>
{{end}}

{{define "icon"}}
<i class="icon-{{.Name}}"></i>
{{end}}

{{define "card"}}
<div class="card">
  <h3>{{.Title}}</h3>
  <div class="content">{{.Content}}</div>
	<button>Plain button</button>
</div>
{{end}}

{{define "lists"}}
<ul>
{{- range .items}}
  <li class="item">{{.}}</li>
{{- end}}
</ul>
{{end}}
`
const tmpl2 = `
<dialog>
	<form method="dialog">
		<p><label>Username: <input></label></p>
		<p><label>Password: <input type="password"></label></p>
		{{ template "button" (dict "Class" "btn-secondary" "Text" "Cancel") }}
		{{ template "button" (dict "Class" "btn-primary outlined" "Text" "Login") }}
	</form>
</dialog>
{{/* Usage */}}
{{template "card" (dict "Title" "My Card" "Content" (render "button" (dict "Class" "btn-secondary" "Text" "Cancel"))) }}

{{render "button" (dict "Text" "Confirm")}}

{{template "lists" (dict "items" (slice "Item 1" 2 "Item 2" "Item 3"))}}
`

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"

func nanoid() string {
	id, _ := gonanoid.Generate(alphabet, 6)
	return id
}

func main() {
	t := template.New("main")
	t = t.Funcs(template.FuncMap{
		"dict": func(values ...any) map[string]any {
			dict := make(map[string]any)
			for i := 0; i < len(values); i += 2 {
				dict[values[i].(string)] = values[i+1]
			}
			return dict
		},
		"uniqueid": nanoid,
		"slice": func(items ...any) any {
			return items
		},
		"render": func(templateName string, data any) template.HTML {
			var buf bytes.Buffer
			err := t.ExecuteTemplate(&buf, templateName, data)
			if err != nil {
				return template.HTML("")
			}
			return template.HTML(buf.String())
		},
	})
	file, _ := os.Open("components.html")
	contents, _ := io.ReadAll(file)
	t = template.Must(t.Parse(string(contents)))
	t = template.Must(t.Parse(tmpl2))
	t.Execute(os.Stdout, nil)

}
