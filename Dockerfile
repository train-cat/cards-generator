FROM eraac/golang

ADD cards-generator /cards-generator

COPY fonts /fonts

CMD ["/cards-generator", "-config", "/config.json"]
