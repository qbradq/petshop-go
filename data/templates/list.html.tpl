{{ template "header.html.tpl" }}

<h1 class="title">List a Pet</h1>

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

<form action="/api/list" enctype="multipart/form-data" method="post">
    <img id="picture-preview" />
    <label for="picture">Picture</label><br />
    <input type="file" id="picture" name="picture" accept="image/*" onchange="preview(event);" />
    <br />
    <label for="name">Name</label><br />
    <input type="text" id="name" name="name" />
    <br />
    <label for="description">Description</label><br />
    <textarea id="description" name="description" rows="10" cols="80"></textarea>
    <br />
    <input type="submit" value="Submit" />
</form>

{{ template "footer.html.tpl" }}
