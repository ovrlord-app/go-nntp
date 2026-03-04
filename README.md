Go NNTP (news) client package
=======

A Go package for interacting with NNTP (Network News Transfer Protocol) servers.
It provides a simple client for connecting, authenticating, browsing groups, and fetching articles.

Forked from [nntp](https://github.com/chrisfarms/nntp/tree/master) which is a fork of [nntp-go](http://code.google.com/p/nntp-go/).

Forked to improve testing and handling of line endings and yEnc support.

## 📚 Guides & Documentation

- 🤝 [Contributing Guide](docs/CONTRIBUTING.md)
- 🔒 [Security Policy](docs/SECURITY.md)

### Example use

```bash
go get github.com/ovrlord-app/go-nntp
```

```go
	// connect to news server
	conn, err := nntp.Dial("tcp", "news.example.com:119")
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}

	// auth
	if err := conn.Authenticate("user", "pass"); err != nil {
		log.Fatalf("Could not authenticate")
	}

	// connect to a news group
	grp := "alt.binaries.pictures"
	_, l, _, err := conn.Group(grp)
	if err != nil {
		log.Fatalf("Could not connect to group %s: %v %d", grp, err, l)
	}

	// fetch an article
	id := "<4c1c18ec$0$8490$c3e8da3@news.astraweb.com>"
	article, err := conn.Article(id)
	if err != nil {
		log.Fatalf("Could not fetch article %s: %v", id, err)
	}

	// read the article contents
	body, err := ioutil.ReadAll(article.Body)
	if err != nil {
		log.Fatalf("error reading reader: %v", err)
	}
```

## License

This project is licensed under a BSD derivative - see the [LICENSE](LICENSE) file for details.

## Support

- 🐛 [Issues](https://github.com/ovrlord-app/go-nntp/issues)