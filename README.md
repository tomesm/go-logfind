# logfind
Tool for searching log files in a given directory for a given words. You can use this tool as a better "grep" command. You can print otuput line by line or in a JSON format. 

Tool provides information about file name, line number where match(es) were found and prints the whole line from the file. All searches ignore cases so that it does not matter if you use lowercase, UPPERCASE or CamelCase.

NOTE: For now matching patterns (regexp) are not supported. Only simple strings are considered"

## Usage

```
logfind [OPTIONS] [STRINGS]

  -dirname string                                                           
      Name of dir to search (default "/var/log")
  -file-type
      Type/suffix of files to be searched (default ".log")
  -match-all                                                                 
      Determine if search should match all searched strings (default "false")
  -format
      Print output in JSON format (default "false")
      
```

### Example
`./logfind -dirname=/var/log -match-all=true Error Java update`

##### Output:

`file.log:34:This is a line where is information about Java ERROR from update`
