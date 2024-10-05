import http.server
import json
import socket


class Handler(http.server.BaseHTTPRequestHandler):
    def do_POST(self):
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()
        
        content_length = int(self.headers['Content-Length'])
        post_data = self.rfile.read(content_length)
        
        response = {'received': json.loads(post_data.decode('utf-8'))}
        
        self.wfile.write(json.dumps(response).encode('utf-8'))


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
