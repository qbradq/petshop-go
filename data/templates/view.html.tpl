{{ template "header.html.tpl" }}

<h1 class="title">{{ .Name }}</h1>
<p>{{ .Description }}</p>
<img id="pet-pic" src="/image/{{ .ID }}{{ .PictureExtension }}" />
<button onclick='window.location.href="/adopt.html?p={{ .ID }}";'></button>

{{ template "footer.html.tpl" }}
