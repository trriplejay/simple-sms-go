# simple-sms-go

a simple package for sending sms via carrier email gateway, using smtp.PlainAuth

# Usage

intialize a client with username, password, and details for your smtp provider

```
emailClient := simplesmsgo.NewClient("username", "password", "smtp.gmail.com", "587")
```

send the sms by providin a number, provider, subject, and message

```
emailClient.Send(5555555432, emailClient.Providers.TMO, "Hello?", "Hello world!")
```

# Note

The carriers seem to be picky about how many emails are sent in this way and what their contents are. This is really meant to be used sparingly for simple occasional things.

This also only supports a limited number of carriers, but I'm happy to accept PRs with additional gateways.
