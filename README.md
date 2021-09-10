# hatebu2shiori
## What is
Generate a shell script to import Hatena bookmark data into [shiori](https://github.com/go-shiori/shiori), a pocket-like web page storage FLOSS server side application.  
I made this to learn the go-lang.  

* [go-shiori/shiori: Simple bookmark manager built with Go](https://github.com/go-shiori/shiori)
* [はてなブックマーク - 設定 - データ管理](https://b.hatena.ne.jp/-/my/config/data_management)

## How to

1. make shell file from atom

	```
	cd ~/Downloads/hatebu2shiori
	go run hatena2shiori.go
	  plz hatena your.bookmarks.atom file FULL path
	  > /Users/user/hatena_export/user.bookmarks.atom
	Success: /Users/user/hatena_export/hatebu2shiori.sh
	```

2. enter docker container (if use docker)

	```
	docker ps
	CONTAINER ID        IMAGE                   NAMES
	df7ec0f3ecaa        radhifadlillah/shiori   shiori

	docker cp ~/hatena_export/hatebu2shiori.sh <YOUR CONTAINER ID>:/go
	docker exec -it shiori sh
	```

3. run shell script

	```
	chmod +x hatebu2shiori.sh
	./hatebu2shiori.sh
	```

## License
MIT license