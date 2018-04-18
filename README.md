# packagesender
Send network packages and stress test

## Usage
* clone https://github.com/joeke80215/packagesender.git
* in src,build ```go build -o psdr .```
* execute ```./psdr -H <target host> -P <target port> -mt tcp```

## Example
```./psdr -H 192.168.100.100 -P 80 -mt tcp```

## Optional
* -H : Target host (default "127.0.0.1")
* -P : Target host port (default "8080")
* -mt : Request method(tcp/udp/http) (default "tcp")
* -pn : Send how many packages (default 10)
* -ps : Package size (byte) (default 1024)
* -rn Execute how many Routines (default 5)

## Note
* send total packages = routines(-rn) * packages(-ps)
* watch out your cpu and memory when setting routines and packages
