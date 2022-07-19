# vscode golang

## go extension
```
https://marketplace.visualstudio.com/items?itemName=golang.go
```

## settings.json
```
{
    "go.goroot": "/path/to/go",
    "go.gotoSymbol.includeGoroot": true,
    "go.toolsManagement.autoUpdate": true,
}
```

## go plugin (go:install/update)
```
go install -v golang.org/x/tools/gopls
go install -v github.com/go-delve/delve/cmd/dlv
```