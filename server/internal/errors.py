class CypherError(Exception):
    status_code = 500

    @property
    def message(self):
        try:
            return self.args[0]
        except IndexError:
            return repr(self)

    @property
    def json(self):
        try:
            return self.args[1]
        except IndexError:
            return {}


class BadRequest(CypherError):
    status_code = 400


class Unauthorized(CypherError):
    status_code = 401


class Forbidden(CypherError):
    status_code = 403


class NotFound(CypherError):
    status_code = 404


class Ratelimited(CypherError):
    status_code = 429
