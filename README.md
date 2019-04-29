# logfind
Tool for searching log files in a direcotry for given text.

## Usage

```
logfind [OPTIONS] [TEXT]

  -dirname string`                                                           
    Name of dir to search (default "/var/log")                           
  -match-all                                                                 
        Determine if search should match all searched strings (default "false")
```

### Example
`./logfind -dirname=/var/log -match-all=true Error Java Update`
