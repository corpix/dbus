dbus
----------

This is a work-in-progress project.

## FAQ

### There is github.com/godbus/dbus, why another implementation?

Godbus package:

* has not much composable and easy to use building blocks to create D-BUS services
* has no building blocks to create your own implementation of dbus-daemon
* is just a library to write very primitive clients not to serve you with protocol primitives
* has hardcoded limits on things like signature length/depth
* contains big amount of complicated and fragile code
* panics when it could just return you an error

So... this package was made with simplicity and composability in mind.

> Well, but why you just don't report this issues?

Because you will need rewrite whole library from the scratch to resolve them.
