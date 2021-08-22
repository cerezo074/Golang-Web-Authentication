package templates

func buildLogin() string {
	return `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>HMAC Example</title>
	</head>
	<body>
		<h1>Welcome to Login</h1>
		<p>Please insert email and pass to sign in</p>
		<form action="/login" method="post">
			<input type="email" name="email" />
			<input type="password" name="password" />
			<input type="submit" />
		</form>
		<br/>
		<a href="/register-view">No account? don't worry click here to register</a>
	</body>
	</html>`
}

func buildRegister() string {
	return `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>HMAC Example</title>
	</head>
	<body>
		<h1>Welcome to Register</h1>
		<br/>
		<p>Please fill fields to sign up</p>
		<form action="/register" method="post">
			<input type="email" name="email" />
			<input type="text" name="username" />
			<input type="password" name="password" />
			<input type="submit" />
		</form>
		<a href="/login-view">Have an account? don't worry click here to sing in</a>
	</body>
	</html>`
}

func buildHome() string {
	return `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>HMAC Example</title>
	</head>
	<body>
		<h1>Welcome to your home</h1>
		<br/>
		<p>This is you homepage(under construction)</p>
		<a href="/login-view">Wanna go out? click here to log out</a>
	</body>
	</html>`
}
