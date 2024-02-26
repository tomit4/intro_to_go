## Why Is It Called A "Mutex"?

Mutex is short for [mutual exclusion](<https://en.wikipedia.org/w/index.php?title=Lock_(computer_science)&useskin=vector>), and the conventional name for the data structure that provides it is "mutex", often abbreviated to "mux".

It's called "mutual exclusion" because a mutex <em>excludes</em> different
threads (or goroutines) from accessing the same data at the same time.
