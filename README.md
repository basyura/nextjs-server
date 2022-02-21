# nextjs-server

## Deploy

```sh
$ npx create-next-app helloworld
$ cd helloworld
$ npm i
$ npm run build
$ npx next export
$ mv out ../public
$ go run main_nextjs_server.go 
```

## Failed to export

```
Error: Image Optimization using Next.js' default loader is not compatible with `next export`.
```

Please remove the Image tag.

