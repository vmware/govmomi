{{ range .Versions }}
<a name="{{ .Tag.Name }}"></a>
## {{ if .Tag.Previous }}[Release {{ .Tag.Name }}]({{ $.Info.RepositoryURL }}/compare/{{ .Tag.Previous.Name }}...{{ .Tag.Name }}){{ else }}{{ .Tag.Name }}{{ end }}

> Release Date: {{ datetime "2006-01-02" .Tag.Date }}

{{ range .CommitGroups -}}
### {{ .Title }}

{{ range .Commits -}}
- [{{ .Hash.Short }}]{{"\t"}}{{ .Subject }}
{{ end }}
{{ end -}}

{{- if .RevertCommits -}}
### ⏮ Reverts

{{ range .RevertCommits -}}
- [{{ .Hash.Short }}]{{"\t"}}{{ .Revert.Header }}
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
- [{{ .Hash.Short }}]{{"\t"}}{{ .Header }}
{{ end -}}
{{ end -}}
{{ end -}}

{{ end -}}