# base
FROM docker:19.03.15-dind AS base
ARG USER=ci-shell
ENV TZ=Asia/Kolkata
ENV HOME=/home/$USER
ENV HOSTNAME="ci-shell"
ENV ZSH_THEME="powerlevel10k/powerlevel10k"


RUN set -ex \
    && apk add --no-cache --update sudo bash bash-completion \
       git vim zsh-autosuggestions zsh-syntax-highlighting bind-tools curl  \
       python3 build-base musl py3-pip python3-dev binutils-dev curl-dev elfutils-libelf \
       jq libc6-compat netcat-openbsd tzdata \
    && rm -rf /var/cache/apk/*

# add new user
RUN set -ex \
    && adduser -D $USER \
    && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
    && chmod 0440 /etc/sudoers.d/$USER

# kcov
FROM alpine:latest as kcov
RUN set -ex \
    && apk add --no-cache build-base cmake ninja python3 \
        binutils-dev curl-dev elfutils-dev
WORKDIR /root
ENV KCOV=https://github.com/SimonKagstrom/kcov/archive/v38.tar.gz
RUN set -ex \
    && wget -q $KCOV -O - | tar xz -C ./ --strip-components 1
RUN set -ex \
    && mkdir build && cd build \
    && CXXFLAGS="-D__ptrace_request=int" cmake -G Ninja .. \
    && cmake --build . --target install

# gh-cli
FROM alpine:latest as gh-cli
WORKDIR /root
ENV GH_CLI=https://github.com/cli/cli/releases/download/v1.11.0/gh_1.11.0_linux_amd64.tar.gz
RUN set -ex \
    && wget -q $GH_CLI -O - | tar xz -C ./ --strip-components 1

# node.js packages
FROM node:alpine as node
RUN set -ex \
    && npm install -g npm@7.16.0 \
    && npm install -g husky commitizen commitlint release-it

# ci-shell
FROM base
USER $USER
WORKDIR $HOME/$USER

# shellspec
RUN set -ex \
    && curl -fsSL https://git.io/shellspec | sh -s -- --yes
ENV PATH="$HOME/.local/bin:$PATH"

# oh-my-zsh
RUN set -ex \
    && sh -c "$(wget https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O -)"
RUN set -ex \
    && echo "source /usr/share/zsh/plugins/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh" >> ~/.zshrc \
    && echo "source /usr/share/zsh/plugins/zsh-autosuggestions/zsh-autosuggestions.zsh" >> ~/.zshrc
RUN set -ex \
    && echo "User $(whoami) running from $PWD with premissions: $(sudo -l)"

# kcov
COPY --from=kcov /usr/local/bin/kcov* /usr/local/bin/
COPY --from=kcov /usr/local/share/doc/kcov /usr/local/share/doc/kcovub

# github cli
COPY --from=gh-cli /root/bin/gh /usr/local/bin/
COPY --from=gh-cli /root/share  /usr/local/share/

# node.js
COPY --from=node /usr/local/ /usr/local

# python packages
RUN set -ex \
    && pip install cruft ggshield pipx pre-commit

# Theme Dot Files
RUN set -ex \
    && git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ~/powerlevel10k \
    && echo "source ~/powerlevel10k/powerlevel10k.zsh-theme"  >> $HOME/.zshrc \
    && echo "# To customize prompt, run `p10k configure` or edit ~/.p10k.zsh."  >> $HOME/.zshrc \
    && echo "[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh" >> $HOME/.zshrc

COPY ci-shell/dotfiles/.p10k.zsh $HOME/.p10k.zsh
COPY ci-shell                    /workspaces/ci-shell

WORKDIR /workspaces

# ENTRYPOINT /bin/zsh
CMD /bin/zsh "$@"
