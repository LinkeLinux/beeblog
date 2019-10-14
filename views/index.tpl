{{template "header"}}
<title>首页 my-beego 博客</title>
</head>
<body>
<div class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        {{template "navdiv" .}}
    </div>
    <div class="container">
        <div class="page-header">
            <h1>我的第一篇博客</h1>
            <h6 class="text-muted">文章发表于{{ .Date}}</h6>
            <p>hello ,这是我的第一篇博客</p>
        </div>
    </div>
    <script type="text/javascript" src="https://cdn.staticfile.org/jquery/3.4.1/jquery.js"></script>
    <script type="text/javascript" src="/static/js/bootstap.min.js"></script>
    </div>
</body>
</html>

