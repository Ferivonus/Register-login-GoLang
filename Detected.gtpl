{{if and (ne .IP "") (eq .Method "POST")}}
		<h2>Dude, why are you trying to f* my website?</h2>
		<h2>I know your IP address which is {{.IP}}</h2>
		<h2>I could find your home, you know that, right?</h2>
	{{end}}