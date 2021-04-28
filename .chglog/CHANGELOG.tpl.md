{{ range .Versions }}
<a name="{{ .Tag.Name }}"></a>
## {{ if .Tag.Previous }}[Release {{ .Tag.Name }}]({{ $.Info.RepositoryURL }}/compare/{{ .Tag.Previous.Name }}...{{ .Tag.Name }}){{ else }}{{ .Tag.Name }}{{ end }}

> Release Date: {{ datetime "2006-01-02" .Tag.Date }}

{{ range .CommitGroups -}}
### {{ .Title }}

{{ range .Commits -}}
- {{ .Subject }}
{{ end }}
{{ end -}}

{{- if .RevertCommits -}}
### ⏮ Reverts

{{ range .RevertCommits -}}
- {{ .Revert.Header }}
{{ end }}
{{ end -}}

{{- if .NoteGroups -}}
{{ range .NoteGroups -}}
### ⚠️ {{ .Title }}

{{ range .Notes }}
{{ .Body }}
{{ end }}
{{ end -}}
{{ end -}}

### 📖 Commits

{{ range .Commits -}}
{{ if not .Merge -}}
{{ if not (contains .Header "Update CHANGELOG for" ) -}}
- {{ .Header }} [{{ .Hash.Short }}]
{{ end -}}
{{ end -}}
{{ end -}}

{{ end -}}