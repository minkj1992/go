# vscode with golang


## go to definition 사용 안되는 경우
> [참고 docs](https://github.com/golang/vscode-go#Set-up-your-environment)

- `go mod`를 사용할 경우 `gopls`를 사용하는 것을 권장하고 있다.

- 해결 방법: settings.json에 아래의 내용을 넣어주면 끝

```json
  
  // https://github.com/golang/tools/blob/master/gopls/doc/vscode.md
  "go.useLanguageServer": true,
  "[go]": {
      "editor.formatOnSave": true,
      "editor.codeActionsOnSave": {
          "source.organizeImports": true,
      },
      // Optional: Disable snippets, as they conflict with completion ranking.
      "editor.snippetSuggestions": "none",
  },
  "[go.mod]": {
      "editor.formatOnSave": true,
      "editor.codeActionsOnSave": {
          "source.organizeImports": true,
      },
  },
  "gopls": {
       // Add parameter placeholders when completing a function.
      "usePlaceholders": true,
  
      // If true, enable additional analyses with staticcheck.
      // Warning: This will significantly increase memory usage.
      "staticcheck": false,
  },
```