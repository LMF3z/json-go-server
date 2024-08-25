# jsonGoServer

## Configurations:

Clone the repository and create an executable according to the operating system

# Linux

1. run
```bash
go build -o jsonGoServer main.go
```

2. Move to a folder. We suggest '/usr/local/bin/' with command

```bash
sudo cp jsonGoServer /usr/local/bin/
```

3. Grant run permissions with the command
```bash
sudo chmod +x /usr/local/bin/jsonGoServer
```

4. In the folder of your choice create a <name>.json file with the following structure
```json
{
    "heroes" : {
        "id" : "uuid",
        ...rest_of_props
    }
}
```

The id (String) is required to filter the data.

5. Open a terminal in the folder and run 
```bash
jsonGoServer -path=./<name>.json
```

By default the server is running on **http://localhost:8000**
to customize the port you can do it with the **-port=port** flag

6. On the endpoint **http://localhost:port/heroes** there will be an example response.
