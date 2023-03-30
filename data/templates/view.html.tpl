{{ template "header.html.tpl" }}

<h1 class="title">{{ .Name }}</h1>
<p>{{ .Description }}</p>
<!-- <img id="picture-preview" src="/api/image?p={{ .ID }}" /> -->
<img src="https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png" />

{{ template "footer.html.tpl" }}
