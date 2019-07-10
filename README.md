# mailCLI

mailCLI is CLI wrapper app written purely in go which use **https://github.com/go-gomail/gomail** and transforms it to command line interface app.

### Example:

#### Command

```
.\mailCLI.exe -to="examplemail@gmail.com,examplemail2@zoho.eu" -attach="file.png" -subj="Subject of mail"
```

#### Usage:
 ```
 -attach string
        files you want to attach, if multiple files use comma as separator ','
  -config string
        location of json config file (default "config.json")
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

3. Create dir named 'config' and then create config.json file inside **(remeber that you will have to store your passowrd and login for STMP serwer)**
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

4. Now go build . or go install (depends on your personal needs) inside */src/github.com/kiwimic/mailCLI*

5. Test and play with your app :) 

with go build .
```
.\mailCLI.exe -to="examplemail@gmail.com,examplemail2@zoho.eu" -attach="file.png" -subj="Subject of mail"
```
or with go install

```
mailCLI -to="examplemail@gmail.com,examplemail2@zoho.eu" -attach="file.png" -subj="Subject of mail"

```

#### To do

1. Add option to take authentication (login, password) from  enviroment variables.
2. Add tests for each function
3. Add some logging and history of operations


Contact me by **michalsiwik[at]gmail.com**
