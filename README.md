## WIKI App on golang

Improved version of https://golang.org/doc/articles/wiki/

### Build

```sh
docker build -t dmbur/wiki/app -f ./deployments/docker/app/Dockerfile .
docker build -t dmbur/wiki/nginx-proxy ./deployments/docker/nginx/
```

Build optimized app image
```sh
docker build -t dmbur/wiki/app-optimized -f ./deployments/docker/app/optimized.Dockerfile .
```

Remove intermediate images
```sh
docker image prune --filter label=stage=intermediate
```

### Further improvements
* add pager
* add navigations
* add breadcrumbs
* add Router.generate()
* 404, 50x pages
* _move external packages into pkg directory_
