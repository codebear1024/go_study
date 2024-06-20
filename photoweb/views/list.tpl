<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>List</title>
</head>
<body>
    <div>
        <h1>简单的相册</h1>
    </div>
    <ol>
      {{range .filename}}
        <li>
            <a href="/view?id={{.|urlquery}}">{{.|html}}</a>
            <a href="/view?id={{.|urlquery}}&action=del">删除</a>
        </li>
      {{end}}
    </ol>
    <div><a href="/upload" >上传</a></div>
</body>
</html>