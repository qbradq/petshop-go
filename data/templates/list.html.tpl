{{ template "header.html.tpl" }}

<h1 class="title">List a Pet</h1>

<style>
img#picture-preview {
    width: 100%;
    display: none;
}
</style>

<script>
    function preview(event) {
        if(event.target.files.length > 0){
            var s = URL.createObjectURL(event.target.files[0]);
            var p = document.getElementById("picture-preview");
            p.src = s;
            p.style.display = "block";
        }
    }
</script>

<form action="/api/list" method="post">
    <img id="picture-preview" />
    <label for="picture">Upload Picture</label>
    <input type="file" name="picture" accept="image/*" onchange="preview(event);" />
    <br />
    <label for="name">Name</label>
    <input type="text" name="name" />
    <br />
    <label for="description">Description</label>
    <textarea name="description" rows="10" cols="80"></textarea>
    <br />
    <input type="submit" value="Submit" />
</form>

{{ template "footer.html.tpl" }}
