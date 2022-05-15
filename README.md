# mp3edit

Simple cross-platform command line tool to edit MP3 ID3 tags and 



`mp3edit -title="Example Title" -album="Example Album" -artist=Steve 01.mp3`


*Example to bulk change all files with new artist*

`mp3edit -artist="New Artist" *.mp3`


## Installation


Linux

`curl https://github.com/szazeski/mp3edit/releases/download/v1.0.0/mp3edit-linux-$(arch) --output mp3edit && chmod +x mp3edit && sudo mv mp3edit /usr/bin`