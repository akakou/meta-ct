# meta-ct

[![Go](https://github.com/akakou/meta-ct/actions/workflows/go.yml/badge.svg)](https://github.com/akakou/meta-ct/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/akakou/meta-ct.svg)](https://pkg.go.dev/github.com/akakou/meta-ct)

This is a library to monitor [CT(Certificate Transparency)](https://www.rfc-editor.org/rfc/rfc9162.html) monitor library via Meta's service.

This library supports:
- [Certificate Transparency API](https://developers.facebook.com/docs/certificate-transparency-api)
- [Certificate Webhook](https://developers.facebook.com/docs/certificate-transparency/certificates-webhook)


To use this library, you must create a `Meta App` and then obtain the `App Id` and `Acess Token`.
- [Web page to request creating Meta App](https://developers.facebook.com/apps/?show_reminder=true)
- [How to obtain access token](https://developers.facebook.com/docs/facebook-login/guides/access-tokens?locale=ja_JP)
