# generate-multifields

![Staging](https://github.com/yanyi/generate-multifields/workflows/Staging/badge.svg?branch=master)

`generate-multifields` (`gmf`) is a CLI library to help generate multiple fields of
a given input format of your GraphQL queries or mutations, by repeating them for
a number of times.

## Intentions of this CLI

There was an instance I had to generate hundreds of queries or mutations that
run by IDs. I needed a tool to generate them with a given input.

>"But you could do this with a small shell script!"

Yes. It was also intended for me to play with [Cobra](https://github.com/spf13/cobra)
and GitHub Actions.

## Downloading

### Released Versions

Get the latest binary for your OS/Architecture from the
[Releases tab](https://github.com/yanyi/generate-multifields/releases/latest).
The following is the `curl` instructions for macOS:

```sh
$ curl -L https://github.com/yanyi/generate-multifields/releases/latest/download/gmf-darwin_x86_64 > gmf && \
  chmod +x gmf && \
  mv gmf /usr/local/bin

$ which gmf
/usr/local/bin/gmf
```

Using Docker for released versions:

```sh
$ docker pull yanyi/generate-multifields:0.1.0
0.1.0: Pulling from yanyi/generate-multifields
```

<!-- Collapse the developmental version instructions -->
<details>
  <summary>Development Version</summary>

#### Go

Clone the repository and run:

```sh
$ go build -i -o gmf && mv gmf /usr/local/bin

$ which gmf
/usr/local/bin/gmf
```

#### Docker

```sh
$ docker pull yanyi/generate-multifields:latest
latest: Pulling from yanyi/generate-multifields
```

</details>
<!-- End of collapsing developmental version instructions -->

## Running

Using a text file, the CLI looks for instances of `$id` and replaces it with
the IDs when running the `mutations` command. Here is an example of a text file:

```txt
hero$id: updateHeroLocation(id: $id, location: "Missing") {
  id
  location
}
```

For non-Docker, running the CLI is as simple as:

```sh
$ gmf mutations -s 10 -e 15 -f hero.txt
```

For Docker, mount the working directory as volume (`-v` flag):

```sh
$ docker run --rm -v $(pwd):/tmp yanyi/generate-multifields:latest \
    mutations -s 10 -e 15 -f /tmp/hero.txt
```

**Tip for Docker users** - Add an `alias` for the Docker run instruction:

```sh
alias gmf="docker run --rm -v $(pwd):/tmp yanyi/generate-multifields:latest"
```

## Commands

### Mutations

As of the current version, `gmf` allows generating for
GraphQL mutations.

Run by entering a start ID and end ID, together with an input file:

```sh
$ gmf mutations -s 10 -e 15 -f hero.txt
```

The CLI generates the output like:

```graphql
mutation {
  hero10: updateHeroLocation(id: 10, location: "Missing") {
    id
    location
  }
  ... (redacted for brevity)
  hero15: updateHeroLocation(id: 15, location: "Missing") {
    id
    location
  }
}
```
