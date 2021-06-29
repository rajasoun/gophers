# gophers

Career guide to become a Awesome Go Developer

Project is aimed at creating a study guide for Early in Career developers to learn Go and gain expertise in handling real world projects

## Dev Containers Inside Visual Studio Code

1. Install [Visual Studio Code](https://code.visualstudio.com/download)

1. [Remote-Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

    - In the terminal `code --install-extension ms-vscode-remote.remote-containers`
    - For robot files extension `code --install-extension tomiturtiainen.rf-intellisense`

1. [Docker](https://www.docker.com/)


## Building Dev Containers Outside of Visual Studio Code

In Terminal 

```
source <(curl -s https://raw.githubusercontent.com/rajasoun/common-lib/main/run_ci_shell.sh)
ln -s ${PWD}/.devcontainer vscode-iaac/go/.devcontainer
./ci.sh build -d
```

## Continously Run Go Tests

In Terminal 

```
gotestsum --watch --format testname
```