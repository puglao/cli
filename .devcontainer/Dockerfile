FROM mcr.microsoft.com/devcontainers/go:1-1.21-bookworm

RUN curl -sL https://github.com/pre-commit/pre-commit/releases/download/v3.3.3/pre-commit-3.3.3.pyz -o /usr/local/bin/pre-commit && \
    chmod +x /usr/local/bin/pre-commit

ARG ACT_VERSION=v0.2.49
RUN case "$(uname -m)" in \
        "x86_64") \
            ARCH="x86_64" ;; \
        "aarch64") \
            ARCH=;"arm64";; \
        *) \
            exit 1;; \
    esac && \
    curl -sL https://github.com/nektos/act/releases/download/${ACT_VERSION}/act_Linux_${ARCH}.tar.gz | tar xz && \
    mv act /usr/local/bin/act

ARG GORELEASER_VERSION=v1.20.0
RUN case "$(uname -m)" in \
        "x86_64") \
            ARCH="x86_64" ;; \
        "aarch64") \
            ARCH=;"arm64";; \
        *) \
            exit 1;; \
    esac && \
    curl -sL https://github.com/goreleaser/goreleaser/releases/download/${GORELEASER_VERSION}/goreleaser_Linux_${ARCH}.tar.gz | tar xz && \
    mv act /usr/local/bin/goreleaser
