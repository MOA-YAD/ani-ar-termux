## Ani-ar

> This is a modified version of Ani-ar to work on Android termux
 
> watch this [video](https://go.screenpal.com/player/cZlhDJnno5o?title=0) To make sure it works
## install
Firstly install go

and clone this repository 

```bash
go run ani-ar-2.0/cmd/ani-ar.go
```
```bash
cd cmd
go build -o ani-ar
```
add this to .bashrc file
```bash
export PATH=$PATH:/data/data/com.termux/files/home/your/directory/ani-ar-2.0/cmd
```

## Interactive search

```bash
ani-ar
```

## search anime title

```bash
ani-ar search -q [search key]
ani-ar search -q "hunter x hunter"
```

## watch anime episode

> you should have either `mpv` or `vlc` to be able to stream the video.

```bash
ani-ar watch [anime-title] [episode-number]
ani-ar watch hunter-x-hunter-2011 50
```


## download anime episode

```bash
ani-ar download [anime-title] [episode-number] [folder-path]
ani-ar download hunter-x-hunter-2011 50 ~/Documents/anime/hunter-x-hunter/
```

## download all anime episodes (use 0 as episode number)

```bash
ani-ar download [anime-title] 0 [folder-path]
ani-ar download hunter-x-hunter-2011 0 ~/Documents/anime/hunter-x-hunter/
```


