<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>Login</title>
</head>
<body>
    <form method="post" action="/login" enctype="multipart/form-data">
        <div>
        <label>账号：
            <input type="text" name="username">
        </label>
        </div>
        <div>
            <label>密码：
                <input type="passWord" name="passWord" >
            </label>
        </div>
        <input type="submit" value="登陆">
    </form>
    <button onclick="window.location.href='/register'">注册</button>
    <div><span style="color: red;">{{.prompt}}</span></div>
</body>
</html>