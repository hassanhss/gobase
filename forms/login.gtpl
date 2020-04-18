<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method="post">
	用户名:<input type="text" name="username"><br>
	密  码:<input type="password" name="password"><br>
	<input type="hidden" name="token" value="{{.}}">
	<input type="submit" value="登录">
</form>
</body>
</html>