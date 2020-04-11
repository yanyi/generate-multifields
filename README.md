# generate-multifields

![Staging](https://github.com/yanyi/generate-multifields/workflows/Staging/badge.svg?branch=master)

`generate-multifields` is a CLI library to help generate multiple fields of
a given input format of your GraphQL queries or mutations, by repeating them for
a number of times.

## Intentions of this CLI

There was an instance I had to generate hundreds of queries or mutations that
run by IDs. I needed a tool to generate them with a given input.

>"But you could do this with a small shell script!"

Yes. It was also intended for me to play with [Cobra](https://github.com/spf13/cobra)
and GitHub Actions.

## Running

### Go

Install the CLI using:

```sh
go install
```

After that, use the CLI:

```sh
generate-multifields mutations -s 10 -e 15 -f hero.txt
```

### Docker

Excluding the `docker pull` command, you can run this immediately:

```sh
docker run --rm -v $(pwd):/usr/ yanyi/generate-multifields:latest \
    mutations -s 10 -e 15 -f /usr/hero.txt
```

## Usage

### Mutations

As of the current version, `generate-multifields` allows generating for
GraphQL mutations.

Run by entering a start ID and end ID, together with an input file:

```sh
generate-multifields mutations -s 10 -e 15 -f hero.txt
```

What the CLI does is to look for instances of `$id` in a given text and replace
it with the IDs. Here is an example of a text file:

```txt
hero$id: updateHeroLocation(id: $id, location: "Missing") {
  id
  location
}
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
