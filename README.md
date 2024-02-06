# webreboot
Reboot your remote device using browser

![](screenshot.gif)

Usage
-----

```bash
webreboot -bind address:port -dry-run
```
Arguments are optional

`-bind` default is `:851` (all interfaces)

`-dry-run` executes `shutdown -k` (warning only, no actual reboot/shutdown)