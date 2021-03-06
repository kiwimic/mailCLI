# mailCLI

mailCLI is CLI wrapper app written purely in go which use **https://github.com/go-mail/mail** and transforms it to command line interface app.

### Example:

#### Command

```
.\mailCLI.exe -to="examplemail@gmail.com,examplemail2@zoho.eu" -attach="file.png" -subj="Subject of mail" -body="Hi<br><b>Cheers<\b><br>bye!"
```

#### Usage:
 ```
  -attach string
        files you want to attach, if multiple files use comma as separator ','
  -body string
        Body-content of mail message or a path to text/html file (default "Example body of message")
  -config string
        location of json config file (only if you want to pass other config file than this from 'mail_config' env var")
  -subj string
        subject of mail message (default "Automated mail sended from mailCLI app")
  -to string
        mail recipient, if multiple recipients use comma as separator ','
```

#### Installation

1. Install dependecies (right now this is the only one), also from creator of gomail package:
 *It requires Go 1.2 or newer. With Go 1.5, no external dependencies are used.*

```
go get gopkg.in/gomail.v2
go get github.com/kiwimic/mailCLI
```

2. In your terminal find location of your go src folder and then find package mailCLI  */src/github.com/kiwimic/mailCLI*

3. Create dir named 'config' and then create config.json file inside **(remeber that you will have to store your password and login for STMP server)**
copy this template and fill it with your personal data.

```
{
    "mailserver"             : {
        "smtp_gate"     : "smtp.gmail.com",
        "smtp_port"     : 587,
        "from_mail"     : "yourmail@gmail.com"
    },
    "authentication"           : {
        "login" : "yourlogin",
        "password" : "yourpassowrd"
    }
}

```
4. Create dir named logs, and then create enviroment variables mail_config and mail_log with exact paths to config.json, and log.txt (log.txt will be created automatically as this file)

for me I used (it's location of my source code)

```
mail_config = C:\goworkspace\src\github.com\mailCLI\config\config.json
mail_log = C:\goworkspace\src\github.com\mailCLI\logs\log.txt
```

5. Now go build . or go install (depends on your personal needs) inside */src/github.com/kiwimic/mailCLI*

6. Test and play with your app :) 

with go build .
```
.\mailCLI.exe -to="examplemail@gmail.com,examplemail2@zoho.eu" -attach="file.png" -subj="Subject of mail"
```
or after go install

```
mailCLI -to="examplemail@gmail.com,examplemail2@zoho.eu" -attach="file.png" -subj="Subject of mail"

```

#### To do

- [ ] Add option to take authentication (login, password) from  enviroment variables.
- [ ] Add tests for each function
- [x] Add some logging and history of operations
- [x] Add location to configuration file, and log file as env vars


Contact me by **michalsiwik[at]gmail.com**
