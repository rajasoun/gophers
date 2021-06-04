# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.177.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version: 1, 1.16, 1.15
ARG VARIANT="1.16"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT} 

# Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"

RUN set -x \
    && CMD="umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1" \
    && if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "$CMD"; fi

# OS packages
RUN set -x \
    && apt-get update \
    && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends bash bash-completion \
       vim git netcat jq curl git-flow wget g++ zsh-autosuggestions \
       zsh-syntax-highlighting \
    && apt-get autoremove -y  \
    && apt-get clean -y \
    && rm -r /var/cache/* /var/lib/apt/lists/*

# go get to install packages
RUN set -x \
    && go get -v golang.org/x/tools/gopls 2>&1 \
    && go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs 2>&1 \
    && go get -v github.com/cweill/gotests/gotests 2>&1 \
    && go get -v github.com/josharian/impl 2>&1 \
    && go get -v honnef.co/go/tools/cmd/staticcheck 2>&1 \
    && rm -rf $GOPATH/pkg/* $GOPATH/src/* $HOME/.cache/go-build 

# RUN set -ex \
#     && curl -sL https://raw.githubusercontent.com/taskctl/taskctl/master/install.sh | sh -s -- -b /tmp \
#     && mv /tmp/taskctl $HOME/.local/bin 

RUN set -ex \
    && curl -fsSL https://git.io/shellspec | sh -s -- --yes 
ENV PATH="$HOME/.local/bin:$PATH"

# Install global node packages.
# RUN set -x \
#     && su vscode -c "source /usr/local/share/nvm/nvm.sh"
#     && su vscode -c "npm install -g <your-package-here>" 2>&1



