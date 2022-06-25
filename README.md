# mp3edit

Simple cross-platform command line tool to edit MP3 ID3 tags and 



`mp3edit -title="Example Title" -album="Example Album" -artist=Steve 01.mp3`


*Example to bulk change all files with new artist*

`mp3edit -artist="New Artist" *.mp3`


## Installation


### Linux

`curl https://github.com/szazeski/mp3edit/releases/download/v1.0.1/mp3edit-linux-$(arch) --output mp3edit && chmod +x mp3edit && sudo mv mp3edit /usr/bin`

(Some distros give strange arch replies - your options are **amd64** for most Intel/Amd 64-bit, **386** for old Intel/Amd 32-bit, **arm** for older arm devices, and **arm64** for newer arm devices)

### Mac

Option 1:

`brew install szazeski/tap/mp3edit`

Option 2:

`curl https://github.com/szazeski/mp3edit/releases/download/v1.0.1/mp3edit-darwin-$(arch) --output mp3edit && chmod +x mp3edit`

Since the file is not signed with Apple, you will need to right click on it and select open. Gatekeeper will ask if you want to run it. You only have to do this once.

`sudo mv mp3edit /usr/bin`

(**amd64** is for intel mac and **arm64** is for M1)

### Windows (Powershell)

`wget https://github.com/szazeski/mp3edit/releases/download/v1.0.1/mp3edit-windows-amd64.exe -outfile mp3edit.exe` 

then move to C:\Windows\mp3edit.exe or other folder in your PATH.
