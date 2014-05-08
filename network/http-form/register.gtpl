<html>
<head>
<title>Register</title>
</head>
<body>
<form action="/register/" method="post">
Name:
<input type="text" name="name"/>
<BR/>
Gender:
<input type="radio" name="gender" value="Male" checked>Male
<input type="radio" name="gender" value="Female">Female
<BR/>
Language:
<input type="checkbox" name="language" value="Java">Java
<input type="checkbox" name="language" value="Objective-C">Objective-C
<input type="checkbox" name="language" value="Golang">Golang
<BR/>
Platform:
<select name="platform">
<option value="Android">Android</option>
<option value="iOS">iOS</option>
<option value="WindowsPhone">WindowsPhone</option>
</select>
<BR/>
<input type="hidden" name="token" value="{{.}}"/>
<input type="submit" value="submit"/>
</form>
</body>
</html>
