FROM eraac/golang

ADD api-stats /api-stats

CMD ["/api-stats", "-config", "/config.json"]

