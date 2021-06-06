# base
FROM docker:19.03.15-dind AS base
ARG USER=ci-shell
ENV HOME /home/$USER

RUN set -ex \
    && apk add --no-cache --update sudo bash bash-completion \
       git vim zsh-autosuggestions zsh-syntax-highlighting bind-tools curl  \
       python3 binutils-dev curl-dev elfutils-libelf jq libc6-compat netcat-openbsd \
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
ENV GH_CLI=https://github.com/cli/cli/releases/download/v1.10.3/gh_1.10.3_linux_amd64.tar.gz
RUN set -ex \
    && wget -q $GH_CLI -O - | tar xz -C ./ --strip-components 1

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

WORKDIR /workspaces
ENTRYPOINT /bin/zsh
CMD /bin/zsh "$@"