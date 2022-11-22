{{ range .Versions }}
<a name="{{ .Tag.Name }}"></a>
## {{ if .Tag.Previous }}[Release {{ .Tag.Name }}]({{ $.Info.RepositoryURL }}/compare/{{ .Tag.Previous.Name }}...{{ .Tag.Name }}){{ else }}{{ .Tag.Name }}{{ end }}

> Release Date: {{ datetime "2006-01-02" .Tag.Date }}

{{ range .CommitGroups -}}
### {{ .Title }}

{{ range .Commits -}}
- [{{ .Hash.Short }}]{{"\t"}}{{ .Subject }}{{ range .Refs }} (#{{ .Ref }}) {{ end }}
{{ end }}
{{ end -}}

{{- if .RevertCommits -}}
### ⏮ Reverts

{{ range .RevertCommits -}}
- [{{ .Hash.Short }}]{{"\t"}}{{ .Revert.Header }}{{ range .Refs }} (#{{ .Ref }}) {{ end }}
{{ end }}
{{ end -}}

### ⚠️ BREAKING

{{ range .Commits -}}
{{ if .Notes -}}
{{ if not .Merge -}}
{{ if not (contains .Header "Update CHANGELOG for" ) -}}
{{ .Subject }} [{{ .Hash.Short }}]:{{"\n"}}{{ range .Notes }}{{ .Body }}
{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}

### 📖 Commits

{{ range .Commits -}}
{{ if not .Merge -}}
{{ if not (contains .Header "Update CHANGELOG for" ) -}}
- [{{ .Hash.Short }}]{{"\t"}}{{ .Header }}{{ range .Refs }} (#{{ .Ref }}) {{ end }}
{{ end -}}
{{ end -}}
{{ end -}}

{{ end -}}