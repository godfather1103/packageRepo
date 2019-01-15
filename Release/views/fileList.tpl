<html class="hb-loaded"><head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta charset="utf-8">
    <link rel="icon" href="/favicon.ico" type="image/x-icon">
    <link rel="shortcut icon" href="/favicon.ico" type="image/x-icon">
    <link rel="stylesheet" type="text/css" href="static/assets/css/font-awesome.min.css">
    <link rel="stylesheet" type="text/css" href="static/assets/css/hljs-vs.css">
    <link rel="stylesheet" type="text/css" href="static/assets/css/main.css?v=24">
    <script src="static/assets/scripts/frameworks.js?v=4"></script>
    <title>{{ .TITLE}} - Package Repo Server</title>
</head>
<body huaban_collector_injected="true">
<header id="header">
    <a href="getFileList">
        <img id="logo" src="static/img/logo.png?v=2" style="width: 45px;" alt="PackageRepo Server">
    </a>
    <a href="upload" style="float: right;margin-right: 2%;margin-top: 0.2%;font-size: 24px;">
        upload
    </a>
</header>
<div id="content">
    <noscript>
        <div id="no-script-msg">
            <h2>浏览器禁用了JS</h2>
            <p>
                页面需要JS的支持，请启动JS
            </p>
        </div>
    </noscript>
    <div id="errors-banner"></div>
    <div>
        <div id="repo-content">
            <header>
                <a href="getFileList" class="home navbar-item">
                    <span class="fa fa-home"></span>
                </a>
                <h1 class="navbar-item">/
                    {{with .BREADCRUMB}}
                    {{range .}}
                    <a href="getFileList?PathFileId={{ .PATHFILEID}}&IdType={{ .IDTYPE}}">{{ .PATHNAME}}</a>/
                    {{end}}
                    {{end}}
                </h1>
            </header>
            <div id="directory-list" class="table">
                <div class="thead">
                    <span class="table-cell"></span>
                    <span class="table-cell">Name</span>
                    <span class="table-cell">Size</span>
                    <span class="table-cell">Rev</span>
                    <span class="table-cell">Author</span>
                    <span class="table-cell">Date</span>
                </div>
                <div class="tbody">
                    {{with .PARENTPATH}}
                    <a class="dir table-row" href="getFileList?PathFileId={{.PATHFILEID}}&IdType={{.IDTYPE}}">
                        <span class="table-cell icon">
                          <i class="fa fa-level-up"></i>
                        </span>
                        <span class="table-cell name">..</span>
                        <span class="table-cell size"></span>
                        <span class="table-cell rev"></span>
                        <span class="table-cell author"></span>
                        <span class="table-cell date"></span>
                    </a>
                    {{end}}

                    {{with .GROUPIDPATHS}}
                    {{range .}}
                        <a class="dir table-row" href="getFileList?PathFileId={{.Id}}&IdType={{if eq .PathType 1}}0{{else}}1{{end}}">
                            <span class="table-cell icon">
                                {{if eq .PathType 2}}
                                    <i class="fa fa-folder-open"></i>
                                {{else}}
                                    <i class="fa fa-folder"></i>
                                {{end}}
                            </span>
                            <span class="table-cell name">{{.PathName}}</span>
                            <span class="table-cell size"></span>
                            <span class="table-cell rev"></span>
                            <span class="table-cell author"></span>
                            <span class="table-cell date">
                               <span title=""></span>
                            </span>
                        </a>
                    {{end}}
                    {{end}}

                    {{with .VERSIONPATHS}}
                    {{range .}}
                    <a class="dir table-row" href="getFileList?PathFileId={{.Id}}&IdType=2">
                    <span class="table-cell icon">
                      <i class="fa fa-folder-o"></i>
                    </span>
                        <span class="table-cell name">{{.Version}}</span>
                        <span class="table-cell size"></span>
                        <span class="table-cell rev"></span>
                        <span class="table-cell author"></span>
                        <span class="table-cell date">
                       <span title=""></span>
                    </span>
                    </a>
                    {{end}}
                    {{end}}

                    {{with .FILEINFOS}}
                    {{range .}}
                    <a class="file table-row" href="{{.STREAMURL}}">
                        <span class="table-cell icon">
                          <i class="fa fa-file-text-o"></i>
                        </span>
                        <span class="table-cell name">{{.FileName}}</span>
                        <span class="table-cell size"></span>
                        <span class="table-cell rev">{{.FileMD5}}</span>
                        <span class="table-cell author"></span>
                        <span class="table-cell date">
                          <span title="{{.LastVersionTime}}">{{.LastVersionTime}}</span>
                        </span>
                    </a>
                    {{end}}
                    {{end}}

                </div>
            </div>
            <div id="directory-list-readme"></div>
        </div>
    </div>
</div>
<script src="/static/js/utils.js"></script>
<script>
    init({{str2html .MSG}})
</script>
<footer>
    Powered by <a href="https://github.com/godfather1103/packageRepo">PackageRepo Server</a>. © 2018 FocusOps Software Ltd.
</footer>
