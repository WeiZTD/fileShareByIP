# fileShareByIP
 Simple IP whitelisted web FTP (Read only)

![](https://i.imgur.com/9DFZzZ2.jpeg)


## Usage
### Host a server

```fileShareByIP.exe  -dir="PATH/TO/DIR"```

Show help message

```-help,-?```

Rirectory to share directory (Required)

```-dir="PATH/TO/DIR"```

Allow users to upload file 

```-upd=true```

Listen at specific port

```-p=number```

### Browse files and folders

```ip:port/file```

```/``` will redirect to ```/file```

### Upload file

```@ip:port/upload```

### IP whitelist control page (admin)

```@ip:port/admin```
