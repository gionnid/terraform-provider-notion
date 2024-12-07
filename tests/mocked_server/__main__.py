import http.server
import json

TOKEN = "secret_1234567890"
VERSION = "2022-06-28"


class Handler(http.server.BaseHTTPRequestHandler):
    version = VERSION
    token = TOKEN

    def fail(self, code, message):
        self.send_response(code)
        self.send_header("Content-type", "application/json")
        self.end_headers()
        response = {"error": message}
        self.wfile.write(json.dumps(response).encode("utf-8"))
        return

    def do_POST(self):

        if self.headers["Notion-Version"] != Handler.version:
            self.fail(400, f"Notion-Version is not {Handler.version}")
            return

        if self.headers["Authorization"] != f"Bearer {Handler.token}":
            self.fail(400, "Unhautorized")
            return

        self.send_response(200)
        self.send_header("Content-type", "application/json")
        self.end_headers()

        try:
            content_length = int(self.headers["Content-Length"])
            post_data = self.rfile.read(content_length)

            response = {"received": json.loads(post_data.decode("utf-8"))}

            self.wfile.write(json.dumps(response).encode("utf-8"))
        except Exception as e:
            self.fail(500, str(e))


class MockedServer:
    def __init__(self):
        self.server = http.server.HTTPServer(
            ("localhost", 8000),
            Handler,
        )

    def start(self):
        self.server.serve_forever()

    def stop(self):
        self.server.shutdown()
        self.server.server_close()


if __name__ == "__main__":
    server = MockedServer()
    server.start()
