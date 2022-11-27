- make

```console
help                  print help
build                 compile wgen
clean                 clean and remove wgen binary
run                   compile and run wgen
install               install wgen into ~/.local/bin
uninstall             unininstall wgen from ~/.local/bin
deps                  resolve dependencies
fmt                   format all
```

- workload (example)

```yaml
workload:
  - day:
      api_name_1:
        rate: 5
        unit: s
      api_name_2:
        rate: 10
        unit: m
  - day:
      api_name_2:
        rate: 2
        unit: h
```

- apispec (example)

```yaml
name: app-name API
baseUrl: http://localhost:8080
api:
  api_name_1:
    relativeUrl: /
    method: GET
  api_name_2:
    relativeUrl: /urlpath
    method: GET
```
