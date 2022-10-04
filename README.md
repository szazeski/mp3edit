# mp3edit

Simple cross-platform command line tool to edit MP3 ID3 tags and 



`mp3edit -title="Example Title" -album="Example Album" -artist=Steve 01.mp3`


*Example to bulk change all files with new artist*

`mp3edit -artist="New Artist" *.mp3`


## Installation

For Linux/Mac:

```
wget https://github.com/szazeski/mp3edit/releases/download/v1.0.3/mp3edit_1.0.3_$(uname -s)_$(uname -m).tar.gz -O mp3edit.tar.gz && tar -xf mp3edit.tar.gz && chmod +x mp3edit && sudo mv mp3edit /usr/bin/
```

Mac

homebrew 
`brew install szazeski/tap/mp3edit`

Windows (Powershell)
Invoke-WebRequest https://github.com/szazeski/mp3edit/releases/download/v1.0.3/mp3edit_1.0.3_Windows_x86_64.tar.gz -outfile mp3edit.tar.gz; tar -xzf mp3edit.tar.gz; echo "if you want, move the file to a PATH directory like WINDOWS folder"
then move to C:\Windows\ or other PATH directory
