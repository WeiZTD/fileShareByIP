# fileShareByIP
 Simple IP whitelisted web FTP (Read only)

![](https://i.imgur.com/9DFZzZ2.jpeg)


## Usage
### Host a server

```fileShareByIP.exe  -dir="PATH/TO/DIR"```

show help message

```-help,-?```

directory to share(Required)

```-dir="PATH/TO/DIR"```

allow users to upload file 

```-upd=true```

listen at specific port

```-p=number```

### Browse files and folders

```@ip:port/file```

```@ip:port/``` will redirect to ```@ip:port/file```

### Upload file

```@ip:port/upload```

### IP whitelist control page (admin)

```@ip:port/admin```
