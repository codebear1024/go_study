<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>Register</title>
</head>
<body>
    <form method="post" action="/register" enctype="multipart/form-data">
        <div><label>账号：<input type="text" name="username"></label></div>
        <div><label>密码：<input type="passWord" name="passWord"></label></div>
        <div><label>再次输入密码：<input type="passWord" name="passwordconfirm"></label></div>
        <div><label>许可码：<input type="text" name="license"></label></div>
        <div><input type="submit" value="注册"></div>
        <div><span style="color: red;">{{.prompt}}</span></div>
    </form>
</body>
</html>