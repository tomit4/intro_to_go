## Setting Up Your Machine

Your machine will contain many version control <em>repositories</em> (managed by
Git, for example).

Each repository contains one or more <em>packages</em>, but will typically be a
single <em>module</em>.

Each package consists of one or more <em>Go source files</em> in a single
directory.

The path to a package's directory determines its <em>import path</em> and where
it can be downloaded from if you decide to host it on a remote version control
system like Github or Gitlab.

### A Note on GOPATH

The $GOPATH environment variable will be set by default somewhere on your
machine (typically in the home directory, `~/go`). Since we will be working in
the new "Go modules" setup, you <em>don't need to worry about that</em>. If you
read something online about setting up your GOTPATH, that documentation is
probably out of date.

These days you should <em>avoid</em> working in the `$GOPATH/src` directory.
Again, that's the old way of doing things and can cause unexpected issues, so
better to just avoid it.

### Get Into Your Workspace

Navigate to a location on your machine where you want to store some code. For
example, I store all my code in `~/workspace`, then organize it into subfolders
based on the remote location. For example,

`~/workspace/github.com/wagslane/go-password-validator` = https://github.com/wagslane/go-password-validator
