package email

import "fmt"

const EmailSubject = "WelCome to petpujaries"

const EmailTemplate = `<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
<meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
</head>
<body>
<br>user name : %v <br>
<br>password : %v  <br>
<div class="moz-signature"><i><br>
<br>
Regards<br>
petpujaries<br>
<i></div>
</body>
</html>
`

func CreateEmail(userName, password string) string {
	return fmt.Sprintf(EmailTemplate, userName, password)
}
