import numpy as np
import random
from dictionary.dictionary import (
    jamPerMingguMapel,
    guruValidMapel,
    blokMapel,
    mapelIds,
    kelasIds,
    kelasTingkatan,
    jumlahKelas,
    slotPerKelas
)

# ===========================
# Generate jadwal untuk satu kelas
# ===========================
def generatePerKelas(tingkatan):
    blok_list = []

    for mapel_id, jam in jamPerMingguMapel().items():
        guru_valid = guruValidMapel().get((mapel_id, tingkatan))
        if not guru_valid:
            continue
        guru = random.choice(guru_valid)
        blok = blokMapel()[mapel_id]

        for b in blok:
            blok_list.append((mapel_id, guru, b))

    random.shuffle(blok_list)

    mapel_arr = []
    guru_arr = []

    for mapel_id, guru, panjang in blok_list:
        mapel_arr.extend([mapel_id] * panjang)
        guru_arr.extend([guru] * panjang)

    mapel_arr = np.array(mapel_arr, dtype=np.int16)
    guru_arr  = np.array(guru_arr, dtype=np.int16)

    # pastikan panjang sesuai slot per kelas
    max_slot = slotPerKelas()
    if len(mapel_arr) > max_slot:
        mapel_arr = mapel_arr[:max_slot]
        guru_arr  = guru_arr[:max_slot]

    while len(mapel_arr) < max_slot:
        mapel = random.choice(mapelIds())
        guru_valid = guruValidMapel().get((mapel, tingkatan))
        if guru_valid:
            guru = random.choice(guru_valid)
            mapel_arr = np.append(mapel_arr, mapel)
            guru_arr  = np.append(guru_arr, guru)

    return mapel_arr, guru_arr

# ===========================
# Generate satu individu jadwal lengkap
# ===========================
def individuConstruct():
    n_kelas = jumlahKelas()
    slot_kelas = slotPerKelas()

    mapel_matrix = np.zeros((n_kelas, slot_kelas), dtype=np.int16)
    guru_matrix  = np.zeros((n_kelas, slot_kelas), dtype=np.int16)

    kelas_ids = kelasIds()
    kelas_tingkatan = kelasTingkatan()

    for i, kelas_id in enumerate(kelas_ids):
        tingkatan = kelas_tingkatan[kelas_id]
        mapel_row, guru_row = generatePerKelas(tingkatan)
        mapel_matrix[i] = mapel_row
        guru_matrix[i]  = guru_row

    return mapel_matrix, guru_matrix

# ===========================
# Generate populasi awal
# ===========================
def generatePopulation(POPULASI):
    n_kelas = jumlahKelas()
    slot_kelas = slotPerKelas()

    pop_mapel = np.zeros((POPULASI, n_kelas, slot_kelas), dtype=np.int16)
    pop_guru  = np.zeros((POPULASI, n_kelas, slot_kelas), dtype=np.int16)

    for i in range(POPULASI):
        mapel_matrix, guru_matrix = individuConstruct()
        pop_mapel[i] = mapel_matrix
        pop_guru[i]  = guru_matrix

    return pop_mapel, pop_guru