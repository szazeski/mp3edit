# mp3edit

Simple cross-platform command line tool to edit MP3 ID3 tags (and one day trim and join mp3 audio too)

```
mp3edit <options> <filename>
 easy mp3 tag editor and simple audio operations
 version 1.0.3 built 2022-06-25
  -clear
  -copyTagsFrom=otherFile.mp3
  -title="New Title"
  -artist="New Artist"
  -album="New Album"
```


`mp3edit -title="Example Title" -album="Example Album" -artist=Steve 01.mp3`


*Example to bulk change all files with new artist but leave all other fields alone*

`mp3edit -artist="New Artist" *.mp3`

*Example to clear all tags*

`mp3eedit -clear file.mp3`


## Installation

**Linux/Mac**

```
wget https://github.com/szazeski/mp3edit/releases/download/v1.0.3/mp3edit_1.0.3_$(uname -s)_$(uname -m).tar.gz -O mp3edit.tar.gz && tar -xf mp3edit.tar.gz && chmod +x mp3edit && sudo mv mp3edit /usr/bin/
```

**Mac**
 
```
brew install szazeski/tap/mp3edit
```

**Windows (Powershell)**
```
Invoke-WebRequest https://github.com/szazeski/mp3edit/releases/download/v1.0.3/mp3edit_1.0.3_Windows_x86_64.tar.gz -outfile mp3edit.tar.gz; tar -xzf mp3edit.tar.gz; echo "if you want, move the file to a PATH directory like WINDOWS folder"
then move to C:\Windows\ or other PATH directory
```
