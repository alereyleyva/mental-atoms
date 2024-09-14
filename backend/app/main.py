from fastapi import FastAPI

app = FastAPI()


@app.get("/health")
async def root():
    return {"message": "Mental atoms API is up and running"}
