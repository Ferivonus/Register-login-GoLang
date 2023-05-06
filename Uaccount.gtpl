<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Account Manager</title>
    <style>

	body {
	font-family: Arial, sans-serif;
}

h1 {
	text-align: center;
	margin-bottom: 20px;
}

h2 {
	text-align: center;
	margin-top: 30px;
}

form {
	margin: 0 auto;
	max-width: 400px;
	padding: 20px;
	border: 1px solid #ccc;
	border-radius: 5px;
	box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
}

label {
	display: block;
	margin-bottom: 10px;
	padding-left: 5px;
}

input[type="email"],
input[type="text"],
input[type="password"] {
	display: block;
	margin-bottom: 20px;
	padding: 10px;
	width: 100%;
	border: 1px solid #ccc;
	border-radius: 5px;
	box-sizing: border-box;
}

input[type="submit"] {
	background-color: #4CAF50;
	color: white;
	padding: 10px;
	border: none;
	border-radius: 5px;
	cursor: pointer;
}

input[type="submit"]:hover {
	background-color: #45a049;
}

.form-error-message {
	color: red;
	margin-top: 10px;
	padding-left: 5px;
}

.form-success-message {
	color: green;
	margin-top: 10px;
	padding-left: 5px;
}

/* Login Form Styles */
.login-form {
	margin: 20px auto;
	max-width: 400px;
	padding: 20px;
	border: 1px solid #ccc;
	border-radius: 5px;
	box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
}

.login-form label {
	display: block;
	margin-bottom: 10px;
}

.login-form input[type="text"],
.login-form input[type="password"] {
	display: block;
	margin-bottom: 20px;
	padding: 10px;
	width: 100%;
	border: 1px solid #ccc;
	border-radius: 5px;
	box-sizing: border-box;
}

.login-form input[type="submit"] {
	background-color: #4CAF50;
	color: white;
	padding: 10px;
	border: none;
	border-radius: 5px;
	cursor: pointer;
}

.login-form input[type="submit"]:hover {
	background-color: #45a049;
}

.login-form .error-message {
	color: red;
	margin-top: 10px;
}

.login-form .success-message {
	color: green;
	margin-top: 10px;
}

/* Sign-up Form Styles */
.signup-form {
	margin: 20px auto;
	max-width: 400px;
	padding: 20px;
	border: 1px solid #ccc;
	border-radius: 5px;
	box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
}

.signup-form label {
	display: block;
	margin-bottom: 10px;
}

.signup-form input[type="text"],
.signup-form input[type="password"] {
	display: block;
	margin-bottom: 20px;
	padding: 10px;
	width: 100%;
	border: 1px solid #ccc;
	border-radius: 5px;
	box-sizing: border-box;
}

.signup-form input[type="submit"] {
	background-color: #4CAF50;
	color: white;
	padding: 10px;
	border: none;
	border-radius: 5px;
	cursor: pointer;
}

.signup-form input[type="submit"]:hover {
	background-color: #45a049;
}

.signup-form .error-message {
	color: red;
	margin-top: 10px;
}

.signup-form .success-message {
	color: green;
	margin-top: 10px;
}

#signup-username {
	display: block;
	margin-bottom: 20px;
	padding: 10px;
	width: 100%;
	border: 1px solid #ccc;
	border-radius: 5px;
	box-sizing: border-box;
}
	</style>
</head>
<body>
	<div style="background-color:#e5e5e5;padding:15px;text-align:center;">
		<h2>Welcome to my homework GUI</h2>
		<h2>Yeah,it is wierd to say that But I just said it.</h2>
	</div>

  <h1>Welcome to our website, would you like to log in or sign up?</h1>
  <h2>Unsafe one</h2>

  <h2>Login</h2>
  	<form action="/UnsafeLogIn" method="POST" class="login-form">
    <label for="login-username">Username:</label>
    <input type="text" id="login-username" name="Login_username" required><br>

    <label for="login-password">Password:</label>
    <input type="password" id="login-password" name="Login_password" required><br>

    <input type="submit" value="Login">
  </form>

  <h2>Sign up</h2>
  <form action="/UnsafeSignIn" method="POST" class="signup-form">
    <label for="signup-username">Username:</label>
	<input type="text" id="signup-username" name="SignIn_username" required><br>

	<label for="signup-password">Password:</label>
	<input type="password" id="signup-password" name="SignIn_password" required><br>

	<label for="signup-email">Email:</label>
	<input type="email" id="signup-email" name="SignIn_email" required><br>


    <input type="submit" value="Sign up">
  </form>
</body>
</html>
