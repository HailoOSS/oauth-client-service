# Oauth Client

This projects aims to provide generic helper endpoints to verify and
obtain information about third parties oauth tokens.

Endpoints available so far:

```
execute verify {"token": "abc", "provider": "google"}

{
    "valid": true
}
```

```
execute info {"token": "abc", "provider": "google"}

{
    "email": "gianni.moschini@bar.com",
    "givenName": "Gianni",
    "familyName": "Moschini"
}
```
