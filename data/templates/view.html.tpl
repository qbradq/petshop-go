{{ template "header.html.tpl" }}

<h1 class="title">{{ .Name }}</h1>
<p>{{ .Description }}</p>
<img id="petPic" src="/image/{{ .ID }}{{ .PictureExtension }}" />
<button onclick='window.location.href="/finalize.html?p={{ .ID }}";'>Adopt Me!</button>

{{ template "footer.html.tpl" }}
