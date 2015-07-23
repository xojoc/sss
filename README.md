# Super simple server
*sss* is small static server written in [golang](http://golang.org) to easily test static site on localhost.

# Usage
First install *sss*:
```
go install xojoc.pw/sss
```

Then run *sss* in the directoy containing your site.
```
(cd site; sss&)
$BROWSER localhost:8080
```
by default *sss* runs on port `8080`. If an argument is provided it will used as the port.
```
(cd site; sss :4567&)
$BROWSER localhost:4567
```

# Who?
*sss* was written by Alexandru cojocaru (http://xojoc.pw).

# License
*sss* is released under the GPLv3 or later, see [COPYING](COPYING).