FROM eraac/golang

ADD cards-generator /cards-generator

CMD ["/cards-generator", "-config", "/config.json"]
