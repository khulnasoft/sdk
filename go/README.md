# Khulnasoft Golang SDK

This directory contains support for writing Khulnasoft programs in the Go language.  There are two aspects to this:

* `khulnasoft/` contains the client language bindings Khulnasoft program's code directly against;
* `khulnasoft-language-go/` contains the language host plugin that the Khulnasoft engine uses to orchestrate updates.

To author a Khulnasoft program in Go, simply say so in your `Khulnasoft.yaml`

    name: <my-project>
    runtime: go

and ensure you have `khulnasoft-language-go` on your path (it is distributed in the Khulnasoft download automatically).

By default, the language plugin will use your project's name, `<my-project>`, as the executable that it loads.  This too
must be on your path for the language provider to load it when you run `khulnasoft preview` or `khulnasoft up`.
