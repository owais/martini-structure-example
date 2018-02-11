UPDATE
=========================

This is probably not the best way to structure Go projects. This was a result of me trying to build a Django project in Go when I had just started with Go. Creating different packages for very small things like routes or views is not that useful.

martini-structure-example
=========================

An example to structure a complex Martini (or any other golang) project. Promotes separation of concerns and app re-usabiity. 


The idea is to put separate concerns into "apps" and make them reusable. Any other martini project should be able mount an app on a given url namespace and it should just work.
