# fileShareByIP
 Simple IP whitelisted web FTP

![](https://i.imgur.com/9DFZzZ2.jpeg)


# Usage

### Host a read only server

```fileShareByIP.exe -dir="PATH/TO/DIR"```

## Parameters

### Show help message

```-help,-?```

### Directory to share (required)

```-dir="PATH/TO/DIR"```

### Allow user to upload file

```-upd=true```

### Listen at specific port

```-p=number```

### Browse files and folders

```@ip:port/file```

```@ip:port/``` will redirect to ```@ip:port/file```

### Upload file

```@ip:port/upload```

### IP whitelist control page (admin)

```@ip:port/admin```
