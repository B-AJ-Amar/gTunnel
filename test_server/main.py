from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse
import uvicorn

app = FastAPI()

@app.api_route("/{path:path}", methods=["GET", "POST", "PUT", "DELETE", "PATCH"])
async def catch_all(request: Request, path: str):
    body = None
    if request.headers.get("content-type", "").startswith("application/json"):
        try:
            body = await request.json()
        except Exception:
            body = None
    else:
        raw_body = await request.body()
        body = raw_body if raw_body else None
    return JSONResponse({
        "status": "ok",
        "endpoint": request.url.path,
        "method": request.method,
        "body": body
    })

'''
uvicorn test_server.main:app --host 0.0.0.0 --port 3000 --reload
'''
