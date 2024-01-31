# Origin

Goland is a programming language created by Google in 2007, and later became open soruced in 2009

# Why use Golang?

Infrastructure changed a lot in the last few years. Multiprocessors became common, and using Cloud Providers with hundreds or thousands of servers to deploy applications became universal practices. This means infrastructure became more <b>scabable, dynamic and with more capacity</b>.

However programming languages at the time did not take fully advantage of this scenario.

Applications would only execute one task at a time in order, and with infrastructure improvements, it became possible to execute multiple tasks in parallel, in a process called <b>Multi-Threading</b>.

Of course it had many challenges, such as concurrency, which means dealing with many things at once, for example acessing and modifying shared data at the same time. Developers needed to write code to prevent such conflicts. Many languages have support for that, but often code could become complex, and with a higher risk of human errors.

Go was design to run on multiple cores and built to support concurrency, in a <b>cheaper</b> and <b>easier</b> way.

# Characteristics

Advantages:
- Simple sintax: it's easy to learn, read and write code
- Fast build time, start up and run
- Requires fewer resources

It's also a compiled language. It compiles into a single binary file, it's faster than interpreted languages, like Python, and you can deploy and run it on different platforms.

# Environment

- Go compiler (comes with Go CLI)
- IDE: Visual Studio Code

Download: https://go.dev/doc/install

Don't forget to install the Go extension for VSCode.
If any windows appear on the corner of your IDE about Go Extension, just go ahead and click "Install All".