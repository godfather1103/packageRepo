<html class="hb-loaded"><head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta charset="utf-8">
    <link rel="icon" href="/favicon.ico" type="image/x-icon">
    <link rel="shortcut icon" href="/favicon.ico" type="image/x-icon">
    <link rel="stylesheet" type="text/css" href="static/assets/css/font-awesome.min.css">
    <link rel="stylesheet" type="text/css" href="static/assets/css/hljs-vs.css">
    <link rel="stylesheet" type="text/css" href="static/assets/css/main.css?v=24">
    <script src="static/assets/scripts/frameworks.js?v=4"></script>
    <script type="text/javascript" src="static/assets/scripts/main.js?v=19"></script>
    <script src="static/assets/scripts/diff.js"></script>
    <script src="static/js/jquery.ajaxfileupload.js"></script>
    <title>Upload - Package Repo Server</title>
</head>
<body huaban_collector_injected="true">
<header id="header">
    <a href="getFileList">
        <img id="logo" src="static/assets/img/logo.svg?v=2" alt="VisualSVN Server">
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
                <h1 class="navbar-item">
                    <a href="getFileList?PathFileId=0&IdType=0"></a>/
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
                    <form action="uploadFile" method="post" enctype="multipart/form-data">
                        <a class="dir table-row" href="javascript:void(0);">
                            <span class="table-cell icon">
                            </span>
                            <span class="table-cell name">GroupId</span>
                            <span class="table-cell size"></span>
                            <span class="table-cell rev"><input type="text" id="groupId" name="groupId"></span>
                            <span class="table-cell author"></span>
                            <span class="table-cell date"></span>
                        </a>
                        <a class="dir table-row" href="javascript:void(0);">
                            <span class="table-cell icon">
                            </span>
                            <span class="table-cell name">ArtifactId</span>
                            <span class="table-cell size"></span>
                            <span class="table-cell rev"><input type="text" id="artifactId" name="artifactId"></span>
                            <span class="table-cell author"></span>
                            <span class="table-cell date"></span>
                        </a>
                        <a class="dir table-row" href="javascript:void(0);">
                            <span class="table-cell icon">
                            </span>
                            <span class="table-cell name">Version</span>
                            <span class="table-cell size"></span>
                            <span class="table-cell rev"><input type="text" id="version" name="version"></span>
                            <span class="table-cell author"></span>
                            <span class="table-cell date"></span>
                        </a>
                        <a class="dir table-row" href="javascript:void(0);">
                            <span class="table-cell icon">
                            </span>
                            <span class="table-cell name">FileExt</span>
                            <span class="table-cell size"></span>
                            <span class="table-cell rev"><input type="text" id="fileExt" name="fileExt"></span>
                            <span class="table-cell author"></span>
                            <span class="table-cell date"></span>
                        </a>
                        <a class="dir table-row" href="javascript:void(0);">
                            <span class="table-cell icon">
                            </span>
                            <span class="table-cell name">File</span>
                            <span class="table-cell size"></span>
                            <span class="table-cell rev"><input type="file" id="file" name="file"></span>
                            <span class="table-cell author"></span>
                            <span class="table-cell date"></span>
                        </a>
                        <a class="dir table-row" href="javascript:void(0);">
                            <span class="table-cell icon">
                            </span>
                            <span class="table-cell name">Oprea</span>
                            <span class="table-cell size"></span>
                            <span class="table-cell rev">
                                <input type="reset" name="reset" id="reset" value="reset">
                                <input type="button" name="upload" id="upload" value="upload">
                            </span>
                            <span class="table-cell author"></span>
                            <span class="table-cell date"></span>
                        </a>
                    </form>
            </div>
            </div>
            <div id="directory-list-readme"></div>
        </div>
    </div>
</div>
<script src="static/js/utils.js"></script>
<footer>
    Powered by <a href="https://github.com/godfather1103/packageRepo">PackageRepo Server</a>. © 2018 FocusOps Software Ltd.
</footer>
