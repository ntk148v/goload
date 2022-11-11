<div align="center">
	<h1>GoLoad</h1>
	<blockquote align="center">Generate load for testing.</blockquote>
	<p>
		<a href="https://github.com/ntk148v/goload/blob/master/LICENSE">
			<img alt="GitHub license" src="https://img.shields.io/github/license/ntk148v/goload?style=for-the-badge">
		</a>
		<a href="https://github.com/ntk148v/goload/stargazers">
			<img alt="GitHub stars" src="https://img.shields.io/github/stars/ntk148v/goload?style=for-the-badge">
		</a>
		<br>
<!--		<a href="https://github.com/ntk148v/goload/actions">
			<img alt="Windows Build Status" src="https://img.shields.io/github/workflow/status/ntk148v/goload/Windows%20Build?style=flat-square&logo=github&label=Windows">
		</a>
		<a href="https://github.com/ntk148v/goload/actions">
			<img alt="GNU/Linux Build Status" src="https://img.shields.io/github/workflow/status/ntk148v/goload/Linux%20Build?style=flat-square&logo=github&label=GNU/Linux">
		</a>
		<a href="https://github.com/ntk148v/goload/actions">
			<img alt="MacOS Build Status" src="https://img.shields.io/github/workflow/status/ntk148v/goload/MacOS%20Build?style=flat-square&logo=github&label=MacOS">
		</a>
		<br>-->
	</p><br>
</div>

## Install

```shell
go get -u github.com/ntk148v/goload
```

## Usage

- Basic:

```shell
# Allocate 1024MB memory in 20MB block
./goload mem -total 1024 -block 20
# Run 100% of all CPU cores for 10 seconds
./goload cpu -time 10
# Run 100% of 2 CPU cores for 10 seconds
./goload cpu -cores 2 -time 10
```

- Docker:

```shell
# Build image
$ docker build -t kiennt26/goload .
# Or simple pull from Hub docker
$ docker pull kiennt26/goload
$ docker run -it --rm -m 8m kiennt26/goload mem -total 10 -block 2
Alloc = 2 MB    TotalAlloc = 2 MiB      Sys = 8 MB      NumGC = 0
Alloc = 4 MB    TotalAlloc = 4 MiB      Sys = 12 MB     NumGC = 1
Alloc = 6 MB    TotalAlloc = 6 MiB      Sys = 16 MB     NumGC = 1
# OOM -> Get killed at 8MB
```
