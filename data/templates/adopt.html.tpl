{{ template "header.html.tpl" }}

<h1 class="title">Adopt a Pet</h1>

{{ range . }}
<div class="wrapperPet">
    <a href="/view.html?p={{ .ID }}"><img class="petPreview" src="/image/{{ .ID }}{{ .PictureExtension }}" /></a>
    <span class="petName">{{ .Name }}</span><br />
    <span class="petDescription">{{ .Description }}</span>
</div>
{{ end }}

{{ template "footer.html.tpl" }}
