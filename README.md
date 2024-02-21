## Free Code Camp's Intro To Go Course With Lane Wagner

These files are simply a following along of [Free Code Camp's Intro To Go Course](https://www.youtube.com/watch?v=un6ZyFkqFKo&pp=ygUaZnJlZSBjb2RlIGNhbXAgaW50cm8gdG8gZ28%3D) with [Lane Wager](https://github.com/wagslane).

This tutorial series utilizes Lane's REPL on his site, Boot.dev. We are simply
following along on our Linux machine. I won't cover how to install Go, as I
already have done so, but you can figure out how to do so on Arch Linux at their
excellent [ArchWiki](https://wiki.archlinux.org/title/Go).

To run a go program, simply make your project directory and then invoke the
following commands:

```bash
go mod init example/name_of_project
```

Then write your go code by creating the file:

```bash
touch name_of_project.go
```

If you wish to run it simply use the go command line interface:

```bash
go run name_of_project.go
# or
go run .
```

If you wish to actually compile it to a binary, it's a simple as:

```bash
go build name_of_project.go
```

This will output an executable binary by the same name as the go file (minus the
.go extension). Thusly you can then run it like any executable:

```bash
./name_of_project
```
