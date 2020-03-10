## TwitExamine

This is a dumb little program that, given the int64 ID of a Tweet, shows the JSON underlying that Tweet,

### Prerequisites

- Go 1.11 or higher, since this uses Go modules. Make sure you have `GO111MODULE` set appropriately for your system. You can probably compile this  with earlier versions of Go, but I've not tried that.
- A Twitter application established for your Twitter account at [Twitter Developers](https://developer.twitter.com/). Your application doesn't need to use "Log In With Twitter" but does need read and write permissions, which are the defaults when you create an app. There will be four values assigned to your application under "Keys and Tokens": two consumer API keys and two access tokens. Copy these and keep them safe until you need them later.

### To Build

- Check this repository out locally.
- In the directory, run `go build .` or `go install .`

If you ran `go build`, you should see a binary, `twitexamine`, created in the repository directory. Leave it there, put it elsewhere on your path, whatever. If you ran `go install`, the `twitexmine` binary has been moved to your Go `bin` directory, which should be in your `$PATH` (which, now that I think about it, is weird when you have Go modules configured, but that's a question for some other day).

### To Run

- Set the following environment variables using the values you generated for  your Twitter app (you kept them safe, right?):

```shell
export TC_TWITTER_CONSUMER_KEY={twitter API key}
export TC_TWITTER_CONSUMER_SECRET={twitter API secret key}
export TC_TWITTER_ACCESS_TOKEN={twitter access token}
export TC_TWITTER_ACCESS_TOKEN_SECRET={twitter secret access token}
```

If you don't set these, `twitexamine` will quit with errors.

Then, run the program:

```shell
# deletes anything older than 30 days
twitexamine -i <tweet_id> 
```

You may want to install [jsonpp](https://jmhodges.github.io/jsonpp/), or your favorite JSON pretty-printer, and pipe the output of `twitexamine` through it. I like that one because it sorts the dictionary keys by default.

### Thanks To

- [Anaconda](https://github.com/ChimeraCoder/anaconda)
- [Logrus](https://github.com/sirupsen/logrus)

### License

This program is licensed under the terms of the [Blue Oak Model License 1.0.0](https://blueoakcouncil.org/license/1.0.0)
