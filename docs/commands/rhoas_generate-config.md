## rhoas generate-config

Generate configurations for the service context

### Synopsis

Generate configuration files for the service context to connect with to be used with various tools and platforms

```
rhoas generate-config [flags]
```

### Examples

```
## Generate configurations for the current service context in json format
$ rhoas generate-config --type json

```

### Options

```
      --client-id string       Client ID of the service account
      --client-secret string   Client secret of the service account
      --client-secret-stdin    Take the client secret from stdin
      --generate-auth          Create service account
      --name string            Name of the context
      --output-file string     Sets a custom file location to save the configurations
      --type string            Type of configuration file to be generated
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI

