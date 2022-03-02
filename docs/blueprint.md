{{ template:title }}

{{ template:logo }}

{{ template:badges }}

{{ template:description }}

Welcome to Fazendaaa's {{ pkg.name }}. This is version {{ pkg.version }}!

Made with:

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)

## Ideia

<div align="center">
  <img src="./assets/gif/demo.gif"/>
</div>

Currently, in the company that I work for we have a CLI (Command Line Interface) made in [Python](https://www.python.org/) called `estat` that you can read more about it right [here](https://github.com/Fazendaaa/Succubus). But one of its features is handeling LaTex packages to each of our projects.

The main ideia is that you have a `shojo.yaml` like the following describing the required packages needed for the project:

```yaml
packages:
  - name: multirow
    revision: 58396
  - name: wrapfig
    revision: 61719
  - name: lastpage
    revision: 60414
  - name: hyphenat
    revision: 15878
  - name: hyphen-portuguese
    revision: 58609
  - name: babel-portuges
    revision: 59883
  - name: fancyhdr
    revision: 57672
  - name: tabu
    revision: 61719
  - name: varwidth
    revision: 24104
```

## Why?

As many that use LaTex in a daily bases know, having to install all packages when you only need a few at most can be troublesome to say at least, besides that:

- Having to maintain an old project
- Many people using a plethora of editors
- Building in a CI/CD environment
- Time consuming
- Updates
- etc.

That's why Shojo was born, to help maintain LaTex projects dependencies in a friendly manner inspired by other tools like:

- [npm](https://www.npmjs.com/)
- [poetry](https://python-poetry.org/)
- [renv](https://rstudio.github.io/renv/index.html)

## Components

What you can do with Shojo:

### install

To install all packages from a Shojo project:

```shell
shojo install
```

### init

To initialize a new Shojo project:

```shell
shojo init
```

### add

To add packages to this project:

```shell
shojo add draftwatermark lastpage tabu ...
```

### rm

To remove packages from this project:

```shell
shojo rm draftwatermark lastpage tabu ...
```

### repo

To change the package's repository:

```shell
shojo repo https://mirror.ctan.org/systems/texlive/tlnet
```

## Installing

### Go

```shell
go install github.com/Fazendaaa/Shojo@latest
```

In case you choose this route, just remember to use `Shojo` instead of `shojo` while using the command.

#### Not loading

Probably missing the following:

```shell
export GOPATH="$HOME/go/"
export PATH="$PATH:$GOPATH/bin/"
```

### Binary

Take a look first at [zyedidia/eget](https://github.com/zyedidia/eget)

```shell
curl https://zyedidia.github.io/eget.sh | sh
./eget Fazendaaa/Shojo
mv Shojo $HOME/.local/bin/shojo
```

### Docker

You don't need to install Go to run this tool, just Docker. And to do so to give it a try, you can do it just by running the following line in your terminal:

```shell
alias shojo='docker run -it --volume ${PWD}:${PWD} --workdir ${PWD} fazenda/shojo-latex'
```

And then running the following to check whether or not is working properly:

```shell
shojo --help
```

After that you can run the following just to fell free to run the following:

```shell
shojo init .
shojo add lastpage
```

## Uninstalling

### Go

```shell
rm $GOPATH/bin/Shojo
```

### Binary

```shell
rm $HOME/.local/bin/shojo
```

### Docker

```shell
docker rmi --force fazenda/shojo-latex
```

## CFD

As `estat` have grown so much and making it available as FOSS (Free and open-source software) was always the idea but the project still in development and not having a properly defined scope, I decided to break its main features in other projects:

- [Succubus](https://github.com/Fazendaaa/Succubus): universal package manager based on cloud-native
- [Jinn](https://github.com/Fazendaaa/Jinn): universal project manager built to expand Succubus capabilities
- [Baba Yaga](https://github.com/Fazendaaa/BabaYaga): universal cloud-native manager built to expand Jinn and Succubus capabilities
- [Wendigo](https://github.com/Fazendaaa/Wendigo): universal project translator from cloud-native projects to other infra technologies
- [Shōjō](https://github.com/Fazendaaa/Shojo): LaTex package manager
- [Hellhound](github.com/Fazendaaa/Hellhound): VSCode extension to integrate Jinn recipes
- [Crocotta](github.com/Fazendaaa/Crocotta): SOC assisted guider

## Author

Only [me](https://github.com/Fazendaaa) because the aforementioned project was implemented by yours only. By knowing each line of that code wrote doing the port would be more easily done this way.

## Contributing

Check more about this in [CONTRIBUTING.md](CONTRIBUTING.md). Here we have a list of some of our contributors:

{{ template:contributors }}

## TODO

Read [TODO.md](./TODO.md) file.

## References

### Docs

- [tlmgr - the native TeX Live Manager](https://www.tug.org/texlive/doc/tlmgr.html)
- [Go YAML](https://zetcode.com/golang/yaml/)

### BlogPosts

- [Set Indentation on New Golang YAML v3 Library](https://hashnode.com/post/set-indentation-on-new-golang-yaml-v3-library-ckbrwl63001skn8s1lj3jmtln)
- [How to Access Interface Fields in Golang?](https://www.geeksforgeeks.org/how-to-access-interface-fields-in-golang/)
- [Build CI/CD pipelines in Go with github actions and Docker](https://dev.to/gopher/build-ci-cd-pipelines-in-go-with-github-actions-and-dockers-1ko7)
- [GitHub Action for release your Go projects as fast and easily as possible](https://dev.to/koddr/github-action-for-release-your-go-projects-as-fast-and-easily-as-possible-20a2)
- [How to test your Go code with Github actions](https://gfgfddgleb.medium.com/how-to-test-your-go-code-with-github-actions-f15881d46089)

{{ template:license }}
