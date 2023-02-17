{{define "views/home/index"}}
{{template "html_begin" .}}
{{template "header" .}}
{{template "navbar" .}}
<body>
<script src="https://cdn.rawgit.com/showdownjs/showdown/2.0.3/dist/showdown.min.js"></script>
<script>
        let fetchRes = fetch("/static/markdown/home.md");
        var converter = new showdown.Converter();
        fetchRes.then(res =>
            res.text()).then(md => {
                document.getElementById('content').innerHTML = converter.makeHtml(md);
                console.log(md)
            })

/*     
  
        // fetchRes is the promise to resolve
        // it by using.then() method
        fetchRes.then(res =>
            res.text()).then(d => {
                document.getElementById('content').innerHTML = marked.parse(d);
                console.log(d)
            })

 */
</script>


<!-- Page content-->
<div class="container">
    <div id="content"></div>
</div>
</body>
 
{{template "footer" .}}
{{template "html_end" .}}
{{end}}

