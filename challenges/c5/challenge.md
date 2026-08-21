# Application Configurator

Your program needs to read a configuration file in JSON format with this structure:

```json
{
  "port": 8080,
  "debug": true,
  "allowed_hosts": ["localhost", "127.0.0.1"],
  "timeout_seconds": 30
}
```

Load this JSON into an appropriate Go structure. If the file does not exist or has an invalid format, the program should display a friendly error message and use default values.

Hint: Use os.Open with error handling, json.Unmarshal, and define a struct with tags. What values ​​could be default constants?