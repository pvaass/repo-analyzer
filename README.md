# repo-analyzer


## How it works
### As a stand-alone application
```shell
repo-analyzer https://github.com/pvaass/repo-analyzer
```

### As a library
```go

platform, _ := getPlatforms().ForURI("https://github.com/pvaass/repo-analyzer") 
repository := repository.New(platform, "https://github.com/pvaass/repo-analyzer")

rules := detector.GetRules('/path/to/my/rules')

var results []detectors.Result := analyze.Run(repository, rules)
// type Result struct {
//    Identifier string
//    Score int
// }
```


# Creating your own rules
You can create your own rules by placing a `*.json` file in the `rules/` directory.
Supported strategies are:
- `file-exist`:
    - Will return a score of 100 when *any* of the files in the argument list are found
    - args: ["file-1", "file-2", ...]
    
- `composer#d`:
    - Will return a score of 100 when the first argument is found as a composer.json dependency
    
    ```javascript
    "args": [
      "symfony/symfony" // Dependency to find
     ]
    ```
    
- `composer#f`:
    - Will return a score of 100 when the requested field matches the value in composer.json
    
    ```javascript
    "args": [
      "name" //field name, 
      "symfony/symfony" //value
     ]
     ```  
     - Supported field names: `name`
     
- `npm#d`:
    - Will return a score of 100 when the first argument is found as a npm (package.json) dependency
    
    ```javascript
    "args": [
      "react" // Dependency to find
     ]
    ```
