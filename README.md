# smtp-testing

Golang Tutorial: Send Email via SMTP With MailHog!

`go install`

In order to run project type:
`go run ./cmd/cli/`

## MailHog

MailHog is an email testing tool that makes it super easy to install and configure a local email server. MailHog sets up a fake SMTP server. You can configure your preferred web applications to use MailHog’s SMTP server to send and receive emails.

### MailHog on Docker

Docker is one of the easiest ways to set up MailHog across any platform.

First, make sure that you have Docker installed on your machine. Next, you can start MailHog directly from Docker Hub by running the following command:

`docker run -d -p 1025:1025 -p 8025:8025 mailhog/mailhog`

Let’s break down the above command and its directives:

- docker run starts a docker container.
- -d directive runs the docker container in the background.
- -p 1025:1025 and -p 8025:8025 directives expose MailHog’s default SMTP ports to your local ones.
- mailhog/mailhog grabs the latest MailHog build from Docker Hub.
Alternately, if you want to mount MailHog’s Maildir to your local filesystem (instead of inside a Docker container), you can run the following command: