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

Contant me by **michalsiwik[at]gmail.com**
