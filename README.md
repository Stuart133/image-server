# image-server

A toy image editing server. Support resizing & rotations

## Dependencies

`image-server` uses [libvips](https://github.com/libvips/libvips) to perform image processing functions. To compile and run this project `libvips` will need to be installed on the local machine.

For example, before building on ubuntu, you will need to run the following command

```
sudo apt install libvips-dev
```

To run the compiled binary only the core library is required. To install this on ubuntu run the following command:

```
sudo apt install libvips
```

## Scaling notes

Currently the server is not optmized for scale. All image processing requests are handled synchronously in process. If the load increased request might be dropped or the entire server could stop functioning. While more servers could be spun up and function independently, each server would remain vulnerable be being overloaded. In that case all in flight requests would be dropped, so a single large job could negatively affect operations for other users.

A more scalable architecture could consist of three components:

- Image upload/download
  - Decoupling the upload & download from the processing requests would stop request or API errors from requiring the upload work to be duplicated
  - Multiple operations could be run without needing to reupload & download each time
- Image processing workers
  - A fleet of independent workers allowing on demand scaling to adapt to incoming traffic
  - Failures in a given worker would require processing to be retried, but would not fail the entire operation
- Image API
  - REST API to serve user requests for uploads/downloads & processing operations
  - Operations are completely asynchronously, the API would allow users to query the current status of operations
- Persistent Datastore
  - Exact format would depend on usage & requirements, but a SQL server would likely fulfil the needs of this application

