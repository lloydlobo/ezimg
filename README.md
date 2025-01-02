# ezimg

> [!WARNING]
> Highly unstable educational project.
> Use at your own risk!

## Build

```bash
go build -gcflags "-m=1" -ldflags="-s" -o ezimg .
```

## Run

### Usage

go run .

### Output

```bash
 ___ ___ _ __ __  __
| __|_  | |  V  |/ _]
| _| / /| | \_/ | [/\
|___|___|_|_| |_|\__/  (0.1.0), built with Go go1.22.5

watching images/
loading image1.jpg
loading image2.jpg
resizing image1.jpg
converting grayscale image1.jpg
saving image1.jpg
success
loading image3.jpg
resizing image2.jpg
resizing image3.jpg
converting grayscale image2.jpg
loading image4.jpg
saving image2.jpg
success
converting grayscale image3.jpg
saving image3.jpg
success
resizing image4.jpg
converting grayscale image4.jpg
saving image4.jpg
success
cleaning...
see you again~
2025/01/02 13:19:23 INFO took 1.651290324s
```

## Credits

- Thanks to [code-heim](https://github.com/code-heim/go_21_goroutines_pipeline)
  for expressing pipelines in a simple manner.
