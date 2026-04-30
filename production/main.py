from fastapi import FastAPI
from dictionary.dictionary import *
from construct.construct import *

app = FastAPI()

@app.get("/")
async def root():
    return {"message": "Hello World"}

@app.get("/individu")
async def individu():
    pop = 25
    pop_mapel, pop_guru = generatePopulation(pop)

    # Konversi ke list agar bisa dijadikan JSON
    pop_mapel_list = pop_mapel.tolist()
    pop_guru_list  = pop_guru.tolist()

    return {
        "message": "ok",
        "data1": pop_mapel_list,
        "data2": pop_guru_list
    }
