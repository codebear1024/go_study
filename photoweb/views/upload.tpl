<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>Upload</title>
</head>
<body>
    <form method="POST" action="/upload" enctype="multipart/form-data">
    <div>Choose an image to upload: <input name="image" type="file" /></div>
    <div><input type="submit" value="Upload" /></div>
    </form>
    <div><a href="/" >首页</a></div>
</body>
</html>