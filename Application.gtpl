<!DOCTYPE html>
<html>
<head>
	<title>Application</title>
</head>
<body>
	{{if ne .Name ""}}
		<h1>Welcome, {{.Name}}!</h1>
    	<p>Your password is {{.Password}}</p>
	{{end}}
	<p>Here is the content of our application.</p>
</body>
</html>
