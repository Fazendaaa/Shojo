<!-- ⚠️ This README has been generated from the file(s) "docs/blueprint.md" ⚠️--><h1 align="center">Shōjō</h1>

<p align="center">
  <img src="https://raw.githubusercontent.com/Fazendaaa/Shojo/master/assets/img/logo.svg" alt="Logo" width="150" height="150" />
</p>

<p align="center">
		<a href="https://github.com/badges/shields"><img alt="Custom badge" src="https://img.shields.io/badge/custom-badge-f39f37.svg" height="20"/></a>
<a href="https://saythanks.io/to/lucas.carotta%40outlook.com"><img alt="Say Thanks!" src="https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg?longCache=true&style=for-the-badge" height="20"/></a>
<a href=""><img alt="" src="" height="20"/></a>
<a href=""><img alt="" src="" height="20"/></a>
<a href=""><img alt="" src="" height="20"/></a>
<a href=""><img alt="" src="" height="20"/></a>
<a href=""><img alt="" src="" height="20"/></a>
<a href=""><img alt="" src="" height="20"/></a>
	</p>


<p align="center">
  <b>LaTex package manager</b></br>
  <sub><sub>
</p>

<br />


Welcome to Fazendaaa's Shōjō. This is version 0.0.0!

Made with:

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#ideia)

## ➤ Ideia

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


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#why)

## ➤ Why?

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


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#components)

## ➤ Components

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


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#installing)

## ➤ Installing

### Arch Linux/Manjaro

```shell
yay -S shojo
```

or:

```shell
pamac install shojo
```

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


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#uninstalling)

## ➤ Uninstalling

### Arch Linux/Manjaro


```shell
yay -R shojo
```

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


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#cfd)

## ➤ CFD

As `estat` have grown so much and making it available as FOSS (Free and open-source software) trough the [CFD](https://github.com/Fazendaaa/CFD) initiative. That's was always the idea but the project still in development and not having a properly defined scope, I decided to break its main features in other projects:

- [Succubus](https://github.com/Fazendaaa/Succubus): universal package manager based on cloud-native
- [Jinn](https://github.com/Fazendaaa/Jinn): universal project manager built to expand Succubus capabilities
- [Baba Yaga](https://github.com/Fazendaaa/BabaYaga): universal cloud-native manager built to expand Jinn and Succubus capabilities
- [Wendigo](https://github.com/Fazendaaa/Wendigo): universal project translator from cloud-native projects to other infra technologies
- [Shōjō](https://github.com/Fazendaaa/Shojo): LaTex package manager
- [Hellhound](github.com/Fazendaaa/Hellhound): VSCode extension to integrate Jinn recipes
- [Crocotta](github.com/Fazendaaa/Crocotta): SOC assisted guider


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#author)

## ➤ Author

Only [me](https://github.com/Fazendaaa) because the aforementioned project was implemented by yours only. By knowing each line of that code wrote doing the port would be more easily done this way.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#contributing)

## ➤ Contributing

Check more about this in [CONTRIBUTING.md](CONTRIBUTING.md). Here we have a list of some of our contributors:


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#contributors)

## ➤ Contributors
	

| [<img alt="Fazendaaa" src="https://avatars2.githubusercontent.com/u/12137236?s=460&u=75ec76d6f0c577de2ebfa4eae77cc4c4ad17ec06&v=4" width="100">](https://twitter.com/the_fznd) | [<img alt="You?" src="https://joeschmoe.io/api/v1/random" width="100">](https://github.com/andreasbm/web-config/blob/master/CONTRIBUTING.md) |
|:--------------------------------------------------:|:--------------------------------------------------:|
| [Fazendaaa](https://twitter.com/the_fznd)        | [You?](https://github.com/andreasbm/web-config/blob/master/CONTRIBUTING.md) |



[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#todo)

## ➤ TODO

Read [TODO.md](./TODO.md) file.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#references)

## ➤ References

### Docs

- [tlmgr - the native TeX Live Manager](https://www.tug.org/texlive/doc/tlmgr.html)
- [Go YAML](https://zetcode.com/golang/yaml/)

### BlogPosts

- [Set Indentation on New Golang YAML v3 Library](https://hashnode.com/post/set-indentation-on-new-golang-yaml-v3-library-ckbrwl63001skn8s1lj3jmtln)
- [How to Access Interface Fields in Golang?](https://www.geeksforgeeks.org/how-to-access-interface-fields-in-golang/)
- [Build CI/CD pipelines in Go with github actions and Docker](https://dev.to/gopher/build-ci-cd-pipelines-in-go-with-github-actions-and-dockers-1ko7)
- [GitHub Action for release your Go projects as fast and easily as possible](https://dev.to/koddr/github-action-for-release-your-go-projects-as-fast-and-easily-as-possible-20a2)
- [How to test your Go code with Github actions](https://gfgfddgleb.medium.com/how-to-test-your-go-code-with-github-actions-f15881d46089)


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#license)

## ➤ License
	
Licensed under [AGPL-3.0](https://opensource.org/licenses/AGPL-3.0).
