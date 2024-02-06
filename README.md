# webreboot
Reboot your remote device using browser

![](screenshot.gif)

# Usage

```bash
webreboot -bind address:port -dry-run
```
Arguments are optional

`-bind` default is `localhost:851`

`-dry-run` executes `shutdown -k` (warning only, no actual reboot/shutdown)

## Behind nginx:

```nginx

# backend
location /power/ {
    proxy_pass http://localhost:851;
}

# embedded ui
location /ui/ {
    rewrite /ui/(.*) /$1 break;
    proxy_pass http://localhost:851;
}
```