{{ template "header.html.tpl" }}

<h1 class="title">{{ .Name }}</h1>
<p>{{ .Description }}</p>
<img id="picture-preview" src="/api/petpic?p={{ .ID }}" />

{{ template "footer.html.tpl" }}
